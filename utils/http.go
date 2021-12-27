package utils

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

func FetchHtml(url string, headers map[string]string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		Get(url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New("http code should be 200")
	}
	return resp.String(), nil
}

func PostForm(url string, data, headers map[string]string) (string, error) {
	client := resty.New()
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
