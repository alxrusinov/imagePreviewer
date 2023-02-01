package service

import (
	"bytes"
	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/stretchr/testify/require"
	"image"
	"image/jpeg"
	"testing"
)

var okRepo = repository.NewMock(true, make([]byte, 10))
var notOkRepo = repository.NewMock(false, struct{}{})
var badValueRepo = repository.NewMock(true, struct{}{})

var cropperService = NewCropperService(okRepo)
var notOkCropperService = NewCropperService(notOkRepo)
var badValueCropperService = NewCropperService(badValueRepo)

func TestCropperService_Fill(t *testing.T) {

	var imgBuf bytes.Buffer
	img := image.NewRGBA(image.Rect(0, 0, 200, 100))
	_ = jpeg.Encode(&imgBuf, img, nil)

	t.Run("decode error", func(t *testing.T) {
		buf := make([]byte, 100)
		params := NewCropperParams("https://foo.bar", 100, 200)

		val, err := cropperService.Fill(buf, params)

		require.Nil(t, val)
		require.ErrorIs(t, err, DecodeImageError)

	})

	t.Run("success filling", func(t *testing.T) {

		params := NewCropperParams("https://foo.bar", 100, 200)

		val, err := cropperService.Fill(imgBuf.Bytes(), params)

		require.NotNil(t, val)
		require.NoError(t, err)

	})
}

func TestCropperService_GetByUrl(t *testing.T) {
	t.Run("url exists", func(t *testing.T) {
		var rawUrl repository.Key = "https://foo.bar"

		val, exists := cropperService.GetByUrl(rawUrl)

		require.NotNil(t, val)
		require.True(t, exists)

	})

	t.Run("url not exists", func(t *testing.T) {
		var rawUrl repository.Key = "https://foo.bar"

		val, exists := notOkCropperService.GetByUrl(rawUrl)

		require.Nil(t, val)
		require.False(t, exists)

	})

	t.Run("wrong value", func(t *testing.T) {
		var rawUrl repository.Key = "https://foo.bar"

		val, exists := badValueCropperService.GetByUrl(rawUrl)

		require.Nil(t, val)
		require.True(t, exists)

	})
}

func TestCropperService_SaveToCache(t *testing.T) {
	t.Run("url exists", func(t *testing.T) {
		var rawUrl repository.Key = "https://foo.bar"
		value := make([]byte, 10)

		exists := cropperService.SaveToCache(rawUrl, value)

		require.True(t, exists)

	})

	t.Run("url not exists", func(t *testing.T) {
		var rawUrl repository.Key = "https://foo.bar"
		value := make([]byte, 10)

		exists := notOkCropperService.SaveToCache(rawUrl, value)

		require.False(t, exists)

	})
}

func TestNewCropperService(t *testing.T) {
	t.Run("create cropper service", func(t *testing.T) {

		cs := NewCropperService(okRepo)

		require.Equal(t, cs.repo, okRepo)
	})
}

func TestNewCropperParams(t *testing.T) {
	t.Run("create cropper service", func(t *testing.T) {
		addr := "http://foo.bar"
		width := 100
		height := 200
		cs := NewCropperParams(addr, width, height)
		expected := &CropperParams{Address: addr, Width: width, Height: height}

		require.EqualValues(t, expected, cs)
	})
}
