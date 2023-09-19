## 1.0.4 (2023-07-22)


### Bug Fixes

* **common**: invalid syntax in ServiceMonitor helm template ([f9a74f50](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/f9a74f50))



## 1.0.3 (2023-07-20)



## 1.0.2 (2023-07-13)


### Documentation

* **chart**: add description and README ([2aee55d9](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/2aee55d9))
* **common**: fix coverage badge in README.md ([ed18e551](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/ed18e551))


### New Features

* **common**: Add flag for prometheus-operator SMon ([648ec9fe](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/648ec9fe))
* **common**: Stronger security context defaults ([dec49d36](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/dec49d36))


### Tasks

* **common**: omit deprecated node-role.kubernetes.io/master taint ([58c08cfd](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/58c08cfd))
  > fixes https://github.com/postfinance/kubelet-csr-approver/issues/123



## 1.0.1 (2023-05-31)


### Documentation

* **README**: add leader-election documentation ([70e8b9b8](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/70e8b9b8))


### New Features

* **common**: enable leader-election ([8aae0967](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/8aae0967))


### Tasks

* **leader-election**: make leader election optional ([9a047cb4](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/9a047cb4))


### Test

* **ci**: setup-envtest with k8s 1.27.x ([3ebd69c8](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/3ebd69c8))
* **leader-election**: default to kube-system namespace when not running in-cluster ([161a3e07](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/161a3e07))



## 1.0.0 (2023-03-31)


### Tasks

* **common**: automate helm versioning ([af8d8b79](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/af8d8b79))
* **common**: bump to Golang 1.20 and replace inet.af with netip{,x} ([89f5330b](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/89f5330b))
* **common**: switch default container registry ([ac65893f](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/ac65893f))
  > #133
* **gh-action**: rename misnamed check-dockerhub-token job ([01d809e3](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/01d809e3))
  > [skip ci]


### Test

* **common**: permit specifying multiple comma-separated SAN ([5bda15bd](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/5bda15bd))
  > [skip ci]



## 0.2.8 (2023-03-10)


### New Features

* **common**: add chart support for custom dnsConfig ([31eeb2f6](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/31eeb2f6))


### Tasks

* **common**: bump chart version to v0.2.8 ([8ab8cbc8](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/8ab8cbc8))



## 0.2.7 (2023-03-07)


### Tasks

* **common**: add CODEOWNERS file ([4edb3fd9](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/4edb3fd9))
* **common**: bump chart version to v0.2.7 ([a0ea8b71](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/a0ea8b71))
  > should fix #130



## 0.2.6 (2023-02-21)


### Tasks

* **common**: add profile.cov to .gitignore ([9ddc539d](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/9ddc539d))
  > [skip ci]



## 0.2.5 (2023-02-21)


### New Features

