package designer

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/define"
	"github.com/kercylan98/woats/core/woats/wtype"
	"github.com/tealeg/xlsx"
	"strings"
)

type BilingualSchoolAffiliatedToEastChinaNormalUniversity struct {
}

func (slf *BilingualSchoolAffiliatedToEastChinaNormalUniversity) Modification(modifier *woats.ModifierFactory) exception.Exception {
	var tss = make([]*wtype.TimeSlot, 55)
	for index := 0; index < 5; index++ {
		tss[index*11+0] = wtype.NewTimeSlot(index*11+1, index+1, 8, 0, 8, 40)
		tss[index*11+1] = wtype.NewTimeSlot(index*11+2, index+1, 8, 50, 9, 30)
		tss[index*11+2] = wtype.NewTimeSlot(index*11+3, index+1, 9, 40, 10, 20)
		tss[index*11+3] = wtype.NewTimeSlot(index*11+4, index+1, 10, 30, 11, 10)
		tss[index*11+4] = wtype.NewTimeSlot(index*11+5, index+1, 11, 20, 12, 0)
		tss[index*11+5] = wtype.NewTimeSlot(index*11+6, index+1, 14, 30, 15, 10)
		tss[index*11+6] = wtype.NewTimeSlot(index*11+7, index+1, 15, 20, 16, 00)
		tss[index*11+7] = wtype.NewTimeSlot(index*11+8, index+1, 16, 10, 16, 0)
		tss[index*11+8] = wtype.NewTimeSlot(index*11+9, index+1, 16, 10, 16, 50)
		tss[index*11+9] = wtype.NewTimeSlot(index*11+10, index+1, 17, 0, 17, 40)
		tss[index*11+10] = wtype.NewTimeSlot(index*11+11, index+1, 17, 50, 18, 30)
	}

	excel, err := xlsx.OpenFile("designer/assets/2019-2020-confusion.xlsx")
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
			sectionSum, err := row.Cells[4].Int()
			if err != nil {
				return define.DataConversionException.Hit().Supplement("err", err)
			}

			var section = make([]int, sectionSum)
			for i := 0; i < sectionSum; i++ {
				section[i] = 1
			}

			class := strings.TrimSpace(row.Cells[1].String())
			ecm := modifier.ElectiveClass(class, section, tss).
				SetDisable(1, 2, 22, 44, 53, 54, 55).
				AddStudent(students[class]...)
			if strings.TrimSpace(row.Cells[6].String()) == "高二" {
				ecm.SetDisable(11, 21)
			}
			ecm.Ok().AddTeacher(strings.TrimSpace(row.Cells[3].String())).Ok()
		}
	}

	return nil
}
