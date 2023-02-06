package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrFileType    = NewError("file is not image")
	ErrReadingFile = NewError("bad reading file")
	ErrClient      = NewError("unable to get image from server")
	ErrServer      = NewError("server error")
)

type Error struct {
	msg string
}

func NewError(msg string) *Error {
	return &Error{msg: msg}
}

func (err *Error) Error() string {
	return err.msg
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}}
}

func (cl *Client) GetWithHeaders(url *url.URL, header http.Header) ([]byte, error) {
	req := &http.Request{URL: url, Method: http.MethodGet, Header: header}

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrClient, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, ErrServer
	}

	contentType := resp.Header.Get("content-type")

	if !isJpegFileType(contentType) {
		return nil, ErrFileType
	}

	wb := make([]byte, 0, resp.ContentLength)
	buf := bytes.NewBuffer(wb)

	_, err = io.Copy(buf, resp.Body)

	if err != nil {
		return nil, ErrReadingFile
	}

	return buf.Bytes(), nil
}
