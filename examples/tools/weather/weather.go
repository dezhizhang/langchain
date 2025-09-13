package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"office-helper/examples/tools/weather/internal"
)

type ReqWeather struct {
	Sheng, Place string
}

type Tool struct {
	client *internal.Client
}

func New(opts ...Option) *Tool {
	opt := &options{
		apkKey: "88888888",
		id:     "88888888",
	}

	for _, option := range opts {
		option(opt)
	}

	return &Tool{client: internal.NewClient(opt.id, opt.apkKey)}
}

func (t *Tool) Name() string {
	return "weather"
}

func (t *Tool) Description() string {
	return `"a weather query interface. "
	"use when you answer information about the weather."
	"when you need to get city weather information, iIt is always the first option.
	"The input should be {"sheng": xx, "place": xx} wwhich can be parsed by json "
	"give an example{"sheng": "广东","广州"}"`
}

func (t *Tool) Call(ctx context.Context, input string) (string, error) {
	sheng, place, err := t.ParseInput(input)
	if err != nil {
		return "", err
	}
	res, err := t.client.Search(sheng, place)
	if err != nil {
		return "", err
	}
	return t.ParseOutPut(res)
}

func (t *Tool) ParseInput(input string) (sheng, place string, err error) {
	var r ReqWeather
	if err := json.Unmarshal([]byte(input), &r); err != nil {
		return "", "", err
	}
	return r.Sheng, r.Place, nil
}

func (t *Tool) ParseOutPut(resp *internal.Response) (string, error) {
	return fmt.Sprintf("某某", resp), nil
}
