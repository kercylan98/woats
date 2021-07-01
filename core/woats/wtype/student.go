package wtype

func NewStudent(studentUniqueSign string, classUniqueSign string) *Student {
	return &Student{
		UniqueSign: studentUniqueSign,
		Class:      classUniqueSign,
	}
}

type Student struct {
	UniqueSign 	string	// 学生标识
	Class 		string  // 班级标识
}
