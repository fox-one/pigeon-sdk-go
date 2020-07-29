package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fox-one/pigeon-sdk-go"
	"github.com/fox-one/pigeon-sdk-go/models"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Hello pigeon")

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name: "all",
			Action: func(c *cli.Context) error {
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinText(context.Background(), "30002", "hello pigeon sdk", "*")
			},
		},
		{
			Name: "text",
			Action: func(c *cli.Context) error {
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinText(context.Background(), "20007", "hello pigeon sdk", "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "post",
			Action: func(c *cli.Context) error {
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinPost(context.Background(), "30002", "## hello pigeon sdk", "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "sticker",
			Action: func(c *cli.Context) error {
				sticker := models.Sticker{
					Name:    "124",
					AlbumID: "123123",
				}
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinSticker(context.Background(), "30002", &sticker, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "location",
			Action: func(c *cli.Context) error {
				location := models.Location{
					Longitude: decimal.NewFromFloat(23.345),
					Latitude:  decimal.NewFromFloat(54.234),
					Name:      "ShangHai",
					Address:   "ShangHai, China",
				}
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinLocation(context.Background(), "30002", &location, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "live",
			Action: func(c *cli.Context) error {
				live := models.Live{
					Width:    decimal.NewFromInt(960),
					Height:   decimal.NewFromInt(640),
					ThumbURL: "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2689960828,3272559041&fm=26&gp=0.jpg",
					URL:      "rtmp://live.hkstv.hk.lxdns.com/live/hks",
				}
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinLive(context.Background(), "30002", &live, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "contact",
			Action: func(c *cli.Context) error {
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinContact(context.Background(), "30002", "8be122b4-596f-4e4f-a307-978bed0ffb75", "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "buttongroup",
			Action: func(c *cli.Context) error {
				buttons := []*models.MsgButton{
					{
						Label:  "直播1",
						Color:  "#1267ab",
						Action: "https://www.fox.one",
					},
					{
						Label:  "直播2",
						Color:  "#1267ab",
						Action: "https://www.fox.one",
					},
				}
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinButtonGroup(context.Background(), "30002", buttons, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "appcard",
			Action: func(c *cli.Context) error {
				appcard := models.AppCard{
					IconURL:     "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=180811274,3179596559&fm=26&gp=0.jpg",
					Title:       "test 001",
					Description: "test description 001 ..........................",
					Action:      "https://www.fox.one",
				}
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinAppcard(context.Background(), "30002", &appcard, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "multi",
			Action: func(c *cli.Context) error {
				textMsg := models.Text{Text: "text msg 001"}
				postMsg := models.Post{Text: "post msg 001"}
				stickerMsg := models.Sticker{Name: "sticker msg", AlbumID: "sticker album id"}
				location := models.Location{Longitude: decimal.Zero, Latitude: decimal.Zero, Name: "SH", Address: "SH, China"}
				live := models.Live{Width: decimal.NewFromInt(960), Height: decimal.NewFromInt(640), ThumbURL: "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2689960828,3272559041&fm=26&gp=0.jpg", URL: "rtmp://live.hkstv.hk.lxdns.com/live/hks"}
				contact := models.Contact{UserID: "8be122b4-596f-4e4f-a307-978bed0ffb75"}
				buttongroup := models.ButtongGroup{
					{
						Label:  "直播1",
						Color:  "#1267ab",
						Action: "https://www.fox.one",
					},
					{
						Label:  "直播2",
						Color:  "#1267ab",
						Action: "https://www.fox.one",
					},
				}
				appcard := models.AppCard{
					IconURL:     "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=180811274,3179596559&fm=26&gp=0.jpg",
					Title:       "test 001",
					Description: "test description 001 ..........................",
					Action:      "https://www.fox.one",
				}

				multiMsg := []models.IMixinMsg{
					textMsg,
					postMsg,
					stickerMsg,
					location,
					live,
					contact,
					buttongroup,
					appcard,
				}

				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendMixinMultiMsg(context.Background(), "30002", multiMsg, "8be122b4-596f-4e4f-a307-978bed0ffb75")
			},
		},
		{
			Name: "sms",
			Action: func(c *cli.Context) error {
				//phone = phoneCode + phoneNumber
				return pigeon.New("http://47.103.148.221:8130", "eddb032c1ac447e7abb287d5a0a50ef5", "eddb032c1ac447e7abb287d5a0a50ef5").SendSMS(context.Background(), "20007", "hello pigeon sdk sms test 001", "8613386016339")
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
