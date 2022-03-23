package cache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	MemcacheServers = "memcache_servers"
)

var (
	cache             *memcache.Client
	memcacheServers   = os.Getenv(MemcacheServers)
	DefaultExpiration = int32(time.Now().Add(time.Minute * 2).Unix())
	ErrCacheMiss      = memcache.ErrCacheMiss
)

func init() {
	r, err := regexp.Compile(" *, *")
	if err != nil {
		panic(err)
	}

	cacheServers := r.Split(memcacheServers, -1)
	cache = memcache.New(cacheServers...)
	err = Ping()
	if err != nil {
		errStr := fmt.Sprintf("failed to connect with Memcached server, %s", cacheServers)
		panic(errStr)
	}

	err = cache.FlushAll()
	if err != nil {
		panic(err)
	}

	log.Println("Memcached successfully configured")
}

func Set(key string, value []byte, expiration int32) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}

	return cache.Set(item)
}

func Replace(key string, value []byte, expiration int32) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}

	return cache.Replace(item)
}

func Get(key string) ([]byte, error) {
	item, err := cache.Get(key)
	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func GetMulti(keys []string) (map[string][]byte, error) {
	keyValues := make(map[string][]byte)
	for _, key := range keys {
		item, err := cache.Get(key)
		if err != nil {
			return nil, err
		}

		keyValues[key] = item.Value
	}
	return keyValues, nil
}

func Delete(key string) error {
	return cache.Delete(key)
}

func DeleteAll() error {
	return cache.DeleteAll()
}

func FlushAll() error {
	return cache.FlushAll()
}

func Ping() error {
	return cache.Ping()
}
