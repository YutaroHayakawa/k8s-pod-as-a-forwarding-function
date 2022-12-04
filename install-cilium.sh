cilium install \
	--helm-set tunnel=disabled \
	--helm-set ipv4NativeRoutingCIDR="10.0.0.0/8" \
	--helm-set bgpControlPlane.enabled=true
