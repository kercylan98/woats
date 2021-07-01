package modifier

type ElectiveClassModifierAddPlaceOptional struct {
	modifier 			*ElectiveClassModifier
	place 				[]string
}

// Ok 完成可选项
func (slf *ElectiveClassModifierAddPlaceOptional) Ok() *ElectiveClassModifier {
	return slf.modifier
}

// PlaceChooseOne 定义该课程的教学场地从提供的数组中抽选一名（默认）
func (slf *ElectiveClassModifierAddPlaceOptional) PlaceChooseOne() *ElectiveClassModifierAddPlaceOptional {
	slf.modifier.placeMode = ElectiveClassModifierPlaceModePlaceChooseOne
	return slf
}

// PlaceWhole 定义该课程的教学场地为提供的数组全体场地
func (slf *ElectiveClassModifierAddPlaceOptional) PlaceWhole() *ElectiveClassModifierAddPlaceOptional {
	slf.modifier.placeMode = ElectiveClassModifierPlaceModePlaceWhole
	return slf
}

// SetLocation 设置教学场地位置
func (slf *ElectiveClassModifierAddPlaceOptional) SetLocation(getLocation func(uniqueSign string) (latitude, longitude float64)) *ElectiveClassModifierAddPlaceOptional {
	for i := 0; i < len(slf.place); i++ {
		place := slf.modifier.place[slf.place[i]]
		place.Latitude, place.Longitude = getLocation(place.UniqueSign)
	}
	return slf
}

// SetConvenientTransportation 设置到达该场地交通是否便利
func (slf *ElectiveClassModifierAddPlaceOptional) SetConvenientTransportation(isConvenientTransportation func(uniqueSign string) bool) *ElectiveClassModifierAddPlaceOptional {
	for i := 0; i < len(slf.place); i++ {
		place := slf.modifier.place[slf.place[i]]
		place.IsConvenientTransportation = isConvenientTransportation(place.UniqueSign)
	}
	return slf
}