package main

import (
	"net/http"

	"github.com/MOZGIII/docker-ps-exporter/internal/collector"
	"github.com/alecthomas/kingpin"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = kingpin.Flag("web.listen-address", "Address on which to expose metrics.").Default(":9491").String()
	metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
)

func boot() error {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	coll := collector.DockerContainers{Client: dockerClient}
	if err := prometheus.Register(&coll); err != nil {
		return err
	}

	http.Handle(*metricsPath, promhttp.Handler())
	return http.ListenAndServe(*listenAddress, nil)
}

func main() {
	kingpin.Parse()
	kingpin.FatalIfError(boot(), "")
}
