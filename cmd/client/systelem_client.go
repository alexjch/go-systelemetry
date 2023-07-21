package main

import "github.com/alexjch/go-systelemetry/internal/pkg/client"

func main() {
	client := client.Client{}
	client.Request()
}
