package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/wtype"
)

// Modifier 修饰器接口
type Modifier interface {

	// ToFactor 转换为排课因子
	ToFactor() []wtype.Factor

	// DataOptimization 对数据进行优化
	DataOptimization()

	// DataVerification 对数据进行校验
	DataVerification() exception.Exception
}
