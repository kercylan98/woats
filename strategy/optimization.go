package strategy

import (
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats"
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/wtype"
)

// Optimization 优选策略
//
// 该策略下会根据排课因子设定的课位优先级进行优先排课。
// 该策略始终会返回可以可用的课位，除非无可排课位。
type Optimization struct {

}

func (slf *Optimization) Initialization() {
}

func (slf *Optimization) Specific(factor wtype.Factor, studio *woats.Studio) bool {

	if slot := studio.GetMatrix().GetAllowFirstSlot(factor); slot != nil {
		studio.FactorPush(factor, slot.Index)
		return false
	}

	return true
}

