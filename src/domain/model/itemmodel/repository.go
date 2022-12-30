package itemmodel

type Repo interface {
	GetAllItems() []Item
}
