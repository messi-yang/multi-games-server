package worldrolemodel

type Repo interface {
	Add(WorldRole) error
	Get(WorldRoleId) (WorldRole, error)
}
