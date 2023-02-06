//go:build integration
// +build integration

package integration_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type ApiSuit struct {
	suite.Suite
	httpClient http.Client
}

func TestApiSuit(t *testing.T) {
	suite.Run(t, new(ApiSuit))
}

func (as *ApiSuit) SetupTest() {

}

func (as *ApiSuit) TestFill_Success() {
	targetURL, _ := url.Parse("http://localhost:3000/fill/200/200/nginxcompose/images/pig_1.jpg")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    targetURL,
	}

	_, err := as.httpClient.Do(req)

	as.Nil(err)
}

func (as *ApiSuit) TestFill_Cached() {
	targetURL, _ := url.Parse("http://localhost:3000/fill/200/200/nginxcompose/images/pig_1.jpg")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    targetURL,
	}
	res1, err := as.httpClient.Do(req)

	as.Nil(err)

	content1, _ := io.ReadAll(res1.Body)
	defer res1.Body.Close()
	as.NotNil(content1)

	res2, err := as.httpClient.Do(req)
	content2, _ := io.ReadAll(res2.Body)
	defer res2.Body.Close()
	as.Nil(err)
	as.NotNil(content2)
	as.EqualValues(content1, content2)
}

func (as *ApiSuit) TestFill_ServerNotExists() {
	notExistsURL, _ := url.Parse("http://localhost:3000/fill/200/200/foo/images/pig_3.jpg")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    notExistsURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusBadGateway, res.StatusCode)
	as.Nil(err)
}

func (as *ApiSuit) TestFill_ImageNotFound() {
	wrongURL, _ := url.Parse("http://localhost:3000/fill/200/200/nginxcompose/images/pig_3.jpg")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    wrongURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusBadGateway, res.StatusCode)
	as.Nil(err)

}

func (as *ApiSuit) TestFill_WrongFileType() {
	wrongURL, _ := url.Parse("http://localhost:3000/fill/200/200/nginxcompose/images/file.txt")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    wrongURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusBadGateway, res.StatusCode)
	as.Nil(err)
}

func (as *ApiSuit) TestFill_ServerError() {
	errURL, _ := url.Parse("http://localhost:3000/fill/200/200/err")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    errURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusBadGateway, res.StatusCode)
	as.Nil(err)
}

func (as *ApiSuit) TestFill_LargeParams() {
	targetURL, _ := url.Parse("http://localhost:3000/fill/5000/5000/nginxcompose/images/pig_1.jpg")
	req := &http.Request{
		Method: http.MethodGet,
		URL:    targetURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusOK, res.StatusCode)
	as.Nil(err)

	content, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	as.NotNil(content)

	contentLength := res.ContentLength

	as.Equal(contentLength, int64(len(content)))

}

func (as *ApiSuit) TestFill_WrongParams() {
	wrongWidthURL, _ := url.Parse("http://localhost:3000/fill/abc/200/nginxcompose/images/pig_1.jpg")
	wrongHeightURL, _ := url.Parse("http://localhost:3000/fill/200/abc/nginxcompose/images/pig_1.jpg")

	req := &http.Request{
		Method: http.MethodGet,
		URL:    wrongWidthURL,
	}
	res, err := as.httpClient.Do(req)

	as.Equal(http.StatusBadRequest, res.StatusCode)
	as.Nil(err)

	req.URL = wrongHeightURL

	res, err = as.httpClient.Do(req)

	as.Equal(http.StatusBadRequest, res.StatusCode)
	as.Nil(err)
}

func (as *ApiSuit) TearDownTestTest() {
	fmt.Println("END")
}
