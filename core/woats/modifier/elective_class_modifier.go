package modifier

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats/define"
	"github.com/kercylan98/woats/core/woats/utils"
	"github.com/kercylan98/woats/core/woats/wtype"
)

func NewElectiveClassModifier(course string, section []int, slot []*wtype.TimeSlot) *ElectiveClassModifier {
	return &ElectiveClassModifier{
		course:      course,
		teachers:    []string{},
		teacherMode: ElectiveClassModifierTeacherModeTeacherWhole,
		students:    []*wtype.Student{},
		place:       map[string]*wtype.Location{},
		placeMode:   ElectiveClassModifierPlaceModePlaceChooseOne,
		section:     section,
		fixed:       []int{},
		disable:     []int{},
		priority:    map[int]float64{},
		slot:        slot,
	}
}

const (
	// ElectiveClassModifierTeacherModeTeacherChooseOne 课程对应任教为选择其中之一
	ElectiveClassModifierTeacherModeTeacherChooseOne = iota
	// ElectiveClassModifierTeacherModeTeacherWhole 课程对应教师为全体成员
	ElectiveClassModifierTeacherModeTeacherWhole
	// ElectiveClassModifierPlaceModePlaceChooseOne 可选教学场地之一
	ElectiveClassModifierPlaceModePlaceChooseOne = iota
	// ElectiveClassModifierPlaceModePlaceWhole 全部教学场地
	ElectiveClassModifierPlaceModePlaceWhole
)

// ElectiveClassModifier 选修班级细节修饰器
type ElectiveClassModifier struct {
	course      string                     // 选修班对应课程
	teachers    []string                   // 选修班任课教师
	teacherMode int                        // 教师和课程关系模型
	students    []*wtype.Student           // 选修班学生
	place       map[string]*wtype.Location // 选修班教学场地
	placeMode   int                        // 教学场地模式
	section     []int                      // 课时数 []int{1, 1, 2} >> 表总共4节，其中2节为连堂课
	fixed       []int                      // 固定课 []int{11, 15} >> 表示11和15号课位为固定课
	disable     []int                      // 该选修班禁排课位号
	priority    map[int]float64            // 特定课位的排课优先级
	slot        []*wtype.TimeSlot          // 上课课位
}

func (slf *ElectiveClassModifier) ToFactor() []wtype.Factor {
	var arr []wtype.Factor
	for _, section := range slf.section {
		for i := 0; i < section; i++ {
			arr = append(arr, &wtype.FactorInfo{
				UniqueSign:  slf.course,
				Course:      slf.course,
				Teacher:     slf.teachers,
				TeacherMode: slf.teacherMode,
				Students:    slf.students,
				Place:       slf.place,
				PlaceMode:   slf.placeMode,
				Section:     section,
				Fixed:       -1,
				Disable:     slf.disable,
				Priority:    slf.priority,
				Slot:        slf.slot,
			})
		}
	}
	for _, fixed := range slf.fixed {
		arr = append(arr, &wtype.FactorInfo{
			UniqueSign:  slf.course,
			Course:      slf.course,
			Teacher:     slf.teachers,
			TeacherMode: slf.teacherMode,
			Students:    slf.students,
			Place:       slf.place,
			PlaceMode:   slf.placeMode,
			Section:     1,
			Fixed:       fixed,
			Disable:     slf.disable,
			Priority:    slf.priority,
			Slot:        slf.slot,
		})
	}
	return arr
}

func (slf *ElectiveClassModifier) DataOptimization() {
	slf.teachers = utils.RemoveRepeatedString(slf.teachers)
	slf.fixed = utils.RemoveRepeatedInt(slf.fixed)
	slf.disable = utils.RemoveRepeatedInt(slf.disable)
}

func (slf *ElectiveClassModifier) DataVerification() exception.Exception {
	if len(slf.slot) == 0 {
		return define.NotFoundException.Hit().Supplement(slf.course, "slot count is zero")
	}
	var slotTemp = map[int]bool{}
	for _, slot := range slf.slot {
		if _, exist := slotTemp[slot.Index]; exist {
			return define.IdenticalSlotIndexException.Hit().Supplement(slf.course, slot.Index)
		}
		slotTemp[slot.Index] = true
	}
	return nil
}

// AddTeacher 添加任课教师
func (slf *ElectiveClassModifier) AddTeacher(uniqueSign ...string) *ElectiveClassModifierAddTeacherOptional {
	slf.teachers = append(slf.teachers, uniqueSign...)
	return &ElectiveClassModifierAddTeacherOptional{slf}
}

// AddPlace 添加教学场地
func (slf *ElectiveClassModifier) AddPlace(uniqueSign ...string) *ElectiveClassModifierAddPlaceOptional {
	for i := 0; i < len(uniqueSign); i++ {
		slf.place[uniqueSign[i]] = &wtype.Location{
			UniqueSign: uniqueSign[i],
		}
	}
	return &ElectiveClassModifierAddPlaceOptional{slf, uniqueSign}
}
