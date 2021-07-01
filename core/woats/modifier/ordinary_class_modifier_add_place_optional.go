package modifier

type OrdinaryClassModifierAddPlaceOptional struct {
	modifier 			*OrdinaryClassModifier
	course 				string
	place				[]string
}

func (slf *OrdinaryClassModifierAddPlaceOptional) Ok() *OrdinaryClassModifier {
	return slf.modifier
}

// PlaceChooseOne 定义该课程的教学场地从提供的数组中抽选一名（默认）
func (slf *OrdinaryClassModifierAddPlaceOptional) PlaceChooseOne() *OrdinaryClassModifierAddPlaceOptional {
	slf.modifier.course[slf.course].placeMode = OrdinaryClassModifierPlaceModePlaceChooseOne
	return slf
}

// PlaceWhole 定义该课程的教学场地为提供的数组全体场地
func (slf *OrdinaryClassModifierAddPlaceOptional) PlaceWhole() *OrdinaryClassModifierAddPlaceOptional {
	slf.modifier.course[slf.course].placeMode = OrdinaryClassModifierPlaceModePlaceWhole
	return slf
}

// SetLocation 设置教学场地位置
func (slf *OrdinaryClassModifierAddPlaceOptional) SetLocation(getLocation func(uniqueSign string) (latitude, longitude float64)) *OrdinaryClassModifierAddPlaceOptional {
	for i := 0; i < len(slf.place); i++ {
		place := slf.modifier.course[slf.course].place[slf.place[i]]
		place.Latitude, place.Longitude = getLocation(place.UniqueSign)
	}
	return slf
}

// SetConvenientTransportation 设置到达该场地交通是否便利
func (slf *OrdinaryClassModifierAddPlaceOptional) SetConvenientTransportation(isConvenientTransportation func(uniqueSign string) bool) *OrdinaryClassModifierAddPlaceOptional {
	for i := 0; i < len(slf.place); i++ {
		place := slf.modifier.course[slf.course].place[slf.place[i]]
		place.IsConvenientTransportation = isConvenientTransportation(place.UniqueSign)
	}
	return slf
}