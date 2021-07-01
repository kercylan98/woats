package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// RollBack 回滚快照，count表示回滚快照数量
//
// 当返回错误的时候表示没有可回退的版本，当回滚成功时，始终都会直接中断后续的策略执行，并且进入一新的因子处理流程
type RollBack func(count int) exception.Exception

// Strategy 排课策略接口
type Strategy interface {

	// Initialization 初始化策略模型，仅在第一次使用时调用
	Initialization()

	// Specific 策略具体内容,当返回false时中断后续策略执行。
	Specific(factor wtype.Factor, studio *Studio) bool
}
