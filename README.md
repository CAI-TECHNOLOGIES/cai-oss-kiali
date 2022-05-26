# Kiali Documentation

## Introduction
______
Kiali is a management console for Istio service mesh.

## Directory Structure

```
.
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── GOVERNANCE.md
├── LICENSE
├── Makefile
├── PULL_REQUEST_TEMPLATE.md
├── README.adoc
├── README.md
├── RELEASING.adoc
├── SECURITY.md
├── STYLE_GUIDE.adoc
├── appstate
├── business
├── config
├── deploy
│   ├── docker
│   ├── get-console.sh
│   ├── jenkins-ci
│   ├── kubernetes
│   └── openshift
├── design
├── doc.go
├── frontend
│   ├── Makefile.jenkins
│   ├── README.adoc
│   ├── RELEASING.adoc
│   ├── cypress
│   ├── cypress.json
│   ├── fetch-k-charted.sh
│   ├── getLabels.py
│   ├── package.json
│   ├── public
│   ├── src
│   ├── tsconfig.json
│   ├── tsconfig.prod.json
│   ├── tsconfig.test.json
│   └── yarn.lock
├── go.mod
├── go.sum
├── graph
├── hack
│   ├── README.adoc
│   ├── aws-openshift.sh
│   ├── check_go_version.sh
│   ├── ci-kind-molecule-tests.sh
│   ├── ci-minikube-molecule-tests.sh
│   ├── ci-openshift-molecule-tests.sh
│   ├── configure-operator.sh
│   ├── crc-openshift.sh
│   ├── docker-io-auth.sh
│   ├── fix_imports.sh
│   ├── hooks
│   ├── ibmcloud-openshift.sh
│   ├── istio
│   ├── jaeger-dep-config.sh
│   ├── jwt-decode.sh
│   ├── jwt-encode.sh
│   ├── k8s-minikube.sh
│   ├── kiali-port-forward.sh
│   ├── purge-kiali-from-cluster.sh
│   ├── run-kiali-config-template.yaml
│   ├── run-kiali.sh
│   ├── run-molecule-tests.sh
│   ├── run-prometheus.sh
│   └── setup-kind-in-ci.sh
├── handlers
├── jaeger
├── kiali.go
├── kiali_api.md
├── kiali_test.go
├── kubernetes
├── log
├── main_test.go
├── make
│   ├── Makefile.build.mk
│   ├── Makefile.cluster.mk
│   ├── Makefile.container.mk
│   ├── Makefile.helm.mk
│   ├── Makefile.molecule.mk
│   ├── Makefile.operator.mk
│   └── Makefile.ui.mk
├── models
├── observability
├── prometheus
├── routing
├── server
├── status
├── swagger.json
├── tests
│   ├── data
│   ├── e2e
│   ├── integration
│   └── testutils
├── tools
└── util

56 directories, 222 files
```

## Important Directories
______
- `frontend` - Contains the UI realted files for Kiali Console. (React)
- `deploy` - Contains the DockerFile for building kiali image and operator image. (Inside `docker` folder)
- `make` - The make scripts that build the image. We use 2 files: `Makefile.build.mk1` and `Makefile.container.mk`.
- `tests` - run tests on the image built.

## Pre-requisites
_____
- NodeJS (Node.js >= 12.22.0 && <16 with the NPM command)
- Go (GO_VERSION_KIALI = 1.17.7)

## Build Commands
____

```bash
export PATH=$PATH:/usr/local/go/bin
echo fs.inotify.max_user_watches=131070 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p

make build test
make build-ui-test
make container-build

```

## Functionality
____

- [Topology](https://kiali.io/docs/features/topology/)
- [Health](https://kiali.io/docs/features/health/)
- [Tracing](https://kiali.io/docs/features/tracing/) (Through Jaeger)
- [Validation and Istio-Configuration](https://kiali.io/docs/features/validations/)
- [Generate application and request routing configuration](https://kiali.io/docs/features/wizards/) 
