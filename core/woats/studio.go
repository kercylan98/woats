package woats

import (
	"fmt"
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/define"
	"github.com/kercylan98/woats/core/woats/wtype"
	"log"
	"math/rand"
	"strconv"
)

func newStudio(group wtype.FactorGroup, matrix ThreeDimensionalMatrix, rotationMax int) *Studio {
	return &Studio{
		matrix:      matrix,
		process:     []wtype.Factor{},
		todo:        group,
		finish:      []wtype.Factor{},
		fgTotal:     len(group),
		snapshot:    map[string]*Studio{},
		rotationMax: rotationMax,
		rotation:    0,
	}
}

// Studio 工作空间
type Studio struct {
	matrix ThreeDimensionalMatrix // 三维课表矩阵(引用)

	process wtype.FactorGroup // 处理中
	todo    wtype.FactorGroup // 待处理
	finish  wtype.FactorGroup // 已放置

	fgTotal      int  // 排课因子总数
	rotationMax  int  // 待处理和处理中状态转换次数上限(避免死循环排课)
	rotation     int  // 待处理和处理中状态转换次数
	lockRotation bool // 为true时rotation将不生效
	logout       bool // 是否开启日志输出

	snapshot map[string]*Studio // 快照

	linger    int // 徘徊次数，当剩余数量减少又增加的时候记为一次徘徊
	remainder int // 上一次剩余数量计数
}

// GetDifficult 获取当前排课困难指数
func (slf *Studio) GetDifficult() float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slf.linger)/((float64(slf.fgTotal))*1.5)), 64)
	if value > 1 {
		return 1
	}
	return value
}

// 追加排课因子
func (slf *Studio) addFactorGroup(group wtype.FactorGroup) {
	slf.todo = append(slf.todo, group...)
	slf.fgTotal += len(group)
	slf.rotation = 0
}

// LockRotation 锁定状态转化次数
func (slf *Studio) LockRotation() {
	slf.lockRotation = true
}

// UnLockRotation 解锁状态转化次数
func (slf *Studio) UnLockRotation() {
	slf.lockRotation = false
}

// GetFactorTotal 获取因子总数
func (slf *Studio) GetFactorTotal() int {
	return slf.fgTotal
}

// GetProcess 获取正在处理的数量
func (slf *Studio) GetProcess() int {
	return len(slf.process)
}

// GetFinish 获取已放置数量
func (slf *Studio) GetFinish() int {
	return len(slf.finish)
}

// GetTodo 获取待办数量
func (slf *Studio) GetTodo() int {
	return len(slf.todo)
}

// RecoverySnapshot 恢复快照
func (slf *Studio) RecoverySnapshot(name string) bool {
	if snapshot := slf.snapshot[name]; snapshot != nil {
		slf.matrix = snapshot.matrix
		slf.process = snapshot.process
		slf.todo = snapshot.todo
		slf.finish = snapshot.finish
		slf.fgTotal = snapshot.fgTotal
		slf.rotationMax = snapshot.rotationMax
		slf.rotation = snapshot.rotation
		slf.lockRotation = snapshot.lockRotation
		slf.logout = snapshot.logout
		slf.linger = snapshot.linger
		slf.remainder = snapshot.remainder
		delete(slf.snapshot, name)
		return true
	}
	return false
}

// clone 克隆
func (slf *Studio) clone() *Studio {
	return &Studio{
		matrix:       slf.matrix.Clone(),
		process:      slf.process.Clone(),
		todo:         slf.todo.Clone(),
		finish:       slf.finish.Clone(),
		fgTotal:      slf.fgTotal,
		rotationMax:  slf.rotationMax,
		rotation:     slf.rotation,
		lockRotation: slf.lockRotation,
		logout:       slf.logout,
		linger:       slf.linger,
		remainder:    slf.remainder,
	}
}

// CreateSnapshot 创建快照
func (slf *Studio) CreateSnapshot(name string) {
	clone := slf.clone()
	slf.snapshot[name] = clone
	clone.snapshot = slf.snapshot
}

