package woats

import (
	"github.com/kercylan98/woats/core/woats/modifier"
	"github.com/kercylan98/woats/core/woats/wtype"
)

type ModifierFactory struct {
	product map[string]Modifier
}

// OrdinaryClass 采用特定修饰定义一个行政班
func (slf *ModifierFactory) OrdinaryClass(classUniqueSign string, stage, leve int, slot []*wtype.TimeSlot) *modifier.OrdinaryCLassModifierInitOptional {
	m := modifier.NewOrdinaryCLassModifier(classUniqueSign, stage, leve, slot)
	slf.product[classUniqueSign] = m
	return modifier.NewOrdinaryCLassModifierInitOptional(m)
}

// ElectiveClass 采用特定修饰定义一个选修班
func (slf *ModifierFactory) ElectiveClass(courseUniqueSign string, section []int, slot []*wtype.TimeSlot) *modifier.ElectiveClassModifierInitOptional {
	m := modifier.NewElectiveClassModifier(courseUniqueSign, section, slot)
	slf.product[courseUniqueSign] = m
	return modifier.NewElectiveClassModifierInitOptional(m)
}

// 例子
//func example()  {
//	new(ModifierFactory).ElectiveClass("阅读选修", []int{1, 1, 1, 2}, []*wtype.TimeSlot{
//		wtype.NewTimeSlot(1, 1, 8, 00, 8, 40),
//		wtype.NewTimeSlot(2, 1, 8, 50, 9, 30),
//	}).
//		AddFixed(1).SetDisable(1, 2, 4, 5).SetPriority(3, 0.3).AddStudent().Ok().
//		AddTeacher("张三", "李四").Ok().
//		AddPlace()
//}
