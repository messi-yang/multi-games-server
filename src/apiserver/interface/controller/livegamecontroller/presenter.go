package livegamecontroller

import (
	"encoding/json"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gzipprovider"
	"github.com/gorilla/websocket"
)

type socketPresenter struct {
	socketConn     *websocket.Conn
	socketConnLock *sync.RWMutex
}

func newSocketPresenter(socketConn *websocket.Conn, socketConnLock *sync.RWMutex) livegameappservice.Presenter {
	return &socketPresenter{socketConn: socketConn, socketConnLock: socketConnLock}
}

func (presenter *socketPresenter) OnSuccess(jsonObj any) {
	presenter.socketConnLock.Lock()
	defer presenter.socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(jsonObj)

	gzipCompressor := gzipprovider.New()
	compressedMessage, _ := gzipCompressor.Gzip(messageJsonInBytes)

	presenter.socketConn.WriteMessage(2, compressedMessage)
}
