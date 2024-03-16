package embedunitmodel

type EmbedUnitRepo interface {
	Add(EmbedUnit) error
	Get(EmbedUnitId) (EmbedUnit, error)
	Update(EmbedUnit) error
	Delete(EmbedUnit) error
}
