package api

import (
	"encoding/json"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/gorilla/websocket"
)

type SocketPresenter struct {
	socketConn     *websocket.Conn
	socketConnLock *sync.RWMutex
}

func NewSocketPresenter(socketConn *websocket.Conn, socketConnLock *sync.RWMutex) gameappservice.Presenter {
	return &SocketPresenter{socketConn: socketConn, socketConnLock: socketConnLock}
}

func (presenter *SocketPresenter) OnMessage(jsonObj any) error {
	presenter.socketConnLock.Lock()
	defer presenter.socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(jsonObj)

	compressedMessage, _ := gziputil.Gzip(messageJsonInBytes)

	return presenter.socketConn.WriteMessage(2, compressedMessage)
}
