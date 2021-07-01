package wtype

import (
	"fmt"
	"github.com/kercylan98/woats/core/woats/utils"
)

// 代表一周七天，从周一开始
const (
	Day1 = iota
	Day2
	Day3
	Day4
	Day5
	Day6
	Day7
)

// NewTimeSlot 创建一个基于时间的课位
func NewTimeSlot(index int, whatDay int, startHour, startMinute, endHour, endMinute int) *TimeSlot {
	ts := &TimeSlot{
		Index:   index,
		WhatDay: whatDay,
		Start:   utils.TimeToNumber(startHour, startMinute),
		End:     utils.TimeToNumber(endHour, endMinute),
	}
	ts.UniqueSign = fmt.Sprintf("%d:%f~%f", ts.WhatDay, ts.Start, ts.End)
	return ts
}

// TimeSlot 时间段数据结构
type TimeSlot struct {
	UniqueSign string
	Index      int     // 课位号
	WhatDay    int     // 第几天
	Start      float64 // 开始时间
	End        float64 // 结束时间
}

// GreaterThan 比较是否大于另一个时间课位（如 8:00～8:40 比 9:00~9:10 小）
func (slf *TimeSlot) GreaterThan(timeSlot *TimeSlot) bool {
	return slf.Start > timeSlot.End
}

// IsCrossed 检查与另一个时间课位是否交叉
func (slf *TimeSlot) IsCrossed(timeSlot *TimeSlot) bool {
	if slf.WhatDay == timeSlot.WhatDay {
		/**
		================================== 1
		slf.start
					TimeSlot.start
					TimeSlot.end
		slf.end
		================================== 1
		slf.start
					TimeSlot.start
		slf.end
					TimeSlot.end
		================================== 2
					TimeSlot.start
		slf.start
					TimeSlot.end
		slf.end
		==================================
		*/
		if (slf.Start < timeSlot.Start && slf.End > timeSlot.Start) ||
			(slf.Start > timeSlot.Start && slf.Start < timeSlot.End) ||
			(slf.UniqueSign == timeSlot.UniqueSign) {
			return true
		}
	}
	return false
}
