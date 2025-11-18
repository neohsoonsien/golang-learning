package redis

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func GetStudent() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       0,
	})

	json, err := json.Marshal(Student{Name: "Adam", Id: "C4321"})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("C4321", string(json), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("C4321").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
