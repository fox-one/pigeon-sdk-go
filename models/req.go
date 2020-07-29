package models

import (
	"github.com/shopspring/decimal"
)

type Text struct {
	Text string `json:"text"`
}

func (t Text) Type() string {
	return MsgTypeText
}

func (t Text) Body() interface{} {
	return t
}

type Post struct {
	Text string `json:"text"`
}

func (t Post) Type() string {
	return MsgTypePost
}

func (t Post) Body() interface{} {
	return t
}

type Sticker struct {
	Name    string `json:"name"`
	AlbumID string `json:"album_id"`
}

func (t Sticker) Type() string {
	return MsgTypeSticker
}

func (t Sticker) Body() interface{} {
	return t
}

type Contact struct {
	UserID string `json:"user_id"`
}

func (t Contact) Type() string {
	return MsgTypeContact
}

func (t Contact) Body() interface{} {
	return t
}

type Location struct {
	Longitude decimal.Decimal `json:"longitude"`
	Latitude  decimal.Decimal `json:"latitude"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
}

func (t Location) Type() string {
	return MsgTypeLocation
}

func (t Location) Body() interface{} {
	return t
}

type Live struct {
	Width    decimal.Decimal `json:"width"`
	Height   decimal.Decimal `json:"height"`
	ThumbURL string          `json:"thumb_url"`
	URL      string          `json:"url"`
}

func (t Live) Type() string {
	return MsgTypeLive
}

func (t Live) Body() interface{} {
	return t
}

type AppCard struct {
	IconURL     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

func (t AppCard) Type() string {
	return MsgTypeAppCard
}

func (t AppCard) Body() interface{} {
	return t
}

type MsgButton struct {
	Label  string `json:"label"`
	Color  string `json:"color"`
	Action string `json:"action"`
}

type ButtongGroup []*MsgButton

func (t ButtongGroup) Type() string {
	return MsgTypeButtongroup
}

func (t ButtongGroup) Body() interface{} {
	return t
}

type MixinMsgWrapper struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type IMixinMsg interface {
	Type() string
	Body() interface{}
}

const (
	MsgTypeText        = "text"
	MsgTypePost        = "post"
	MsgTypeSticker     = "sticker"
	MsgTypeLocation    = "location"
	MsgTypeLive        = "live"
	MsgTypeContact     = "contact"
	MsgTypeButtongroup = "buttongroup"
	MsgTypeAppCard     = "appcard"
)
