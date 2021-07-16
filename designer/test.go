package designer

import (
	"github.com/kercylan98/exception"
	"github.com/kercylan98/woats/core/woats"
	"github.com/kercylan98/woats/core/woats/wtype"
)

type Test struct {
}

func (slf *Test) Modification(modifier *woats.ModifierFactory) exception.Exception {
	tss := []*wtype.TimeSlot{
		wtype.NewTimeSlot(1, 1, 8, 00, 8, 40),
		wtype.NewTimeSlot(2, 1, 8, 50, 9, 30),
		wtype.NewTimeSlot(3, 1, 9, 40, 10, 20),
		wtype.NewTimeSlot(4, 1, 10, 30, 11, 10),
		wtype.NewTimeSlot(5, 1, 11, 20, 12, 00),
		wtype.NewTimeSlot(6, 1, 14, 30, 15, 10),
		wtype.NewTimeSlot(7, 1, 15, 20, 16, 00),
		wtype.NewTimeSlot(8, 1, 16, 10, 16, 50),

		wtype.NewTimeSlot(9, 2, 8, 00, 8, 40),
		wtype.NewTimeSlot(10, 2, 8, 50, 9, 30),
		wtype.NewTimeSlot(11, 2, 9, 40, 10, 20),
		wtype.NewTimeSlot(12, 2, 10, 30, 11, 10),
		wtype.NewTimeSlot(13, 2, 11, 20, 12, 00),
		wtype.NewTimeSlot(14, 2, 14, 30, 15, 10),
		wtype.NewTimeSlot(15, 2, 15, 20, 16, 00),
		wtype.NewTimeSlot(16, 2, 16, 10, 16, 50),

		wtype.NewTimeSlot(17, 3, 8, 00, 8, 40),
		wtype.NewTimeSlot(18, 3, 8, 50, 9, 30),
		wtype.NewTimeSlot(19, 3, 9, 40, 10, 20),
		wtype.NewTimeSlot(20, 3, 10, 30, 11, 10),
		wtype.NewTimeSlot(21, 3, 11, 20, 12, 00),
		wtype.NewTimeSlot(22, 3, 14, 30, 15, 10),
		wtype.NewTimeSlot(23, 3, 15, 20, 16, 00),
		wtype.NewTimeSlot(24, 3, 16, 10, 16, 50),

		wtype.NewTimeSlot(25, 4, 8, 00, 8, 40),
		wtype.NewTimeSlot(26, 4, 8, 50, 9, 30),
		wtype.NewTimeSlot(27, 4, 9, 40, 10, 20),
		wtype.NewTimeSlot(28, 4, 10, 30, 11, 10),
		wtype.NewTimeSlot(29, 4, 11, 20, 12, 00),
		wtype.NewTimeSlot(30, 4, 14, 30, 15, 10),
		wtype.NewTimeSlot(31, 4, 15, 20, 16, 00),
		wtype.NewTimeSlot(32, 4, 16, 10, 16, 50),

		wtype.NewTimeSlot(33, 5, 8, 00, 8, 40),
		wtype.NewTimeSlot(34, 5, 8, 50, 9, 30),
		wtype.NewTimeSlot(35, 5, 9, 40, 10, 20),
		wtype.NewTimeSlot(36, 5, 10, 30, 11, 10),
		wtype.NewTimeSlot(37, 5, 11, 20, 12, 00),
		wtype.NewTimeSlot(38, 5, 14, 30, 15, 10),
		wtype.NewTimeSlot(39, 5, 15, 20, 16, 00),
		wtype.NewTimeSlot(40, 5, 16, 10, 16, 50),
	}

	modifier.OrdinaryClass("小学一年级1班", 1, 1, tss).
		AddLeader("李鸣宇").
		AddStudent("小红", "小明", "小狼").Ok().
		//AddCourse("语文", "李玉涛").
		//Section(1, 1).
		//AddFixed(5, 2, 3, 28).
		//SetDisable(9).Ok().
		AddCourse("数学", "张华").Section(1).Ok().
		AddCourse("英语", "刘雨依").Section(1).Ok()
	//AddCourse("补刀课位").Section(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).Ok()

	//modifier.OrdinaryClass("小学一年级2班", 1, 1, tss).
	//	AddLeader("李鸣宇").
	//	AddStudent("小王", "小学生", "小东西").Ok().
	//	//AddCourse("语文", "李玉涛").Section(1, 1, 1, 2).AddFixed(4).SetDisable(9).Ok().
	//	AddCourse("数学", "张华").Section(1, 1).Ok().
	//	AddCourse("英语", "刘雨依").Section(1, 1).Ok()
	//AddCourse("补刀课位").Section(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).Ok()

	//modifier.OrdinaryClass("小学二年级1班", 1, 2, tss).
	//	AddLeader("张涛").
	//	AddStudent("大红", "大明", "大狼").Ok().
	//	AddCourse("语文", "王宇").Section(1, 1, 1, 2).AddFixed(7).SetDisable(9).Ok().
	//	AddCourse("数学", "宇智博").Section(1, 1, 1, 2).Ok().
	//	AddCourse("英语", "汪欣").Section(1, 1, 1, 1, 1, 1, 1).Ok().
	//	AddCourse("补刀课位").Section(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).Ok()
	//
	//modifier.OrdinaryClass("小学二年级2班", 1, 2, tss).
	//	AddLeader("张涛").
	//	AddStudent("大王", "大学生", "大东西").Ok().
	//	AddCourse("语文", "王宇").Section(1, 1, 1, 2).AddFixed(21).SetDisable(9).Ok().
	//	AddCourse("数学", "宇智博").Section(1, 1, 1, 2).SetPriority(3, 0.8).Ok().
	//	AddCourse("英语", "汪欣").Section(1, 1, 1, 1, 1, 1, 1).Ok().
	//	AddCourse("补刀课位").Section(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).Ok()
	//
	//modifier.ElectiveClass("绘画精进", []int{1, 1, 1, 1}, tss). //AddFixed(1).
	//							AddStudent("小王", "小学生", "小东西").Ok().
	//							AddTeacher("王宇").Ok()

	return nil
}
