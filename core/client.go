package core

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bububa/tiktok-business/core/internal/debug"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	client  *http.Client
	appID   string
	secret  string
	preReqs []PreRequest
	debug   bool
	sandbox bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
		client: defaultHttpClient(),
	}
}

func (c *SDKClient) AppID() string {
	return c.appID
}

func (c *SDKClient) Secret() string {
	return c.secret
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

func (c *SDKClient) WithPreRequests(reqs ...PreRequest) {
	c.preReqs = reqs
}

func (c *SDKClient) AddPreRequests(reqs ...PreRequest) {
	c.preReqs = append(c.preReqs, reqs...)
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		appID:   c.appID,
		secret:  c.secret,
		debug:   c.debug,
		sandbox: c.sandbox,
		client:  c.client,
		preReqs: c.preReqs,
	}
}

// Post post api
func (c *SDKClient) Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	base := BASE_URL
	if c.sandbox {
		base = SANDBOX_URL
	}
	return c.post(ctx, base, gw, req, resp, accessToken)
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	base := BASE_URL
	if c.sandbox {
		base = SANDBOX_URL
	}
	return c.get(ctx, base, gw, req, resp, accessToken)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(ctx context.Context, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	base := BASE_URL
	if c.sandbox {
		base = SANDBOX_URL
	}
	return c.upload(ctx, base, gw, req, resp, accessToken)
}

func (c *SDKClient) post(ctx context.Context, base string, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := util.StringsJoin(base, gw)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.fetch(httpReq, resp)
}

func (c *SDKClient) get(ctx context.Context, base string, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	reqUrl := util.StringsJoin(base, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	return c.fetch(httpReq, resp)
}

func (c *SDKClient) upload(ctx context.Context, base string, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	mw := multipart.NewWriter(buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			builder := util.NewStringsBuilder()
			builder.WriteString("@")
			builder.WriteString(v.Value)
			mp[v.Key] = builder.String()
			util.ReleaseStringsBuilder(builder)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	reqUrl := util.StringsJoin(base, gw)
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}

	return c.fetch(httpReq, resp)
}

type PreRequest func(httpReq *http.Request) error

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	if len(c.preReqs) > 0 {
		for _, req := range c.preReqs {
			if err := req(httpReq); err != nil {
				return err
			}
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	body, err := debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if httpResp.ContentLength <= 0 {
		httpResp.ContentLength = int64(len(body))
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return errors.Join(err, model.BaseResponse{
			Code:    httpResp.StatusCode,
			Message: string(body),
		})
	}
	// if throttle := httpResp.Header.Get("X-Tt-Ads-Throttle"); throttle != "" {
	//
	// }
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return errors.Join(err, model.BaseResponse{
			Code:    httpResp.StatusCode,
			Message: string(body),
		})
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
