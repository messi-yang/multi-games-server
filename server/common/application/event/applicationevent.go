package event

type ApplicationEvent interface {
	Serialize() []byte
}
