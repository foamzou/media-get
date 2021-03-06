package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const TimeOut = 3 * time.Second
const TcpConnectTimeout = 1500 * time.Millisecond

func HttpGet(url string, headers map[string]string) (string, error) {
	client := createClient()
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

func GetLocation(url string, headers map[string]string) (string, error) {
	client := createClient()
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

func GetCookie(url string, headers map[string]string) (string, error) {
	client := createClient()
	resp, err := client.R().
		SetHeaders(headers).
		Get(url)

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

func PostForm(url string, data, headers map[string]string) (string, error) {
	client := createClient()
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

func createClient() resty.Client {
	return *(resty.New().SetTimeout(TimeOut).SetTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout: TcpConnectTimeout,
		}).DialContext,
	}))
}
