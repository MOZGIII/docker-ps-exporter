package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MOZGIII/docker-ps-exporter/internal/collector"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getEnv(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}
	return "", fmt.Errorf("environment variable %s is not set", key)
}

func getAddr() (string, error) {
	addr, err := getEnv("ADDR")
	if err != nil {
		return "", fmt.Errorf("reading listen address: %s", err)
	}
	return addr, nil
}

func boot() error {
	addr, err := getAddr()
	if err != nil {
		return err
	}

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	coll := collector.DockerContainers{Client: dockerClient}
	if err := prometheus.Register(&coll); err != nil {
		return err
	}

	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addr, nil)
}

func main() {
	if err := boot(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
