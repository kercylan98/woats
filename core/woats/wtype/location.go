package wtype

// Location 位置数据结构
type Location struct {
	UniqueSign 	string
	Latitude 	float64	// 纬度
	Longitude 	float64	// 经度
	IsConvenientTransportation bool // 交通是否便利
}