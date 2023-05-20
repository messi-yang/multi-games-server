package worldhttphandler

type createWorldRequestBody struct {
	Name string `json:"name"`
}

type updateWorldRequestBody struct {
	Name string `json:"name"`
}
