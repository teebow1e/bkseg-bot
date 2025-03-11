package ctftime

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type CTFTimeClient struct {
	Client  *fasthttp.Client
	Timeout time.Duration
}

func NewCTFTimeClient(timeout time.Duration) *CTFTimeClient {
	return &CTFTimeClient{
		Client: &fasthttp.Client{
			Name:            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
			MaxConnsPerHost: 100,
			MaxConnDuration: 30,
		},
		Timeout: timeout,
	}
}

func (ctc *CTFTimeClient) CallAPI(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	err := ctc.Client.DoTimeout(req, resp, ctc.Timeout)

	if resp.Header.StatusCode() < 200 || resp.Header.StatusCode() >= 300 {
		bodyBytes := append([]byte{}, resp.Body()...)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.Header.StatusCode(), string(bodyBytes))
	}

	return append([]byte{}, resp.Body()...), err
}

func (ctc *CTFTimeClient) CallAndParseAPI(url string, target interface{}) error {
	data, err := ctc.CallAPI(url)
	if err != nil {
		return err
	}

	if target != nil {
		if err := json.Unmarshal(data, target); err != nil {
			return fmt.Errorf("error during unmarshalling data: %v", err)
		}
	}
	return nil
}
