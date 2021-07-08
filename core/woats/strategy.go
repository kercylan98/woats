package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/wtype"
)

var SkipStrategy = exception.Reg(0, "skip")

// Strategy 排课策略接口
type Strategy interface {

	// Initialization 初始化策略模型，仅在第一次使用时调用
	Initialization()

	// Specific 策略具体内容,当返回空时中断后续策略执行。
	Specific(factor wtype.Factor, studio *Studio) exception.Exception

	// GroupSpecific 策略针对连堂课的具体内容,当返回空时中断后续策略执行。
	GroupSpecific(factor wtype.Factor, studio *Studio) exception.Exception

	// OnPush 当某个因子被推入的时候是否需要该策略做些什么
	OnPush(factor wtype.Factor, slot *wtype.TimeSlot, studio *Studio)

	// OnPop 当某个因子被弹出的时候是否需要该策略做些什么
	OnPop(factor wtype.Factor, slot *wtype.TimeSlot, studio *Studio)
}
