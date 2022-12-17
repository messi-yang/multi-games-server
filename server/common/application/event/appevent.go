package event

type AppEvent interface {
	Serialize() []byte
}
