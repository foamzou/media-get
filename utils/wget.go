package utils

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/foamzou/audio-get/logger"
)

type WgetBinaryProgressReader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *WgetBinaryProgressReader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	logger.Infof("\rdownloading percent: %.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}

type TCPProgressReader struct {
	net.Conn
	Init    bool
	Current int64
	Total   int64
}

func (r *TCPProgressReader) Read(p []byte) (n int, err error) {
	n, err = r.Conn.Read(p)
	if !r.Init {
		//fmt.Println(string(p[:]))
		r.Init = true
	}

	// TODO: parse Total from first package
	//logger.Infof("\rdownloading percent: %.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}

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
		return errors.New("should return 2xx, but got " + resp.Status)
	}

	if len(resp.Header["Content-Type"]) > 0 && strings.Contains(resp.Header["Content-Type"][0], "text") {
		return errors.New("target is not a binary")
	}

	reader := &WgetBinaryProgressReader{
		Reader: resp.Body,
		Total:  resp.ContentLength,
	}
	_, err = io.Copy(out, reader)
	logger.Info("\n")
	if err != nil {
		return err
	}
	return nil
}

func DownloadBinaryWithTCP(inputUrl string, downloadTo string, headers map[string]string) error {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", u.Host+":80")
	if err != nil {
		return err
	}

	u.Path = strings.ReplaceAll(u.Path, "+", "%2B")
	rt := fmt.Sprintf("GET %v HTTP/1.1\r\n", u.Path)
	rt += fmt.Sprintf("Host: %v\r\n", u.Host)
	for k, v := range headers {
		rt += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	rt += "Connection: close\r\n"
	rt += "\r\n"

	_, err = conn.Write([]byte(rt))
	if err != nil {
		return err
	}
	reader := &TCPProgressReader{
		Conn: conn,
	}
	resp, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = os.WriteFile(downloadTo, resp, 0600)
	if err != nil {
		return err
	}

	_ = conn.Close()

	return nil
}
