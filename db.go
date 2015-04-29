package main

import (
	"gopkg.in/redis.v2"
)

func NewClient() *redis.Client {
	client := redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
