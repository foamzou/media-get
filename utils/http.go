package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

const TimeOut = 3 * time.Second
const TcpConnectTimeout = 1500 * time.Millisecond
const HttpRetryCnt = 3
const ConfigFile = "/$HOME/.media-get.json"

func HttpGet(source, url string, headers map[string]string) (string, error) {
	client := createClient(source)
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

func GetLocation(source, url string, headers map[string]string) (string, error) {
	client := createClient(source)
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
	client := createClient(source)
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
	client := createClient(source)
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

func createClient(source string) resty.Client {
	// match source from proxy config
	proxyAddr := getProxyConfig(source)
	restCli := *(resty.New().SetTimeout(TimeOut).SetRetryCount(HttpRetryCnt).SetTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout: TcpConnectTimeout,
		}).DialContext,
	}))
	// match source from proxy config
	if proxyAddr != "" {
		restCli.SetProxy(proxyAddr)
	}
	return restCli
}

func getProxyConfig(source string) string {
	f, err := os.Open(ConfigFile)
	if err != nil {
		if err != os.ErrNotExist {
			fmt.Println("open file err = ", err)
		}
		return ""
	}

	defer f.Close()

	proxyInfos := make(map[string]string)
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&proxyInfos)
	if err != nil {
		fmt.Printf("json decode has error:%v\n", err)
		return ""
	}
	value, ok := proxyInfos[source]
	if ok {
		return value
	}
	return ""
}
