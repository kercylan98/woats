package strategy

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/utils"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// QuantitativeCombination 定量组合策略
//
// From：华东师范大学附属双语学校（高中部）
//
// 该策略将把特定一组的课位号进行联合，最终决定该组合内相同特质的因子仅能出现特定次数。
// 如：第一节课同一位老师最多只能出现一次
type QuantitativeCombination struct {
	Combination  []int // 课位号组合
	Count        int   // 最大出现次数
	CheckTeacher bool  // 校验教师
	CheckCourse  bool  // 校验课程
	CheckStudent bool  // 校验学生
	CheckPlace   bool  // 校验教学场地
}

func (slf *QuantitativeCombination) Initialization() {
	slf.Combination = utils.RemoveRepeatedInt(slf.Combination)
	if slf.Count < 0 {
		slf.Count = 0
	}
}

func (slf *QuantitativeCombination) Specific(factor wtype.Factor, studio *woats.Studio) bool {
	var exec = func(factor wtype.Factor) bool {
		slot := studio.GetMatrix().GetAllowFirstSlot(factor, studio.GetAllSlotNumber(factor, slf.Combination...)...)
		if slot != nil {
			if err := studio.FactorPush(factor, slot.Index); err != nil {
				return true
			}
			return false
		}
		return true
	}

	if slf.CheckTeacher && slf.Count > 0 && studio.GetSameTeacher(factor, slf.Combination) < slf.Count {
		return exec(factor)
	} else if slf.CheckCourse && slf.Count > 0 && studio.GetSameCourse(factor, slf.Combination) < slf.Count {
		return exec(factor)
	} else if slf.CheckStudent && slf.Count > 0 && studio.GetSameStudent(factor, slf.Combination) < slf.Count {
		return exec(factor)
	} else if slf.CheckPlace && slf.Count > 0 && studio.GetSamePlace(factor, slf.Combination) < slf.Count {
		return exec(factor)
	}
	return true
}

func (slf *QuantitativeCombination) GroupSpecific(factor wtype.Factor, studio *woats.Studio) bool {
	return true
}

func (slf *QuantitativeCombination) OnPush(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	// 阻止其他策略破坏该策略
	var exec = func() {
		if utils.IsContainInt(slf.Combination, slot.Index) {
			if err := studio.FactorPop(factor, slot.Index); err != nil {
				panic(err)
			}
		}
	}

	if slf.CheckTeacher && slf.Count > 0 && studio.GetSameTeacher(factor, slf.Combination) > slf.Count {
		exec()
	} else if slf.CheckCourse && slf.Count > 0 && studio.GetSameCourse(factor, slf.Combination) > slf.Count {
		exec()
	} else if slf.CheckStudent && slf.Count > 0 && studio.GetSameStudent(factor, slf.Combination) > slf.Count {
		exec()
	} else if slf.CheckPlace && slf.Count > 0 && studio.GetSamePlace(factor, slf.Combination) > slf.Count {
		exec()
	}
	return
}

func (slf *QuantitativeCombination) OnPop(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	return
}
