# Changelog

## 2.2.0 (2025-06-04)

Full Changelog: [v2.1.1...v2.2.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.1.1...v2.2.0)

### Features

* **api:** update via SDK Studio ([860f76b](https://github.com/goperigon/perigon-go-sdk/commit/860f76b336656e61025527a5447b39fa36c55af8))
* **api:** update via SDK Studio ([2fc33ae](https://github.com/goperigon/perigon-go-sdk/commit/2fc33ae1c3370c666c5bef5fc73f904a090eeb58))
* **client:** allow overriding unions ([93dde3a](https://github.com/goperigon/perigon-go-sdk/commit/93dde3a391ceaef1aafc1cbb9b5f4cfebb0b8afd))


### Chores

* make go mod tidy continue on error ([cec4f3e](https://github.com/goperigon/perigon-go-sdk/commit/cec4f3e26a1cc6c6e47e89a92dcc9f50b61bd912))

## 2.1.1 (2025-05-30)

Full Changelog: [v2.1.0...v2.1.1](https://github.com/goperigon/perigon-go-sdk/compare/v2.1.0...v2.1.1)

### Bug Fixes

* **client:** correctly set stream key for multipart ([72d0ce6](https://github.com/goperigon/perigon-go-sdk/commit/72d0ce64e80465193c2d884826fcfdce12ca0833))
* **client:** don't panic on marshal with extra null field ([beac890](https://github.com/goperigon/perigon-go-sdk/commit/beac890f4f4fcae989fc62802f517e06e9cfa991))
* fix error ([29eaf73](https://github.com/goperigon/perigon-go-sdk/commit/29eaf739b68e5c3ebe07daa7ba91320f456d02be))


### Chores

* **docs:** grammar improvements ([f376994](https://github.com/goperigon/perigon-go-sdk/commit/f376994468180a0620a20204e53f5c70a586b1a5))
* improve devcontainer setup ([5b7f17b](https://github.com/goperigon/perigon-go-sdk/commit/5b7f17b676b3e9027c70b312ff25dbbba2653f7b))

## 2.1.0 (2025-05-14)

Full Changelog: [v2.0.0...v2.1.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.0.0...v2.1.0)

### Features

* **client:** add support for endpoint-specific base URLs in python ([1e064b1](https://github.com/goperigon/perigon-go-sdk/commit/1e064b18946da36468ee91e50721cff6c9138217))


### Chores

* **internal:** version bump ([c9ba3f4](https://github.com/goperigon/perigon-go-sdk/commit/c9ba3f4c6a809e7249517e0fb33a9e9f452be1c5))

## 2.0.0 (2025-05-08)

Full Changelog: [v1.0.1...v2.0.0](https://github.com/goperigon/perigon-go-sdk/compare/v1.0.1...v2.0.0)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names

### Features

* **api:** update via SDK Studio ([bca7862](https://github.com/goperigon/perigon-go-sdk/commit/bca78628e16c4f1691d360cfd0841e28bbd0e150))
* **client:** experimental support for unmarshalling into param structs ([3e34248](https://github.com/goperigon/perigon-go-sdk/commit/3e34248acf3b0778995ea2cf3d887c50e5ba20e6))
* **client:** rename resp package ([4bf8999](https://github.com/goperigon/perigon-go-sdk/commit/4bf8999e1de8e551e12cd9b861a68d5d3922b1e7))


### Bug Fixes

* bad example in readme ([#6](https://github.com/goperigon/perigon-go-sdk/issues/6)) ([b34bed5](https://github.com/goperigon/perigon-go-sdk/commit/b34bed5b21b13fcc50acecdc11ce76e09e1ec1ec))
* **client:** clean up reader resources ([f8469b5](https://github.com/goperigon/perigon-go-sdk/commit/f8469b583111797eb22a1f48e23928c36063cebd))
* **client:** correctly update body in WithJSONSet ([24fbc3c](https://github.com/goperigon/perigon-go-sdk/commit/24fbc3cd10dcfa185ea044b0c95927ecdc881d07))
* **client:** improve core function names ([da48a8a](https://github.com/goperigon/perigon-go-sdk/commit/da48a8a702c0fb94233a2e0c3fbcf2f9c23ec3cb))
* **client:** unmarshal responses properly ([8255c0a](https://github.com/goperigon/perigon-go-sdk/commit/8255c0a1809bba231ff623d3ad0be2dd84e9d707))


### Chores

* **docs:** update respjson package name ([83a1435](https://github.com/goperigon/perigon-go-sdk/commit/83a143524951f69493a835acaa0dc70ab8c34e74))

## 1.0.1 (2025-05-05)

Full Changelog: [v0.1.0-alpha.3...v1.0.1](https://github.com/goperigon/perigon-go-sdk/compare/v0.1.0-alpha.3...v1.0.1)

### Chores

* configure new SDK language ([926ec97](https://github.com/goperigon/perigon-go-sdk/commit/926ec97919d30206ca36dc16542abda83fdda1de))

## 0.1.0-alpha.3 (2025-05-03)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/goperigon/perigon-go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* add examples ([98beafa](https://github.com/goperigon/perigon-go-sdk/commit/98beafab2b17d61fb58e605b70a55cb7bde9d62e))
* **api:** update via SDK Studio ([f6d9498](https://github.com/goperigon/perigon-go-sdk/commit/f6d9498fba0807541113922a1a4415c0da4a4352))
* **api:** update via SDK Studio ([3ba78e4](https://github.com/goperigon/perigon-go-sdk/commit/3ba78e49a641285a336011c212adf4bc6de220a1))
* **api:** update via SDK Studio ([4fed9e3](https://github.com/goperigon/perigon-go-sdk/commit/4fed9e34867be9cb43fb91e24d199df3f2b64e54))
* **api:** update via SDK Studio ([ae88ff1](https://github.com/goperigon/perigon-go-sdk/commit/ae88ff1fc376fcf655ae84b3ff7b556e50dc90cf))

## 0.1.0-alpha.2 (2025-05-03)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/goperigon/perigon-go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([244340a](https://github.com/goperigon/perigon-go-sdk/commit/244340a5f266846768d23aee14f73e78cac95ce1))

## 0.1.0-alpha.1 (2025-05-02)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/goperigon/perigon-go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([b66a90a](https://github.com/goperigon/perigon-go-sdk/commit/b66a90ae8156c3cd3013dfb370bbd3289d5645fa))
* **api:** update via SDK Studio ([e9f5abd](https://github.com/goperigon/perigon-go-sdk/commit/e9f5abd27f00a215210ee7b272ef09446a0df80d))


### Chores

* configure new SDK language ([3e2c3be](https://github.com/goperigon/perigon-go-sdk/commit/3e2c3bea7877599cbb14292464b43dac15abc3e9))
* update SDK settings ([7ac7a94](https://github.com/goperigon/perigon-go-sdk/commit/7ac7a94949cfb0ba00384e542b8679a965ed2633))
