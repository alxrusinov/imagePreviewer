package repository

type LRU struct{}

func NewLRU() *LRU {
	return &LRU{}
}

func (lru *LRU) Save() error {
	return nil
}

func (lru *LRU) Get() error {
	return nil
}

func Foo(r *Repo) {}
