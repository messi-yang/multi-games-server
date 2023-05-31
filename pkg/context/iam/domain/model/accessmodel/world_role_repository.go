package accessmodel

type WorldRoleRepo interface {
	Add(WorldRole) error
	Get(WorldRoleId) (WorldRole, error)
}
