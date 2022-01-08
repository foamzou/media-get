package utils

import (
	"io"
	"net/http"
	"os"
)

func Wget(url string, downloadTo string, headers map[string]string) error {
	out, err := os.Create(downloadTo)
	if err != nil {
		return err
	}
	defer out.Close()

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

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
