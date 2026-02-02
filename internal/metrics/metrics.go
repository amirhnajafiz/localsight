package metrics

import "github.com/prometheus/client_golang/prometheus"

// constant values for namespace and subsystem
const (
	NS                 = "ls_ex"
	SSEphemeralStorage = "ephemeral_storage"
	SSContainerMemory  = "container_memory"
	SSContainerRootFS  = "container_rootfs"
	SSContainerLogs    = "container_logs"
	SSPodVolume        = "pod_volume"
)

// Metrics holds the Prometheus metrics for the exporter.
type Metrics struct {
	// API Metrics
	apiStatus  *prometheus.GaugeVec
	apiLatency *prometheus.GaugeVec

	// Ephemeral Storage
	ephemeralStorageAvailableBytes *prometheus.GaugeVec
	ephemeralStorageCapacityBytes  *prometheus.GaugeVec
	ephemeralStorageUsageBytes     *prometheus.GaugeVec
	ephemeralStorageInodes         *prometheus.GaugeVec
	ephemeralStorageInodesFree     *prometheus.GaugeVec
	ephemeralStorageInodesUsed     *prometheus.GaugeVec

	// Container Memory
	containerMemoryAvailableBytes *prometheus.GaugeVec
	containerMemoryCapacityBytes  *prometheus.GaugeVec
	containerMemoryUsageBytes     *prometheus.GaugeVec

	// Container RootFS
	containerRootfsAvailableBytes *prometheus.GaugeVec
	containerRootfsCapacityBytes  *prometheus.GaugeVec
	containerRootfsUsageBytes     *prometheus.GaugeVec
	containerRootfsInodes         *prometheus.GaugeVec
	containerRootfsInodesFree     *prometheus.GaugeVec
	containerRootfsInodesUsed     *prometheus.GaugeVec

	// Container Logs
	containerLogsAvailableBytes *prometheus.GaugeVec
	containerLogsCapacityBytes  *prometheus.GaugeVec
	containerLogsUsageBytes     *prometheus.GaugeVec
	containerLogsInodes         *prometheus.GaugeVec
	containerLogsInodesFree     *prometheus.GaugeVec
	containerLogsInodesUsed     *prometheus.GaugeVec

	// Pod Volume
	podVolumeAvailableBytes *prometheus.GaugeVec
	podVolumeCapacityBytes  *prometheus.GaugeVec
	podVolumeUsageBytes     *prometheus.GaugeVec
	podVolumeInodes         *prometheus.GaugeVec
	podVolumeInodesFree     *prometheus.GaugeVec
	podVolumeInodesUsed     *prometheus.GaugeVec
}

// NewMetrics initializes and registers the Prometheus metrics for the exporter.
func NewMetrics() (*Metrics, error) {
	// create Prometheus metrics using the newGaugeVec helper function in utils.go
	return &Metrics{
		apiStatus: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Name:      "api_status",
			Help:      "Summary API status on the target node (0 is down, 1 is up)",
		}, []string{"exported_node"}),
		apiLatency: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Name:      "api_latency",
			Help:      "Summary API response time in seconds",
		}, []string{"exported_node"}),
		ephemeralStorageAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "available_bytes",
			Help:      "Ephemeral storage available space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		ephemeralStorageCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "capacity_bytes",
			Help:      "Ephemeral storage space capacity in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		ephemeralStorageUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "used_bytes",
			Help:      "Ephemeral storage used space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		ephemeralStorageInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "inodes_total",
			Help:      "Ephemeral storage number of total inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		ephemeralStorageInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "inodes_free",
			Help:      "Ephemeral storage number of free inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		ephemeralStorageInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSEphemeralStorage,
			Name:      "inodes_used",
			Help:      "Ephemeral storage number of used inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node"}),
		containerMemoryAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerMemory,
			Name:      "available_bytes",
			Help:      "Container memory available space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerMemoryCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerMemory,
			Name:      "capacity_bytes",
			Help:      "Container memory capacity in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerMemoryUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerMemory,
			Name:      "usage_bytes",
			Help:      "Container memory used space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "available_bytes",
			Help:      "Container root file system available space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "capacity_bytes",
			Help:      "Container root file system capacity in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "usage_bytes",
			Help:      "Container root file system used space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "inodes_total",
			Help:      "Container root file system total number of inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "inodes_free",
			Help:      "Container root file system number of free inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerRootfsInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerRootFS,
			Name:      "inodes_used",
			Help:      "Container root file system number of used inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "available_bytes",
			Help:      "Container logs space available in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "capacity_bytes",
			Help:      "Container logs capacity in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "usage_bytes",
			Help:      "Container logs used space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "inodes_total",
			Help:      "Container logs total number of inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "inodes_free",
			Help:      "Container logs number of free inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		containerLogsInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSContainerLogs,
			Name:      "inodes_used",
			Help:      "Container logs number of used inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_container"}),
		podVolumeAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "available_bytes",
			Help:      "Pod volume space available in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
		podVolumeCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "capacity_bytes",
			Help:      "Pod volume capacity in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
		podVolumeUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "usage_bytes",
			Help:      "Pod volume used space in bytes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
		podVolumeInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "inodes_total",
			Help:      "Pod volume total number of inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
		podVolumeInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "inodes_free",
			Help:      "Pod volume number of free inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
		podVolumeInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SSPodVolume,
			Name:      "inodes_used",
			Help:      "Pod volume number of used inodes",
		}, []string{"exported_pod", "exported_namespace", "exported_node", "exported_volume"}),
	}, nil
}
