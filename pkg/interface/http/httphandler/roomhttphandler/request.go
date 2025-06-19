package roomhttphandler

type createRoomRequestBody struct {
	Name string `json:"name"`
}

type updateRoomRequestBody struct {
	Name string `json:"name"`
}
