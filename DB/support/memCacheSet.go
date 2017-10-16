package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	var err error
	var TTL = 10
	// Connect to our memcache instance
	mc := memcache.New("127.0.0.1:11211")
	keyWithNumber := "number"
	err = mc.Add(&memcache.Item{
		Key:        keyWithNumber,
		Value:      []byte("10"),
		Expiration: int32(TTL),
	})

	if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		checkErr(err)
	}

	afterIncrement, err := mc.Increment(keyWithNumber, uint64(42))
	checkErr(err)
	fmt.Println("After Increment by 42 ", keyWithNumber, "is", afterIncrement)

	afterDecrement, err := mc.Decrement(keyWithNumber, uint64(1))
	checkErr(err)
	fmt.Println("afterDecrement by 1 ", keyWithNumber, "is", afterDecrement)

	afterIncrement, err = mc.Increment("_ot_exist", uint64(1))
	fmt.Println("After Increment not existing record ", afterIncrement)
	if err == memcache.ErrCacheMiss {
		fmt.Println("Record not exist")
	} else if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		checkErr(err)
	} else {
		fmt.Println("After Decrement by 1 ", keyWithNumber, "is", afterDecrement)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
