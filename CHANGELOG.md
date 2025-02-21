
## [1.2.5] - 2025-01-08

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.2.4...v1.2.5

### 🐛 Bug Fixes

- Typo in env var name in helm chart (MAX_EXPIRATION_SEC) - ([ce86095](https://github.com/postfinance/kubelet-csr-approver/commit/ce86095408801fb5521ae28bde3ad1c009f89f37))

### Build

- *(helm-action)* Update helm-test dependencies - ([a8521d3](https://github.com/postfinance/kubelet-csr-approver/commit/a8521d33cca0afdbb07cdad9ef9fa8c78351d993))




## New Contributors
* @NotWearingPants made their first contribution in [#296](https://github.com/postfinance/kubelet-csr-approver/pull/296)## [1.2.4] - 2024-12-30

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.2.3...v1.2.4

### 🚀 Features

- Add flag to disable CSR denial - ([f440f6f](https://github.com/postfinance/kubelet-csr-approver/commit/f440f6f1f6934cd861a3e4ba19d879d96e8309ad))

### 🧪 Testing

- *(skip-deny)* Add test covering 'skip-deny-csr' flag - ([68e4d8a](https://github.com/postfinance/kubelet-csr-approver/commit/68e4d8ae0ac5da02ab2f08f2404fadc23096779a))

### ⚙️ Miscellaneous Tasks

- Document skip-deny-step and add helm parameter - ([2bcaf1e](https://github.com/postfinance/kubelet-csr-approver/commit/2bcaf1eeb6f04801c88a2ff968195d87f31fe3e0))

### Build

- *(deps)* Bump k8s.io/client-go from 0.31.1 to 0.31.2 - ([d696824](https://github.com/postfinance/kubelet-csr-approver/commit/d696824ac52b4419d7e6d55d7424eb738ff3ca32))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.19.0 to 0.19.1 - ([cbb12e6](https://github.com/postfinance/kubelet-csr-approver/commit/cbb12e6fd50c0203bfd83807f723788e7484fe4a))
- *(deps)* Bump k8s.io/api from 0.31.2 to 0.31.3 - ([987c539](https://github.com/postfinance/kubelet-csr-approver/commit/987c539d79730ac090d5fa799a37e62377e08aab))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.19.1 to 0.19.2 - ([5c535df](https://github.com/postfinance/kubelet-csr-approver/commit/5c535df4e23e58fa36ef756bd793e68b061a6536))
- *(deps)* Bump k8s.io/client-go from 0.31.2 to 0.31.3 - ([bf5dc76](https://github.com/postfinance/kubelet-csr-approver/commit/bf5dc76983697fefb98b3f9d8448b835a821510f))
- *(deps)* Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 - ([d0c6f5e](https://github.com/postfinance/kubelet-csr-approver/commit/d0c6f5e43c3d5e81ad94c48976dc910db95264fc))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.19.2 to 0.19.3 - ([ad010d1](https://github.com/postfinance/kubelet-csr-approver/commit/ad010d13ddbd7596049f34e9d0d5ec1ac515c818))
- *(deps)* Bump helm/kind-action from 1.10.0 to 1.11.0 - ([ce657a6](https://github.com/postfinance/kubelet-csr-approver/commit/ce657a6fba4dc9c3fd8847d4da34e1e4f15a1728))
- *(deps)* Bump helm/kind-action from 1.11.0 to 1.12.0 - ([9edf2c6](https://github.com/postfinance/kubelet-csr-approver/commit/9edf2c6d9dde6a6072f82ca96df0f474768ff108))


## [1.2.3] - 2024-10-07

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.2.2...v1.2.3

### ⚙️ Miscellaneous Tasks

- Add v1.2.2 changelog - ([c5ca70d](https://github.com/postfinance/kubelet-csr-approver/commit/c5ca70db40ca5002e9d7c047eb7126049b97dbf6))

### Build

- *(deps)* Bump ko-build/setup-ko from 0.6 to 0.7 - ([f5b863f](https://github.com/postfinance/kubelet-csr-approver/commit/f5b863f77e0b71a8030323b0ed1d6df0601b1df0))
- *(deps)* Bump k8s.io/client-go from 0.30.2 to 0.30.3 - ([584183f](https://github.com/postfinance/kubelet-csr-approver/commit/584183f5329e82d10c392efc5cff3c8819f5a032))
- *(deps)* Bump golangci/golangci-lint-action from 6.0.1 to 6.1.0 - ([47f3776](https://github.com/postfinance/kubelet-csr-approver/commit/47f37761384d627cd20c5df75064bf7d682cadc1))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.18.4 to 0.19.0 - ([501df7f](https://github.com/postfinance/kubelet-csr-approver/commit/501df7ffe3b8ee6aaf9eeb478a4bd18e7faa4fd6))
- *(deps)* Bump k8s.io/client-go from 0.31.0 to 0.31.1 - ([bf8d08b](https://github.com/postfinance/kubelet-csr-approver/commit/bf8d08ba3e4b5e231d988442a8d0bb788a3ac2e5))
- *(deps)* Bump golangci/golangci-lint-action from 6.1.0 to 6.1.1 - ([4a0096c](https://github.com/postfinance/kubelet-csr-approver/commit/4a0096c41c08751ef4b6b660fa79565d61693251))


## [1.2.2] - 2024-06-19

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.2.1...v1.2.2

### 🚀 Features

- Allow specifying (optional) priorityClass - ([44e2c43](https://github.com/postfinance/kubelet-csr-approver/commit/44e2c43d2db00ff5a56e3634760bed6105150b47))

### Build

- *(deps)* Bump github.com/go-logr/logr from 1.4.1 to 1.4.2 - ([c33c45d](https://github.com/postfinance/kubelet-csr-approver/commit/c33c45d70cf96fa8bdb3abc799a1d1adc3934bbf))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.18.2 to 0.18.3 - ([3b56613](https://github.com/postfinance/kubelet-csr-approver/commit/3b566133d5ec28c93463d7931a12c5a3bc71af6b))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.18.3 to 0.18.4 - ([846d113](https://github.com/postfinance/kubelet-csr-approver/commit/846d1138ebf0e0e835854f173f29cd46afc9afdb))
- *(deps)* Bump k8s.io/client-go from 0.30.1 to 0.30.2 - ([f91bfba](https://github.com/postfinance/kubelet-csr-approver/commit/f91bfbaf0bc3bc9e4237aa146f0b29b85c806ece))


## New Contributors
* @jaccolo made their first contribution in PR #263

## [1.2.1] - 2024-05-21

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.2.0...1.2.1

### 🐛 Bug Fixes

- X509 CR SAN IPs properly ignored when bypassing DNS Resolution - ([4a53481](https://github.com/postfinance/kubelet-csr-approver/commit/4a5348140bbac00716bbc9b6e47c2a01edeef6a2))

### 🧪 Testing

- Add testcase covering issue #253 - ([4118bc0](https://github.com/postfinance/kubelet-csr-approver/commit/4118bc045fca8e491c50ef694cd3908e63156817))

### ⚙️ Miscellaneous Tasks

- Add (cliff) changelog for v1.2.0 - ([fb0ae32](https://github.com/postfinance/kubelet-csr-approver/commit/fb0ae32e3e2d70f3633ea6def59f113df9bca808))
- Add changelog for v1.2.1 - ([4622254](https://github.com/postfinance/kubelet-csr-approver/commit/46222547cca9ffc597a5236225e28ef7ccc49b7b))

### Build

- *(deps)* Bump k8s.io/api from 0.30.0 to 0.30.1 - ([c2effac](https://github.com/postfinance/kubelet-csr-approver/commit/c2effac08baeaf542e7e132b9f043d8ab5ba9484))


## [1.2.0] - 2024-05-13

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.1.0...v1.2.0

### 🐛 Bug Fixes

- Also validate DNS/SAN names when DNS resolution is bypassed - ([f4654d3](https://github.com/postfinance/kubelet-csr-approver/commit/f4654d3fe3013185e598e9601e09d60490717b58))

### 🧪 Testing

- Update to k8s v1.30 for envtest - ([7eb8ddd](https://github.com/postfinance/kubelet-csr-approver/commit/7eb8ddddda5c02c628c91f353edf12c626a5bd66))
- Add testcase covering issue #247 - ([f91b9a1](https://github.com/postfinance/kubelet-csr-approver/commit/f91b9a1781929314daf8f0b212c92d6f375a5d42))

### Build

- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.17.2 to 0.17.3 - ([eef91f8](https://github.com/postfinance/kubelet-csr-approver/commit/eef91f851f1388e7bfb1e1fd531b26fff77df6a4))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.17.3 to 0.18.0 - ([e1261e8](https://github.com/postfinance/kubelet-csr-approver/commit/e1261e801354e0ece7913bfe6a6f5261a77bacba))
- *(deps)* Bump helm/kind-action from 1.9.0 to 1.10.0 - ([1bc7816](https://github.com/postfinance/kubelet-csr-approver/commit/1bc781683f18c5827ee9d32b38179bc26040baec))
- *(deps)* Bump golangci/golangci-lint-action from 4.0.0 to 5.0.0 - ([7a51b11](https://github.com/postfinance/kubelet-csr-approver/commit/7a51b11ce01115f0b40f673fb37a0f5e5a748426))
- *(deps)* Bump golangci/golangci-lint-action from 5.0.0 to 5.3.0 - ([c45d1f8](https://github.com/postfinance/kubelet-csr-approver/commit/c45d1f80209313d8d8d4b36eef7add4207df9647))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.18.0 to 0.18.1 - ([1690a7e](https://github.com/postfinance/kubelet-csr-approver/commit/1690a7e08d9a10b4c825b921688c112c3ad52e04))
- *(deps)* Bump golangci/golangci-lint-action from 5.3.0 to 6.0.1 - ([1d87285](https://github.com/postfinance/kubelet-csr-approver/commit/1d872856888543f645d661f8f9374fade0859fb3))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.18.1 to 0.18.2 - ([52f32bc](https://github.com/postfinance/kubelet-csr-approver/commit/52f32bc5c4ac1e54475012e3e7e6637e7a5c1c2f))
- *(kind)* Use default kind version from gh-action - ([c7d04b3](https://github.com/postfinance/kubelet-csr-approver/commit/c7d04b3e9db1f42f869a4910af596dde33af4ad0))


## [1.1.0] - 2024-04-04

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.7...v1.1.0

### 🚀 Features

- Allow adding extra objects to helm chart - ([302ad58](https://github.com/postfinance/kubelet-csr-approver/commit/302ad582ae1954fc0e7904fde98cd684bd3bf202))

### ⚙️ Miscellaneous Tasks

- Rename values.extraObjects field - ([b443620](https://github.com/postfinance/kubelet-csr-approver/commit/b443620f50c92a875516d358d4ade50f8c4d4c0a))
- Update to go 1.22 and envtest with k8s 1.29 - ([f525ea6](https://github.com/postfinance/kubelet-csr-approver/commit/f525ea61ab18f690b220baacfd18fa18163cad03))
- Update changelog with 1.1.0 release - ([07f9318](https://github.com/postfinance/kubelet-csr-approver/commit/07f9318bea95684bd07d69cb5538790e40140ff8))

### Build

- *(chart)* Fix servicemonitor namespaceSelector - ([093e538](https://github.com/postfinance/kubelet-csr-approver/commit/093e538618bd85ba358283c30f3ebe4d4e25caf8))
- *(deps)* Bump golangci/golangci-lint-action from 3.7.0 to 4.0.0 - ([3dddd7e](https://github.com/postfinance/kubelet-csr-approver/commit/3dddd7ef1e383a87c019d155fe62564601615a19))
- *(deps)* Bump helm/kind-action from 1.8.0 to 1.9.0 - ([7233407](https://github.com/postfinance/kubelet-csr-approver/commit/723340723e8405223676b86300b22882abc1404b))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.17.0 to 0.17.1 - ([bc67d62](https://github.com/postfinance/kubelet-csr-approver/commit/bc67d62aaea576a2bdbce3ee1c4d62ee0140f726))
- *(deps)* Bump github.com/foxcpp/go-mockdns from 1.0.0 to 1.1.0 - ([dd251c2](https://github.com/postfinance/kubelet-csr-approver/commit/dd251c2b240d35eb8ff1a54959a25a28106d82a1))
- *(deps)* Bump k8s.io/client-go from 0.29.1 to 0.29.2 - ([4a47742](https://github.com/postfinance/kubelet-csr-approver/commit/4a477423d27d0ba687e83f572bf657700ac4dc0d))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.17.1 to 0.17.2 - ([0c16855](https://github.com/postfinance/kubelet-csr-approver/commit/0c168555c92d73290aa480afb0f9057f39d4850a))
- *(deps)* Bump go.uber.org/zap from 1.26.0 to 1.27.0 - ([c8c8bfc](https://github.com/postfinance/kubelet-csr-approver/commit/c8c8bfcc86b19946ba3cecc4525d858569e7d21d))
- *(deps)* Bump azure/setup-helm from 3 to 4 - ([94f2e16](https://github.com/postfinance/kubelet-csr-approver/commit/94f2e16a9ee3f6f94081dfda62a6393933b72e83))
- *(deps)* Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 - ([962232a](https://github.com/postfinance/kubelet-csr-approver/commit/962232acb861c60f6755379c85813b37d18c151f))
- *(deps)* Bump k8s.io/apimachinery from 0.29.2 to 0.29.3 - ([9a32fe5](https://github.com/postfinance/kubelet-csr-approver/commit/9a32fe5b3c5d11b33952347b292211c2aece7109))
- *(deps)* Bump k8s.io/api from 0.29.2 to 0.29.3 - ([9d86e35](https://github.com/postfinance/kubelet-csr-approver/commit/9d86e35d7dbd43399d1a5ec943e4bb5addf2e046))
- *(deps)* Bump k8s.io/client-go from 0.29.2 to 0.29.3 - ([00d958d](https://github.com/postfinance/kubelet-csr-approver/commit/00d958d29bcd7237e7a8ac1305234b6b2b3f20b9))
- Update golangci-lint to v1.56.2 - ([79218a5](https://github.com/postfinance/kubelet-csr-approver/commit/79218a5e3b2450c450b9c0987e36349683a0d6b1))




## New Contributors
* @LeTT00r made their first contribution## [1.0.7] - 2024-01-29

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.6...v1.0.7

### 🚀 Features

- *(helm)* Add configuratble podLabels - ([b158e44](https://github.com/postfinance/kubelet-csr-approver/commit/b158e440b30453ced522cb02b500990473266d04))

### ⚙️ Miscellaneous Tasks

- Update changelog with 1.0.7 release - ([d08719c](https://github.com/postfinance/kubelet-csr-approver/commit/d08719c06e99a82ee153f0be4bc79d626a7d27be))

### Build

- *(deps)* Bump github.com/go-logr/logr from 1.3.0 to 1.4.1 - ([eb5a5fe](https://github.com/postfinance/kubelet-csr-approver/commit/eb5a5fe6a7f1fb276ea88aa4fef1c8dffb3987b0))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.16.3 to 0.17.0 - ([cb449bb](https://github.com/postfinance/kubelet-csr-approver/commit/cb449bbe8015f44f615b4cb7f0cca5a2a84913e7))
- *(deps)* Bump k8s.io/api from 0.29.0 to 0.29.1 - ([8de18aa](https://github.com/postfinance/kubelet-csr-approver/commit/8de18aa3d406f1d6b9c234be73b16590ef9524ea))
- *(deps)* Bump k8s.io/client-go from 0.29.0 to 0.29.1 - ([fa92830](https://github.com/postfinance/kubelet-csr-approver/commit/fa928301b6d2859050dbb79b53e0e791ee98270e))


## [1.0.6] - 2023-12-21

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.5...v1.0.6

### ⚙️ Miscellaneous Tasks

- Render build commits - ([f0a7d00](https://github.com/postfinance/kubelet-csr-approver/commit/f0a7d0018d9184177e0957c641479384be24b822))
- Update changelog with 1.0.6 release - ([c503194](https://github.com/postfinance/kubelet-csr-approver/commit/c503194a6c5f7b8e7cb2fd1b5a09dc0747151acb))

### Build

- *(deps)* Bump k8s.io/apimachinery from 0.28.2 to 0.28.3 - ([e41d4b0](https://github.com/postfinance/kubelet-csr-approver/commit/e41d4b0f00669a389cc7ef5c4cedce774ea4396c))
- *(deps)* Bump k8s.io/client-go from 0.28.2 to 0.28.3 - ([b3a8167](https://github.com/postfinance/kubelet-csr-approver/commit/b3a81673f45959798da830fb8d4dc5c48d905909))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.16.2 to 0.16.3 - ([4abe5fc](https://github.com/postfinance/kubelet-csr-approver/commit/4abe5fce2d5ffd8d10dab7518848d03a70e57e76))
- *(deps)* Bump github.com/go-logr/logr from 1.2.4 to 1.3.0 - ([6e687c7](https://github.com/postfinance/kubelet-csr-approver/commit/6e687c7723af090dfc9fdddf7aebdddb77675093))
- *(deps)* Bump helm/chart-testing-action from 2.4.0 to 2.6.1 - ([3f7e4de](https://github.com/postfinance/kubelet-csr-approver/commit/3f7e4deffd13ac474b3f6c16fa3cb1649cb239db))
- *(deps)* Bump github.com/go-logr/zapr from 1.2.4 to 1.3.0 - ([dc0e5fd](https://github.com/postfinance/kubelet-csr-approver/commit/dc0e5fd31470e7f398dab77c7eda99717cc46ff3))
- *(deps)* Bump k8s.io/api from 0.28.3 to 0.28.4 - ([3084cb6](https://github.com/postfinance/kubelet-csr-approver/commit/3084cb686f3771cc1fcb10041180ca539700845a))
- *(deps)* Bump k8s.io/client-go from 0.28.3 to 0.28.4 - ([a00eeed](https://github.com/postfinance/kubelet-csr-approver/commit/a00eeed18efb9fede116b154043d49a0651eef89))
- *(deps)* Bump actions/setup-python from 4 to 5 - ([80da607](https://github.com/postfinance/kubelet-csr-approver/commit/80da607fb87c221fe3a0e7cdf3a287163c61b724))
- *(deps)* Bump actions/setup-go from 4 to 5 - ([5089f3d](https://github.com/postfinance/kubelet-csr-approver/commit/5089f3d8829ef24a7136293925330aaca1a1fcd6))
- *(deps)* Bump k8s.io/apimachinery from 0.28.4 to 0.29.0 - ([a537683](https://github.com/postfinance/kubelet-csr-approver/commit/a537683d5658725cec791dedb1f24d85dbc3aaa8))
- *(deps)* Bump k8s.io/client-go from 0.28.4 to 0.29.0 - ([41beabf](https://github.com/postfinance/kubelet-csr-approver/commit/41beabf564b9f0d5e6adf2e3873c3e8eef6dd40d))




## New Contributors
* @remmen-io made their first contribution## [1.0.5] - 2023-09-19

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.4...v1.0.5

### 🚀 Features

- Add changelog config and CHANGELOG.md - ([bfb1801](https://github.com/postfinance/kubelet-csr-approver/commit/bfb1801263f35dea2e1103a171bbf9e7bf9c5a66))

### 🐛 Bug Fixes

- *(k8s-deploy)* Set image tag to latest - ([69e9249](https://github.com/postfinance/kubelet-csr-approver/commit/69e9249c5ab87cc1b30e09098d2f06924d6c2803))

### 📚 Documentation

- *(readme)* Add `--bypass-dns-resolution` information - ([cb4d9bb](https://github.com/postfinance/kubelet-csr-approver/commit/cb4d9bb6b29fb309e65bf02262dc0d3e64de0739))

### ⚙️ Miscellaneous Tasks

- Update controller-runtime and fix linting errors - ([9abb30e](https://github.com/postfinance/kubelet-csr-approver/commit/9abb30edf93112e4fe34575bf25df8404d105dd3))
- Update changelog with 1.0.5 release - ([f935b2a](https://github.com/postfinance/kubelet-csr-approver/commit/f935b2a58470fbc07725106de5dc1bb72d20e0c1))

### Build

- *(deps)* Bump k8s.io/api from 0.27.3 to 0.27.4 - ([ac4b5f3](https://github.com/postfinance/kubelet-csr-approver/commit/ac4b5f362e98627c1fd723b3e036018fa48fc7d9))
- *(deps)* Bump github.com/peterbourgon/ff/v3 from 3.3.2 to 3.4.0 - ([12b6a65](https://github.com/postfinance/kubelet-csr-approver/commit/12b6a65a3455504cfaa046c414e4cfcda15061c0))
- *(deps)* Bump k8s.io/client-go from 0.27.3 to 0.27.4 - ([0998887](https://github.com/postfinance/kubelet-csr-approver/commit/0998887715259a023dd1fae45adc681905217bf6))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.15.0 to 0.15.1 - ([984e933](https://github.com/postfinance/kubelet-csr-approver/commit/984e9336a13d6069e60206bbf3eab046ebf0d872))
- *(deps)* Bump go.uber.org/zap from 1.24.0 to 1.25.0 - ([e628fcf](https://github.com/postfinance/kubelet-csr-approver/commit/e628fcfb0e00fadc21299cb150f73a635fda7ec5))
- *(deps)* Bump golangci/golangci-lint-action from 3.6.0 to 3.7.0 - ([fa34e3a](https://github.com/postfinance/kubelet-csr-approver/commit/fa34e3aab34b23f89a62c52f589794d0de301461))
- *(deps)* Bump actions/checkout from 3 to 4 - ([5371707](https://github.com/postfinance/kubelet-csr-approver/commit/5371707e274c1b1f4b3ed04f52db62ad3ad2c93d))
- *(deps)* Bump docker/login-action from 2 to 3 - ([b8c2208](https://github.com/postfinance/kubelet-csr-approver/commit/b8c2208e59eed80ff604ddc3901531707d97941a))
- *(deps)* Bump go.uber.org/zap from 1.25.0 to 1.26.0 - ([6c679f5](https://github.com/postfinance/kubelet-csr-approver/commit/6c679f54f73bd51bac3b5e35735e681c57131254))
- *(deps)* Bump k8s.io/client-go from 0.27.4 to 0.28.2 - ([fd6a0dd](https://github.com/postfinance/kubelet-csr-approver/commit/fd6a0dd3b17080d3b88dca857b4c982ac9e9435d))


## [1.0.4] - 2023-07-22

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.3...v1.0.4

### 🐛 Bug Fixes

- Invalid syntax in ServiceMonitor helm template - ([f9a74f5](https://github.com/postfinance/kubelet-csr-approver/commit/f9a74f50610698284048375aaaf7b2bf17446dc4))




## New Contributors
* @networkhermit made their first contribution in [#173](https://github.com/postfinance/kubelet-csr-approver/pull/173)## [1.0.3] - 2023-07-20

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.2...v1.0.3

### Build

- *(deps)* Bump helm/kind-action from 1.7.0 to 1.8.0 - ([b3b0e00](https://github.com/postfinance/kubelet-csr-approver/commit/b3b0e008978941389710607ac2c3025f2756b91a))
- Add .github/release.yml - ([fa99d97](https://github.com/postfinance/kubelet-csr-approver/commit/fa99d97c4f5c67c45623c25d8ef2cedbaca60891))


## [1.0.2] - 2023-07-13

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.1...v1.0.2

### 🚀 Features

- Add flag for prometheus-operator SMon - ([648ec9f](https://github.com/postfinance/kubelet-csr-approver/commit/648ec9fe67c619e76f2ecb26986d91d6493a3ee8))
- Stronger security context defaults - ([dec49d3](https://github.com/postfinance/kubelet-csr-approver/commit/dec49d3636dbbf68f76351056d2257ae537074d7))

### 🐛 Bug Fixes

- *(helm)* Missing `and` in ServiceMonitor template - ([5452898](https://github.com/postfinance/kubelet-csr-approver/commit/5452898e0c6d1b13dabef6e93c014aa7dcb325b4))

### 📚 Documentation

- *(chart)* Add description and README - ([2aee55d](https://github.com/postfinance/kubelet-csr-approver/commit/2aee55d9d3531cf36bf730b401b23cb55594565f))
- Fix coverage badge in README.md - ([ed18e55](https://github.com/postfinance/kubelet-csr-approver/commit/ed18e551fcf967ac3003f420e598c8893a2a1ce1))

### ⚙️ Miscellaneous Tasks

- Omit deprecated node-role.kubernetes.io/master taint - ([58c08cf](https://github.com/postfinance/kubelet-csr-approver/commit/58c08cfd3c3b3db265f40d76619498f88f74cd6b))

### Build

- *(deps)* Bump github.com/stretchr/testify from 1.8.3 to 1.8.4 - ([849aa42](https://github.com/postfinance/kubelet-csr-approver/commit/849aa42e3e59aa8481da5e0fd029ec1b8e5af40d))
- *(deps)* Bump golangci/golangci-lint-action from 3.4.0 to 3.5.0 - ([1b0b018](https://github.com/postfinance/kubelet-csr-approver/commit/1b0b01839d6a07d83711897419724d82bd571ef7))
- *(deps)* Bump github.com/peterbourgon/ff/v3 from 3.3.1 to 3.3.2 - ([44a6af0](https://github.com/postfinance/kubelet-csr-approver/commit/44a6af0cebadf8f40315394a03aec5269d3a4921))
- *(deps)* Bump k8s.io/api from 0.27.2 to 0.27.3 - ([4437965](https://github.com/postfinance/kubelet-csr-approver/commit/4437965280cdc17864e50a0ed2f6a4fc85bdc366))
- *(deps)* Bump golangci/golangci-lint-action from 3.5.0 to 3.6.0 - ([57773ba](https://github.com/postfinance/kubelet-csr-approver/commit/57773ba97cca1fa299cab3f56544583f34eb4bc6))
- *(deps)* Bump k8s.io/client-go from 0.27.2 to 0.27.3 - ([6a9f76f](https://github.com/postfinance/kubelet-csr-approver/commit/6a9f76f7496aefef6bdc89ee36b71e9bfe9fe292))




## New Contributors
* @jcpunk made their first contribution## [1.0.1] - 2023-05-31

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v1.0.0...v1.0.1

### 🚀 Features

- Enable leader-election - ([8aae096](https://github.com/postfinance/kubelet-csr-approver/commit/8aae0967b1f5f87a55be981cb37b921e832ed11a))

### 📚 Documentation

- *(README)* Add leader-election documentation - ([70e8b9b](https://github.com/postfinance/kubelet-csr-approver/commit/70e8b9b8f1ee7523b07cd2c2e6e86268dd443ccf))

### 🧪 Testing

- *(ci)* Setup-envtest with k8s 1.27.x - ([3ebd69c](https://github.com/postfinance/kubelet-csr-approver/commit/3ebd69c807511a578ff5535b437928127ce8b834))
- *(leader-election)* Default to kube-system namespace when not running in-cluster - ([161a3e0](https://github.com/postfinance/kubelet-csr-approver/commit/161a3e07aac0c8767fc75c51d65c557832350268))

### ⚙️ Miscellaneous Tasks

- *(leader-election)* Make leader election optional - ([9a047cb](https://github.com/postfinance/kubelet-csr-approver/commit/9a047cb4ffc02bd5a8ae4d5720abaddb252cd985))

### Build

- *(deps)* Bump github.com/go-logr/logr from 1.2.3 to 1.2.4 - ([ede88c4](https://github.com/postfinance/kubelet-csr-approver/commit/ede88c4f5c4f4b66cd6cfae2569c8a3c43a625ba))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.14.5 to 0.14.6 - ([881b0ed](https://github.com/postfinance/kubelet-csr-approver/commit/881b0edb92336c8fe377da12c24f67c5e70f3a9e))
- *(deps)* Bump github.com/thanhpk/randstr from 1.0.4 to 1.0.5 - ([e3e65f3](https://github.com/postfinance/kubelet-csr-approver/commit/e3e65f396869431cf8766f8f0eb7e13f2a4b6b81))
- *(deps)* Bump github.com/postfinance/flash from 0.5.0 to 0.5.1 - ([74a59de](https://github.com/postfinance/kubelet-csr-approver/commit/74a59de21554bba99827d8e15b9b939cf192fc92))
- *(deps)* Bump k8s.io/apimachinery from 0.26.3 to 0.27.1 - ([9936533](https://github.com/postfinance/kubelet-csr-approver/commit/9936533dc48f334435fe30ce676886150e7aada1))
- *(deps)* Bump github.com/go-logr/zapr from 1.2.3 to 1.2.4 - ([e581bc3](https://github.com/postfinance/kubelet-csr-approver/commit/e581bc39a6ebef96f4ce44a77559cc894e3f09e5))
- *(deps)* Bump github.com/peterbourgon/ff/v3 from 3.3.0 to 3.3.1 - ([01b9f08](https://github.com/postfinance/kubelet-csr-approver/commit/01b9f085e27743ed24687c71981de78fbeef4d67))
- *(deps)* Bump helm/kind-action from 1.5.0 to 1.7.0 - ([6d8591e](https://github.com/postfinance/kubelet-csr-approver/commit/6d8591e7d7581bedfc96e20bfff1cd51acbcef1b))
- *(deps)* Bump github.com/stretchr/testify from 1.8.2 to 1.8.3 - ([d0091d2](https://github.com/postfinance/kubelet-csr-approver/commit/d0091d2d044fbfb6556558ac3ed7e6bc604bf2c6))
- *(deps)* Bump k8s.io/apimachinery from 0.27.1 to 0.27.2 - ([c5afe15](https://github.com/postfinance/kubelet-csr-approver/commit/c5afe15b66a34d9f364266fd43f0259082a9ffc2))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.14.6 to 0.15.0 - ([e7e06b6](https://github.com/postfinance/kubelet-csr-approver/commit/e7e06b65fa70dd036ea840f4cdf760ab4c9e2172))
- *(deps)* Bump github.com/thanhpk/randstr from 1.0.5 to 1.0.6 - ([9bb44d0](https://github.com/postfinance/kubelet-csr-approver/commit/9bb44d0f958be4565a35a030ab8928a65f867b1d))
- *(leader-election)* Configure Helm and manifests - ([0a20837](https://github.com/postfinance/kubelet-csr-approver/commit/0a20837d4f7e3c04671050e2882b0aa74c98e2b5))


## [1.0.0] - 2023-03-30

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.8...v1.0.0

### 🧪 Testing

- Permit specifying multiple comma-separated SAN - ([5bda15b](https://github.com/postfinance/kubelet-csr-approver/commit/5bda15bdd66ac4246b16ba3d60db62bb6175228a))

### ⚙️ Miscellaneous Tasks

- *(gh-action)* Rename misnamed check-dockerhub-token job - ([01d809e](https://github.com/postfinance/kubelet-csr-approver/commit/01d809e33a9404be9356e01c6336bfc0fd08829a))
- Switch default container registry - ([ac65893](https://github.com/postfinance/kubelet-csr-approver/commit/ac65893fdc427637ee276fe5a90d5358130d3e5f))
- Bump to Golang 1.20 and replace inet.af with netip{,x} - ([89f5330](https://github.com/postfinance/kubelet-csr-approver/commit/89f5330bafa30599bdb2b12e4679014ac50cec22))
- Automate helm versioning - ([af8d8b7](https://github.com/postfinance/kubelet-csr-approver/commit/af8d8b79cd0e5594f8031442a0a1204ddb2d9baa))

### Build

- *(deps)* Bump actions/setup-go from 3 to 4 - ([19b4ef9](https://github.com/postfinance/kubelet-csr-approver/commit/19b4ef96ce698331e1aeffec4a3e8d20b6260e49))
- *(deps)* Bump helm/chart-testing-action from 2.3.1 to 2.4.0 - ([d4807a9](https://github.com/postfinance/kubelet-csr-approver/commit/d4807a9b76b812db2e09e0ee56b9b18324562d4d))
- *(deps)* Bump k8s.io/api from 0.26.2 to 0.26.3 - ([73e7c21](https://github.com/postfinance/kubelet-csr-approver/commit/73e7c21671bcb5df72ae95a463945b7d73383932))
- *(deps)* Bump k8s.io/client-go from 0.26.2 to 0.26.3 - ([8afe9ba](https://github.com/postfinance/kubelet-csr-approver/commit/8afe9bab631e8b35d9c7aa8f0049a5ea94fc9b9b))
- *(gh-actions)* Fix helm-test job - ([9b6f499](https://github.com/postfinance/kubelet-csr-approver/commit/9b6f499a6d30f47f711b963afefde0a13177677e))
- *(ghcr.io)* Log in to ghcr.io before pushing containers - ([2a5d8c1](https://github.com/postfinance/kubelet-csr-approver/commit/2a5d8c1fe1fb524a376766a947f9b50f02daabb5))
- *(ghcr.io)* Also push for dev and tagged images - ([2cc4280](https://github.com/postfinance/kubelet-csr-approver/commit/2cc4280100e32cdb4d343a4528de128bb020772b))
- *(ko)* Remove defaultBaseImage and keep using ko's default - ([c04cdf3](https://github.com/postfinance/kubelet-csr-approver/commit/c04cdf32af53e3c8c01abfd4eba540915b8c930b))
- Specify ko default base image image - ([c2cbeab](https://github.com/postfinance/kubelet-csr-approver/commit/c2cbeabf2fba7962927894cdf25ddb36a509d73e))
- Also push containers to ghcr.io - ([67988ab](https://github.com/postfinance/kubelet-csr-approver/commit/67988ab3f7d515fe09ec196e8ad386388d62f2d2))


## [0.2.8] - 2023-03-10

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.7...v0.2.8

### 🚀 Features

- Add chart support for custom dnsConfig - ([31eeb2f](https://github.com/postfinance/kubelet-csr-approver/commit/31eeb2f6b1ad353b61ff33635902df30b7a40fd1))

### ⚙️ Miscellaneous Tasks

- Bump chart version to v0.2.8 - ([8ab8cbc](https://github.com/postfinance/kubelet-csr-approver/commit/8ab8cbc8319e8b251123e8ee9114c3f58d53bd49))




## New Contributors
* @j4m3s-s made their first contribution in [#132](https://github.com/postfinance/kubelet-csr-approver/pull/132)## [0.2.7] - 2023-03-07

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.6...v0.2.7

### ⚙️ Miscellaneous Tasks

- Add CODEOWNERS file - ([4edb3fd](https://github.com/postfinance/kubelet-csr-approver/commit/4edb3fd9c65a0fd0b952b5fdb90ae924cbc3a417))
- Bump chart version to v0.2.7 - ([a0ea8b7](https://github.com/postfinance/kubelet-csr-approver/commit/a0ea8b71cab577367e1927b8155f0767f31be7ba))

### Build

- *(deps)* Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 - ([1c3023f](https://github.com/postfinance/kubelet-csr-approver/commit/1c3023f7134a8d957b92b2b05c55a0faf0d45c9b))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.14.4 to 0.14.5 - ([4125e60](https://github.com/postfinance/kubelet-csr-approver/commit/4125e6084331db7217f2dc7fc9c697ef1c432ac1))
- *(deps)* Bump k8s.io/api from 0.26.1 to 0.26.2 - ([c24a75a](https://github.com/postfinance/kubelet-csr-approver/commit/c24a75a14055c64cacff7127c190526e5a6f140a))
- *(deps)* Bump k8s.io/client-go from 0.26.1 to 0.26.2 - ([6edfd9a](https://github.com/postfinance/kubelet-csr-approver/commit/6edfd9a4ea37e29dcc410eaa9c64f99bac0c7d8f))
- *(helm)* Bump chart version - ([0bb0a7e](https://github.com/postfinance/kubelet-csr-approver/commit/0bb0a7e2c5952ad7a0e73853704d8692505090e6))


## [0.2.6] - 2023-02-21

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.5...v0.2.6

### ⚙️ Miscellaneous Tasks

- Add profile.cov to .gitignore - ([9ddc539](https://github.com/postfinance/kubelet-csr-approver/commit/9ddc539d635baaf31fe6b6f001da3c9033e09423))

### Build

- Fix commit and ref embedding through ldflags - ([33043fd](https://github.com/postfinance/kubelet-csr-approver/commit/33043fdb75021a38c67f5ef8d7d4c2443f1a5c5a))


## [0.2.5] - 2023-02-21

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.4...v0.2.5

### 🚀 Features

- Support deploy in out-of-cluster mode (#116) - ([3c55012](https://github.com/postfinance/kubelet-csr-approver/commit/3c55012dece72a3210ae42dd34badee260c22c22))

### ⚙️ Miscellaneous Tasks

- Reformat LICENSE.md - ([2fd4edd](https://github.com/postfinance/kubelet-csr-approver/commit/2fd4edd4e9c7fc8d20b33c520948d3393f849859))
- Fix linting errors - ([aea094f](https://github.com/postfinance/kubelet-csr-approver/commit/aea094f86279a2525a9faa4fbb0302da6f838520))

### Build

- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.12.3 to 0.13.0 - ([c68183f](https://github.com/postfinance/kubelet-csr-approver/commit/c68183f09a00aad2ad83e1091ff2b9b9e235fd0f))
- *(deps)* Bump k8s api to v0.25.1 - ([afccb65](https://github.com/postfinance/kubelet-csr-approver/commit/afccb659d4eeea30d8b14f02e169a69e9cdadd49))
- *(deps)* Bump github.com/postfinance/flash from 0.3.0 to 0.4.0 - ([4c2afdc](https://github.com/postfinance/kubelet-csr-approver/commit/4c2afdcba207973ef6254a00ffa0cf16dd7bf1c4))
- *(deps)* Bump k8s.io/api from 0.25.1 to 0.25.2 - ([359cfab](https://github.com/postfinance/kubelet-csr-approver/commit/359cfab3927073e0ea40610e6b9d264f6e91239c))
- *(deps)* Bump k8s.io/client-go from 0.25.1 to 0.25.2 - ([481f790](https://github.com/postfinance/kubelet-csr-approver/commit/481f790fad7a60dff4089f59604893fcd44af580))
- *(deps)* Bump k8s.io/apimachinery from 0.25.2 to 0.25.3 - ([9c4e29b](https://github.com/postfinance/kubelet-csr-approver/commit/9c4e29b487c33929a771a6f5fd2de06c12cd63ec))
- *(deps)* Bump k8s.io/client-go from 0.25.2 to 0.25.3 - ([af22cf0](https://github.com/postfinance/kubelet-csr-approver/commit/af22cf09f633b798057458e834bc3116336e2662))
- *(deps)* Bump github.com/stretchr/testify from 1.8.0 to 1.8.1 - ([5506909](https://github.com/postfinance/kubelet-csr-approver/commit/5506909e7183b3b82357a8a8a1fd73bd363eef70))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.13.0 to 0.13.1 - ([baf60d6](https://github.com/postfinance/kubelet-csr-approver/commit/baf60d63b476cd451c3c7c4cd238c431929d92b0))
- *(deps)* Bump k8s.io/api from 0.25.3 to 0.25.4 - ([673c8be](https://github.com/postfinance/kubelet-csr-approver/commit/673c8be26834b08498c9cfa61b24448c0332bf86))
- *(deps)* Bump k8s.io/client-go from 0.25.3 to 0.25.4 - ([b6c1b8a](https://github.com/postfinance/kubelet-csr-approver/commit/b6c1b8ae491fbd1e82dd971bbcf2221301ad312a))
- *(deps)* Bump github.com/postfinance/flash from 0.4.0 to 0.5.0 - ([f974262](https://github.com/postfinance/kubelet-csr-approver/commit/f974262a1b9c48caecee5196b22d454baf3cd3a1))
- *(deps)* Bump go.uber.org/zap from 1.23.0 to 1.24.0 - ([ed18e5c](https://github.com/postfinance/kubelet-csr-approver/commit/ed18e5cf752ba675ecf39a62a8400b6238c59f58))
- *(deps)* Bump k8s.io/api from 0.25.4 to 0.26.0 - ([b9e8205](https://github.com/postfinance/kubelet-csr-approver/commit/b9e82052a67a89266f8007ce33355b7c79c017fa))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.13.1 to 0.14.0 - ([81fd215](https://github.com/postfinance/kubelet-csr-approver/commit/81fd215b9eb3929ee2f7832b7f250c683c6dcb08))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.14.0 to 0.14.1 - ([bf7ed88](https://github.com/postfinance/kubelet-csr-approver/commit/bf7ed8811e817a588cb47f41053f4575be90c1b9))
- *(deps)* Bump k8s.io/api from 0.26.0 to 0.26.1 - ([8567a85](https://github.com/postfinance/kubelet-csr-approver/commit/8567a8506c55e9d41172d801fd5b139496cad344))
- *(deps)* Bump k8s.io/client-go from 0.26.0 to 0.26.1 - ([ebbc05a](https://github.com/postfinance/kubelet-csr-approver/commit/ebbc05a5a6542b1586b3fd27c969e34d63881960))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.14.1 to 0.14.4 - ([04aed3a](https://github.com/postfinance/kubelet-csr-approver/commit/04aed3a27c04d207e5cb4bc44b31568193b6d40a))
- *(deps)* Bump azure/setup-helm from 1 to 3 (#121) - ([1baffa5](https://github.com/postfinance/kubelet-csr-approver/commit/1baffa5f47a668df28f25a5c26008e0880a96dd6))
- *(deps)* Bump helm/kind-action from 1.2.0 to 1.5.0 (#120) - ([1bba6ca](https://github.com/postfinance/kubelet-csr-approver/commit/1bba6cad0ff9df3c1c1aa77ff00fba6ed38ebcf2))
- *(deps)* Bump docker/login-action from 1 to 2 (#118) - ([d6c4815](https://github.com/postfinance/kubelet-csr-approver/commit/d6c48156dfd628dc74aea280b4e292baf6f042e5))
- *(deps)* Bump actions/setup-python from 2 to 4 (#117) - ([85a7462](https://github.com/postfinance/kubelet-csr-approver/commit/85a7462477558ecb9b9cecb537fbe447df647e20))
- *(deps)* Bump helm/chart-testing-action from 2.0.1 to 2.3.1 (#119) - ([ac74aa5](https://github.com/postfinance/kubelet-csr-approver/commit/ac74aa535a4c5c835f92403c51a0b4e07fa3ec25))
- *(ko)* Upgrade setup-ko version to fix pipelines - ([626d05a](https://github.com/postfinance/kubelet-csr-approver/commit/626d05a8bd5f941d996d3bc40146fd9650e2dba8))
- Update gh-action components and go to 1.19 - ([c47466f](https://github.com/postfinance/kubelet-csr-approver/commit/c47466fe788d63e35f5d0cd9ed1d73729821864b))
- Remove deprecated field - ([9f6ebb9](https://github.com/postfinance/kubelet-csr-approver/commit/9f6ebb9c919b2a10a5ae604ef74c28aad816a669))
- Bump k8s envtest to 1.26 - ([9ea0f2a](https://github.com/postfinance/kubelet-csr-approver/commit/9ea0f2aea4dc10d512fc5d69dae9e48a8eae7cf0))




## New Contributors
* @hansedong made their first contribution in [#116](https://github.com/postfinance/kubelet-csr-approver/pull/116)## [0.2.4] - 2022-08-30

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.3...v0.2.4

### 🚀 Features

- Permit using multiple DNS names in CSRs - ([056b5fd](https://github.com/postfinance/kubelet-csr-approver/commit/056b5fd2ed53055c0cac1afe776a6a46ad2bb4a8))
- Bypass hostname check - ([f7f44c2](https://github.com/postfinance/kubelet-csr-approver/commit/f7f44c25e4af9df39aa3eef4f1290cb2111311a4))
- Add `--bypass-hostname-check` config flag - ([f22fefb](https://github.com/postfinance/kubelet-csr-approver/commit/f22fefb0adcf9418cfa94fc1ab61ca3971e13700))

### 📚 Documentation

- Document allowed-dns-names configuration flag - ([f8e2ef2](https://github.com/postfinance/kubelet-csr-approver/commit/f8e2ef2372ff606e22272d8939b41c3efd75da4b))
- Bypass-hostname-check explanation - ([73ed5af](https://github.com/postfinance/kubelet-csr-approver/commit/73ed5af2c8f52b5f2ece94005ee176b477683f41))

### 🧪 Testing

- Bump k8s version for envtest to 1.24 - ([9c44133](https://github.com/postfinance/kubelet-csr-approver/commit/9c44133325fc7285188a3e41ac90ee1bfc307476))
- Validate bypass hostname check behaviour - ([68c6657](https://github.com/postfinance/kubelet-csr-approver/commit/68c66571432c6b1bedf91939947d80eb0c0502fa))

### ⚙️ Miscellaneous Tasks

- Refactor code with embedded struct configs - ([fcf5ead](https://github.com/postfinance/kubelet-csr-approver/commit/fcf5ead94b5e8f5929441180a9811c8bd850bb67))
- Cleanup commented test - ([7a79eae](https://github.com/postfinance/kubelet-csr-approver/commit/7a79eae8cd326a9a4e4c6ec1dc252909e4e1cd0f))
- Fix linting error for nolint comment - ([9562c44](https://github.com/postfinance/kubelet-csr-approver/commit/9562c44f67dfc7ab11130b7e7e9a955f95a0e7b2))

### Build

- *(deps)* Bump github.com/stretchr/testify from 1.7.5 to 1.8.0 - ([ae86c8b](https://github.com/postfinance/kubelet-csr-approver/commit/ae86c8bfae862c1517eef46780fe249b552cdb0e))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.12.2 to 0.12.3 - ([667e7cb](https://github.com/postfinance/kubelet-csr-approver/commit/667e7cb14f89e8ee955ec5c13681c01e576cd15a))
- *(deps)* Bump k8s.io/api from 0.24.2 to 0.24.3 - ([f0602a9](https://github.com/postfinance/kubelet-csr-approver/commit/f0602a9c2b8fa260ccb5183c0bbde02f0b74b23a))
- *(deps)* Bump k8s.io/client-go from 0.24.2 to 0.24.3 - ([61a05a1](https://github.com/postfinance/kubelet-csr-approver/commit/61a05a1e7c3abd0930a50cec270df5763cdf2033))
- *(deps)* Bump github.com/peterbourgon/ff/v3 from 3.1.2 to 3.3.0 - ([d5dab5b](https://github.com/postfinance/kubelet-csr-approver/commit/d5dab5bf99de4956b8693cb51ba3306ad7675c7e))
- *(deps)* Bump go.uber.org/zap from 1.21.0 to 1.22.0 - ([65fe494](https://github.com/postfinance/kubelet-csr-approver/commit/65fe49479df0783a4658f8e2a5f3fe960a12453c))
- *(deps)* Bump k8s.io/client-go from 0.24.3 to 0.24.4 - ([e180b33](https://github.com/postfinance/kubelet-csr-approver/commit/e180b33c1459ab1fa6c0b8199aaaec9499744ea0))
- *(deps)* Bump inet.af/netaddr and uber.org/zap - ([5094607](https://github.com/postfinance/kubelet-csr-approver/commit/50946070bf90d4034cc99b7eb0e18aff7ae4a0e1))
- *(deps)* Bump to k8s 1.25.0 api/clientgo - ([a9437d4](https://github.com/postfinance/kubelet-csr-approver/commit/a9437d43ed97acaa26db8adc49f47d86d8ec32f2))
- *(gh-actions)* Bump go-version and golangci-lint - ([a92b62c](https://github.com/postfinance/kubelet-csr-approver/commit/a92b62cd845e83d419d02db8aeb6794b4f077eb9))

### Deploy

- *(helm)* Bump chart config and version - ([0424e3d](https://github.com/postfinance/kubelet-csr-approver/commit/0424e3d5b84e5cde33273442421a71665b1407a7))


## [0.2.3] - 2022-06-29

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.2...v0.2.3

### 🚀 Features

- Permit CSR with non-system-node usernames - ([decc89a](https://github.com/postfinance/kubelet-csr-approver/commit/decc89afd1fa39e22a8f173d99dbc6294978c2cd))

### ⚙️ Miscellaneous Tasks

- *(helm)* Make it possible to specify logging level - ([7a663b8](https://github.com/postfinance/kubelet-csr-approver/commit/7a663b81c0c29ae53c9da3b403aa467f2fb12823))
- Document `ignore-non-system-node` and add to helm chart - ([156e4c3](https://github.com/postfinance/kubelet-csr-approver/commit/156e4c34422f016da086600a45138c1c072d4af0))

### Build

- *(deps)* Bump k8s.io/client-go from 0.23.5 to 0.23.6 - ([b91636e](https://github.com/postfinance/kubelet-csr-approver/commit/b91636e727fe9a7d6a5a3b216a5c409d082c52c1))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.11.2 to 0.12.0 - ([d85bc93](https://github.com/postfinance/kubelet-csr-approver/commit/d85bc93298090c819e9c41db97f7dd44698a5b7d))
- *(deps)* Bump k8s.io/api from 0.24.0 to 0.24.1 - ([d8a29d6](https://github.com/postfinance/kubelet-csr-approver/commit/d8a29d6fc599afc6c99ad390a778542799769490))
- *(deps)* Bump k8s.io/client-go from 0.24.0 to 0.24.1 - ([3fe5763](https://github.com/postfinance/kubelet-csr-approver/commit/3fe5763fc9a7bb55b0f92d3112580bdcefc93767))
- *(deps)* Bump controller-runtime and gopkg.in/yaml.v3 - ([8e75952](https://github.com/postfinance/kubelet-csr-approver/commit/8e759529f5bed83348899681f098e311fcec930b))
- *(deps)* Bump github.com/stretchr/testify from 1.7.1 to 1.7.2 - ([713b88c](https://github.com/postfinance/kubelet-csr-approver/commit/713b88c8aafca94ceb565f45c1dba9c19dd0058c))
- *(deps)* Bump k8s.io/api from 0.24.1 to 0.24.2 - ([b98ea1f](https://github.com/postfinance/kubelet-csr-approver/commit/b98ea1f311b92ab4e3490b5899687d7dd960f83f))
- *(deps)* Bump k8s.io/client-go from 0.24.1 to 0.24.2 - ([8de48d2](https://github.com/postfinance/kubelet-csr-approver/commit/8de48d274a93096e518ce698f055eb0ecb42641d))
- *(deps)* Bump github.com/stretchr/testify from 1.7.2 to 1.7.5 - ([23dec44](https://github.com/postfinance/kubelet-csr-approver/commit/23dec44b3eb2cccf2e37f4f6646c80b15125933c))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.12.1 to 0.12.2 - ([2d6cb52](https://github.com/postfinance/kubelet-csr-approver/commit/2d6cb524a0973920d4e2263aaac54e1319254eaf))


## [0.2.2] - 2022-04-11

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.2.1...v0.2.2

### ⚙️ Miscellaneous Tasks

- Simplify GH actions - ([dfc4154](https://github.com/postfinance/kubelet-csr-approver/commit/dfc4154f4c8b32a233530b9380915b44dcdf9caf))
- Switch to proper SemVer for Helm chart - ([7f2cb27](https://github.com/postfinance/kubelet-csr-approver/commit/7f2cb27a1775c672b256ceeaa9949f67ece94d1a))


## [0.2.0] - 2022-04-01

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.1.2...v0.2.0

### 🚀 Features

- Start implementing IP whitelist vaildation - ([6b49658](https://github.com/postfinance/kubelet-csr-approver/commit/6b496588323c9b10fd5cc1b517cb2ae88ca239c1))
- Prefix based IP whitelisting - ([accb23e](https://github.com/postfinance/kubelet-csr-approver/commit/accb23e9f158d1315ff1f0da6a6025e02d1cb639))
- Dns-resolved IPs must fall within whitelisted IP set - ([d09363b](https://github.com/postfinance/kubelet-csr-approver/commit/d09363bad1badc0b816a4df3114eebb0f23ca91f))

### 📚 Documentation

- Update helm install instruction - ([79def97](https://github.com/postfinance/kubelet-csr-approver/commit/79def978d9b74ce8c5d65f9f9b0f2f3fdd3b10c8))
- Update README.md to describe new IP whitelisting feature - ([36cbcdc](https://github.com/postfinance/kubelet-csr-approver/commit/36cbcdc289dfcbb3fb86902cac37817d78253aa1))

### 🧪 Testing

- *(helm)* Add providerIpPrefixes values.yaml test - ([7bd70c0](https://github.com/postfinance/kubelet-csr-approver/commit/7bd70c00496cb4ed4b377de8253dbd1aeddb7b67))
- Add ipv6 tests as well - ([a80342e](https://github.com/postfinance/kubelet-csr-approver/commit/a80342e6536ee60001c0b24e8fde074029ce2497))

### ⚙️ Miscellaneous Tasks

- Bump dependencies - ([27df807](https://github.com/postfinance/kubelet-csr-approver/commit/27df807d6491b3a733a653de092e6d07a8b596c5))
- Fix critical conditional handling for whitelistedIPCheck - ([1c3aaef](https://github.com/postfinance/kubelet-csr-approver/commit/1c3aaefd986e8508ec1630b83542a0e01237f74e))
- Fix documentation and implement v4 and v6 whitelisted IP tests - ([d8d9d5a](https://github.com/postfinance/kubelet-csr-approver/commit/d8d9d5ab9aecee0f982fa9bd3fe8db7bc569dd5c))
- Switch default branch to main - ([fb425f5](https://github.com/postfinance/kubelet-csr-approver/commit/fb425f587f6525986f30beecb1bcd58ccdecdcd0))
- Bump helm chart version to v0.2.0 - ([4ad6173](https://github.com/postfinance/kubelet-csr-approver/commit/4ad61732e3663f406dde7b3a2bbd8daa92e90794))

### Build

- *(deps)* Bump k8s.io/api from 0.23.1 to 0.23.2 - ([75756ac](https://github.com/postfinance/kubelet-csr-approver/commit/75756ac6140eb72e6d4c7a00d335359140374fdd))
- *(deps)* Bump k8s.io/client-go from 0.23.1 to 0.23.2 - ([dbb54a7](https://github.com/postfinance/kubelet-csr-approver/commit/dbb54a70a75d1bcb60039a4ad63dd127618a4a6c))
- *(deps)* Bump k8s.io/client-go from 0.23.4 to 0.23.5 - ([5ee999d](https://github.com/postfinance/kubelet-csr-approver/commit/5ee999d50ce11ca09c74e578c4e52e5dea35c241))
- *(deps)* Bump github.com/stretchr/testify from 1.7.0 to 1.7.1 - ([0a43422](https://github.com/postfinance/kubelet-csr-approver/commit/0a434222521fe8af7ef218d17939c9b0b9d1afba))
- *(deps)* Bump github.com/postfinance/flash from 0.2.0 to 0.3.0 - ([6ef1e62](https://github.com/postfinance/kubelet-csr-approver/commit/6ef1e62dc7406ac920221624e5d28f06c7536e81))
- *(deps)* Bump sigs.k8s.io/controller-runtime from 0.11.1 to 0.11.2 - ([32d070c](https://github.com/postfinance/kubelet-csr-approver/commit/32d070ca792528bb706b69a1d86810fe6dfe89a2))
- Feature images are tagged with branch name - ([5aaf31c](https://github.com/postfinance/kubelet-csr-approver/commit/5aaf31c3b3c65bec98c6d3b7f89798bf31f5aecc))
- Make publish-untagged push to :dev tag - ([4e7ebff](https://github.com/postfinance/kubelet-csr-approver/commit/4e7ebff6a91418bc35e65fce0d5a940c735d72e4))


## [0.1.2] - 2022-01-18

**Full Changelog**: https://github.com/postfinance/kubelet-csr-approver/compare/v0.1.1...v0.1.2

### 🚀 Features

- Implement maxExpirationSeconds check - ([158beba](https://github.com/postfinance/kubelet-csr-approver/commit/158bebaa97894d8d516288dde9450b4a7291102a))
- Implement DNS resolution bypass config flag - ([1652c90](https://github.com/postfinance/kubelet-csr-approver/commit/1652c90c265c6151c3c5454ac199716fd7daa5e6))

### 📚 Documentation

- Add build badge to the README - ([0a60929](https://github.com/postfinance/kubelet-csr-approver/commit/0a60929501efad08890f3fdc34493e9ce7ceb525))
- Improve documentation with MAX_EXPIRATION_SEC section - ([1901c6b](https://github.com/postfinance/kubelet-csr-approver/commit/1901c6b5eaec0829135c23a29f6feebcf8623c62))
- Improve parameters section with new cmdline-env options - ([0bb1705](https://github.com/postfinance/kubelet-csr-approver/commit/0bb1705d1658b7bb95c0252a172d23b734ce91bc))
- Document bypass-dns-resolution config flag - ([1af62c8](https://github.com/postfinance/kubelet-csr-approver/commit/1af62c8054ef69eeeb8641324ca86a99f5cb9dac))

### 🧪 Testing

- Improve test cases - ([47e843f](https://github.com/postfinance/kubelet-csr-approver/commit/47e843f368b28af5481fb74cc141e92ad50285c7))

### ⚙️ Miscellaneous Tasks

- *(helm)* Update chart to v0.1.2 - ([908812e](https://github.com/postfinance/kubelet-csr-approver/commit/908812e7d52a2740bb443b948a56d9c6d79d1e44))
- Fix incorrect image in deployment.yaml - ([af3061e](https://github.com/postfinance/kubelet-csr-approver/commit/af3061e772e92b76b110936c53a0bfc26a184ca6))
- Improve README.md formatting and .gitignore - ([786ac0d](https://github.com/postfinance/kubelet-csr-approver/commit/786ac0d3702b901133b8a45b1595ee29b0b5887f))
- Bump dependencies (k8s api, controller-runtime, logr) - ([02b14f6](https://github.com/postfinance/kubelet-csr-approver/commit/02b14f6a2fa2b967950d005af87838e92d794ac3))
- Add back the `-race` test flag - ([0537fe3](https://github.com/postfinance/kubelet-csr-approver/commit/0537fe3c0d9a46e76dcae43ba4fa6276661673fa))
- Add helm configuration option and improve doc - ([95bbc10](https://github.com/postfinance/kubelet-csr-approver/commit/95bbc1032feb219ed4108ad0c9ff407a22b683b7))
- Fix chart linting issues - ([7cc3eb1](https://github.com/postfinance/kubelet-csr-approver/commit/7cc3eb1627434f4594c9affb0b6e51b5b86c8af7))
- Fix helm-testing value to string - ([483d029](https://github.com/postfinance/kubelet-csr-approver/commit/483d029ccd1c3bf8c7eb228f06257fabc492f7ed))
- Quote maxExpirationSeconds env var in helm chart - ([406f6a1](https://github.com/postfinance/kubelet-csr-approver/commit/406f6a1f1b2616c3389cf70ad31cb02f1b8c41e6))
- Fix duplicate env key in helm chart - ([7555d8e](https://github.com/postfinance/kubelet-csr-approver/commit/7555d8eeba93cc56ffd424a7c1f801772e77a269))
- Refactoring - ([c12e4a3](https://github.com/postfinance/kubelet-csr-approver/commit/c12e4a353deecda196ac79fb646a652074114c03))
- Cleanup startup code and use ff/v3 for flags/environ - ([f1695f0](https://github.com/postfinance/kubelet-csr-approver/commit/f1695f0ee8a56c7a95cb25ca12cecbaac5139cbe))
- Make k8s auth providers opt-in with build tag - ([6be1ca4](https://github.com/postfinance/kubelet-csr-approver/commit/6be1ca4f9dd5e89afcfe002b620fd7a1e4d36e15))
- Refactor test package to reuse Cmd functions - ([1a6d4af](https://github.com/postfinance/kubelet-csr-approver/commit/1a6d4af5c91c133d1834236aef9bb46507e202bf))
- Add example max_expiration_sec env variable to k8s manifests - ([7708571](https://github.com/postfinance/kubelet-csr-approver/commit/7708571b6caa01f754e09572faac56f635bffe01))
- Implement test for bypassDNSResolution flag - ([4a9a9ec](https://github.com/postfinance/kubelet-csr-approver/commit/4a9a9ec6267e7852c978daebf12c5714434e86cb))
- Remove unused vars - ([3b00bbf](https://github.com/postfinance/kubelet-csr-approver/commit/3b00bbf561652194bd9371954cd4e8aed3f4ca0a))

### Build

- *(deps)* Bump k8s.io/api from 0.23.0 to 0.23.1 - ([a534ae7](https://github.com/postfinance/kubelet-csr-approver/commit/a534ae749c7b238b99515483ae935afd31d0c393))
- *(deps)* Bump k8s.io/client-go from 0.23.0 to 0.23.1 - ([0ab962a](https://github.com/postfinance/kubelet-csr-approver/commit/0ab962ad43fc9c3b6a332d5a8f40f7d05c278920))
- *(deps)* Bump github.com/go-logr/zapr from 1.2.0 to 1.2.2 - ([b765a62](https://github.com/postfinance/kubelet-csr-approver/commit/b765a62f64220537ad51833953f7ed07600a93b1))
- *(deps)* Bump go.uber.org/zap from 1.19.1 to 1.20.0 - ([a7049da](https://github.com/postfinance/kubelet-csr-approver/commit/a7049da9436d2d890def380440d00f3a476cb376))
- Make dependabot target the dev branch for PR [skip ci] - ([9049f4b](https://github.com/postfinance/kubelet-csr-approver/commit/9049f4b75ee6044c3569c554b72c52239494a806))
- Increase golangci timeout - ([302a50f](https://github.com/postfinance/kubelet-csr-approver/commit/302a50fd8d4107d74c28311d6913958b6ecb5a16))
- Bump k8s envtest to k8s 1.22 - ([dcd7ad3](https://github.com/postfinance/kubelet-csr-approver/commit/dcd7ad33eb220063a4184f5cd3879ec26a9d39db))
- Prevent `ko publish' to run on pull request - ([a7477a4](https://github.com/postfinance/kubelet-csr-approver/commit/a7477a4ab01081da52e2544c5dede215f5b44679))
- Prevent publishing without GH_TOKEN - ([a606957](https://github.com/postfinance/kubelet-csr-approver/commit/a6069573971c1cf7b5d7cc0cde87cdcfb83d714c))
- Only publish helm/container when GH_TOKEN available - ([ef8c20f](https://github.com/postfinance/kubelet-csr-approver/commit/ef8c20f70d9f9150d6feef97531cff93360bba4a))
- Prevent publishing without GH_TOKEN - ([a4fd0b8](https://github.com/postfinance/kubelet-csr-approver/commit/a4fd0b871a97b6bb40007a4b6b0a58cf26075477))


## [0.1.1] - 2021-12-01



### 📚 Documentation

- Improve the readme and add a quickstart guide - ([0346f22](https://github.com/postfinance/kubelet-csr-approver/commit/0346f22ed4032836da190c476d6c16b5b6cb79c8))

### ⚙️ Miscellaneous Tasks

- Repository clean-up and multi-platform build - ([e004aa3](https://github.com/postfinance/kubelet-csr-approver/commit/e004aa355079e64df716bce9e3f9c3f611496271))

### Build

- Add initial build workflow - ([ea54f5e](https://github.com/postfinance/kubelet-csr-approver/commit/ea54f5e7dcf259320dfbdf66d530d4140e0406b5))
- Fix the build action - ([5a55337](https://github.com/postfinance/kubelet-csr-approver/commit/5a5533768f36c99f6b9360f562e91b3174afb4e3))
- Add ko publish action - ([6b8d203](https://github.com/postfinance/kubelet-csr-approver/commit/6b8d203996b2b81bf7718d07321460f1adf823f8))
- Fix ko main.go action - ([b194536](https://github.com/postfinance/kubelet-csr-approver/commit/b194536cd0de1df697f9dd4feab0a39594762b78))
- Improve GH actions and move cmd/ folder - ([1020102](https://github.com/postfinance/kubelet-csr-approver/commit/1020102bf2fa59408e1af4f29c9e8e3af9fda1d6))
- Fix code coverage upload - ([6269ea1](https://github.com/postfinance/kubelet-csr-approver/commit/6269ea1eb3af50e262d41fe7bc3a1c23eb32f9d7))
- Correct ko build path for proper container name - ([5efd64f](https://github.com/postfinance/kubelet-csr-approver/commit/5efd64ff79aede2c01c31e5e284a90a61f1c5f0a))
- Fix local directory for ko publish - ([7c6d8eb](https://github.com/postfinance/kubelet-csr-approver/commit/7c6d8eb7538da2cdf08dd90a25fa0829e0b8249b))
- Add linting action before publishing - ([4c0c74d](https://github.com/postfinance/kubelet-csr-approver/commit/4c0c74d27336cd0a23c5e2567f3f97b227c6ed88))
- Add ldflags -X commit to ko settings - ([acbd014](https://github.com/postfinance/kubelet-csr-approver/commit/acbd014a6a2b4bdd0437137c15747edbb7008cc5))
- Add descriptive name for gh actions job steps - ([53fedac](https://github.com/postfinance/kubelet-csr-approver/commit/53fedacd189f34c9f5a90c4c4f308fc94f670702))
- Add TAG env before ko publish - ([9ceaa32](https://github.com/postfinance/kubelet-csr-approver/commit/9ceaa327c091165c8be44bc914b029e28d28b849))
- Removing race condition detection for tests - ([2c6b3cf](https://github.com/postfinance/kubelet-csr-approver/commit/2c6b3cfeb55399889a3140865d0b94f6c6c5c7c3))
- Configure dependabot to check go  modules - ([8a74be5](https://github.com/postfinance/kubelet-csr-approver/commit/8a74be53f8c02965af8b725a47d7a7a8650845d3))
- Fix mising commit sha and add proper ref - ([bb76361](https://github.com/postfinance/kubelet-csr-approver/commit/bb7636118921ea4485ccc0a0348665bd6fcf99b6))
- Fix missing TAG env variable for tagged release - ([140eba2](https://github.com/postfinance/kubelet-csr-approver/commit/140eba24c960ffdb6e78199899910bf57c770b17))
- Add coverage badge on README.md - ([f4f5c5d](https://github.com/postfinance/kubelet-csr-approver/commit/f4f5c5daddea8113f8bcf0a109340edbf3c0edb1))
- Prevent `ko publish` on forks - ([356d806](https://github.com/postfinance/kubelet-csr-approver/commit/356d80687aa91d15c2e8e558d350b9106d2a5279))
- Prevent external forks from publishing - ([b0754ed](https://github.com/postfinance/kubelet-csr-approver/commit/b0754ed5087fc27878c8353cb5f494d9490103fd))
- Prevent external forks from publishing - ([8abdc99](https://github.com/postfinance/kubelet-csr-approver/commit/8abdc99e1f9a1561ac480f6e9cead51461bfee35))
- Change helm chart releaser to official GH action - ([5ec3aeb](https://github.com/postfinance/kubelet-csr-approver/commit/5ec3aebe861d73ce7bead8648ffc9c3d9c382bd2))
- Push helm chart-releaser to separate step - ([cbf2328](https://github.com/postfinance/kubelet-csr-approver/commit/cbf2328db6e74c34b9edc31518d6ac69ba10ce41))



