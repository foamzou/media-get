package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func WgetBinary(url string, downloadTo string, headers map[string]string) error {
	out, err := os.Create(downloadTo)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil) // nolint
	if err != nil {
		return err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("should return 2xx")
	}

	if len(resp.Header["Content-Type"]) > 0 && strings.Contains(resp.Header["Content-Type"][0], "text") {
		return errors.New("target is not a binary")
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
