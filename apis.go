package pigeon

import (
	"context"
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

type MsgError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (m MsgError) Error() string {
	return m.Msg
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

func (sdk *SDK) SendSMS(ctx context.Context, appID string, msgID, content string, phoneNumbers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/sms", sdk.BaseURL, appID)

	var req struct {
		Content   string   `json:"content"`
		MessageID string   `json:"message_id,omitempty"`
		Phones    []string `json:"phones,omitempty"`
	}
	req.Content = content
	req.MessageID = msgID
	req.Phones = phoneNumbers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinText(ctx context.Context, appID string, msgID, text string, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/text", sdk.BaseURL, appID)

	var req struct {
		Text      string   `json:"text"`
		MessageID string   `json:"message_id,omitempty"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Text = text
	req.MessageID = msgID
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinPost(ctx context.Context, appID string, msgID, post string, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/post", sdk.BaseURL, appID)

	var req struct {
		Text      string   `json:"text"`
		MessageID string   `json:"message_id,omitempty"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Text = post
	req.MessageID = msgID
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinSticker(ctx context.Context, appID string, sticker *models.Sticker, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/sticker", sdk.BaseURL, appID)

	var req struct {
		models.Sticker
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Sticker = *sticker
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinLocation(ctx context.Context, appID string, location *models.Location, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/location", sdk.BaseURL, appID)

	var req struct {
		models.Location
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Location = *location
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinLive(ctx context.Context, appID string, live *models.Live, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/live", sdk.BaseURL, appID)

	var req struct {
		models.Live
		Receivers []string `json:"receivers,omitempty"`
	}
	req.Live = *live
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinContact(ctx context.Context, appID string, msgID, mixinUserID string, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/contact", sdk.BaseURL, appID)

	var req struct {
		UserID    string   `json:"user_id"`
		MessageID string   `json:"message_id"`
		Receivers []string `json:"receivers,omitempty"`
	}
	req.UserID = mixinUserID
	req.MessageID = msgID
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinButtonGroup(ctx context.Context, appID string, msgID string, buttons []*models.MsgButton, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/buttongroup", sdk.BaseURL, appID)

	var req struct {
		Buttons   []*models.MsgButton `json:"buttons"`
		MessageID string              `json:"message_id,omitempty"`
		Receivers []string            `json:"receivers,omitempty"`
	}
	req.Buttons = buttons
	req.MessageID = msgID
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinAppcard(ctx context.Context, appID string, appcard *models.AppCard, receivers []string) *MsgError {
	url := fmt.Sprintf("%s/apps/%s/mixin/msg/appcard", sdk.BaseURL, appID)

	var req struct {
		models.AppCard
		Receivers []string `json:"receivers,omitempty"`
	}
	req.AppCard = *appcard
	req.Receivers = receivers

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}

func (sdk *SDK) SendMixinMultiMsg(ctx context.Context, appID string, messages []models.IMixinMsg, receivers []string) *MsgError {
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

	var okResp models.OkResponse
	_, e := fhttp.Execute(sdk.requestWithBasicAuth(ctx), "POST", url, &req, &okResp)
	if e != nil && !e.IsSuccessful() {
		fmt.Println(url, ":error:", e, "code:", e.Code(), ";msg:", e.Message())
		return &MsgError{Code: e.Code(), Msg: e.Message()}
	}

	return nil
}
