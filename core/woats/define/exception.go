package define

import "github.com/kercylan98/exception"

// 系统异常(10xxx)
var (
	OpenFileException       = exception.Reg(10500, "open file failed", "文件打开失败")
	DataConversionException = exception.Reg(10501, "data conversion failed", "数据转换失败")
)

// 排课过程中异常(100xxx)
var (
	CourseConflictException  = exception.Reg(100100, "course conflict", "课程冲突")
	TeacherConflictException = exception.Reg(100101, "teacher conflict", "教师冲突")
	StudentConflictException = exception.Reg(100102, "student conflict", "学生冲突")
	PlaceConflictException   = exception.Reg(100103, "place conflict", "教学场地冲突")
	GroupConflictException   = exception.Reg(100104, "group course conflict", "连堂课冲突")

	NotFoundException         = exception.Reg(100400, "not found element", "未找到元素")
	SnapshotRecoveryException = exception.Reg(100401, "snapshot recover failed", "三维课表快照恢复失败，没有可恢复的快照")

	IdenticalSlotIndexException = exception.Reg(100500, "there are unreasonable slot", "存在相同的课位号")
	DisableSlotException        = exception.Reg(100501, "disabled slot", "被禁用的排课课位")
	FactorGroupPushException    = exception.Reg(100502, "factor group push failed", "连堂课放置失败")
)
