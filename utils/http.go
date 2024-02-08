package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/foamzou/audio-get/logger"
	"github.com/go-resty/resty/v2"
)

const TimeOut = 3 * time.Second
const TcpConnectTimeout = 1500 * time.Millisecond
const HttpRetryCnt = 3

func HttpGet(source, url string, headers map[string]string) (string, error) {
	client := createClient(source, true)
	resp, err := client.R().
		SetHeaders(headers).
		Get(url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http code should be 200, got %d", resp.StatusCode())
	}
	return resp.String(), nil
}

func HttpHead(source, url string, headers map[string]string) error {
	client := createClient(source, true)
	resp, err := client.R().
		SetHeaders(headers).
		Head(url)

	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("http code should be 2xx, got %d", resp.StatusCode())
	}
	return nil
}

func GetLocation(source, url string, headers map[string]string) (string, error) {
	client := createClient(source, false)
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	// Should ignore the error since the redirect policy would return error when found redirect
	resp, _ := client.R().
		SetHeaders(headers).
		Get(url)

	if resp.StatusCode() < 300 || resp.StatusCode() > 399 {
		return "", fmt.Errorf("should redirect")
	}

	location := resp.RawResponse.Header.Get("Location")
	if location == "" {
		return "", fmt.Errorf("location is empty")
	}

	return location, nil
}

func GetCookie(source, url string, headers map[string]string, isHead bool) (string, error) {
	client := createClient(source, true)
	var resp *resty.Response
	var err error
	if isHead {
		resp, err = client.R().
			SetHeaders(headers).
			Head(url)
	} else {
		resp, err = client.R().
			SetHeaders(headers).
			Get(url)
	}

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http code should be 200, got %d", resp.StatusCode())
	}
	cookie := resp.RawResponse.Header.Get("Set-Cookie")
	if cookie == "" {
		return "", fmt.Errorf("cookie is empty")
	}

	return cookie, nil
}

func PostForm(source, url string, data, headers map[string]string) (string, error) {
	client := createClient(source, true)
	resp, err := client.R().
		SetHeaders(headers).
		SetFormData(data).
		Post(url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New("http code should be 200")
	}
	return resp.String(), nil
}

func createClient(source string, retry bool) resty.Client {
	// match source from proxy config
	proxyAddr := getProxyConfig(source)
	restCli := *(resty.New().SetTimeout(TimeOut).SetTransport(&http.Transport{

		DialContext: (&net.Dialer{
			Timeout: TcpConnectTimeout,
		}).DialContext,
	}))
	if retry {
		restCli.SetRetryCount(HttpRetryCnt)
	}
	//restCli.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// match source from proxy config
	if proxyAddr != "" {
		restCli.SetProxy(proxyAddr)
	}

	restCli.SetLogger(logger.MyLogger{})
	return restCli
}

func getProxyConfig(source string) string {
	//return "http://127.0.0.1:8888"
	mediaGetConfig := GetConfig()
	if mediaGetConfig == nil {
		logger.Debug("config not found")
		return ""
	}

	value, ok := mediaGetConfig.Proxy[source]
	if ok {
		return value
	}
	return ""
}
