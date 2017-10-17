package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// START1 OMIT
	mc := memcache.New("127.0.0.1:11211")
	keyWithNumber := "number"
	err := mc.Add(&memcache.Item{
		Key:        keyWithNumber,
		Value:      []byte("10"),
		Expiration: int32(10),
	})

	if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		checkErr(err)
	}
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
