module github.com/openshift/gcp-pd-csi-driver-operator

go 1.13

require (
	github.com/bertinatto/csi-driver-controller v0.0.0-20200618075311-97c61ba58279 // indirect
	github.com/jteeuwen/go-bindata v3.0.8-0.20151023091102-a0ff2567cfb7+incompatible
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/openshift/api v0.0.0-20200521101457-60c476765272
	github.com/openshift/build-machinery-go v0.0.0-20200424080330-082bf86082cc
	github.com/openshift/client-go v0.0.0-20200521150516-05eb9880269c // indirect
	github.com/openshift/library-go v0.0.0-20200623125929-17bb296f39b8
	github.com/prometheus/client_golang v1.4.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	golang.org/x/tools v0.0.0-20200513154647-78b527d18275 // indirect
	google.golang.org/genproto v0.0.0-20191220175831-5c49e3ecc1c1 // indirect
	k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/code-generator v0.18.3
	k8s.io/component-base v0.18.3
	k8s.io/klog v1.0.0
)

replace github.com/openshift/library-go => ../library-go
