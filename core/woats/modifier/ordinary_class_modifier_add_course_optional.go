package modifier

type OrdinaryClassModifierAddCourseOptional struct {
	modifier 			*OrdinaryClassModifier
	course 				string
}

func (slf *OrdinaryClassModifierAddCourseOptional) Ok() *OrdinaryClassModifier {
	return slf.modifier
}

// TeacherChooseOne 定义该课程的任教从提供的数组中抽选一名
func (slf *OrdinaryClassModifierAddCourseOptional) TeacherChooseOne() *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].teacherMode = OrdinaryClassModifierTeacherModeTeacherChooseOne
	return slf
}

// TeacherWhole 定义该课程的任教为提供的数组全体教师（默认）
func (slf *OrdinaryClassModifierAddCourseOptional) TeacherWhole() *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].teacherMode = OrdinaryClassModifierTeacherModeTeacherWhole
	return slf
}

// Section 定义该课程特定连堂课的周课时数
func (slf *OrdinaryClassModifierAddCourseOptional) Section(section ...int) *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].section = append(slf.modifier.course[slf.course].section, section...)
	return slf
}

// AddFixed 添加该课程固定课
func (slf *OrdinaryClassModifierAddCourseOptional) AddFixed(slot ...int) *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].fixed = append(slf.modifier.course[slf.course].fixed, slot...)
	return slf
}

// SetPriority 设置该课程对应课位的优先级
func (slf *OrdinaryClassModifierAddCourseOptional) SetPriority(slot int, priority float64) *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].priority[slot] = priority
	return slf
}

// SetDisable 设置该课程禁排课位
func (slf *OrdinaryClassModifierAddCourseOptional) SetDisable(slot ...int) *OrdinaryClassModifierAddCourseOptional {
	slf.modifier.course[slf.course].disable = append(slf.modifier.course[slf.course].disable, slot...)
	return slf
}