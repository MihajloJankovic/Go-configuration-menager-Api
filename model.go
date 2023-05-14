package main

type Config struct {
	Id      string            `json:"id"`
	Version string            `json:"version"`
	Entries map[string]string `json:"entries"`
}

type ConfigGroup struct {
	Id      string                        `json:"id"`
	Configs map[string]map[string]*Config `json:"configs"`
}
