package lru

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"sync"
	"testing"

	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("aaa", 100)
		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)
		c.Set("bbb", 200)
		c.Set("ccc", 300)
		c.Set("ddd", 400)

		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		c.Get("bbb")
		c.Get("ddd")
		wasInCache := c.Set("bbb", 222)
		require.True(t, wasInCache)

		wasInCache = c.Set("eee", 500)
		require.False(t, wasInCache)
		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clearing cache", func(t *testing.T) {
		c := NewCache(2)

		c.Set("aaa", 100)
		c.Set("bbb", 200)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		c.Clear()

		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(repository.Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			random, _ := rand.Int(rand.Reader, big.NewInt(1_000_000))
			c.Get(repository.Key(strconv.FormatInt(random.Int64(), 10)))
		}
	}()

	wg.Wait()
}
