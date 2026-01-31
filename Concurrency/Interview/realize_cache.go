package main

import "errors"

type KVdatabase interface {
	Get(key string) (string, error)
	GetKeys() ([]string, error)
}

type Redis struct {
	data map[string]string
}

var ErrorNothingFound = errors.New("nothing found")

func (d *Redis) Get(key string) (string, error) {
	if val, ok := d.data[key]; ok {
		return val, nil
	}
	return "", ErrorNothingFound
}

func (d *Redis) GetKeys() ([]string, error) {
	var res []string
	for key := range d.data {
		res = append(res, key)
	}
	if len(res) == 0 {
		return res, ErrorNothingFound
	}
	return res, nil
}

type RedisCache struct {
	redis KVdatabase
	data  map[string]string
}

func NewRedisCache(redis KVdatabase) *RedisCache {
	return &RedisCache{
		data:  make(map[string]string),
		redis: redis,
	}
}

func (r *RedisCache) Get(key string) (string, error) {
	if val, ok := r.data[key]; ok {
		return val, nil
	} else if val, err := r.redis.Get(key); err == nil {
		r.data[key] = val
		return val, nil
	}
	return "", ErrorNothingFound
}

func (r *RedisCache) GetKeys() ([]string, error) {
	var res []string
	for key := range r.data {
		res = append(res, key)
	}
	keys, err := r.redis.GetKeys()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		v, err := r.Get(key)
		if err != nil {
			r.data[key] = v
			res = append(res, v)
		}
	}

	if len(res) == 0 {
		return res, ErrorNothingFound
	}

	return res, nil
}

func main() {

}
