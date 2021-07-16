package designer

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/define"
	"github.com/kercylan98/woats/core/woats/wtype"
	"github.com/tealeg/xlsx"
	"strings"
)

type BilingualSchoolAffiliatedToEastChinaNormalUniversity2021 struct {
}

func (slf *BilingualSchoolAffiliatedToEastChinaNormalUniversity2021) Modification(modifier *woats.ModifierFactory) exception.Exception {
	var tss = make([]*wtype.TimeSlot, 50)
	for index := 0; index < 5; index++ {
		tss[index*10+0] = wtype.NewTimeSlot(index*10+1, index+1, 8, 0, 8, 40)
		tss[index*10+1] = wtype.NewTimeSlot(index*10+2, index+1, 8, 50, 9, 30)
		tss[index*10+2] = wtype.NewTimeSlot(index*10+3, index+1, 9, 40, 10, 20)
		tss[index*10+3] = wtype.NewTimeSlot(index*10+4, index+1, 10, 30, 11, 10)
		tss[index*10+4] = wtype.NewTimeSlot(index*10+5, index+1, 11, 20, 12, 0)
		tss[index*10+5] = wtype.NewTimeSlot(index*10+6, index+1, 14, 30, 15, 10)
		tss[index*10+6] = wtype.NewTimeSlot(index*10+7, index+1, 15, 20, 16, 00)
		tss[index*10+7] = wtype.NewTimeSlot(index*10+8, index+1, 16, 10, 16, 0)
		tss[index*10+8] = wtype.NewTimeSlot(index*10+9, index+1, 16, 10, 16, 50)
		tss[index*10+9] = wtype.NewTimeSlot(index*10+10, index+1, 17, 0, 17, 40)
	}

	excel, err := xlsx.OpenFile("designer/assets/2021-2022-confusion.xlsx")
	if err != nil {
		return define.OpenFileException.Hit().Supplement("err", err)
	}

	var students = map[string][]string{}
	for _, row := range excel.Sheets[1].Rows {
		class := strings.TrimSpace(row.Cells[0].String())
		stu := strings.TrimSpace(row.Cells[3].String())
		if stuArr, exist := students[class]; !exist {
			students[class] = []string{stu}
		} else {
			students[class] = append(stuArr, stu)
		}
	}

	for i, row := range excel.Sheets[0].Rows {
		if i != 0 {
			if strings.TrimSpace(row.Cells[4].String()) == "" {
				break
			}
			sectionSum, err := row.Cells[4].Int()
			if err != nil {
				return define.DataConversionException.Hit().Supplement("err", err)
			}

			var section = make([]int, 0)
			for i := 0; i < sectionSum; i++ {
				switch row.Cells[1].String() {
				case "DP2 VA: HL":
					//"DP2 Physics HL",
					//"DP1 Physics HL+SL",
					//"DP2 Physics HL+SL",
					//"DP1 Chemistry HL",
					//"DP1 Chemistry HL+SL",
					//"DP2 Chemistry HL",
					//"DP2 Chemistry HL+SL",
					//"DP1 Biology HL+SL",
					//"DP2 Biology HL+SL",
					//"DP1 Biology HL",
					//"DP2 Biology HL",
					//"DP1 Biology":
					section = append(section, 2)
					i++
				default:
					section = append(section, 1)
				}
			}

			class := strings.TrimSpace(row.Cells[1].String())
			ecm := modifier.ElectiveClass(class, section, tss).
				//SetDisable(1, 2, 22, 44, 53, 54, 55).
				AddStudent(students[class]...)
			if strings.TrimSpace(row.Cells[6].String()) == "高二" {
				ecm.SetDisable(9, 10)
			}
			ecm.Ok().AddTeacher(strings.TrimSpace(row.Cells[3].String())).Ok()
		}
	}

	return nil
}
