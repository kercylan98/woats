package modifier

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/define"
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/utils"
	"github.com/kercylan98/work-out-a-teaching-schedule/core/woats/wtype"
)

func NewOrdinaryCLassModifier(class string, stage, level int, slot []*wtype.TimeSlot) *OrdinaryClassModifier {
	return &OrdinaryClassModifier{
		class:  class,
		stage:  stage,
		level:  level,
		leader: []string{},
		course: map[string]*OrdinaryClassModifierCourse{},
		slot:   slot,
		students: []*wtype.Student{},
	}
}

const (
	// OrdinaryClassModifierTeacherModeTeacherChooseOne 课程对应任教为选择其中之一
	OrdinaryClassModifierTeacherModeTeacherChooseOne = iota
	// OrdinaryClassModifierTeacherModeTeacherWhole 课程对应教师为全体成员
	OrdinaryClassModifierTeacherModeTeacherWhole
	// OrdinaryClassModifierPlaceModePlaceChooseOne 可选教学场地之一
	OrdinaryClassModifierPlaceModePlaceChooseOne = iota
	// OrdinaryClassModifierPlaceModePlaceWhole 全部教学场地
	OrdinaryClassModifierPlaceModePlaceWhole
)

// OrdinaryClassModifier 行政班级细节修饰器
type OrdinaryClassModifier struct {
	class  string                                  // 班级
	stage  int                                     // 学段
	level  int                                     // 年级
	leader []string                                // 班主任
	course map[string]*OrdinaryClassModifierCourse // 课程信息
	slot   []*wtype.TimeSlot                       // 上课课位
	students    []*wtype.Student           		   // 行政班学生
}

// OrdinaryClassModifierCourse 行政班修饰器课程详情数据结构
type OrdinaryClassModifierCourse struct {
	course      string                     // 课程
	teacher     []string                   // 教师
	teacherMode int                        // 教师和课程关系模型
	section     []int                      // 课时数 []int{1, 1, 2} >> 表总共4节，其中2节为连堂课
	fixed       []int                      // 固定课 []int{11, 15} >> 表示11和15号课位为固定课
	disable     []int                      // 该课程禁排课位号
	priority    map[int]float64            // 特定课位的排课优先级
	place       map[string]*wtype.Location // 选修班教学场地
	placeMode   int                        // 教学场地模式
}

func (slf *OrdinaryClassModifier) ToFactor() []wtype.Factor {
	var arr []wtype.Factor
	for _, course := range slf.course {
		for _, section := range course.section {
			for i := 0; i < section; i++ {
				arr = append(arr, &wtype.FactorInfo{
					UniqueSign:  slf.class,
					Course:      course.course,
					Teacher:     course.teacher,
					TeacherMode: course.teacherMode,
					Students: 	 slf.students,
					Place:       course.place,
					PlaceMode:   course.placeMode,
					Section:     section,
					Fixed:       -1,
					Disable:     course.disable,
					Priority:    course.priority,
					Slot:  		 slf.slot,
				})
			}
		}
		for _, fixed := range course.fixed {
			arr = append(arr, &wtype.FactorInfo{
				UniqueSign:  slf.class,
				Course:      course.course,
				Teacher:     course.teacher,
				TeacherMode: course.teacherMode,
				Students: 	 slf.students,
				Place:       course.place,
				PlaceMode:   course.placeMode,
				Section:     1,
				Fixed:       fixed,
				Disable:     course.disable,
				Priority:    course.priority,
				Slot:  		 slf.slot,
			})
		}
	}
	return arr
}

func (slf *OrdinaryClassModifier) DataOptimization() {
	slf.leader = utils.RemoveRepeatedString(slf.leader)
	for _, course := range slf.course {
		course.teacher = utils.RemoveRepeatedString(course.teacher)
		course.fixed = utils.RemoveRepeatedInt(course.fixed)
		course.disable = utils.RemoveRepeatedInt(course.disable)
	}
}

func (slf *OrdinaryClassModifier) DataVerification() exception.Exception {
	if len(slf.slot) == 0 {
		return define.NotFoundException.Hit().Supplement(slf.class, "slot count is zero")
	}
	var slotTemp = map[int]bool{}
	for _, slot := range slf.slot {
		if _, exist := slotTemp[slot.Index]; exist {
			return define.IdenticalSlotIndexException.Hit().Supplement(slf.class, slot.Index)
		}
		slotTemp[slot.Index] = true
	}

	return nil
}


// AddLeader 添加班主任
func (slf *OrdinaryClassModifier) AddLeader(uniqueSign ...string) *OrdinaryClassModifier {
	slf.leader = append(slf.leader, uniqueSign...)
	return slf
}

// AddCourse 添加班级课程
func (slf *OrdinaryClassModifier) AddCourse(courseUniqueSign string, teacherUniqueSign ...string) *OrdinaryClassModifierAddCourseOptional {
	slf.course[courseUniqueSign] = &OrdinaryClassModifierCourse{
		course:      courseUniqueSign,
		teacher:     teacherUniqueSign,
		teacherMode: OrdinaryClassModifierTeacherModeTeacherWhole,
		section:     []int{},
		fixed:       []int{},
		disable:     []int{},
		priority:    map[int]float64{},
		place:       map[string]*wtype.Location{},
		placeMode:   OrdinaryClassModifierPlaceModePlaceChooseOne,
	}
	return &OrdinaryClassModifierAddCourseOptional{slf, courseUniqueSign}
}

// AddPlace 添加教学场地
func (slf *OrdinaryClassModifier) AddPlace(courseUniqueSign string, placeUniqueSign ...string) *OrdinaryClassModifierAddPlaceOptional {
	for i := 0; i < len(placeUniqueSign); i++ {
		slf.course[courseUniqueSign].place[placeUniqueSign[i]] = &wtype.Location{
			UniqueSign: placeUniqueSign[i],
		}
	}
	return &OrdinaryClassModifierAddPlaceOptional{slf, courseUniqueSign, placeUniqueSign}
}