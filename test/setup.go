package test

import "github.com/guatom999/backend-challenge/config"

func NewTestConfig() *config.Config {
	cfg := config.GetTestConfig()
	return cfg
}
