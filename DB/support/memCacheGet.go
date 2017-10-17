package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// START1 OMIT
	mc := memcache.New("127.0.0.1:11211")
	keyOne := "key_one"
	keyTwo := "key_two"
	keyWithNumber := "number"

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
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
