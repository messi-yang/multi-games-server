package itemmodel

type ItemIdVo struct {
	id int16
}

func NewItemIdVo(id int16) ItemIdVo {
	return ItemIdVo{
		id: id,
	}
}

func (id ItemIdVo) IsEqual(anotherId ItemIdVo) bool {
	return id.id == anotherId.id
}

func (id ItemIdVo) ToInt16() int16 {
	return id.id
}
