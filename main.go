package main

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/designer"
	"github.com/kercylan98/woats/strategy"
)

func main() {
	w := woats.New(new(designer.BilingualSchoolAffiliatedToEastChinaNormalUniversity), 10,
		//&strategy.Continuity{
		//	Optimum: 3,
		//	Min:     2,
		//	Max:     4,
		//},
		&strategy.OptimizationLink{SlotA: 3, SlotB: 4},
		&strategy.OptimizationLink{SlotA: 14, SlotB: 15},
		&strategy.OptimizationLink{SlotA: 25, SlotB: 26},
		&strategy.OptimizationLink{SlotA: 36, SlotB: 37},
		&strategy.OptimizationLink{SlotA: 47, SlotB: 48},
		new(strategy.Optimization),
		&strategy.RandomSeek{
			StopLoss: 0.96,
		},
	)

	if view, es := w.Run(); len(es) > 0 {
		for _, e := range es {
			e.Print()
		}
	} else {
		if err := view.CreateClassViewExcel("DP1 Physics HL", "out"); err != nil {
			panic(err)
		}
		if err := view.CreateClassViewExcel("DP1 Mathematics-Pure", "out"); err != nil {
			panic(err)
		}
		if err := view.CreateAllViewExcel("out"); err != nil {
			panic(err)
		}
	}
}
