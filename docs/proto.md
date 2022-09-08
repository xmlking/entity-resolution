# Protobuf


## Install

1. Install [prerequisites](docs/prerequisites.md)
2. Read [Buf](./buf.md)
3. Setup [Editor Integration](https://docs.buf.build/editor-integration)

![buf-goland.png](docs/images/goland-buf-watch.png)

## Usage

### Maintenance
Whenever you change _protos_, **format** it before commit.

```shell
buf format -w --exit-code
```

```shell
buf lint
#buf lint --error-format=config-ignore-yaml
```

### Generate
Then generate code

```shell
rm -rf gen
buf generate
```

### Publish
and publish to [Buf Schema Registry (BSR)](https://buf.build/explore)

ONE TIME ACTION:

For the fist time, go to `https://buf.build/micro` and click `Create repository` button and add  `micro/entityapis`

> **Note:** use Release process described below, where once `cog bump`, the [GitHub Action](.github/workflows/release.yml) will publish to BSR and GitHub

NOTE: `policyapis` depends of `paymentapis`

```shell
buf mod update proto/entityapis # may update proto/entityapis/buf.lock
buf push proto/entityapis
#buf push proto/entityapis --tag <TAG_NAME>
# commit changed buf.lock files.
```

Above step will publish artifacts at **Threat Zero's** [BSR](https://buf.build/threat-zero)

## Release

Following command bump **VERSION** number and push `changes` and `tag` to remote<br/>
Then, [GitHub Action](.github/workflows/release.yml) trigger `GoReleaser` process.

> NOTE: make sure you commit all changes before running this command.

```shell
### 
```shell
# dry-run: calculate the next version based on the commit types since the latest tag
cog bump --auto --dry-run 
# calculate the next version based on the commit types since the latest tag
cog bump --auto
```

* check [cog](https://docs.cocogitto.io/guide/#automatic-versioning) docs
* check [buf.build GitHub Actions](https://docs.buf.build/ci-cd/github-actions) docs
