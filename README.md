# mmesh-cli

[![Release](https://img.shields.io/github/release/mmesh/m-cli.svg?style=flat)](https://github.com/mmesh/m-cli/releases/latest)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmmesh%2Fm-cli.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmmesh%2Fm-cli?ref=badge_shield)

This repository contains the `mmeshctl` CLI for managing the mmesh multi-cloud integration platform.

## Minimun Requirements

`mmeshctl` has the same [minimum requirements][] as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

[minimum requirements]: https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements

## Installation

View the [changelog](/CHANGELOG.md) for the latest updates and changes by version.

### Binary Downloads

Linux, macOS and Windows binary downloads are available from the [mmesh downloads](https://mmesh.io/downloads)
page or from each [GitHub Releases](https://github.com/mmesh/m-cli/releases) page.

Download the pre-compiled binaries and install them with the appropriate tools.

### Package Repository

mmesh also provides a package repo that contains both DEB and RPM downloads.

For DEB-based platforms (e.g. Ubuntu and Debian) run the following to setup a new APT sources.list entry and install `mmesh-cli`:

```shell
echo 'deb [trusted=yes] https://repo.mmesh.io/apt/ /' | sudo tee /etc/apt/sources.list.d/mmesh.list
sudo apt update
sudo apt install mmesh-cli
```

For RPM-based platforms (e.g. RHEL, CentOS) use the following to create a repo file and install `mmesh-cli`:

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/mmesh.repo
[influxdata]
name=mmesh Repository - Stable
baseurl=https://repo.mmesh.io/yum
enabled=1
gpgcheck=0
EOF
sudo yum install mmesh-cli
```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want, and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [releases](https://github.com/mmesh/mmesh-cli/releases) page:

    ```shell
    wget https://github.com/mmesh/m-cli/releases/download/v0.1.0/checksums.txt
    wget https://github.com/mmesh/m-cli/releases/download/v0.1.0/checksums.txt.sig
    wget https://github.com/mmesh/m-cli/releases/download/v0.1.0/checksums.txt.pem
    ```

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match with the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```shell
COSIGN_EXPERIMENTAL=1 cosign verify mmeshdev/mmeshctl
```

## Getting Started

### Generating the mmeshctl.yml config file

The first time you run `mmeshctl`, you will be assisted to generate your `mmeshctl.yml`. This config file will be located by default at the `$HOME/.mmesh` directory.

### Usage

See usage with:

```shell
mmeshctl help
```

## Running with Docker

You can also run `mmeshctl` as a Docker container. See examples below.

Registries:

- `mmeshdev/mmeshctl`
- `ghcr.io/mmesh/mmeshctl`

Example usage:

```shell
docker run --rm -ti -v $HOME/.mmesh:/root/.mmesh:ro mmeshdev/mmeshctl help
```

## Documentation

Documentation site: <https://mmesh.io/docs>

## Community

Have questions, need support and or just want to talk about mmesh?

Get in touch with the mmesh community!

[![Discord](https://img.shields.io/badge/Join_Us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://mmesh.io/discord)
[![Twitter](https://img.shields.io/badge/Follow_Us_on_twitter-1DA1F2?style=flat&logo=twitter&logoColor=white)](https://twitter.com/mmesh_io)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussion-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/mmesh/discussions)
