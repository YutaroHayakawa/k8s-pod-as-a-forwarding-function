# Using k8s Pod as a forwarding function (LB, NAT, VPN-GW, etc...)

What could go wrong?

## Topology

```mermaid
flowchart TB
  Router0---Server0["Server0 (k8s)"]
  Router0---Server1["Server1 (k8s)"]
  Router0---Server2["Server2 (k8s)"]
  Router0---Server3
  Router0---Server4
  Router0---Server5
```

## How to use `make`

- Create Lab: `make deploy`
- Destroy Lab: `make destroy`
- Recreate Lab: `make redeploy`
