package main

type Config struct {
	Id      string             `json:"id"`
	entries map[string]*string `json:"entries"`
}
