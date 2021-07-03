package woats

import "github.com/kercylan98/woats/core/woats/wtype"

// IsContainFactor 一组因子中是否包含另一组因子
func IsContainFactor(group wtype.FactorGroup, factorGroup wtype.FactorGroup) bool {
	for _, factor := range group {
		for _, f := range factorGroup {
			if factor == f {
				return true
			}
		}
	}
	return false
}

// RemoveFactor 返回删除特定因子后的因子组
func RemoveFactor(group wtype.FactorGroup, factor ...wtype.Factor) wtype.FactorGroup {
	var result = make(wtype.FactorGroup, 0)
	for _, f := range group {
		match := false
		for _, w := range factor {
			if f == w {
				match = true
				break
			}
		}
		if !match {
			result = append(result, f)
		}
	}
	return result
}
