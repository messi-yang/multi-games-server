package messageservice

import "fmt"

type subscribeCallback func(bytes []byte)
type subscriber struct {
	callback subscribeCallback
	token    string
}

type MessageService interface {
	Publish(topic string, message []byte)
	Subscribe(topic string, callback subscribeCallback) (subsriptionToken string)
	Unsubscribe(topic string, subsriptionToken string)
}

type messageServiceImpl struct {
	subscribers map[string][]subscriber
}

var messageService MessageService

func GetMessageService() MessageService {
	if messageService == nil {
		messageService = &messageServiceImpl{
			subscribers: make(map[string][]subscriber),
		}
		return messageService
	} else {
		return messageService
	}
}

func (msi *messageServiceImpl) Publish(topic string, bytes []byte) {
	if msi.subscribers[topic] == nil {
		return
	}

	for _, subscriber := range msi.subscribers[topic] {
		subscriber.callback(bytes)
	}
}

func (msi *messageServiceImpl) Subscribe(topic string, callback subscribeCallback) (subsriptionToken string) {
	if msi.subscribers[topic] == nil {
		msi.subscribers[topic] = []subscriber{}
	}

	newSubsriptionToken := generateRandomHash(10)
	msi.subscribers[topic] = append(msi.subscribers[topic], subscriber{
		callback: callback,
		token:    newSubsriptionToken,
	})

	return newSubsriptionToken
}

func (msi *messageServiceImpl) Unsubscribe(topic string, subsriptionToken string) {
	if msi.subscribers[topic] == nil {
		return
	}

	var subscriberWithGivenTokenIndex int = -1
	for subscriberIdx, subscriber := range msi.subscribers[topic] {
		if subscriber.token == subsriptionToken {
			fmt.Println("Found it")
			fmt.Println(subscriber.token)
			subscriberWithGivenTokenIndex = subscriberIdx
		}
	}

	if subscriberWithGivenTokenIndex > -1 {
		msi.subscribers[topic] = append(
			msi.subscribers[topic][:subscriberWithGivenTokenIndex],
			msi.subscribers[topic][subscriberWithGivenTokenIndex+1:]...,
		)
	}
}
