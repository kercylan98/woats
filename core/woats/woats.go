package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/wtype"
	"sync"
)

// New 根据指定的设计器和策略进行排课
func New(designer Designer, rotationMax int, strategy ...Strategy) *Woats {
	woats := &Woats{
		designer:         designer,
		strategy:         strategy,
		strategyInitLock: map[Strategy]*sync.Once{},

		rotationMax: rotationMax,
	}
	for _, s := range strategy {
		woats.strategyInitLock[s] = new(sync.Once)
	}
	return woats
}

type Woats struct {
	designer         Designer                // 采用的数据设计器
	strategy         []Strategy              // 采用的排课策略
	strategyInitLock map[Strategy]*sync.Once // 排课策略初始化锁

	rotationMax int // 流入Studio轮转次数上限的配置
}

// Run 运行排课
func (slf *Woats) Run() (*DataView, []exception.Exception) {
	var (
		mf         = &ModifierFactory{product: map[string]Modifier{}}
		exceptions []exception.Exception
		factors    = make(wtype.FactorGroup, 0)
	)

	// 数据初始化
	if err := slf.designer.Modification(mf); err != nil {
		return nil, []exception.Exception{err}
	}
	// 数据优化并校验
	for _, modifier := range mf.product {
		modifier.DataOptimization()
		if e := modifier.DataVerification(); e != nil {
			exceptions = append(exceptions, e)
		}
	}
	// 校验异常
	if len(exceptions) > 0 {
		return nil, exceptions
	}

	// 排课因子转化
	for _, modifier := range mf.product {
		factors = append(factors, modifier.ToFactor()...)
	}
	tdm := newTDM(factors)
	return tdm.DataView(), slf.start(factors, tdm)
}

func (slf *Woats) start(factors wtype.FactorGroup, matrix ThreeDimensionalMatrix) []exception.Exception {
	var fixed = factors.PopAllFixed()
	var normal = factors
	var studio = newStudio([]wtype.Factor{}, matrix, slf.strategy, slf.rotationMax)
	studio.logout = true
	slf.handleFixed(fixed, studio)
	slf.handleNormal(normal, studio)

	return []exception.Exception{}
}

func (slf *Woats) handleFixed(factors wtype.FactorGroup, studio *Studio) {
	studio.addFactorGroup(factors)
	studio.Run(func(factor wtype.Factor, studio *Studio) exception.Exception {
		if err := studio.FactorPush(factor, factor.GetFixed()); err != nil {
			// 无法被放置的固定课
			return err
		}
		factor.SetDisableChange()
		return nil
	})
}

func (slf *Woats) handleNormal(factors wtype.FactorGroup, studio *Studio) {
	// 按照优先级排序（最好的 -> 最差的 -> 未设置的）
	fg := factors.SortWithPriorityAVE()
	ng := fg.GetAllNotGroupFactor()
	g := fg.GetAllGroupFactor()
	studio.addFactorGroup(append(g, ng...))
	studio.Run(func(factor wtype.Factor, studio *Studio) exception.Exception {
		var se exception.Exception = nil
		for _, strategy := range slf.strategy {
			// 初始化阶段
			slf.strategyInitLock[strategy].Do(func() {
				strategy.Initialization()
			})
			// 逻辑
			var err exception.Exception
			if factor.IsGroup() {
				err = strategy.GroupSpecific(factor, studio)
			} else {
				err = strategy.Specific(factor, studio)
			}
			if err != nil {
				se = err
			} else {
				break
			}
		}

		return se
	})

}
