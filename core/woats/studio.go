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

func newStudio(group wtype.FactorGroup, matrix ThreeDimensionalMatrix, strategy []Strategy, rotationMax int) *Studio {
	return &Studio{
		matrix:      matrix,
		process:     []wtype.Factor{},
		todo:        group,
		finish:      []wtype.Factor{},
		fgTotal:     len(group),
		snapshot:    map[string]*Studio{},
		rotationMax: rotationMax,
		rotation:    0,
		strategy:    strategy,
	}
}

// Studio 工作空间
type Studio struct {
	matrix   ThreeDimensionalMatrix // 三维课表矩阵(引用)
	strategy []Strategy

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

// PushSameFactor 放置一个相同特征的因子到特定课位，无可用因子时候返回错误信息
func (slf *Studio) PushSameFactor(factor wtype.Factor, slot int) exception.Exception {
	var target wtype.Factor = nil
	for _, f := range append(slf.process, slf.todo...) {
		if f == factor {
			continue
		}
		if f.GetUniqueSign() == factor.GetUniqueSign() && f.GetCourse() == factor.GetCourse() {
			target = f
			break
		}
	}
	if target == nil {
		// 尝试从其他课位拔取
		for _, f := range slf.finish {
			if f.GetUniqueSign() == factor.GetUniqueSign() && f.GetCourse() == factor.GetCourse() {
				if f.GetTimeSlot().Index == slot || f == factor {
					continue
				}
				target = f
				if err := slf.FactorPop(f, f.GetTimeSlot().Index); err != nil {
					return err
				}
				break
			}
		}
		if target != nil {
			if err := slf.FactorPop(target, target.GetTimeSlot().Index); err != nil {
				return err
			}
		} else {
			return define.NotFoundException.Hit().
				Supplement("err", "not found same factor").
				Supplement("class", factor.GetUniqueSign()).
				Supplement("course", factor.GetCourse()).
				Supplement("teacher", factor.GetTeacher())
		}
	}

	for _, f := range slf.matrix.GetConflictFactor(target, slot) {
		if err := slf.FactorPop(f, slot); err != nil {
			return err
		}
	}

	if len(slf.todo)+len(slf.process)+len(slf.finish) > slf.fgTotal {
		panic("!!!")
	}
	if err := slf.FactorPush(target, slot); err != nil {
		return err
	}

	slf.todo = slf.todo.RemoveFactor(target)
	slf.process = slf.process.RemoveFactor(target)
	slf.finish = append(slf.finish, target)
	slf.finish = slf.finish.Unrepeated()

	return nil
}

// PopSameFactor 在特定课位弹出一个相同特征的因子，无可用因子时候返回错误信息
func (slf *Studio) PopSameFactor(factor wtype.Factor, slot int) exception.Exception {
	var target wtype.Factor = nil
	for _, f := range slf.finish {
		if f == factor {
			continue
		}
		if f.GetUniqueSign() == factor.GetUniqueSign() && f.GetCourse() == factor.GetCourse() {
			target = f
			break
		}
	}
	if target == nil {
		return define.NotFoundException.Hit().
			Supplement("err", "not found same factor").
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse()).
			Supplement("teacher", factor.GetTeacher())
	}

	if err := slf.FactorPop(target, slot); err != nil {
		return err
	}

	//slf.todo = append(slf.todo, target)
	//slf.process = slf.process.RemoveFactor(target)
	//slf.finish = slf.finish.RemoveFactor(target)

	return nil
}

// GetDifficult 获取当前排课困难指数
func (slf *Studio) GetDifficult() float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slf.linger)/((float64(slf.fgTotal))*1.5)), 64)
	if value > 1 {
		return 1
	}
	return value
}

