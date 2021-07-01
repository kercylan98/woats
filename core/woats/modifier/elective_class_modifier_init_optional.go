package modifier

import "github.com/kercylan98/woats/core/woats/wtype"

func NewElectiveClassModifierInitOptional(modifier *ElectiveClassModifier) *ElectiveClassModifierInitOptional {
	return &ElectiveClassModifierInitOptional{modifier: modifier}
}

type ElectiveClassModifierInitOptional struct {
	modifier *ElectiveClassModifier
}

// Ok 完成可选项
func (slf *ElectiveClassModifierInitOptional) Ok() *ElectiveClassModifier {
	return slf.modifier
}

// AddFixed 添加该课程特定课位的固定课
func (slf *ElectiveClassModifierInitOptional) AddFixed(slot ...int) *ElectiveClassModifierInitOptional {
	slf.modifier.fixed = append(slf.modifier.fixed, slot...)
	return slf
}

// SetPriority 设置该课程对应课位的优先级
func (slf *ElectiveClassModifierInitOptional) SetPriority(slot int, priority float64) *ElectiveClassModifierInitOptional {
	slf.modifier.priority[slot] = priority
	return slf
}

// SetDisable 设置该课程禁排课位
func (slf *ElectiveClassModifierInitOptional) SetDisable(slot ...int) *ElectiveClassModifierInitOptional {
	slf.modifier.disable = append(slf.modifier.disable, slot...)
	return slf
}

// AddStudent 添加学生
func (slf *ElectiveClassModifierInitOptional) AddStudent(student ...string) *ElectiveClassModifierInitOptional {
	for _, s := range student {
		slf.modifier.students = append(slf.modifier.students, wtype.NewStudent(s, slf.modifier.course))
	}
	return slf
}
