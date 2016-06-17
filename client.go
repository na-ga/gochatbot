package gochatbot

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	//
	defaultEndPoint = "https://chatbot-api.userlocal.jp/api"

	//
	defaultAPIKey = "sample"
)

type (

	// ChatbotConfig is the configuration for initializing chatbot client
	ChatbotConfig struct {
		URL     string
		APIKey  string
		TimeOut time.Duration
	}

	// ChatbotClient is the client to access chatbot API
	ChatbotClient struct {
		ChatbotConfig
	}
)

// NewClient is creating chatbot client and returns it.
func NewClient() (*ChatbotClient, error) {

	url := os.Getenv("CHATBOT_ENDPOINT_URL")
	if url == "" {
		url = defaultEndPoint
	}

	apiKey := os.Getenv("CHATBOT_API_KEY")
	if apiKey == "" {
		apiKey = defaultAPIKey
	}

	return &ChatbotClient{ChatbotConfig{
		URL:     url,
		APIKey:  apiKey,
		TimeOut: 5 * time.Second,
	}}, nil
}

// NewClientWithConfig is creating chatbot client with config and returns it.
func NewClientWithConfig(config ChatbotConfig) (*ChatbotClient, error) {

	if config.URL == "" {
		config.URL = defaultEndPoint
	}

	if config.APIKey == "" {
		config.APIKey = defaultAPIKey
	}

	if config.TimeOut == 0 {
		config.TimeOut = 5 * time.Second
	}

	return &ChatbotClient{config}, nil
}

//
func (c *ChatbotClient) do(req *gorequest.SuperAgent, data interface{}) error {
	res, body, errs := req.End()
	if errs != nil {
		return &ErrorResponse{Message: errs[0].Error()}
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return &ErrorResponse{Code: res.StatusCode, Message: body}
	}
	if data == nil {
		return nil
	}
	return json.NewDecoder(strings.NewReader(body)).Decode(data)
}

//
func (c *ChatbotClient) get(path string) *gorequest.SuperAgent {
	return c.initRequest(gorequest.New().Get(c.URL + path))
}

//
func (c *ChatbotClient) post(path string) *gorequest.SuperAgent {
	return c.initRequest(gorequest.New().Post(c.URL + path))
}

//
func (c *ChatbotClient) put(path string) *gorequest.SuperAgent {
	return c.initRequest(gorequest.New().Put(c.URL + path))
}

//
func (c *ChatbotClient) delete(path string) *gorequest.SuperAgent {
	return c.initRequest(gorequest.New().Delete(c.URL + path))
}

//
func (c *ChatbotClient) initRequest(req *gorequest.SuperAgent) *gorequest.SuperAgent {
	return req.Timeout(c.TimeOut)
}
