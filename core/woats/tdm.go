package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/define"
	"github.com/kercylan98/woats/core/woats/utils"
	"github.com/kercylan98/woats/core/woats/wtype"
	"math"
)

// ThreeDimensionalMatrix 三维矩阵课表模型数据结构
type ThreeDimensionalMatrix map[string] /*class*/ map[int] /*slot*/ wtype.FactorGroup

func newTDM(factors wtype.FactorGroup) ThreeDimensionalMatrix {
	var tdm ThreeDimensionalMatrix = map[string]map[int]wtype.FactorGroup{}

	for _, factor := range factors {
		slots, exist := tdm[factor.GetUniqueSign()]
		if !exist {
			slots = map[int]wtype.FactorGroup{}
		}
		for _, slot := range factor.GetSlot() {
			slots[slot.Index] = []wtype.Factor{}
		}
		tdm[factor.GetUniqueSign()] = slots
	}
	return tdm
}

// GetPlaceSlotWithCourse 根据课程获取特定因子同班级下已放置的课位
func (slf ThreeDimensionalMatrix) GetPlaceSlotWithCourse(factor wtype.Factor) []*wtype.TimeSlot {
	var slots []*wtype.TimeSlot
	for slot, fg := range slf[factor.GetUniqueSign()] {
		for _, f := range fg {
			if f.GetCourse() == factor.GetCourse() {
				slots = append(slots, factor.GetSlotWithIndex(slot))
				break
			}
		}
	}
	return slots
}

// GetPlaceSlotWithTeacher 根据教师获取特定因子同班级下已放置的课位(即便只有一名教师满足)
func (slf ThreeDimensionalMatrix) GetPlaceSlotWithTeacher(factor wtype.Factor) []*wtype.TimeSlot {
	var slots []*wtype.TimeSlot
	for slot, fg := range slf[factor.GetUniqueSign()] {
		for _, f := range fg {
			found := false
			for _, s := range f.GetTeacher() {
				if utils.IsContainString(factor.GetTeacher(), s) {
					found = true
					break
				}
			}
			if found {
				slots = append(slots, factor.GetSlotWithIndex(slot))
				break
			}
		}
	}
	return slots
}

// Clone 克隆一份
func (slf ThreeDimensionalMatrix) Clone() ThreeDimensionalMatrix {
	var tdm ThreeDimensionalMatrix = map[string]map[int]wtype.FactorGroup{}
	for class, slots := range slf {
		tdm[class] = map[int]wtype.FactorGroup{}
		for index, factorGroup := range slots {
			tdm[class][index] = []wtype.Factor{}
			for _, factor := range factorGroup {
				tdm[class][index] = append(tdm[class][index], factor)
			}
		}
	}
	return tdm
}

// GetAllFactor 获取所有因子
func (slf ThreeDimensionalMatrix) GetAllFactor() wtype.FactorGroup {
	var result = make(wtype.FactorGroup, 0)
	for _, slot := range slf {
		for _, group := range slot {
			result = append(result, group...)
		}
	}
	return result
}

// GetAllSlot 获取所有班级课位
func (slf ThreeDimensionalMatrix) GetAllSlot() []*wtype.TimeSlot {
	var result = make([]*wtype.TimeSlot, 0)
	var existed = map[string]bool{}
	for _, factor := range slf.GetAllFactor() {
		if _, exist := existed[factor.GetUniqueSign()]; exist {
			continue
		}
		existed[factor.GetUniqueSign()] = true
		result = append(result, factor.GetSlot()...)
	}
	return result
}

// GetSlotExcludeClass 获取不包含特定班级的所有课位
func (slf ThreeDimensionalMatrix) GetSlotExcludeClass(class ...string) []*wtype.FactorTimeSlot {
	var result = make([]*wtype.FactorTimeSlot, 0)
	var existed = map[string]bool{}
	for _, factor := range slf.GetAllFactor() {
		for _, s := range class {
			if factor.GetUniqueSign() == s {
				continue
			}
		}
		if _, exist := existed[factor.GetUniqueSign()]; exist {
			continue
		}
		existed[factor.GetUniqueSign()] = true
		for _, slot := range factor.GetSlot() {
			result = append(result, &wtype.FactorTimeSlot{
				Factor: factor,
				Slot:   slot,
			})
		}
	}
	return result
}

// GetSlots 获取特定班级的所有课位及存放的课程
func (slf ThreeDimensionalMatrix) GetSlots(class string) map[int]wtype.FactorGroup {
	return slf[class]
}

// GetTimeSlots 获取特定班级的所有课位
func (slf ThreeDimensionalMatrix) GetTimeSlots(class string) []*wtype.TimeSlot {
	for _, group := range slf[class] {
		for _, factor := range group {
			if factor.GetUniqueSign() == class {
				return factor.GetSlot()
			}
		}
	}
	return []*wtype.TimeSlot{}
}

