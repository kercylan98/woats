package wtype

import (
	"math/rand"
	"sort"
	"time"
)

// FactorGroup 对一组排课因子的封装
type FactorGroup []Factor

// IsContainFactor 该组因子中是否包含另一组因子
func (slf *FactorGroup) IsContainFactor(factorGroup FactorGroup) bool {
	for _, factor := range *slf {
		for _, f := range factorGroup {
			if factor.GetUniqueSign() == f.GetUniqueSign() {
				return true
			}
		}
	}
	return false
}

// RemoveFactor 返回删除特定因子后的因子组
func (slf *FactorGroup) RemoveFactor(factor ...Factor) FactorGroup {
	var result = make(FactorGroup, 0)
	for _, f := range *slf {
		match := false
		for _, w := range factor {
			if f.GetUniqueSign() == w.GetUniqueSign() {
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

// Unrepeated 返回去重后的因子组
func (slf *FactorGroup) Unrepeated() (newArr FactorGroup) {
	newArr = make(FactorGroup, 0)
	for i := 0; i < len(*slf); i++ {
		repeat := false
		for j := i + 1; j < len(*slf); j++ {
			if (*slf)[i].GetUniqueSign() == (*slf)[j].GetUniqueSign() {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, (*slf)[i])
		}
	}
	return
}

// PopAllFixed 弹出所有固定课
func (slf *FactorGroup) PopAllFixed() FactorGroup {
	var replace = make(FactorGroup, 0)
	var result = make(FactorGroup, 0)
	for i := 0; i < len(*slf); i++ {
		f := (*slf)[i]
		if f.IsFixed() {
			result = append(result, f)
		} else {
			replace = append(replace, f)
		}
	}
	*slf = replace
	return result
}

// PushFixed 推入固定课
func (slf *FactorGroup) PushFixed(factor ...Factor) {
	for _, f := range factor {
		if f.IsFixed() {
			*slf = append(*slf, f)
		}
	}
}

// Upset 将返回一个新的被打乱的因子组
func (slf *FactorGroup) Upset() FactorGroup {
	// 打乱顺序
	var newMap = map[int]Factor{}
	for i, factor := range *slf {
		newMap[i] = factor
	}
	// 存放至新数组
	var newFG = make(FactorGroup, 0)
	for _, factor := range newMap {
		newFG = append(newFG, factor)
	}
	return newFG
}

// GetAllGroupFactor 获取所有连堂课
func (slf *FactorGroup) GetAllGroupFactor() FactorGroup {
	var newFG = make(FactorGroup, 0)
	for _, factor := range *slf {
		if factor.IsGroup() {
			newFG = append(newFG, factor)
		}
	}
	return newFG
}

// GetAllNotGroupFactor 获取所有非连堂课
func (slf *FactorGroup) GetAllNotGroupFactor() FactorGroup {
	var newFG = make(FactorGroup, 0)
	for _, factor := range *slf {
		if !factor.IsGroup() {
			newFG = append(newFG, factor)
		}
	}
	return newFG
}

// SortWithPriorityAVE 返回一个根据优先级均值排序后的因子组
func (slf *FactorGroup) SortWithPriorityAVE() FactorGroup {
	var newFG = slf.Upset()
	// 排序
	sort.Slice(newFG, func(i, j int) bool {
		ia, ja := newFG[i].GetPriorityAVE(), newFG[j].GetPriorityAVE()
		if ia == ja {
			// 确保公平
			rand.Seed(time.Now().UnixNano())
			return (0 + rand.Intn(2-0)) == 0
		} else {
			return ia > ja
		}
	})
	return newFG
}

// SortWithPrioritySUM 返回一个根据优先级总和排序后的因子组
func (slf *FactorGroup) SortWithPrioritySUM() FactorGroup {
	var newFG = slf.Upset()
	// 排序
	sort.Slice(newFG, func(i, j int) bool {
		ia, ja := newFG[i].GetPrioritySUM(), newFG[j].GetPrioritySUM()
		if ia == ja {
			// 确保公平
			rand.Seed(time.Now().UnixNano())
			return (0 + rand.Intn(2-0)) == 0
		} else {
			return ia > ja
		}
	})
	return newFG
}

// Clone 克隆一份
func (slf *FactorGroup) Clone() FactorGroup {
	var newFG = make(FactorGroup, len(*slf))
	for i := 0; i < len(newFG); i++ {
		newFG[i] = (*slf)[i]
	}
	return newFG
}
