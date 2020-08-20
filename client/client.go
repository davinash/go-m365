package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

// LoginBaseURL represents the basic url used to acquire a token for the msgraph api
const LoginBaseURL string = "https://login.microsoftonline.com"

// BaseURL represents the URL used to perform all ms graph API-calls
const BaseURL string = "https://graph.microsoft.com"

type Auth struct {
	TenantId          string
	ApplicationId     string
	ClientSecurityKey string
}

type Client struct {
	lock  sync.Mutex
	token Token // the current token to be used
	auth  Auth
}

func NewClient(c Auth) (*Client, error) {
	client := Client{
		auth: c,
	}
	client.lock.Lock()
	defer client.lock.Unlock()

	return &client, client.refreshToken()
}

func (c *Client) refreshToken() error {
	if c.auth.TenantId == "" {
		return fmt.Errorf("TenantId is empty")
	}
	resource := fmt.Sprintf("/%v/oauth2/token", c.auth.TenantId)
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", c.auth.ApplicationId)
	data.Add("client_secret", c.auth.ClientSecurityKey)
	data.Add("resource", BaseURL)

	u, err := url.ParseRequestURI(LoginBaseURL)
	if err != nil {
		return fmt.Errorf("URL Parsing failed: %w", err)
	}
	u.Path = resource
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(data.Encode()))

	if err != nil {
		return fmt.Errorf("HTTP Request Error: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	var newToken Token
	_, err = c.performRequest(req, &newToken)
	if err != nil {
		return fmt.Errorf("error on getting msgraph Token: %w", err)
	}
	c.token = newToken

	return nil
}

func (c *Client) performRequest(req *http.Request, v interface{}) (int, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   100 * time.Second,
			ResponseHeaderTimeout: 100 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
			IdleConnTimeout:       1000 * time.Second,
		},
		Timeout: time.Second * 500,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		if resp == nil {
			return 400, fmt.Errorf("HTTP response error: %w ", err)
		}
		return resp.StatusCode, fmt.Errorf("HTTP response error: %w ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp.StatusCode, fmt.Errorf("StatusCode is not OK: %v. Body: %v ", resp.StatusCode, string(body))
	} else if resp.StatusCode == 204 {
		return resp.StatusCode, err
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("HTTP response read error: %w of http.Request: %v", err, req.URL)
	}
	return resp.StatusCode, json.Unmarshal(body, &v)
}
