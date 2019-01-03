package main

type Config struct {
	Token          string   `json:"token"`
	Prefix         string   `json:"prefix"`
	Owners         []string `json:"owners"`
	Statuses       []string `json:"statuses"`
	StatusInterval string   `json:"status_interval"`
}
