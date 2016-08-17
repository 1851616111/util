package relationship

func (p *RelationShip) Convert() interface{} {
	return &struct {
		RelationShip
		Type string `json:"type"`
	}{
		*p,
		"relationship",
	}
}

func (l *RelationShipList) Convert() interface{} {
	converter := []interface{}{}
	for _, relationShipt := range *l {
		converter = append(converter, relationShipt.Convert())
	}

	return converter
}

func Merge(l ...*RelationShipList) *RelationShipList {

	tmp := []*RelationShip{}
	for _, list := range l {
		tmp = append(tmp, []*RelationShip(*list)...)
	}

	ret := RelationShipList(tmp)
	return &ret
}
