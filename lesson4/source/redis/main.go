package main

import (
	"context"
	"encoding/json"
	"fmt"
	"iitu/db/redis/model"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("connected")
	err := rdb.Set(ctx, "key1", "value233", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("key set")
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	user := model.User{
		ID:        15,
		FirstName: "Joe",
		LastName:  "Biden",
	}
	encoded, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	rdb.Set(ctx, strconv.Itoa(15), string(encoded), 2*time.Second)

	val, err = rdb.Get(ctx, "15").Result()
	if err != nil {
		panic(err)
	}
	var userRead model.User
	json.Unmarshal([]byte(val), &userRead)
	fmt.Println(15, userRead)
}
