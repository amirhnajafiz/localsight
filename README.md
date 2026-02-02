# LocalSight

Even though local storage is a resource you can set in Kubernetes manifests, there aren’t tools (exporters) that show how much of it is actually being used. When this local storage fills up, it causes problems which makes the Kubelet remove (evict) pods. An exporter that reports how much local storage each pod is using would help with monitoring and setting up alerts before issues happen. That's where LocalSight comes in use.