// GetConflictFactorAll 获取所有与特定因子冲突的课位及因子映射(不包含该因子所在课位，可指定最大返回课位数量)
func (slf ThreeDimensionalMatrix) GetConflictFactorAll(factor wtype.Factor, maxSlot ...int) map[int]wtype.FactorGroup {
	var max = math.MinInt
	if len(maxSlot) > 0 {
		max = maxSlot[0]
	}

	var conflict = map[int]wtype.FactorGroup{}

	for _, slot := range factor.GetSlotMap() {
		if slf.FactorIsExistTo(factor, slot.Index) {
			continue
		}

		if cs := slf.GetConflictFactor(factor, slot.Index); len(cs) > 0 {
			conflict[slot.Index] = cs
		}

		if len(conflict) == max {
			return conflict
		}
	}
	return conflict
}

// GetFactors 获取特定班级特定课位号上的因子
func (slf ThreeDimensionalMatrix) GetFactors(class string, slot int) wtype.FactorGroup {
	return slf[class][slot]
}

// GetConflictFactor 获取特定课位上不包含自身的冲突因子
func (slf ThreeDimensionalMatrix) GetConflictFactor(factor wtype.Factor, slot int) wtype.FactorGroup {
	var conflict = make(wtype.FactorGroup, 0)

	for class, _ := range slf {
		for _, s := range slf.GetTimeSlots(class) {
			if s.IsCrossed(factor.GetSlotWithIndex(slot)) {

				for _, sf := range slf[class][s.Index] {
					if sf.GetCourse() != factor.GetCourse() {
						pass := true
						for _, t := range factor.GetTeacher() {
							if utils.IsContainString(sf.GetTeacher(), t) {
								pass = false
								break
							}
						}
						if pass {
							for _, s := range factor.GetStudent() {
								if utils.IsContainString(sf.GetStudentUniqueSign(), s.UniqueSign) {
									pass = false
									break
								}
							}
						}
						if pass {
							for _, p := range factor.GetPlace() {
								if utils.IsContainString(sf.GetPlaceUniqueSign(), p.UniqueSign) {
									pass = false
									break
								}
							}
						}
						if !pass {
							conflict = append(conflict, sf)
						}

					}
				}

			}
		}
	}

	return conflict
}

// FactorIsExistTo 检查特定因子是否存在特定课位上
func (slf ThreeDimensionalMatrix) FactorIsExistTo(factor wtype.Factor, slot int) bool {
	for _, f := range slf[factor.GetUniqueSign()][slot] {
		if f.GetCourse() == factor.GetCourse() {
			return true
		}
	}
	return false
}

// IsConflict 检查一个因子在特定课位是否冲突
func (slf ThreeDimensionalMatrix) IsConflict(factor wtype.Factor, slot int) bool {
	return slf.IsConflictErr(factor, slot) != nil
}

// IsConflictErr 检查一个因子在特定课位是否冲突
func (slf ThreeDimensionalMatrix) IsConflictErr(factor wtype.Factor, slot int) exception.Exception {
	slfSlot := factor.GetSlotWithIndex(slot)
	// 检查该课位是否禁排
	if utils.IsContainInt(factor.GetDisable(), slot) {
		return define.DisableSlotException.Hit().
			Supplement("class", factor.GetUniqueSign()).
			Supplement("course", factor.GetCourse()).
			Supplement("teacher", factor.GetTeacher())
	}
	// 检查该课位是否存在该课程
	for _, f := range slf.GetFactors(factor.GetUniqueSign(), slot) {
		if f.GetCourse() == factor.GetCourse() {
			return define.CourseConflictException.Hit().
				Supplement("class", factor.GetUniqueSign()).
				Supplement("course", factor.GetCourse()).
				Supplement("teacher", factor.GetTeacher())
		}
	}

	for _, timeSlot := range slf.GetSlotExcludeClass(factor.GetUniqueSign()) {
		// 满足时间交叉时候进行冲突检查
		if slfSlot.IsCrossed(timeSlot.Slot) {
			// 教师冲突
			for _, t := range factor.GetTeacher() {
				for _, f := range slf[timeSlot.Factor.GetUniqueSign()][slot] {
					if utils.IsContainString(f.GetTeacher(), t) {
						return define.TeacherConflictException.Hit().
							Supplement("class", factor.GetUniqueSign()).
							Supplement("course", factor.GetCourse()).
							Supplement("teacher", t)
					}
				}
			}
			// 学生冲突
			for _, s := range factor.GetStudent() {
				for _, f := range slf[timeSlot.Factor.GetUniqueSign()][slot] {
					if utils.IsContainString(f.GetStudentUniqueSign(), s.UniqueSign) {
						return define.StudentConflictException.Hit().
							Supplement("class", factor.GetUniqueSign()).
							Supplement("course", factor.GetCourse()).
							Supplement("teacher", factor.GetTeacher()).
							Supplement("student", s).
							Supplement("He is in where", f.GetUniqueSign())
					}
				}
			}
			// 场地冲突
			for _, p := range factor.GetPlace() {
				for _, f := range slf[timeSlot.Factor.GetUniqueSign()][slot] {
					if utils.IsContainString(f.GetPlaceUniqueSign(), p.UniqueSign) {
						return define.PlaceConflictException.Hit().
							Supplement("class", factor.GetUniqueSign()).
							Supplement("course", factor.GetCourse()).
							Supplement("teacher", factor.GetTeacher()).
							Supplement("place", factor.GetPlace())
					}
				}
			}
		}
	}
	return nil
}

