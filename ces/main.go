package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

type Student struct {
	Number  string
	Name    string
	Courses []string
	Level   string
}

func (slf *Student) ToLine() []string {
	var lines []string
	for _, course := range slf.Courses {
		lines = append(lines, course+","+slf.Number+","+slf.Name)
	}
	return lines
}

func main() {
	var students = map[string]*Student{}

	excel, err := xlsx.OpenFile("/Users/kercylan/Documents/Work/上海市师悦信息科技有限公司/2.客户信息留档/2.华东师范大学附属双语学校/2021-2022年第一学期/学生选课结果/G12 subject selection-2020-2021_20210630.xlsx")
	if err != nil {
		panic(err)
	}

	for i, row := range excel.Sheets[0].Rows {
		if i == 0 {
			continue
		}
		number := row.Cells[1].String()
		name := row.Cells[2].String()
		courses := make([]string, 0)
		for i := 5; i < len(row.Cells); i++ {
			course := strings.TrimSpace(row.Cells[i].String())
			if course != "" {
				courses = append(courses, course)
			}
		}
		students[number] = &Student{
			Number:  number,
			Name:    name,
			Courses: courses,
			Level:   "DP2",
		}
	}

	for _, student := range students {
		for _, line := range student.ToLine() {
			fmt.Println(line)
		}
	}
}
