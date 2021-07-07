package strategy

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// OptimizationLink 优选链接策略
//
// 将两个课位进行关联的优选策略，当课程放置到某个课位时，同时向另一个课位进行放置。当课位A和课位B相同时忽略该策略
type OptimizationLink struct {
	SlotA int
	SlotB int
}

func (slf *OptimizationLink) OnPush(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	return
}

func (slf *OptimizationLink) OnPop(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	if slot.Index == slf.SlotA {
		_ = studio.PopSameFactor(factor, slf.SlotB)
	}
	if slot.Index == slf.SlotB {
		_ = studio.PopSameFactor(factor, slf.SlotA)
	}
	return
}

func (slf *OptimizationLink) Initialization() {

}

func (slf *OptimizationLink) Specific(factor wtype.Factor, studio *woats.Studio) bool {
	if slf.SlotA == slf.SlotB {
		return true
	}

	if slot := studio.GetMatrix().GetAllowFirstSlot(factor); slot != nil {
		if slot.Index == slf.SlotA {
			if err := studio.FactorPush(factor, slot.Index); err != nil {
				return true
			}
			if err := studio.PushSameFactor(factor, slf.SlotB); err != nil {
				return true
			}
			return false
		}
		if slot.Index == slf.SlotB {
			if err := studio.FactorPush(factor, slot.Index); err != nil {
				return true
			}
			if err := studio.PushSameFactor(factor, slf.SlotA); err != nil {
				return true
			}
			return false
		}
	}

	return true
}

func (slf *OptimizationLink) GroupSpecific(factor wtype.Factor, studio *woats.Studio) bool {
	return true
}
