package main

import (
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/designer"
	"github.com/kercylan98/woats/strategy"
)

func main() {
	w := woats.New(new(designer.BilingualSchoolAffiliatedToEastChinaNormalUniversity), 10,
		&strategy.Continuity{
			Optimum: 3,
			Min:     2,
			Max:     4,
		},
		new(strategy.Optimization),
		&strategy.RandomSeek{
			StopLoss: 0.96,
		},
	)

	if es := w.Run(); len(es) > 0 {
		for _, e := range es {
			e.Print()
		}
	}
}
