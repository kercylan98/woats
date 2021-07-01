package woats

import (
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/wtype"
	"sort"
)

type DataViewRelation struct {
	slots 				[]*wtype.TimeSlot
	slotFactorGroupR 	map[int]wtype.FactorGroup
}

func newDataViewRelation(s []*wtype.TimeSlot, m map[int]wtype.FactorGroup) *DataViewRelation {
	dvr := &DataViewRelation{
		slots: s,
		slotFactorGroupR: m,
	}
	sort.Slice(dvr.slots, func(i, j int) bool {
		return !dvr.slots[i].GreaterThan(dvr.slots[j])
	})
	return dvr
}