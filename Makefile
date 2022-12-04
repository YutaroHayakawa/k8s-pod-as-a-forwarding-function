CLUSTERNAME=clab-k8s-pod-as-a-forwarding-function

deploy:
	kind create cluster --config cluster.yaml
	sudo containerlab -t topo.yaml deploy
	./install-cilium.sh
	kubectl apply -f k8s-resources.yaml

redeploy:
	kind delete clusters $(CLUSTERNAME)
	kind create cluster --config cluster.yaml
	sudo containerlab -t topo.yaml deploy --reconfigure
	./install-cilium.sh
	kubectl apply -f k8s-resources.yaml

redeploy-clab:
	sudo containerlab -t topo.yaml deploy --reconfigure

destroy:
	sudo containerlab -t topo.yaml destroy
	kind delete clusters $(CLUSTERNAME)

router0-summary:
	docker exec -it clab-k8s-pod-as-a-forwarding-function-router0 vtysh -c "show bgp ipv4 summary"

router0-routes:
	docker exec -it clab-k8s-pod-as-a-forwarding-function-router0 vtysh -c "show bgp ipv4"

server3-summary:
	docker exec -it clab-k8s-pod-as-a-forwarding-function-server3 vtysh -c "show bgp ipv4 summary"

server3-routes:
	docker exec -it clab-k8s-pod-as-a-forwarding-function-server3 vtysh -c "show bgp ipv4"
