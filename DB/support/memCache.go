package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// START1 OMIT
	var err error
	mc := memcache.New("127.0.0.1:11211")
	keyOne := "key_one"
	keyTwo := "key_two"
	keyWithNumber := "number"

	var TTL = 90000000
	err = mc.Set(&memcache.Item{Key: keyOne, Value: []byte("Banana")})
	err = mc.Set(&memcache.Item{Key: keyTwo, Value: []byte("Apple")})
	err = mc.Set(&memcache.Item{Key: keyWithNumber, Value: []byte("5"), Expiration: int32(TTL)})

	checkErr(err)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
