package modifier

type ElectiveClassModifierAddTeacherOptional struct {
	modifier 			*ElectiveClassModifier
}

// Ok 完成可选项
func (slf *ElectiveClassModifierAddTeacherOptional) Ok() *ElectiveClassModifier {
	return slf.modifier
}

// TeacherChooseOne 定义该课程的任教从提供的数组中抽选一名
func (slf *ElectiveClassModifierAddTeacherOptional) TeacherChooseOne() *ElectiveClassModifierAddTeacherOptional {
	slf.modifier.teacherMode = ElectiveClassModifierTeacherModeTeacherChooseOne
	return slf
}

// TeacherWhole 定义该课程的任教为提供的数组全体教师（默认）
func (slf *ElectiveClassModifierAddTeacherOptional) TeacherWhole() *ElectiveClassModifierAddTeacherOptional {
	slf.modifier.teacherMode = ElectiveClassModifierTeacherModeTeacherWhole
	return slf
}
