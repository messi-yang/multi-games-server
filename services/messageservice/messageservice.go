package messageservice

type SubscribeCallback func(bytes []byte)

type MessageService interface {
	Publish(topic string, message []byte)
	Subscribe(topic string, callback SubscribeCallback)
}

type messageServiceImpl struct {
	subscriberCallbacks map[string][]SubscribeCallback
}

var messageService MessageService

func GetMessageService() MessageService {
	if messageService == nil {
		messageService = &messageServiceImpl{
			subscriberCallbacks: make(map[string][]SubscribeCallback),
		}
		return messageService
	} else {
		return messageService
	}
}

func (msi *messageServiceImpl) Publish(topic string, bytes []byte) {
	if msi.subscriberCallbacks[topic] == nil {
		return
	}

	for _, callback := range msi.subscriberCallbacks[topic] {
		callback(bytes)
	}
}

func (msi *messageServiceImpl) Subscribe(topic string, callback SubscribeCallback) {
	if msi.subscriberCallbacks[topic] == nil {
		msi.subscriberCallbacks[topic] = []SubscribeCallback{}
	}

	msi.subscriberCallbacks[topic] = append(msi.subscriberCallbacks[topic], callback)
}
