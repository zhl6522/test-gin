package session

type Session struct {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string)	error
	Save() error
}

