package strategy

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/wtype"
	"math"
)

// RandomSeek 随机寻求策略
//
// 随机寻求策略将对传入的因子进行正常排课，当发生无法排课的情况下，会逐渐扩大搜索范围将冲突课程移除后再尝试排课。
// 当达到寻求上限时，恢复最高分快照。
// 在设置止损点的时候，当总体进度大于止损点则停止继续寻求
// TODO: 存在很大的优化空间，该策略赋予了更大随机的可能性
type RandomSeek struct {
	StopLoss    float64 // 止损点(0~1)，当值为0的时候不生效
	ExcludeSlot []int   // 排除课位

	basic    map[wtype.Factor]int // 寻求基数
	basicMax int                  // 寻求基数上限
	top      float64              // 最高分
}

func (slf *RandomSeek) OnPush(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	return
}

func (slf *RandomSeek) OnPop(factor wtype.Factor, slot *wtype.TimeSlot, studio *woats.Studio) {
	return
}

func (slf *RandomSeek) Initialization() {
	slf.basic = map[wtype.Factor]int{}
	slf.basicMax = 10
	slf.top = math.MaxInt

	if slf.StopLoss < 0 {
		slf.StopLoss = 0
	}
	if slf.StopLoss > 1 {
		slf.StopLoss = 1
	}
}

func (slf *RandomSeek) GroupSpecific(factor wtype.Factor, studio *woats.Studio) exception.Exception {
	return woats.SkipStrategy.Hit()
}

func (slf *RandomSeek) Specific(factor wtype.Factor, studio *woats.Studio) exception.Exception {
	if slf.StopLoss != 0 {
		if studio.GetProgress()/100.0 >= slf.StopLoss {
			return nil
		}
	}
seek:
	{
	}
	studio.LockRotation()
	if slot := studio.GetMatrix().GetAllowFirstSlot(factor, slf.ExcludeSlot...); slot != nil {
		if err := studio.FactorPush(factor, slot.Index); err != nil {
			return err
		}
		top := studio.GetProgress()
		if top >= slf.top {
			slf.top = top
			studio.CreateSnapshot("RANDOM_SEEK_SNAPSHOT")
		} else {
			studio.RecoverySnapshot("RANDOM_SEEK_SNAPSHOT")
		}
		studio.UnLockRotation()
		return nil
	}

	// 初始化基数
	if slf.basic[factor] == 0 {
		slf.basic[factor] = 1
	}
	basic := slf.basic[factor]

	if conflict := studio.GetMatrix().GetConflictFactorAll(factor, basic); len(conflict) != 0 {
		for _, fg := range conflict {
			for _, f := range fg {
				if f.GetTimeSlot() != nil {
					if err := studio.FactorPop(f, f.GetTimeSlot().Index); err != nil {
						return err
					}
				}
			}
		}
		slf.basic[factor] = basic + 1
	} else {
		// 寻求不到冲突还无法排课的情况不可能出现，应该直接跳过该策略
		panic("error code logic")
	}

	if basic <= slf.basicMax {
		goto seek
	}

	// 很难到这
	studio.RecoverySnapshot("RANDOM_SEEK_SNAPSHOT")
	studio.UnLockRotation()

	slf.basic[factor] = 1
	return woats.SkipStrategy.Hit()
}
