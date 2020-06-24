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
