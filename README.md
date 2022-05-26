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

Follow these [pre-requisites](https://github.com/CAI-TECHNOLOGIES/kiali/blob/master/README_old.adoc#developer-setup) to set things up.
Use the folllowing commands to build image.

```bash
git clone https://github.com/CAI-TECHNOLOGIES/cai-oss-kiali.git

# Using Go
export PATH=$PATH:/usr/local/go/bin

make build test
make build-ui-test
make container-build

# Increasing Max Watchers incase of error
echo fs.inotify.max_user_watches=131070 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p

```

## Functionality
____

- [Topology](https://kiali.io/docs/features/topology/)
- [Health](https://kiali.io/docs/features/health/)
- [Tracing](https://kiali.io/docs/features/tracing/) (Through Jaeger)
- [Validation and Istio-Configuration](https://kiali.io/docs/features/validations/)
- [Generate application and request routing configuration](https://kiali.io/docs/features/wizards/) 

## Enabling RBAC in Kiali
____

Refer to [this documentation](https://developer.okta.com/blog/2021/10/08/secure-access-to-aws-eks) for associating the EKS cluster on which kiali is to be brought up with Okta. However use client secret method instead of pkce. Configure `.kube/config` accordingly to include client secret.

```bash
# For using oidc login
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"

kubectl oidc-login setup --oidc-issuer-url=ISSUER_URL --oidc-client-secret=CLIENT_SECRET --oidc-client-id=CLIENT_ID
```

Before bringing up kiali, make sure client secret for kiali is updated. Use the following command to do so, and restart kiali if already up.

```bash
kubectl create secret generic kiali --from-literal="oidc-secret=meJz3nTIybneGksCvHcpT2UnvIDPX_1ThElOU9Zb" -n istio-system
```
