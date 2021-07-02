package woats

import (
	"fmt"
	"github.com/kercylan98/woats/core/woats/wtype"
	"github.com/tealeg/xlsx"
	"math"
	"sort"
)

type DataView struct {
	matrix ThreeDimensionalMatrix
}

// GetClassDayView 获取特定班级特定日视图
func (slf *DataView) GetClassDayView(class string, day int) *DataViewRelation {
	var slots []*wtype.TimeSlot
	slotFactorGroupR := map[int]wtype.FactorGroup{}
	ts := slf.matrix.GetTimeSlots(class)

	for _, slot := range ts {
		if slot.WhatDay == day {
			slots = append(slots, slot)
			slotFactorGroupR[slot.Index] = slf.matrix[class][slot.Index]
		}
	}
	return newDataViewRelation(slots, slotFactorGroupR)
}

// GetClassView 获取特定班级周视图
func (slf *DataView) GetClassView(class string) []*DataViewRelation {
	var dvrs []*DataViewRelation
	for i := 1; i < 7; i++ {
		if dvr := slf.GetClassDayView(class, i); len(dvr.slots) > 0 {
			dvrs = append(dvrs, dvr)
		}
	}
	return dvrs
}

// GetView 获取总视图
func (slf *DataView) GetView() map[int]map[string]wtype.FactorGroup {
	time := map[int]map[string]wtype.FactorGroup{}
	for class, _ := range slf.matrix {
		for _, relation := range slf.GetClassView(class) {
			for _, slot := range relation.slots {
				if _, exist := time[slot.WhatDay]; !exist {
					time[slot.WhatDay] = map[string]wtype.FactorGroup{}
					time[slot.WhatDay][slot.UniqueSign] = make(wtype.FactorGroup, 0)
				} else {
					if _, exist = time[slot.WhatDay][slot.UniqueSign]; !exist {
						time[slot.WhatDay][slot.UniqueSign] = make(wtype.FactorGroup, 0)
					}
				}
				for _, factor := range relation.slotFactorGroupR[slot.WhatDay] {
					if factor.GetTimeSlot().UniqueSign == slot.UniqueSign {
						time[slot.WhatDay][slot.UniqueSign] = append(time[slot.WhatDay][slot.UniqueSign], factor)
					}
				}
			}
		}
	}
	return time
}

// CreateClassViewExcel 创建班级视图表格
func (slf *DataView) CreateClassViewExcel(class string, path string) error {
	excel := xlsx.NewFile()
	shell, err := excel.AddSheet("表")
	if err != nil {
		return err
	}

	classView := slf.GetClassView(class)
	title := shell.AddRow()
	for i, _ := range classView {
		title.AddCell().SetInt(i + 1)
	}
	maxRow := 0
	for _, relation := range classView {
		if len(relation.slots) > maxRow {
			maxRow = len(relation.slots)
		}
	}

	for i := 0; i < maxRow; i++ {
		row := shell.AddRow()
		for ci := 0; ci < len(classView); ci++ {
			content := ""
			for _, factor := range classView[ci].slotFactorGroupR[classView[ci].slots[i].Index] {
				content += fmt.Sprintf("%s, %s\r\n%s", factor.GetCourse(), factor.GetTeacher(), classView[ci].slots[i].UniqueSign)
			}
			row.AddCell().SetString(content)
		}
	}

	return excel.Save(path + "/" + class + ".xlsx")
}

// CreateAllViewExcel 创建总视图表格
func (slf *DataView) CreateAllViewExcel(path string) error {
	excel := xlsx.NewFile()
	shell, err := excel.AddSheet("表")
	if err != nil {
		return err
	}

	title := shell.AddRow()
	title.AddCell().SetString("*")
	for i := 0; i < 7; i++ {
		title.AddCell().SetString(fmt.Sprintf("周 %d", i+1))
	}

	style := xlsx.NewStyle()
	style.Alignment.WrapText = true
	for i := 0; i < 50; i++ {
		r := shell.AddRow()
		r.SetHeight(170)
		r.AddCell()
		for i := 0; i < 7; i++ {
			r.AddCell().SetStyle(style)
		}
	}
	shell.SetColWidth(1, 7, 80)
	shell.SetColWidth(0, 0, 16)

	content := map[int]string{}

	for class, _ := range slf.matrix {
		ss := slf.matrix.GetTimeSlots(class)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Index < ss[j].Index
		})
		for _, timeSlot := range ss {
			var fContent string
			for _, factor := range slf.matrix[class][timeSlot.Index] {
				fContent += fmt.Sprintf("%s, %s, %s\r\n", factor.GetUniqueSign(), factor.GetCourse(), factor.GetTeacher())
			}
			content[timeSlot.Index] = content[timeSlot.Index] + fContent
		}
	}

	for class, _ := range slf.matrix {
		ss := slf.matrix.GetTimeSlots(class)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Index < ss[j].Index
		})
		for _, timeSlot := range ss {
			// 当天最大最小index
			var min, max = math.MaxInt, math.MinInt
			for _, s := range ss {
				if s.WhatDay == timeSlot.WhatDay {
					if s.Index > max {
						max = s.Index
					}
					if s.Index < min {
						min = s.Index
					}
				}
			}

			shell.Rows[max-(max-timeSlot.Index)-min+1].Cells[timeSlot.WhatDay].SetString(content[timeSlot.Index])
			shell.Rows[max-(max-timeSlot.Index)-min+1].Cells[0].SetString(fmt.Sprintf("%d:%d~%d:%d",
				timeSlot.StartHour, timeSlot.StartMinute, timeSlot.EndHour, timeSlot.EndMinute))
		}
	}

	return excel.Save(path + "/all.xlsx")
}
