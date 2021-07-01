package main

import (
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats"
	"github.com/kercylan98/work-out-a-teaching-schedule/designer"
	"github.com/kercylan98/work-out-a-teaching-schedule/strategy"
)

// 测试入口
//func main() {
//	es := woats.New(new(designer.Test),
//		&strategy.Continuity{
//			Optimum: 3,
//			Min:     2,
//			Max:     4,
//			ExcludeClass: []string{"小学一年级1班"},
//		},
//		new(strategy.Optimization),
//		).Run()
//
//
//	if len(es) > 0 {
//		for _, e := range es {
//			e.Print()
//		}
//	}
//}

// 双语入口
func main() {
	w := woats.New(new(designer.BilingualSchoolAffiliatedToEastChinaNormalUniversity), 10,

		&strategy.RandomSeek{
			StopLoss: 0.96,
		},
		new(strategy.Optimization),
		&strategy.Continuity{
			Optimum: 3,
			Min:     2,
			Max:     4,
		},

	)

	if es := w.Run(); len(es) > 0 {
		for _, e := range es {
			e.Print()
		}
	}
}
