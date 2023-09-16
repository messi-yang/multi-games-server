package userhttphandler

type updateMyUserRequestBody struct {
	Username     string `json:"username"`
	FriendlyName string `json:"friendlyName"`
}
