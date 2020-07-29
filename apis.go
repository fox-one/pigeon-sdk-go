package pigeon

import (
	"context"
	"errors"
	"fmt"

	"github.com/fox-one/pigeon-sdk-go/fhttp"
	"github.com/fox-one/pigeon-sdk-go/models"
	"github.com/go-resty/resty/v2"
)

type SDK struct {
	AppKey    string
	AppSecret string
	BaseURL   string
}

func New(baseURL, appKey, appSecret string) *SDK {
	return &SDK{
		AppKey:    appKey,
		AppSecret: appSecret,
		BaseURL:   baseURL,
	}
}

func (sdk *SDK) requestWithBasicAuth(ctx context.Context) *resty.Request {
	return fhttp.Request(ctx).SetBasicAuth(sdk.AppKey, sdk.AppSecret)
}

func (sdk *SDK) SendSMS(ctx context.Context, appID string, content string, phoneNumber ...string) error {
	url := fmt.Sprintf("%s/apps/%s/sms", sdk.BaseURL, appID)

	var req struct {
		Content string   `json:"content"`
		Phones  []string `json:"phones,omitempty"`
	}
	req.Content = content
	req.Phones = phoneNumber

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send sms error")
	}

	return nil
}

func (sdk *SDK) SendMixinText(ctx context.Context, appID string, text string, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/text", sdk.BaseURL, appID)

	var req struct {
		Text      string   `json:"text"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Text = text
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinPost(ctx context.Context, appID string, post string, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/post", sdk.BaseURL, appID)

	var req struct {
		Text      string   `json:"text"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Text = post
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinSticker(ctx context.Context, appID string, sticker *models.Sticker, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/sticker", sdk.BaseURL, appID)

	var req struct {
		models.Sticker
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Sticker = *sticker
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinLocation(ctx context.Context, appID string, location *models.Location, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/location", sdk.BaseURL, appID)

	var req struct {
		models.Location
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Location = *location
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinLive(ctx context.Context, appID string, live *models.Live, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/live", sdk.BaseURL, appID)

	var req struct {
		models.Live
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Live = *live
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinContact(ctx context.Context, appID string, mixinUserID string, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/contact", sdk.BaseURL, appID)

	var req struct {
		UserID    string   `json:"user_id"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.UserID = mixinUserID
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinButtonGroup(ctx context.Context, appID string, buttons []*models.MsgButton, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/buttongroup", sdk.BaseURL, appID)

	var req struct {
		Buttons   []*models.MsgButton `json:"buttons"`
		Receivers []string            `json:"receivers,omitempty"`
	}
	req.Buttons = buttons
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinAppcard(ctx context.Context, appID string, appcard *models.AppCard, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/appcard", sdk.BaseURL, appID)

	var req struct {
		models.AppCard
		Receivers []string `json:"receivers,omitempty"`
	}
	req.AppCard = *appcard
	req.Receivers = receivers

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}

func (sdk *SDK) SendMixinMultiMsg(ctx context.Context, appID string, messages []models.IMixinMsg, receivers ...string) error {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/more", sdk.BaseURL, appID)

	var req struct {
		Items     []*models.MixinMsgWrapper `json:"items"`
		Receivers []string                  `json:"receivers,omitempty"`
	}
	req.Receivers = receivers
	items := make([]*models.MixinMsgWrapper, 0)
	for _, m := range messages {
		items = append(items, &models.MixinMsgWrapper{
			Type: m.Type(),
			Body: m.Body(),
		})
	}

	req.Items = items

	// reqb, reqE := json.Marshal(&req)
	// if reqE != nil {
	// 	return reqE
	// }

	// fmt.Println(string(reqb))

	var okResp models.OkResponse
	var errResp models.ErrorResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp, &errResp)
	if e != nil || errResp.Code != models.SuccessCode {
		fmt.Println(url, ":error:", e, "code:", errResp.Code, ";msg:", errResp.Message)
		return errors.New("send mixin text error")
	}

	return nil
}