// DataView 返回数据视图
func (slf ThreeDimensionalMatrix) DataView() *DataView {
	return &DataView{slf}
}

// GetAllowSlot 获取所有可排课位
func (slf ThreeDimensionalMatrix) GetAllowSlot(factor wtype.Factor) []*wtype.TimeSlot {
	var target []*wtype.TimeSlot
	for _, slot := range factor.GetSlot() {
		if !slf.IsConflict(factor, slot.Index) {
			target = append(target, slot)
		}
	}
	return target
}

// GetAllowFirstSlot 获取一个最佳可放置课位，支持排除特定课位
func (slf ThreeDimensionalMatrix) GetAllowFirstSlot(factor wtype.Factor, exclude ...int) *wtype.TimeSlot {
	var target *wtype.TimeSlot = nil
	var priority = -99999.0
	for _, slot := range factor.GetSlotMap() {
		if !slf.IsConflict(factor, slot.Index) && !utils.IsContainInt(exclude, slot.Index) {
			p := factor.GetPriority()[slot.Index]
			if p > priority {
				priority = p
				target = slot
			}
		}
	}
	return target
}

// GetAllowGreaterThanAVESlot 获取所有大于或等于优先级均值的课位
func (slf ThreeDimensionalMatrix) GetAllowGreaterThanAVESlot(factor wtype.Factor) []*wtype.TimeSlot {
	var targets []*wtype.TimeSlot
	var ave = factor.GetPriorityAVE()
	for _, slot := range factor.GetSlot() {
		if !slf.IsConflict(factor, slot.Index) {
			if factor.GetPriority()[slot.Index] >= ave {
				targets = append(targets, slot)
			}
		}
	}
	return targets
}

// GetNextSlot 获取特定课位的下一个课位
func (slf *ThreeDimensionalMatrix) GetNextSlot(factor wtype.Factor, slot int) *wtype.TimeSlot {
	var width = math.MinInt
	var minWidth = math.MaxInt
	var index int
	for i, timeSlot := range factor.GetSlot() {
		if timeSlot.Index == slot {
			continue
		}
		// 1 2 3 4 5 6 7

		nowWidth := timeSlot.Index - slot
		if nowWidth > width && nowWidth > 0 {
			width = nowWidth
			if width < minWidth {
				minWidth = width
				index = i
			}
		}
	}
	if minWidth == math.MaxInt {
		return nil
	}
	return factor.GetSlot()[index]
}

// GetPrevSlot 获取特定课位的上一个课位
func (slf *ThreeDimensionalMatrix) GetPrevSlot(factor wtype.Factor, slot int) *wtype.TimeSlot {
	var width = math.MinInt
	var index int
	for i, timeSlot := range factor.GetSlot() {
		if timeSlot.Index == slot {
			continue
		}

		// TODO: 或许 nowWidth < 0 满足时就可以直接返回
		nowWidth := timeSlot.Index - slot
		if nowWidth > width && nowWidth < 0 {
			index = i
			width = nowWidth
		}
	}
	if width == math.MinInt {
		return nil
	}
	return factor.GetSlot()[index]
}

// GetNextSlotSameDay 获取下一个同一天课位
func (slf *ThreeDimensionalMatrix) GetNextSlotSameDay(factor wtype.Factor, slot int) *wtype.TimeSlot {
	target := slf.GetNextSlot(factor, slot)
	if target == nil {
		return nil
	}
	if target.WhatDay != factor.GetSlotWithIndex(slot).WhatDay {
		return nil
	}
	return target
}

// GetPrevSlotSameDay 获取上一个同一天课位
func (slf *ThreeDimensionalMatrix) GetPrevSlotSameDay(factor wtype.Factor, slot int) *wtype.TimeSlot {
	target := slf.GetPrevSlot(factor, slot)
	if target == nil {
		return nil
	}
	if target.WhatDay != factor.GetSlotWithIndex(slot).WhatDay {
		return nil
	}
	return target
}