// DeleteSnapshot 删除快照
func (slf *Studio) DeleteSnapshot(match func(name string) bool) {
	var deletes []string
	for s, _ := range slf.snapshot {
		if match(s) {
			deletes = append(deletes, s)
		}
	}
	for _, s := range deletes {
		delete(slf.snapshot, s)
	}
}

// FactorPop 将特定因子从特定课位中弹出到代办中的前面
func (slf *Studio) FactorPop(factor wtype.Factor, slot int) {
	if factor.IsDisableChange() {
		return
	}
	if factor.IsNoChange() {
		// 命中0～100000中产生的少部分随机数，则允许调整
		if !(rand.Intn(1000001) <= int(1000000-(slf.GetDifficult()*1000000))) {
			return
		}
		factor.SetNoChange(false)
	}
	log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "Pop!")

	var replace = make(wtype.FactorGroup, 0)
	var replaceFinish = make(wtype.FactorGroup, 0)

	count := 0
	for i := 0; i < len(slf.matrix[factor.GetUniqueSign()][slot]); i++ {
		f := slf.matrix[factor.GetUniqueSign()][slot][i]
		if f == factor {
			slf.todo = append(slf.todo, f)
			// 查找finish
			for _, finish := range slf.finish {
				if finish != f {
					replaceFinish = append(replaceFinish, finish)
				}
			}
			slf.finish = replaceFinish
			// 从课表中拔出
			for _, mf := range slf.matrix[factor.GetUniqueSign()][slot] {
				if mf != f {
					mf.(*wtype.FactorInfo).TimeSlot = nil
					replace = append(replace, mf)
				}
			}
			slf.matrix[factor.GetUniqueSign()][slot] = replace
			count++
		} else {
			replace = append(replace, f)
		}
	}

	slf.matrix[factor.GetUniqueSign()][slot] = replace
	if count > 0 && factor.IsGroup() {
		for _, f := range factor.GetGroup() {
			if f.GetTimeSlot() != nil {
				slf.FactorPop(f, f.GetTimeSlot().Index)
			}
		}
	}
}

// FactorGroupPush 放置连堂课
func (slf *Studio) FactorGroupPush(factor wtype.Factor, slot int) exception.Exception {
	if !factor.IsGroup() {
		return define.FactorGroupPushException.Hit().
			Supplement("factor", "not is a group").
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse())
	}

	min, max := slf.matrix.GetGroupSlotIndex(factor.GetGroup(), slot)
	if slot >= min || slot <= max {
		var push = func(factor wtype.Factor, slot int) {
			log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "Group Posh!")
			slf.matrix[factor.GetUniqueSign()][slot] = append(slf.matrix[factor.GetUniqueSign()][slot], factor)
			for _, timeSlot := range factor.GetSlot() {
				if timeSlot.Index == slot {
					factor.(*wtype.FactorInfo).TimeSlot = timeSlot
					break
				}
			}
		}
		push(factor, slot)
		slot++
		for _, f := range factor.GetGroup() {
			push(f, slot)
			slot++
		}
		return nil
	}
	return define.FactorGroupPushException.Hit().
		Supplement("err", fmt.Sprintf("slot min and max is: %d ~ %d", min, max)).
		Supplement("slot", slot).
		Supplement("class", factor.GetUniqueSign()).
		Supplement("course", factor.GetCourse())
}

// FactorPush 将特定因子放置到特定课位
func (slf *Studio) FactorPush(factor wtype.Factor, slot int) {
	if factor.IsGroup() {
		panic("please use FactorGroupPush!")
	}
	log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "Posh!")

	slf.matrix[factor.GetUniqueSign()][slot] = append(slf.matrix[factor.GetUniqueSign()][slot], factor)
	for _, timeSlot := range factor.GetSlot() {
		if timeSlot.Index == slot {
			factor.(*wtype.FactorInfo).TimeSlot = timeSlot
			break
		}
	}

}

