package repository

type mock struct {
	ok  bool
	val interface{}
}

func NewMock(ok bool, val interface{}) Repo {
	return &mock{ok: ok, val: val}
}

func (m mock) Set(key Key, value interface{}) bool {
	_ = key
	_ = value

	return m.ok
}

func (m mock) Get(key Key) (interface{}, bool) {
	_ = key

	if !m.ok {
		return nil, false
	}

	if result, ok := m.val.([]byte); ok {
		return result, ok
	}

	return nil, true
}

func (m mock) Clear() {}
