package messageservice

type subscribeCallback func(bytes []byte)
type subscriber struct {
	callback subscribeCallback
	token    string
}

type MessageService interface {
	Publish(topic string, message []byte)
	Subscribe(topic string, callback subscribeCallback)
}

type messageServiceImpl struct {
	subscriberCallbacks map[string][]subscribeCallback
}

var messageService MessageService

func GetMessageService() MessageService {
	if messageService == nil {
		messageService = &messageServiceImpl{
			subscriberCallbacks: make(map[string][]subscribeCallback),
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

func (msi *messageServiceImpl) Subscribe(topic string, callback subscribeCallback) {
	if msi.subscriberCallbacks[topic] == nil {
		msi.subscriberCallbacks[topic] = []subscribeCallback{}
	}

	msi.subscriberCallbacks[topic] = append(msi.subscriberCallbacks[topic], callback)
}
