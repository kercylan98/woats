package modifier

import "github.com/kercylan98/work-out-a-teaching-schedule/core/woats/wtype"

func NewOrdinaryCLassModifierInitOptional(modifier *OrdinaryClassModifier) *OrdinaryCLassModifierInitOptional {
	return &OrdinaryCLassModifierInitOptional{
		modifier: modifier,
	}
}

// OrdinaryCLassModifierInitOptional 行政班级细节修饰器
type OrdinaryCLassModifierInitOptional struct {
	modifier 	*OrdinaryClassModifier
}

func (slf *OrdinaryCLassModifierInitOptional) Ok() *OrdinaryClassModifier {
	return slf.modifier
}

// AddLeader 添加班主任
func (slf *OrdinaryCLassModifierInitOptional) AddLeader(uniqueSign ...string) *OrdinaryCLassModifierInitOptional {
	slf.modifier.leader = append(slf.modifier.leader, uniqueSign...)
	return slf
}

// AddStudent 添加学生
func (slf *OrdinaryCLassModifierInitOptional) AddStudent(student ...string) *OrdinaryCLassModifierInitOptional {
	for _, s := range student {
		slf.modifier.students = append(slf.modifier.students, wtype.NewStudent(s, slf.modifier.class))
	}
	return slf
}