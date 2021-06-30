package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VkPhotoSize struct {
	Width  int
	Height int
	Url    string
}

type VkPhoto struct {
	Id         int64
	Photo_1280 string
	Photo_807  string
	Photo_2560 string
}

type VkAttachment struct {
	Photo VkPhoto
}

type VkPost struct {
	Id          int
	Text        string
	Attachments []VkAttachment
}

type VkWall struct {
	Count int
	Items []VkPost
}

type VkWallResp struct {
	Response VkWall
}

var (
	access_token string = "6cd4af8e6cd4af8e6cd4af8ee46cacc68066cd46cd4af8e0c0e49a1b3c939d6de610803"
)

func GetVkPublicPosts(ownerId string) ([]VkPost, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%s&count=100&v=5.52&access_token=%s", ownerId, access_token),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	var value VkWallResp
	if err := json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	return value.Response.Items, nil
}
