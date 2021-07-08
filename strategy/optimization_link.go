package strategy

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// OptimizationLink 优选链接策略
//
// From：华东师范大学附属双语学校（高中部）
//
// 2/3和8/9节课必须连排，不限制具体课程。只要在这个课位上就必须连排(具体如下表)，2/3和8/9实际上是两门课程，但是
// 中间没有休息时间，所以必须是连堂课，主要问题在于，什么课程被放到这里就必须是连堂。如果把2/3和8/9拆分为两个节次
// 来看，那么所有的选修课都必须设置连堂课（冲突增加选修课本身是可以单节排课的）
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

func (slf *OptimizationLink) Specific(factor wtype.Factor, studio *woats.Studio) exception.Exception {
	if slf.SlotA == slf.SlotB {
		return woats.SkipStrategy.Hit()
	}

	if studio.GetSameFactorCount(factor) == 1 {
		return nil
	}

	if slot := studio.GetMatrix().GetAllowFirstSlot(factor, studio.GetAllSlotNumber(factor, slf.SlotA, slf.SlotB)...); slot != nil {
		if slot.Index == slf.SlotA {
			if err := studio.FactorPush(factor, slot.Index); err != nil {
				return err
			}
			if err := studio.PushSameFactor(factor, slf.SlotB); err != nil {
				return err
			}
			return nil
		}
		if slot.Index == slf.SlotB {
			if err := studio.FactorPush(factor, slot.Index); err != nil {
				return err
			}
			if err := studio.PushSameFactor(factor, slf.SlotA); err != nil {
				return err
			}
			return nil
		}

	}
	return woats.SkipStrategy.Hit()
}

func (slf *OptimizationLink) GroupSpecific(factor wtype.Factor, studio *woats.Studio) exception.Exception {
	return woats.SkipStrategy.Hit()
}
