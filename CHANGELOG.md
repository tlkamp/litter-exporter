# [2.0.0](https://github.com/tlkamp/litter-exporter/compare/v1.5.1...v2.0.0) (2022-07-23)


### Features

* support environment variables ([#22](https://github.com/tlkamp/litter-exporter/issues/22)) ([7f90542](https://github.com/tlkamp/litter-exporter/commit/7f90542bbf885c1588e14dceee488416e0d6721d)), closes [#17](https://github.com/tlkamp/litter-exporter/issues/17)


### BREAKING CHANGES

* Pflag does not support multi-character shorthands for flags, so double hyphens are
now required for all flags being passed via CLI.

## [1.5.1](https://github.com/tlkamp/litter-exporter/compare/v1.5.0...v1.5.1) (2022-07-21)


### Bug Fixes

* **ci:** fix goreleaser again ([#20](https://github.com/tlkamp/litter-exporter/issues/20)) ([33c4ea3](https://github.com/tlkamp/litter-exporter/commit/33c4ea368cab592743e800a14fd4152de4493e00))

# [1.5.0](https://github.com/tlkamp/litter-exporter/compare/v1.4.2...v1.5.0) (2022-07-21)


### Features

* **cli:** add defaults for some args ([#19](https://github.com/tlkamp/litter-exporter/issues/19)) ([1ab4e0e](https://github.com/tlkamp/litter-exporter/commit/1ab4e0e35d8853a5136f91a57cfe128f86d4af58)), closes [#17](https://github.com/tlkamp/litter-exporter/issues/17)

## [1.4.2](https://github.com/tlkamp/litter-exporter/compare/v1.4.1...v1.4.2) (2021-12-30)


### Bug Fixes

* **api:** update api version to 1.0.3 ([93d9734](https://github.com/tlkamp/litter-exporter/commit/93d9734ff1f59443bad8ed75379342f527410eac))

## [1.4.1](https://github.com/tlkamp/litter-exporter/compare/v1.4.0...v1.4.1) (2021-04-08)


### Bug Fixes

* **litter-api:** upgrades litter-api ([594f82b](https://github.com/tlkamp/litter-exporter/commit/594f82b05d3ae6942b1c6013cbe6fbe99209646b))

# [1.4.0](https://github.com/tlkamp/litter-exporter/compare/v1.3.0...v1.4.0) (2021-04-07)


### Features

* **collector:** adds cycles after full metric ([2fb432d](https://github.com/tlkamp/litter-exporter/commit/2fb432dbce23f3b6e1cff056d37d084314d0e128))

# [1.3.0](https://github.com/tlkamp/litter-exporter/compare/v1.2.0...v1.3.0) (2021-04-06)


### Features

* **collector:** adds cycles until full metric ([010ede6](https://github.com/tlkamp/litter-exporter/commit/010ede68d368b0e4bfda97680533f7a6089898a5))

# [1.2.0](https://github.com/tlkamp/litter-exporter/compare/v1.1.0...v1.2.0) (2021-04-06)


### Features

* **collector:** adds dfi cycle count metric ([2fafbf4](https://github.com/tlkamp/litter-exporter/commit/2fafbf4c19e3261d99bafd43093985ad20832c6b))

# [1.1.0](https://github.com/tlkamp/litter-exporter/compare/v1.0.0...v1.1.0) (2021-04-06)


### Features

* **panel lock:** adds panel locked metric ([#5](https://github.com/tlkamp/litter-exporter/issues/5)) ([e9442a5](https://github.com/tlkamp/litter-exporter/commit/e9442a554b90fd59070ec493bb9e9886b0996573))

# 1.0.0 (2021-04-06)


### Features

* initial commit ([cd6e97d](https://github.com/tlkamp/litter-exporter/commit/cd6e97d1697329694ea045562f3ad4e2a368f7a8))
