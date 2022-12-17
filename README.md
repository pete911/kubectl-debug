# kubectl-debug

Main aim of this project is to serve as template for writing kubectl plugins. This plugin has currently only one
command `kubectl debug pod` which displays more information about failing pods.

Project needs to be built and binary placed on path.

## build

 - `make build` or `make install`
 - place built binary on path, or in case of `make install` make sure `GOBIN` is exported

## use

 - get all pods under a namespace `kubectl debug pod -n kube-system`
 - select pods by label `kubectl debug pod -n kube-system -l integration-test=storage-provisioner`
 - specific pod `kubectl debug pod -n kube-system coredns-66bff467f8-7mz46`

## example

```
kubectl debug pod -n kube-system coredns-66bff467f8-7mz46

--- [Namespace: kube-system Pod: coredns-565d847f94-r9wvh] ---
Labels:
  k8s-app: kube-dns
  pod-template-hash: 565d847f94
Annotations:
  -
Events:
  -
Container: coredns
  Image:        registry.k8s.io/coredns/coredns:v1.9.3
  Ready:        true
  Started:      true
  Restart:      0
  State:        Running
  Last State:   
  Logs:
    .:53
    [INFO] plugin/reload: Running configuration SHA512 = 591cf328cccc12bc490481273e738df59329c62c0b729d94e8b61db9961c2fa5f046dd37f1cf888b953814040d180f52594972691cd6ff41be96639138a43908
    CoreDNS-1.9.3
    linux/arm64, go1.18.2, 45b0a11
```
