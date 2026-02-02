package main

import (
	"time"

	"github.com/amirhnajafiz/localsight/internal/collector"
	"github.com/amirhnajafiz/localsight/internal/configs"
	"github.com/amirhnajafiz/localsight/internal/logr"
	"github.com/amirhnajafiz/localsight/internal/metrics"

	"go.uber.org/zap"
)

func main() {
	// load the configuration from the environment variables
	conf, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	// initialize a zap logger
	logger := logr.NewZapLogger(conf.Debug, conf.JSONLog)

	// create a new metrics instance
	mtx, err := metrics.NewMetrics()
	if err != nil {
		logger.Fatal("failed to create metrics instance", zap.Error(err))
	}

	// start the metrics server on port 8080
	metrics.StartMetricsServer(logger.Named("metrics-server"), conf.Port)

	// create a new collector instance with the metrics
	col := &collector.Collector{
		CertFile: conf.CertFile,
		KeyFile:  conf.KeyFile,
		EndPoint: conf.K8SLocalAPI,
		Logr:     logger.Named("collector"),
		Metrics:  mtx,
		Interval: time.Duration(conf.Interval) * time.Second,
	}

	// start the collector to fetch and update metrics
	if err := col.Start(); err != nil {
		logger.Fatal("failed to start collector", zap.Error(err))
	}
}
