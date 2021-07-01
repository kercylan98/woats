package woats

import "github.com/kercylan98/exception"

// Designer 排课整体设计器实现
type Designer interface {

	// Modification 排课数据修饰
	Modification(modifier *ModifierFactory) exception.Exception

}