* **common**: support deploy in out-of-cluster mode (#116) ([3c55012d](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/3c55012d))
  > * feat: support deploy in out-of-cluster mode


### Tasks

* **common**: fix linting errors ([aea094f8](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/aea094f8))
* **common**: reformat LICENSE.md ([2fd4edd4](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/2fd4edd4))



## 0.2.4 (2022-08-30)


### Documentation

* **common**: bypass-hostname-check explanation ([73ed5af2](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/73ed5af2))
  > cf. #76
  > [skip ci]
* **common**: document allowed-dns-names configuration flag ([f8e2ef23](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/f8e2ef23))
  > [skip ci]


### New Features

* **common**: add `--bypass-hostname-check` config flag ([f22fefb0](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/f22fefb0))
  > cf. #76
* **common**: bypass hostname check ([f7f44c25](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/f7f44c25))
  > relates to #76
  > [skip ci]
* **common**: permit using multiple DNS names in CSRs ([056b5fd2](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/056b5fd2))
  > related to #70


### Tasks

* **common**: cleanup commented test ([7a79eae8](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7a79eae8))
* **common**: fix linting error for nolint comment ([9562c44f](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/9562c44f))
  > https://github.com/golangci/golangci-lint/issues/3109#issuecomment-1218872255
* **common**: refactor code with embedded struct configs ([fcf5ead9](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/fcf5ead9))


### Test

* **common**: bump k8s version for envtest to 1.24 ([9c441333](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/9c441333))
* **common**: validate bypass hostname check behaviour ([68c66571](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/68c66571))



## 0.2.3 (2022-06-29)


### New Features

* **common**: permit CSR with non-system-node usernames ([#61](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/61), [decc89af](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/decc89af))


### Tasks

* **common**: document `ignore-non-system-node` and add to helm chart ([156e4c34](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/156e4c34))
* **helm**: make it possible to specify logging level ([#60](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/60), [7a663b81](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7a663b81))



## 0.2.2 (2022-04-11)


### Tasks

* **common**: simplify GH actions ([dfc4154f](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/dfc4154f))
* **common**: switch to proper SemVer for Helm chart ([#46](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/46), [7f2cb27a](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7f2cb27a))



## 0.2.1 (2022-04-07)



## 0.2.0 (2022-04-01)


### Documentation

* **common**: update README.md to describe new IP whitelisting feature ([36cbcdc2](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/36cbcdc2))
  > [skip ci]


### New Features

* **common**: dns-resolved IPs must fall within whitelisted IP set ([d09363ba](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/d09363ba))
* **common**: prefix based IP whitelisting ([accb23e9](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/accb23e9))
  > https://github.com/postfinance/kubelet-csr-approver/issues/28
* **common**: start implementing IP whitelist vaildation ([6b496588](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/6b496588))
  > https://github.com/postfinance/kubelet-csr-approver/issues/28


### Tasks

* **common**: bump dependencies ([27df807d](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/27df807d))
* **common**: bump helm chart version to v0.2.0 ([4ad61732](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/4ad61732))
* **common**: fix critical conditional handling for whitelistedIPCheck ([1c3aaefd](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/1c3aaefd))
* **common**: fix documentation and implement v4 and v6 whitelisted IP tests ([d8d9d5ab](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/d8d9d5ab))
* **common**: switch default branch to main ([fb425f58](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/fb425f58))


### Test

* **common**: add ipv6 tests as well ([a80342e6](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/a80342e6))
* **helm**: add providerIpPrefixes values.yaml test ([7bd70c00](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7bd70c00))



## 0.1.2 (2022-01-18)


### Documentation

* **common**: add build badge to the README ([0a609295](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/0a609295))
* **common**: document bypass-dns-resolution config flag ([1af62c80](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/1af62c80))
  > [skip ci]
* **common**: improve documentation with MAX_EXPIRATION_SEC section ([1901c6b5](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/1901c6b5))
* **common**: improve parameters section with new cmdline-env options ([0bb1705d](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/0bb1705d))


### New Features

* **common**: implement DNS resolution bypass config flag ([1652c90c](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/1652c90c))
  > permits to ignore  DNS resolution and will permit to fix #15
  > some refactoring with this commit to make it easier to implement this
  > feature
* **common**: implement maxExpirationSeconds check ([158bebaa](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/158bebaa))
  > related to #10
  > I implemented a strict maximum of 367 days for all CSR where the
  > expirationSeconds field is set
  > 
  > the user can optionally reduce this maximum by setting the
  > MAX_EXPIRATION_SEC environment variable


### Tasks

* **common**: add back the `-race` test flag ([#1](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/1), [0537fe3c](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/0537fe3c))
* **common**: add example max_expiration_sec env variable to k8s manifests ([7708571b](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7708571b))
* **common**: add helm configuration option and improve doc ([95bbc103](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/95bbc103))
* **common**: bump dependencies (k8s api, controller-runtime, logr) ([02b14f6a](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/02b14f6a))
* **common**: cleanup startup code and use ff/v3 for flags/environ ([f1695f0e](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/f1695f0e))
* **common**: fix chart linting issues ([7cc3eb16](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7cc3eb16))
* **common**: fix duplicate env key in helm chart ([7555d8ee](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/7555d8ee))
* **common**: fix helm-testing value to string ([483d029c](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/483d029c))
* **common**: fix incorrect image in deployment.yaml ([#8](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/8), [af3061e7](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/af3061e7))
* **common**: implement test for bypassDNSResolution flag ([#15](https://github.com/github.com/postfinance/kubelet-csr-approver/issues/15), [4a9a9ec6](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/4a9a9ec6))
* **common**: improve README.md formatting and .gitignore ([786ac0d3](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/786ac0d3))
* **common**: make k8s auth providers opt-in with build tag ([6be1ca4f](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/6be1ca4f))
* **common**: quote maxExpirationSeconds env var in helm chart ([406f6a1f](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/406f6a1f))
* **common**: refactor test package to reuse Cmd functions ([1a6d4af5](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/1a6d4af5))
* **common**: refactoring ([c12e4a35](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/c12e4a35))
* **common**: remove unused vars ([3b00bbf5](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/3b00bbf5))
* **helm**: update chart to v0.1.2 ([908812e7](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/908812e7))


### Test

* **common**: improve test cases ([47e843f3](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/47e843f3))



## 0.1.1 (2021-12-01)


### Tasks

* **common**: repository clean-up and multi-platform build ([e004aa35](https://github.com/github.com/postfinance/kubelet-csr-approver/commit/e004aa35))



