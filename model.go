package main

// swagger:model Config
type Config struct {

	// Id of the config
	// in: string
	Id      string            `json:"id"`

	// Version of the config
	// in: string
	Version string            `json:"version"`

	// Entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"`
}

// swagger: model ConfigGroup
type ConfigGroup struct {

	// Id of the configGroup
	// in: string
	Id      string                        `json:"id"`

	// Configs of the configGroup
	// in map[string]map[string]*Config
	Configs map[string]map[string]*Config `json:"configs"`
}