// FactorMove 将特定因子从一个课位移动到另一个课位(不考虑冲突)
func (slf *Studio) FactorMove(factor wtype.Factor, slot int, targetSlot int) {
	if factor.IsDisableChange() {
		return
	}
	if factor.IsNoChange() {
		// 命中0～100000中产生的少部分随机数，则允许调整
		if !(rand.Intn(1000001) <= int(1000000-(slf.GetDifficult()*1000000))) {
			return
		}
		factor.SetNoChange(false)
	}
	log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "=>", targetSlot)
	slf.FactorPop(factor, slot)
	slf.FactorPush(factor, targetSlot)
}

// Run 开始执行
func (slf *Studio) Run(handle func(factor wtype.Factor, studio *Studio) bool) {
	// 真实循环次数
	var realLoopCount = 0
	for {
		// 当处理中的因子组为空且代办不为空的时候，交换处理中的因子组
		if len(slf.process) == 0 && len(slf.todo) > 0 {
			if slf.rotation == slf.rotationMax && !slf.lockRotation {
				slf.log()
				return
			}
			slf.process = slf.todo.SortWithPriorityAVE()
			slf.todo = []wtype.Factor{}
			if !slf.lockRotation {
				slf.rotation++
			}
			realLoopCount++
			log.Println("Loop count added, now real loop count is:", realLoopCount)
		} else if len(slf.finish) == slf.fgTotal || len(slf.process) == 0 {
			slf.log()
			return
		} else {
			factor := slf.process[0]
			if success := handle(factor, slf); success {
				if !factor.IsGroup() {
					slf.finish = append(slf.finish, factor)
				} else {
					for _, f := range factor.GetGroup() {
						slf.finish = append(slf.finish, f)
					}
				}
			} else {
				if !factor.IsGroup() {
					slf.todo = append(slf.todo, factor)
				} else {
					for _, f := range factor.GetGroup() {
						slf.todo = append(slf.todo, f)
					}
				}
			}
			if !factor.IsGroup() {
				slf.process = slf.process[1:]
			} else {
				slf.process = slf.process[1:]
				var replace = make(wtype.FactorGroup, 0)
				for _, process := range slf.process {
					add := true
					for _, f := range factor.GetGroup() {
						if process == f {
							add = false
							break
						}
					}
					if add {
						replace = append(replace, process)
					}
				}
				slf.process = replace
			}
			slf.log(factor)

			// 徘徊检测
			var finishCount = len(slf.finish)
			if finishCount <= slf.remainder {
				slf.linger++
				log.Println("The process is wandering, Linger:", slf.linger, "Difficult:", slf.GetDifficult())
			}
			slf.remainder = len(slf.finish)

			// 死循环怀疑检测
			if realLoopCount > 100 {
				log.Println("There may be an infinite loop that is not unlocked!!!!!!")
			}
		}
	}
}

// GetMatrix 获取课表矩阵
func (slf *Studio) GetMatrix() ThreeDimensionalMatrix {
	return slf.matrix
}

// GetProgress 获取当前进度
func (slf *Studio) GetProgress() float64 {
	if len(slf.finish) == 0 {
		return 0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(len(slf.finish))/float64(slf.fgTotal)*100.0), 64)
	return value
}

// log 打印日志记录
func (slf *Studio) log(factor ...wtype.Factor) {
	if !slf.logout {
		return
	}
	if len(factor) > 0 {
		var f = factor[0]
		log.Println(fmt.Sprintf("%.2f%s [%s] %s Process:%4d Todo:%4d Unfinish:%4d Finish:%4d Total:%4d",
			slf.GetProgress(), "%", f.GetUniqueSign(), f.GetCourse(),
			len(slf.process), len(slf.todo), len(slf.todo)+len(slf.process), len(slf.finish), slf.fgTotal))
	} else {
		log.Println(fmt.Sprintf("%.2f%s Process:%4d Todo:%4d Unfinish:%4d Finish:%4d Total:%4d",
			slf.GetProgress(), "%",
			len(slf.process), len(slf.todo), len(slf.todo)+len(slf.process), len(slf.finish), slf.fgTotal))
	}
}
