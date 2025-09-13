package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
	Id     string
	Key    string
}

func NewClient(id, key string) *Client {
	return &Client{
		client: http.DefaultClient,
		Id:     id,
		Key:    key,
	}
}

func (c *Client) Search(sheng, place string) (*Response, error) {
	baseURL := "https://cn.apihz.cn/api/tianqi/tqyb.php"
	//?id=88888888&key=88888888&sheng=广东&place=广州
	url := fmt.Sprintf("%s?id=%s&key=%s&sheng=%s&place=%s", baseURL, c.Id, c.Key, sheng, place)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
