package gspreed

type Room struct {
	DisplayName string `json:"displayName"`
	Token       string `json:"token"`
}

type CreateRoomResult struct {
	Token string `json:"token"`
}
