package strategy

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// Link 链接策略
//
// 将两个特定因子关联起来，
type Link struct {
}

func (slf *Link) Initialization() {

}

func (slf *Link) Specific(factor wtype.Factor, studio *woats.Studio) bool {

	panic("implement me")
}

func (slf *Link) GroupSpecific(factor wtype.Factor, studio *woats.Studio) bool {
	panic("implement me")
}
