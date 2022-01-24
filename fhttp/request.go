package fhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fox-one/pigeon-sdk-go/models"
	"github.com/go-resty/resty/v2"
)

var runOnce sync.Once
var restyClient *resty.Client

func Client() *resty.Client {
	runOnce.Do(func() {
		restyClient = resty.New().
			SetHeader("Content-Type", "application/json").
			SetHeader("Charset", "utf-8").
			SetTimeout(10 * time.Second).
			OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
				ctx := r.Context()

				if values := r.Header.Values(HeaderKeyRequestID); len(values) == 0 {
					r.Header.Set(HeaderKeyRequestID, RequestIdFromContext(ctx))
				}
				return nil
			}).
			OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
				expect := r.Request.Header.Get(HeaderKeyRequestID)
				got := r.Header().Get(HeaderKeyRequestID)
				if expect != "" && got != "" && expect != got {
					return fmt.Errorf("%s mismatched, expect %q but got %q", HeaderKeyRequestID, expect, got)
				}

				return nil
			})
	})

	return restyClient
}

func Request(ctx context.Context) *resty.Request {
	return Client().R().SetContext(ctx)
}

func RequestWithID(ctx context.Context, requestID string) *resty.Request {
	return Request(ctx).SetHeader(HeaderKeyRequestID, requestID)
}

func Execute(request *resty.Request, method, url string, body interface{}, resp interface{}) (int, models.Err) {
	fmt.Printf("url:%s\n", url)

	if body != nil {
		request = request.SetBody(body)
	}

	r, err := request.Execute(strings.ToUpper(method), url)
	if err != nil {
		return r.StatusCode(), models.NewErr(-1, err.Error())
	}

	return r.StatusCode(), ParseResponse(r, resp)
}

func ParseResponse(r *resty.Response, obj interface{}) models.Err {
	//fail
	if !r.IsSuccess() {
		return models.NewErrWithBytes(r.Body())
	}

	//success
	if obj != nil {
		e := json.Unmarshal(r.Body(), obj)
		if e != nil {
			fmt.Printf("parseResponse:%s", e.Error())
			return models.NewErr(-1, e.Error())
		}
		return nil
	}
	return nil
}

type DataWrapperResponse struct {
	FError
	Data *json.RawMessage `json:"data"`
}

type FError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *FError) Error() string {
	b, je := json.Marshal(e)
	if je != nil {
		return je.Error()
	}
	return string(b)
}
