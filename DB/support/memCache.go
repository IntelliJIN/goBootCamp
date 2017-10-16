package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	var err error
	// Connect to our memcache instance
	mc := memcache.New("127.0.0.1:11211")
	keyOne := "key_one"
	keyTwo := "key_two"
	keyWithNumber := "number"
	// Set some values
	var TTL = 10
	err = mc.Set(&memcache.Item{Key: keyOne, Value: []byte("Banana")})
	err = mc.Set(&memcache.Item{Key: keyTwo, Value: []byte("Apple")})
	err = mc.Set(&memcache.Item{Key: keyWithNumber, Value: []byte("5"), Expiration: int32(TTL)})

	checkErr(err)

	val, err := mc.Get(keyOne)

	if err == memcache.ErrCacheMiss {
		fmt.Println("Key: ", keyOne, " is missed")
		return
	} else if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Key: %s , Value: %s\n", keyOne, val.Value)

	it, err := mc.GetMulti([]string{keyOne, keyTwo, keyWithNumber})

	checkErr(err)

	for k, v := range it {
		fmt.Printf("Key: %s , Value: %s\n", k, v.Value)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
