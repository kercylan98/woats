package strategy

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/utils"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// Continuity 连续性排课策略(需要确保后续存在其他策略，否则可能导致跳过课程)
//
// 该策略用于满足非强制性的连排要求，如尽量三节连排等
// 默认均为0表示无效
type Continuity struct {
	Optimum        int      // 最佳连排数
	Min            int      // 最小尽量满足连排数
	Max            int      // 最大能接受的连排数
	ExcludeClass   []string // 忽略班级
	ExcludeCourse  []string // 忽略课程
	ExcludeTeacher []string // 忽略教师
	ExcludePlace   []string // 忽略场地
}

func (slf *Continuity) Initialization() {

}

func (slf *Continuity) Specific(factor wtype.Factor, studio *woats.Studio) bool {
	// 忽略内容检查
	if utils.IsContainString(slf.ExcludeClass, factor.GetUniqueSign()) ||
		utils.IsContainString(slf.ExcludeCourse, factor.GetCourse()) {
		return true
	}
	for _, s := range factor.GetTeacher() {
		if utils.IsContainString(slf.ExcludeTeacher, s) {
			return true
		}
	}
	for l, _ := range factor.GetPlace() {
		if utils.IsContainString(slf.ExcludePlace, l) {
			return true
		}
	}

	var mapper = map[*wtype.TimeSlot][]*wtype.TimeSlot{}
	var slots []*wtype.TimeSlot

	// 当未预先放置任何该因子关联课程时，交由下一个策略处理
	if slots = studio.GetMatrix().GetPlaceSlotWithCourse(factor); len(slots) == 0 {
		return true
	} else {
		for _, slot := range slots {
			mapper[slot] = factor.GetSlotWithSameDay(slot)
		}
	}

	// 排序后的索引
	var slotIndex = make([]int, len(slots))
	for i, slot := range slots {
		slotIndex[i] = slot.Index
	}

	// 连续数映射包含量连续量和可用课位索引的关系
	var continuousQuantityMapper = map[int][]int{}
	for slot, timeSlots := range mapper {
		// 连续量
		var continuousQuantity int
		var hit []int
		for _, timeSlot := range timeSlots {
			if utils.IsContainInt(slotIndex, timeSlot.Index) || slot.Index != timeSlot.Index && continuousQuantity < slf.Max {
				continuousQuantity++
			} else if !studio.GetMatrix().IsConflict(factor, timeSlot.Index) && continuousQuantity < slf.Max {
				continuousQuantity++
				hit = append(hit, timeSlot.Index)
			} else {
				if continuousQuantity > slf.Min && continuousQuantity <= slf.Max {
					continuousQuantityMapper[continuousQuantity] = hit
					continuousQuantity = 0
				}
			}
		}
	}

	// 查找合适连续量的可用课位并放置
	var size int
	var target []int
	for i, is := range continuousQuantityMapper {
		if i > size {
			size = i
			target = is
		}
		if size == slf.Optimum {
			break
		}
	}

	if len(target) > 0 {
		studio.FactorPush(factor, target[0])
		return false
	}

	return true
}
