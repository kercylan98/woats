package wtype

import (
	"sort"
)

// Factor 排课因子，包含了完整的一条教学计划信息
type Factor interface {
	// GetUniqueSign 获取唯一标识
	GetUniqueSign() string
	// GetCourse 获取课程标识
	GetCourse() string
	// GetTeacher 获取任课教师
	GetTeacher() []string
	// GetTeacherMode 获取任课教师选择模式
	GetTeacherMode() int
	// GetStudent 获取学生
	GetStudent() []*Student
	// GetStudentUniqueSign 获取学生唯一标识
	GetStudentUniqueSign() []string
	// GetPlace 获取教学场地
	GetPlace() map[string]*Location
	// GetPlaceUniqueSign 获取场地唯一标识
	GetPlaceUniqueSign() []string
	// GetPlaceMode 获取教学场地选择模式
	GetPlaceMode() int
	// GetContinuity 获取连堂课数量
	GetContinuity() int
	// IsFixed 是否是固定课
	IsFixed() bool
	// GetFixed 获取固定课位号
	GetFixed() int
	// GetDisable 获取禁排课位号
	GetDisable() []int
	// GetPriority 获取所有课位优先级设置
	GetPriority() map[int]float64
	// GetPriorityAVE 获取所有课位优先级均值
	GetPriorityAVE() float64
	// GetPrioritySUM 获取所有课位优先级总和
	GetPrioritySUM() float64
	// GetSlot 获取所有课位
	GetSlot() []*TimeSlot
	// GetSlotMap 获取所有课位map形式
	GetSlotMap() map[int]*TimeSlot
	// GetSlotWithIndex 根据索引获取课位
	GetSlotWithIndex(index int) *TimeSlot
	// GetTimeSlot 获取自身所在课位
	GetTimeSlot() *TimeSlot
	// GetSlotWithSameDay 获取相同天的课位
	GetSlotWithSameDay(slot *TimeSlot) []*TimeSlot
	// IsDisableChange 是否禁止调整(使用Pop和Move将无效)
	IsDisableChange() bool
	// IsNoChange 是否不应该调整(应该在策略中进行考虑)
	IsNoChange() bool
	// SetDisableChange 设置禁止调整，设置后无法取消
	SetDisableChange()
	// SetNoChange 设置是否不应该再调整，可重新设置
	SetNoChange(noChange bool)
	// IsGroup 是否是连堂课
	IsGroup() bool
	// GetGroup 获取其他连堂对象
	GetGroup() FactorGroup
}

// FactorInfo 排课因子数据结构定义
type FactorInfo struct {
	UniqueSign  string               // 唯一标识(选修班或行政班标识)
	Course      string               // 课程标识
	Teacher     []string             // 一组任课教师
	TeacherMode int                  // 教师选择模式
	Students    []*Student           // 一组学生
	Place       map[string]*Location // 一组教学场地
	PlaceMode   int                  // 教学场地选择模式
	Section     int                  // 课时数（大于1表示为连堂课）
	Fixed       int                  // 固定课位号码(-1不考虑)，永远占1课时
	Disable     []int                // 禁排课位号码
	Priority    map[int]float64      // 课位优先级(-1和不设置等同)
	Slot        []*TimeSlot          // 课位信息
	Group       FactorGroup          // 连堂课组

	TimeSlot      *TimeSlot // 当前所在课位
	DisableChange bool      // 该因子放上课位后是否禁止再调整
	NoChange      bool      // 该因子放上课位后是否不应该再调整（允许在实在困难的情况下被调整）
}

func (slf *FactorInfo) IsGroup() bool {
	return slf.Group != nil && len(slf.Group) > 0
}

func (slf *FactorInfo) GetGroup() FactorGroup {
	if len(slf.Group) == 0 {
		return []Factor{slf}
	}
	group := make(FactorGroup, 0)
	for _, factor := range slf.Group {
		if factor != nil || factor != slf {
			group = append(group, factor)
		}
	}
	return group
}

func (slf *FactorInfo) IsDisableChange() bool {
	return slf.DisableChange
}

func (slf *FactorInfo) IsNoChange() bool {
	return slf.NoChange
}

func (slf *FactorInfo) SetDisableChange() {
	slf.DisableChange = true
}

func (slf *FactorInfo) SetNoChange(noChange bool) {
	slf.NoChange = noChange
}

func (slf *FactorInfo) GetSlotWithSameDay(slot *TimeSlot) []*TimeSlot {
	var result []*TimeSlot
	for _, timeSlot := range slf.Slot {
		if timeSlot.WhatDay == slot.WhatDay {
			result = append(result, timeSlot)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Index < result[j].Index
	})
	return result
}

func (slf *FactorInfo) GetSlotMap() map[int]*TimeSlot {
	var m = map[int]*TimeSlot{}
	for _, slot := range slf.GetSlot() {
		m[slot.Index] = slot
	}
	return m
}

func (slf *FactorInfo) GetSlotWithIndex(index int) *TimeSlot {
	for _, slot := range slf.GetSlot() {
		if slot.Index == index {
			return slot
		}
	}
	return nil
}

func (slf *FactorInfo) GetPrioritySUM() float64 {
	var sum = 0.0
	for _, f := range slf.Priority {
		sum += f
	}
	return sum
}

func (slf *FactorInfo) GetPriorityAVE() float64 {
	var sum = 0.0
	for _, f := range slf.Priority {
		sum += f
	}
	if sum == 0 {
		return -1
	}

	// 保留两位小数
	//f,_ := strconv.ParseFloat(fmt.Sprint(sum / float64(len(slf.Priority))),64)
	//pow10N := math.Pow10(2)
	//return math.Trunc((f + 0.5 / pow10N) * pow10N) / pow10N

	return sum / float64(len(slf.Priority))
}

func (slf *FactorInfo) GetTimeSlot() *TimeSlot {
	return slf.TimeSlot
}

func (slf *FactorInfo) GetFixed() int {
	return slf.Fixed
}

func (slf *FactorInfo) GetPlaceUniqueSign() []string {
	var result []string
	for _, p := range slf.Place {
		result = append(result, p.UniqueSign)
	}
	return result
}

func (slf *FactorInfo) GetCourse() string {
	return slf.Course
}

func (slf *FactorInfo) GetStudentUniqueSign() []string {
	var result = make([]string, len(slf.Students))
	for i := 0; i < len(result); i++ {
		result[i] = slf.Students[i].UniqueSign
	}
	return result
}

func (slf *FactorInfo) GetSlot() []*TimeSlot {
	return slf.Slot
}

func (slf *FactorInfo) GetUniqueSign() string {
	return slf.UniqueSign
}

func (slf *FactorInfo) GetTeacher() []string {
	return slf.Teacher
}

func (slf *FactorInfo) GetTeacherMode() int {
	return slf.TeacherMode
}

func (slf *FactorInfo) GetStudent() []*Student {
	return slf.Students
}

func (slf *FactorInfo) GetPlace() map[string]*Location {
	return slf.Place
}

func (slf *FactorInfo) GetPlaceMode() int {
	return slf.PlaceMode
}

func (slf *FactorInfo) GetContinuity() int {
	return slf.Section
}

func (slf *FactorInfo) IsFixed() bool {
	return slf.Fixed != -1
}

func (slf *FactorInfo) GetDisable() []int {
	return slf.Disable
}

func (slf *FactorInfo) GetPriority() map[int]float64 {
	return slf.Priority
}
