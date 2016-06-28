package redisstore

import (
	nSessions "github.com/acoshift/negroni-sessions"
	"github.com/boj/redistore"
	"github.com/garyburd/redigo/redis"
	gSessions "github.com/gorilla/sessions"
)

//New returns a new Redis store
func New(size int, network, address, password string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

// NewWithDB returns a new Redis store
func NewWithDB(size int, network, address, password, DB string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStoreWithDB(size, network, address, password, DB, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

// NewWithPool returns a new Redis store
func NewWithPool(pool *redis.Pool, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStoreWithPool(pool, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

type rediStore struct {
	*redistore.RediStore
}

func (c *rediStore) Options(options nSessions.Options) {
	c.RediStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}
