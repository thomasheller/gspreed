package gspreed

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/thomasheller/gncloud"
)

const (
	RoomType = 3

	baseURL = "apps/spreed/api/v1/room"
)

type Spreed struct {
	nc *gncloud.Nextcloud
}

func NewSpreed(nextcloud *gncloud.Nextcloud) *Spreed {
	return &Spreed{nc: nextcloud}
}

func (s Spreed) ListRooms() (*[]Room, error) {
	var result []Room

	err := s.nc.Get(baseURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s Spreed) FindRoomTokens(name string) ([]string, error) {
	roomList, err := s.ListRooms()
	if err != nil {
		return nil, err
	}

	var tokens []string

	for _, room := range *roomList {
		if room.DisplayName == name {
			tokens = append(tokens, room.Token)
		}
	}

	return tokens, nil
}

func (s Spreed) CreateRoom(name string) (token string, err error) {
	var result CreateRoomResult

	data := url.Values{}
	data.Set("roomName", name)
	data.Set("roomType", strconv.Itoa(RoomType))

	err = s.nc.Post(baseURL, data, &result)
	if err != nil {
		return
	}

	token = result.Token
	return
}

func (s Spreed) SetRoomPassword(roomToken, password string) error {
	data := url.Values{}
	data.Set("password", password)

	url := s.internalRoomURL(roomToken) + "/password"
	err := s.nc.Put(url, data, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s Spreed) RemoveRoomByToken(roomToken string) error {
	url := s.internalRoomURL(roomToken)
	err := s.nc.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s Spreed) RoomURL(roomToken string) string {
	return fmt.Sprintf("%s/index.php/call/%s", s.nc.BaseURL, roomToken)
}

func (s Spreed) internalRoomURL(roomToken string) string {
	return fmt.Sprintf("%s/%s", baseURL, roomToken)
}
