package gamesocket

import (
	"encoding/json"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/gorilla/websocket"
)

type presenter struct {
	socketConn     *websocket.Conn
	socketConnLock *sync.RWMutex
}

func newPresenter(socketConn *websocket.Conn, socketConnLock *sync.RWMutex) gamesocketappservice.Presenter {
	return &presenter{socketConn: socketConn, socketConnLock: socketConnLock}
}

func (presenter *presenter) OnMessage(jsonObj any) {
	presenter.socketConnLock.Lock()
	defer presenter.socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(jsonObj)

	compressedMessage, _ := gziputil.Gzip(messageJsonInBytes)

	presenter.socketConn.WriteMessage(2, compressedMessage)
}
