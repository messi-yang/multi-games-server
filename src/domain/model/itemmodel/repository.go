package itemmodel

type Repo interface {
	GetAllItems() []ItemAgg
}
