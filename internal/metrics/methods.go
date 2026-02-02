package metrics

// SetAPIValues sets the summary API status on the target node.
func (m *Metrics) SetAPIStatus(node string, status int) {
	m.apiStatus.WithLabelValues(node).Set(float64(status))
}

// SetAPIValues sets the summary API metrics on the target node.
func (m *Metrics) SetAPIValues(node string, latency float64) {
	m.apiLatency.WithLabelValues(node).Set(latency)
}

// SetEphemeralStorageValues sets the ephemeral storage metrics for a specific pod, namespace, and node.
func (m *Metrics) SetEphemeralStorageValues(
	pod, namespace, node string,
	used, available, capacity float64,
) {
	m.ephemeralStorageUsageBytes.WithLabelValues(pod, namespace, node).Set(used)
	m.ephemeralStorageAvailableBytes.WithLabelValues(pod, namespace, node).Set(available)
	m.ephemeralStorageCapacityBytes.WithLabelValues(pod, namespace, node).Set(capacity)
}

// SetEphemeralStorageInodes sets the ephemeral storage inode metrics for a specific pod, namespace, and node.
func (m *Metrics) SetEphemeralStorageInodes(
	pod, namespace, node string,
	used, available, capacity float64,
) {
	m.ephemeralStorageInodesUsed.WithLabelValues(pod, namespace, node).Set(used)
	m.ephemeralStorageInodesFree.WithLabelValues(pod, namespace, node).Set(available)
	m.ephemeralStorageInodes.WithLabelValues(pod, namespace, node).Set(capacity)
}

// SetContainerMemoryValues sets the memory metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerMemoryValues(
	pod, namespace, node, container string,
	used, available, capacity float64,
) {
	m.containerMemoryUsageBytes.WithLabelValues(pod, namespace, node, container).Set(used)
	m.containerMemoryAvailableBytes.WithLabelValues(pod, namespace, node, container).Set(available)
	m.containerMemoryCapacityBytes.WithLabelValues(pod, namespace, node, container).Set(capacity)
}

// SetContainerRootfsValues sets the root filesystem metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerRootfsValues(
	pod, namespace, node, container string,
	used, available, capacity float64,
) {
	m.containerRootfsUsageBytes.WithLabelValues(pod, namespace, node, container).Set(used)
	m.containerRootfsAvailableBytes.WithLabelValues(pod, namespace, node, container).Set(available)
	m.containerRootfsCapacityBytes.WithLabelValues(pod, namespace, node, container).Set(capacity)
}

// SetContainerRootfsInodes sets the root filesystem inode metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerRootfsInodes(
	pod, namespace, node, container string,
	used, available, capacity float64,
) {
	m.containerRootfsInodesUsed.WithLabelValues(pod, namespace, node, container).Set(used)
	m.containerRootfsInodesFree.WithLabelValues(pod, namespace, node, container).Set(available)
	m.containerRootfsInodes.WithLabelValues(pod, namespace, node, container).Set(capacity)
}

// SetContainerLogsValues sets the logs metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerLogsValues(
	pod, namespace, node, container string,
	used, available, capacity float64,
) {
	m.containerLogsUsageBytes.WithLabelValues(pod, namespace, node, container).Set(used)
	m.containerLogsAvailableBytes.WithLabelValues(pod, namespace, node, container).Set(available)
	m.containerLogsCapacityBytes.WithLabelValues(pod, namespace, node, container).Set(capacity)
}

// SetContainerLogsInodes sets the logs inode metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerLogsInodes(
	pod, namespace, node, container string,
	used, available, capacity float64,
) {
	m.containerLogsInodesUsed.WithLabelValues(pod, namespace, node, container).Set(used)
	m.containerLogsInodesFree.WithLabelValues(pod, namespace, node, container).Set(available)
	m.containerLogsInodes.WithLabelValues(pod, namespace, node, container).Set(capacity)
}

// SetPodVolumeValues sets the pod volume metrics for a specific volume in a pod, namespace, and node.
func (m *Metrics) SetPodVolumeValues(
	pod, namespace, node, volume string,
	used, available, capacity float64,
) {
	m.podVolumeUsageBytes.WithLabelValues(pod, namespace, node, volume).Set(used)
	m.podVolumeAvailableBytes.WithLabelValues(pod, namespace, node, volume).Set(available)
	m.podVolumeCapacityBytes.WithLabelValues(pod, namespace, node, volume).Set(capacity)
}

// SetPodVolumeInodes sets the pod volume inode metrics for a specific volume in a pod, namespace, and node.
func (m *Metrics) SetPodVolumeInodes(
	pod, namespace, node, volume string,
	used, available, capacity float64,
) {
	m.podVolumeInodesUsed.WithLabelValues(pod, namespace, node, volume).Set(used)
	m.podVolumeInodesFree.WithLabelValues(pod, namespace, node, volume).Set(available)
	m.podVolumeInodes.WithLabelValues(pod, namespace, node, volume).Set(capacity)
}
