package commandmodel

type CommandRepo interface {
	Add(Command) error
}
