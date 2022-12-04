TAG=k8s-pod-as-a-forwarding-function

cilium install \
	--chart-directory $CHART \
	--helm-set tunnel=disabled \
	--helm-set ipv4NativeRoutingCIDR="10.0.0.0/8" \
	--helm-set bgpControlPlane.enabled=true \
	--helm-set image.repository=localhost:5000/cilium/cilium \
	--helm-set image.tag=$TAG \
	--helm-set image.pullPolicy=Always \
	--helm-set image.useDigest=false \
	--helm-set operator.image.repository=localhost:5000/cilium/operator \
	--helm-set operator.image.tag=$TAG \
	--helm-set operator.image.suffix=""
