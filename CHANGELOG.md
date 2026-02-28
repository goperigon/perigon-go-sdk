# Changelog

## 2.12.2 (2026-02-28)

Full Changelog: [v2.12.1...v2.12.2](https://github.com/goperigon/perigon-go-sdk/compare/v2.12.1...v2.12.2)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([85a09f7](https://github.com/goperigon/perigon-go-sdk/commit/85a09f781e8c15b822b5a62d4678005077a73409))


### Chores

* **internal:** move custom custom `json` tags to `api` ([11a13f6](https://github.com/goperigon/perigon-go-sdk/commit/11a13f63e7e51f756a4e2856bc475fc9e41ccc06))
* update mock server docs ([3c9db42](https://github.com/goperigon/perigon-go-sdk/commit/3c9db42096fc3f7d4563d1d449526be6ca9677ee))

## 2.12.1 (2026-02-11)

Full Changelog: [v2.12.0...v2.12.1](https://github.com/goperigon/perigon-go-sdk/compare/v2.12.0...v2.12.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([aedc958](https://github.com/goperigon/perigon-go-sdk/commit/aedc95810107bc6c96655554216447f42379bee0))

## 2.12.0 (2026-02-09)

Full Changelog: [v2.11.0...v2.12.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.11.0...v2.12.0)

### Features

* **api:** manual updates ([79253e4](https://github.com/goperigon/perigon-go-sdk/commit/79253e40df6d25fc993d71dda3126b12230b1403))
* **client:** add a convenient param.SetJSON helper ([1744116](https://github.com/goperigon/perigon-go-sdk/commit/1744116bed82d912b40d8f141434cdeee01f94f8))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([8eaabc9](https://github.com/goperigon/perigon-go-sdk/commit/8eaabc9dd0a732c33cb513b9a4db09e3b05f810f))
* **docs:** fix mcp installation instructions for remote servers ([d893fce](https://github.com/goperigon/perigon-go-sdk/commit/d893fceb2bbd4d31d69ee9677b6a343c20ca1405))


### Chores

* **internal:** update `actions/checkout` version ([3518696](https://github.com/goperigon/perigon-go-sdk/commit/35186960fe1d0a30dd59804d1a29e72c031ddbc2))

## 2.11.0 (2026-01-05)

Full Changelog: [v2.10.0...v2.11.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.10.0...v2.11.0)

### Features

* **api:** api update ([83636e0](https://github.com/goperigon/perigon-go-sdk/commit/83636e0000e80e17b5ffb57b446d4f20917add20))
* **api:** api update ([67c5e07](https://github.com/goperigon/perigon-go-sdk/commit/67c5e07ea90a2b8bdae05e69a5497ed3bcc867df))


### Chores

* add float64 to valid types for RegisterFieldValidator ([93700c1](https://github.com/goperigon/perigon-go-sdk/commit/93700c169c732bb2d7910bba9a94dcc951eb0a1e))

## 2.10.0 (2025-12-15)

Full Changelog: [v2.9.0...v2.10.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.9.0...v2.10.0)

### Features

* **api:** api update ([8c7cebe](https://github.com/goperigon/perigon-go-sdk/commit/8c7cebe8b701591f19ba43c33785bce199361ceb))
* **encoder:** support bracket encoding form-data object members ([6fc005e](https://github.com/goperigon/perigon-go-sdk/commit/6fc005ecac8e52bdf7d71fee7e3b8bac12fc0858))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([a882a2c](https://github.com/goperigon/perigon-go-sdk/commit/a882a2cbc32607f5418d781bf0f927859f42d2d9))
* rename param to avoid collision ([da5ae42](https://github.com/goperigon/perigon-go-sdk/commit/da5ae4270943192f68108ba84f8d512f52aa24fe))


### Chores

* elide duplicate aliases ([7f1b4fc](https://github.com/goperigon/perigon-go-sdk/commit/7f1b4fc4b3703e1a3940ba95f449089be822b21a))
* **internal:** codegen related update ([36f4470](https://github.com/goperigon/perigon-go-sdk/commit/36f4470421733087ff5d35d4e14015e846bb72d4))

## 2.9.0 (2025-11-14)

Full Changelog: [v2.8.1...v2.9.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.8.1...v2.9.0)

### Features

* **api:** api update ([5772320](https://github.com/goperigon/perigon-go-sdk/commit/577232077f881ef93da54ac04a533beec35c98a5))


### Chores

* bump gjson version ([c8167cf](https://github.com/goperigon/perigon-go-sdk/commit/c8167cf15570e634993cf2c78c7c4a253e6cd7b4))

## 2.8.1 (2025-11-06)

Full Changelog: [v2.8.0...v2.8.1](https://github.com/goperigon/perigon-go-sdk/compare/v2.8.0...v2.8.1)

### Chores

* **internal:** grammar fix (it's -&gt; its) ([d6506a5](https://github.com/goperigon/perigon-go-sdk/commit/d6506a547ffdb053be250d12e0abf0aec11bdacc))

## 2.8.0 (2025-10-10)

Full Changelog: [v2.7.0...v2.8.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.7.0...v2.8.0)

### Features

* **api:** api update ([edec337](https://github.com/goperigon/perigon-go-sdk/commit/edec3375c8e729e6bfeb5bcac9e387d2d611e82a))
* **client:** handle recursive schemas properly ([baafdcb](https://github.com/goperigon/perigon-go-sdk/commit/baafdcbc5caff1016aeabbc2c341184ee582f2ad))

## 2.7.0 (2025-09-26)

Full Changelog: [v2.6.0...v2.7.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.6.0...v2.7.0)

### Features

* **api:** api update ([b8ed800](https://github.com/goperigon/perigon-go-sdk/commit/b8ed800e0842bbd82e55d0080eecd3dceadb7ad4))
* **api:** api update ([ee132db](https://github.com/goperigon/perigon-go-sdk/commit/ee132db9e39d62c73f63d2949c2fb94c9409aebc))
* **api:** api update ([9292ad6](https://github.com/goperigon/perigon-go-sdk/commit/9292ad6c16b64ceef475f54e77f4ef83f794b604))
* **api:** api update ([eebf091](https://github.com/goperigon/perigon-go-sdk/commit/eebf09145d8dda17af7996bd307f79d90a9350b2))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([0779e86](https://github.com/goperigon/perigon-go-sdk/commit/0779e8619cbd2ccbb62225a83596abd83ef256ee))
* close body before retrying ([e27d271](https://github.com/goperigon/perigon-go-sdk/commit/e27d27104be3893877ea205ca54df8336b398caa))
* **internal:** unmarshal correctly when there are multiple discriminators ([0616fab](https://github.com/goperigon/perigon-go-sdk/commit/0616fabc84cd33f698cde144e6b6df29585dfbb1))
* use slices.Concat instead of sometimes modifying r.Options ([750a6b0](https://github.com/goperigon/perigon-go-sdk/commit/750a6b07b94a6446edf381f72f7e67a36e71397a))


### Chores

* bump minimum go version to 1.22 ([8bb93db](https://github.com/goperigon/perigon-go-sdk/commit/8bb93db1cf79cd513c5985b1cb4d062626168149))
* do not install brew dependencies in ./scripts/bootstrap by default ([dcc099c](https://github.com/goperigon/perigon-go-sdk/commit/dcc099c605ed0c2a726d7562159d12fae26d5104))
* **internal:** codegen related update ([e8f0c15](https://github.com/goperigon/perigon-go-sdk/commit/e8f0c15dfe03b5e4b9c0a4c0075561e2db328053))
* **internal:** improve examples ([a3ef5e8](https://github.com/goperigon/perigon-go-sdk/commit/a3ef5e819a7d8120d402591c1e520fae3682533d))
* **internal:** use some smaller example values ([1c87e17](https://github.com/goperigon/perigon-go-sdk/commit/1c87e17d9019fccd013703929e16557c63b94c87))
* update more docs for 1.22 ([6205199](https://github.com/goperigon/perigon-go-sdk/commit/620519946e724c5a794a78394078d241b34953d1))

## 2.6.0 (2025-08-18)

Full Changelog: [v2.5.0...v2.6.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.5.0...v2.6.0)

### Features

* **api:** api update ([06a4c4e](https://github.com/goperigon/perigon-go-sdk/commit/06a4c4ef53e6f40c683c221ad2f81b9b18f624e9))
* **client:** support optional json html escaping ([73538f2](https://github.com/goperigon/perigon-go-sdk/commit/73538f242c8a56d22cd959a312bbdd7a5000e82b))


### Chores

* **internal:** update comment in script ([63f0013](https://github.com/goperigon/perigon-go-sdk/commit/63f0013b653d60b0bcd5a337d04e4238a3158136))
* update @stainless-api/prism-cli to v5.15.0 ([68fc890](https://github.com/goperigon/perigon-go-sdk/commit/68fc890aa08a52d10d0b58eacc5ae0f8b7563e66))

## 2.5.0 (2025-08-01)

Full Changelog: [v2.4.1...v2.5.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.4.1...v2.5.0)

### Features

* **api:** api update ([0302dfa](https://github.com/goperigon/perigon-go-sdk/commit/0302dfa6d7470c37f76f18cbf7e7db14a28b38f2))

## 2.4.1 (2025-07-29)

Full Changelog: [v2.4.0...v2.4.1](https://github.com/goperigon/perigon-go-sdk/compare/v2.4.0...v2.4.1)

## 2.4.0 (2025-07-23)

Full Changelog: [v2.3.2...v2.4.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.3.2...v2.4.0)

### Features

* **api:** add wikipedia endpoints ([73bf2a5](https://github.com/goperigon/perigon-go-sdk/commit/73bf2a583c67f56da04009961f19f0d2f840a762))
* **api:** api update ([3c973f7](https://github.com/goperigon/perigon-go-sdk/commit/3c973f7834483a27deb8027c2f4737ae1f8d777e))

## 2.3.2 (2025-07-22)

Full Changelog: [v2.3.1...v2.3.2](https://github.com/goperigon/perigon-go-sdk/compare/v2.3.1...v2.3.2)

### Bug Fixes

* **client:** process custom base url ahead of time ([6352f4d](https://github.com/goperigon/perigon-go-sdk/commit/6352f4d07046f9c19b56e95eb50063df721c109c))

## 2.3.1 (2025-07-11)

Full Changelog: [v2.3.0...v2.3.1](https://github.com/goperigon/perigon-go-sdk/compare/v2.3.0...v2.3.1)

### Chores

* **internal:** fix lint script for tests ([47eae65](https://github.com/goperigon/perigon-go-sdk/commit/47eae65b69d6b6e60d26975a6cb6e5ab1af47cf7))
* lint tests ([9e2661a](https://github.com/goperigon/perigon-go-sdk/commit/9e2661a9384fe741dce782011d0c8281f5449b2f))
* lint tests in subpackages ([270fec6](https://github.com/goperigon/perigon-go-sdk/commit/270fec68975193089bb7790cdaf6f80116900e61))

## 2.3.0 (2025-06-28)

Full Changelog: [v2.2.0...v2.3.0](https://github.com/goperigon/perigon-go-sdk/compare/v2.2.0...v2.3.0)

### Features

* **client:** add debug log helper ([e330b9c](https://github.com/goperigon/perigon-go-sdk/commit/e330b9c5fbff9fee6d7bcc61a313c2225a14f4a0))
* **client:** add escape hatch for null slice & maps ([c19b4a7](https://github.com/goperigon/perigon-go-sdk/commit/c19b4a77484664f519e3f7a1abdbed141f60e78d))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([98641ec](https://github.com/goperigon/perigon-go-sdk/commit/98641ecd0feb09be8022033a158dfc11d7e0e8ea))


### Chores

* **ci:** enable for pull requests ([55a4d17](https://github.com/goperigon/perigon-go-sdk/commit/55a4d17c9a357bcd6f4b91c0ab06a7413e019d63))
* **ci:** only run for pushes and fork pull requests ([e07881f](https://github.com/goperigon/perigon-go-sdk/commit/e07881f8b4d79a0149242a666eaff0086cab3d07))
* fix documentation of null map ([22e2b39](https://github.com/goperigon/perigon-go-sdk/commit/22e2b39e30ba349bb027c50b908854f90594a1eb))

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
