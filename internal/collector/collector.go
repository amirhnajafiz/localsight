package collector

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/localsight/internal/metrics"
	"github.com/amirhnajafiz/localsight/pkg/fetch"
	"github.com/amirhnajafiz/localsight/pkg/types"

	"go.uber.org/zap"
)

// Collector is responsible for collecting storage usage metrics from the kubelet summary endpoint
// and updating the provided metrics instance with the collected data.
type Collector struct {
	Logr     *zap.Logger
	Metrics  *metrics.Metrics
	Interval time.Duration
}

// Start initiates the process of fetching storage usage metrics from the kubelet summary endpoint
// and updates the provided metrics instance with the data.
func (c *Collector) Start(endpoint string) error {
	// build the HTTP request to the kubelet summary endpoint
	req, err := buildHTTPRequest(endpoint)
	if err != nil {
		return fmt.Errorf("failed to build HTTP request: %w", err)
	}

	c.Logr.Info(
		"starting kubelet summary collector",
		zap.String("endpoint", req.URL.String()),
		zap.Duration("interval", c.Interval),
	)

	for {
		// wait for the specified interval before fetching metrics
		time.Sleep(c.Interval)
		c.Logr.Debug("fetching kubelet summary for storage usage metrics")

		// perform the HTTP GET request
		resp, err := fetch.GET(req)
		if err != nil {
			c.Logr.Error("failed to fetch kubelet summary", zap.Error(err))
			continue
		}

		// decode the JSON response into a summary structure
		var summary types.Summary
		if err := fetch.JSON(resp, &summary); err != nil {
			c.Logr.Error("failed to decode kubelet summary JSON", zap.Error(err))
			continue
		}

		// process the summary data and update the metrics
		for _, pod := range summary.Pods {
			c.setPodStorageUsage(pod, summary.Node.NodeName)
			c.setVolumeStorageUsage(pod, summary.Node.NodeName)
			c.setContainerStorageUsage(pod, summary.Node.NodeName)
		}

		c.Logr.Debug("successfully updated storage usage metrics", zap.String("node", summary.Node.NodeName))
	}
}

// setPodStorageUsage sets the ephemeral storage usage for a pod in the provided metrics instance.
func (c *Collector) setPodStorageUsage(pod types.PodSummary, nodeName string) {
	// set the ephemeral storage usage for the pod
	c.Metrics.SetEphemeralStorageValues(
		pod.PodRef.Name,
		pod.PodRef.Namespace,
		nodeName,
		float64(pod.EphemeralStorage.UsedBytes),
		float64(pod.EphemeralStorage.AvailableBytes),
		float64(pod.EphemeralStorage.CapacityBytes),
	)

	// set the ephemeral storage inodes for the pod
	c.Metrics.SetEphemeralStorageInodes(
		pod.PodRef.Name,
		pod.PodRef.Namespace,
		nodeName,
		float64(pod.EphemeralStorage.InodesUsed),
		float64(pod.EphemeralStorage.InodesFree),
		float64(pod.EphemeralStorage.Inodes),
	)
}

// setVolumeStorageUsage sets the volume usage for a volume in the provided metrics instance.
func (c *Collector) setVolumeStorageUsage(pod types.PodSummary, nodeName string) {
	for _, volume := range pod.Volume {
		c.Metrics.SetPodVolumeValues(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			volume.Name,
			float64(volume.UsedBytes),
			float64(volume.AvailableBytes),
			float64(volume.CapacityBytes),
		)

		c.Metrics.SetPodVolumeInodes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			volume.Name,
			float64(volume.InodesUsed),
			float64(volume.InodesFree),
			float64(volume.Inodes),
		)
	}
}

// setContainerStorageUsage sets the storage usage for each container in a pod in the provided metrics instance.
func (c *Collector) setContainerStorageUsage(pod types.PodSummary, nodeName string) {
	for _, container := range pod.Containers {
		// set the memory usage for the container
		c.Metrics.SetContainerMemoryValues(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			float64(container.Memory.UsageBytes),
			float64(container.Memory.AvailableBytes),
			float64(container.Memory.CapacityBytes),
		)

		// set the root filesystem usage for the container
		c.Metrics.SetContainerRootfsValues(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			float64(container.Rootfs.UsedBytes),
			float64(container.Rootfs.AvailableBytes),
			float64(container.Rootfs.CapacityBytes),
		)
		c.Metrics.SetContainerRootfsInodes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			float64(container.Rootfs.InodesUsed),
			float64(container.Rootfs.InodesFree),
			float64(container.Rootfs.Inodes),
		)

		// set the logs usage for the container
		c.Metrics.SetContainerLogsValues(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			float64(container.Logs.UsedBytes),
			float64(container.Logs.AvailableBytes),
			float64(container.Logs.CapacityBytes),
		)
		c.Metrics.SetContainerLogsInodes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			nodeName,
			container.Name,
			float64(container.Logs.InodesUsed),
			float64(container.Logs.InodesFree),
			float64(container.Logs.Inodes),
		)
	}
}
