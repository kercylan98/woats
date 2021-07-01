package woats

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/wtype"
	"sync"
)

// New 根据指定的设计器和策略进行排课
func New(designer Designer, rotationMax int, strategy ...Strategy) *Woats {
	woats := &Woats{
		designer: designer,
		strategy: strategy,
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

	rotationMax 	 int 					 // 流入Studio轮转次数上限的配置
}

// Run 运行排课
func (slf *Woats) Run() []exception.Exception {
	var (
		mf         = &ModifierFactory{product: map[string]Modifier{}}
		exceptions []exception.Exception
		factors    = make(wtype.FactorGroup, 0)
	)

	// 数据初始化
	if err := slf.designer.Modification(mf); err != nil {
		return []exception.Exception{err}
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
		return exceptions
	}

	// 排课因子转化
	for _, modifier := range mf.product {
		factors = append(factors, modifier.ToFactor()...)
	}

	return slf.start(factors, newTDM(factors))
}

func (slf *Woats) start(factors wtype.FactorGroup, matrix ThreeDimensionalMatrix) []exception.Exception {
	var fixed = factors.PopAllFixed()
	var normal = factors
	var studio = newStudio([]wtype.Factor{}, matrix, slf.rotationMax)
	// 终校验
	if es := slf.check(fixed, normal, studio); len(es) > 0 {
		return es
	}
	studio.logout = true
	slf.handleFixed(fixed, studio)
	slf.handleNormal(normal, studio)

	return []exception.Exception{}
}

func (slf *Woats) check(fixed, normal wtype.FactorGroup, studio *Studio) []exception.Exception {
	var exceptions []exception.Exception

	if len(fixed) > 0 {
		fixedStudio := studio.clone()
		fixedStudio.addFactorGroup(fixed)
		fixedStudio.Run(func(factor wtype.Factor, studio *Studio) bool {
			// TODO: 固定课检查内容
			if e := studio.matrix.IsConflictErr(factor, factor.GetFixed()); e != nil {
				exceptions = append(exceptions, e)
			}
			studio.FactorPush(factor, factor.GetFixed())
			return true
		})
	}

	if len(normal) > 0 {
		normalStudio := studio.clone()
		normalStudio.addFactorGroup(normal)
		normalStudio.Run(func(factor wtype.Factor, studio *Studio) bool {
			// TODO: 普通课检查内容
			return true
		})
	}
	return exceptions
}

func (slf *Woats) handleFixed(factors wtype.FactorGroup, studio *Studio) {
	studio.addFactorGroup(factors)
	studio.Run(func(factor wtype.Factor, studio *Studio) bool {
		studio.FactorPush(factor, factor.GetFixed())
		return true
	})
}

func (slf *Woats) handleNormal(factors wtype.FactorGroup, studio *Studio) {
	// 按照优先级排序（最好的 -> 最差的 -> 未设置的）
	studio.addFactorGroup(factors)
	studio.Run(func(factor wtype.Factor, studio *Studio) bool {

		for _, strategy := range slf.strategy {
			// 初始化阶段
			slf.strategyInitLock[strategy].Do(func() {
				strategy.Initialization()
			})
			// 逻辑
			if isContinue := strategy.Specific(factor, studio); isContinue == false {
				return true
			}
		}

		return false
	})

}






















