package main

import (
	"time"
)

type Config struct {
	Token          string   `json:"token"`
	Prefix         string   `json:"prefix"`
	Owners         []string `json:"owners"`
	Statuses       []string `json:"statuses"`
	StatusInterval string   `json:"status_interval"`
}

type AFKEntry struct {
	Message string
	Set     time.Time
}