// addFactorGroup 追加排课因子
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
		slf.strategy = snapshot.strategy
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
		strategy:     slf.strategy,
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
func (slf *Studio) FactorPop(factor wtype.Factor, slot int) exception.Exception {
	if factor.IsDisableChange() {
		return define.DisableSlotException.Hit().
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse())
	}
	if factor.IsNoChange() {
		// 命中0～100000中产生的少部分随机数，则允许调整
		if !(rand.Intn(1000001) <= int(1000000-(slf.GetDifficult()*1000000))) {
			return define.NoChangeException.Hit().
				Supplement("class", factor.GetUniqueSign()).
				Supplement("course", factor.GetCourse())
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
			for _, strategy := range slf.strategy {
				strategy.OnPop(f, f.GetSlotMap()[slot], slf)
			}
		} else {
			replace = append(replace, f)
		}
	}

	slf.matrix[factor.GetUniqueSign()][slot] = replace
	if count > 0 && factor.IsGroup() {
		for _, f := range factor.GetGroup() {
			if f.GetTimeSlot() != nil {
				if err := slf.FactorPop(f, f.GetTimeSlot().Index); err != nil {
					return err
				}
				for _, strategy := range slf.strategy {
					strategy.OnPop(f, f.GetSlotMap()[slot], slf)
				}
			}
		}
	}

	return nil
}

// FactorPush 将特定因子放置到特定课位
func (slf *Studio) FactorPush(factor wtype.Factor, slot int) exception.Exception {
	if factor.IsGroup() {
		min, max := slf.matrix.GetGroupSlotIndex(factor.GetGroup(), slot)
		if slot >= min || slot <= max {
			var push = func(factor wtype.Factor, slot int) {
				slf.matrix[factor.GetUniqueSign()][slot] = append(slf.matrix[factor.GetUniqueSign()][slot], factor)
				for _, timeSlot := range factor.GetSlot() {
					if timeSlot.Index == slot {
						factor.(*wtype.FactorInfo).TimeSlot = timeSlot
						for _, strategy := range slf.strategy {
							strategy.OnPush(factor, timeSlot, slf)
						}
						log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "Group Push!")
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
	} else {
		slf.matrix[factor.GetUniqueSign()][slot] = append(slf.matrix[factor.GetUniqueSign()][slot], factor)
		for _, timeSlot := range factor.GetSlot() {
			if timeSlot.Index == slot {
				factor.(*wtype.FactorInfo).TimeSlot = timeSlot
				for _, strategy := range slf.strategy {
					strategy.OnPush(factor, timeSlot, slf)
				}
				log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "Push!")
				break
			}
		}
	}
	return nil
}

// FactorMove 将特定因子从一个课位移动到另一个课位(不考虑冲突)
func (slf *Studio) FactorMove(factor wtype.Factor, slot int, targetSlot int) exception.Exception {
	if factor.IsDisableChange() {
		return define.DisableSlotException.Hit().
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse())
	}
	if factor.IsNoChange() {
		// 命中0～100000中产生的少部分随机数，则允许调整
		if !(rand.Intn(1000001) <= int(1000000-(slf.GetDifficult()*1000000))) {
			return define.NoChangeException.Hit().
				Supplement("class", factor.GetUniqueSign()).
				Supplement("course", factor.GetCourse())
		}
		factor.SetNoChange(false)
	}
	log.Println(factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher(), "=>", slot, "=>", targetSlot)
	if !factor.IsGroup() {
		if err := slf.FactorPop(factor, slot); err != nil {
			return err
		}
		return slf.FactorPush(factor, targetSlot)
	} else {
		min, max := slf.matrix.GetGroupSlotIndex(factor.GetGroup(), slot)
		if slot >= min || slot <= max {
			if err := slf.FactorPop(factor, slot); err != nil {
				return err
			}
			return slf.FactorPush(factor, targetSlot)
		}
		return define.FactorGroupPushException.Hit().
			Supplement("err", fmt.Sprintf("slot min and max is: %d ~ %d", min, max)).
			Supplement("slot", slot).
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse())
	}
}

// Run 开始执行
func (slf *Studio) Run(handle func(factor wtype.Factor, studio *Studio) bool) {
	var (
		successHandle = func(factor ...wtype.Factor) {
			slf.todo = slf.todo.RemoveFactor(factor...)
			slf.finish = append(slf.finish, factor...)
			slf.process = slf.process.RemoveFactor(factor...)
		}
		failedHandle = func(factor ...wtype.Factor) {
			slf.todo = append(slf.todo, factor...)
			slf.finish = slf.finish.RemoveFactor(factor...)
			slf.process = slf.process.RemoveFactor(factor...)
		}
	)

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
				successHandle(factor)
				successHandle(factor.GetGroup()...)
			} else {
				failedHandle(factor)
				failedHandle(factor.GetGroup()...)
			}

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
			slf.log(factor)

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
