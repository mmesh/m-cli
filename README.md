[![mmesh.io](https://github.com/mmesh/assets/blob/HEAD/images/logo/mmesh_logo_v5_180x30.png)](https://mmesh.io)

[![Discord](https://img.shields.io/discord/654291649572241408?color=%236d82cb&style=flat&logo=discord&logoColor=%23ffffff&label=Chat)](https://mmesh.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/mmesh/discussions)
[![Twitter](https://img.shields.io/badge/Follow_on_Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white)](https://twitter.com/mmesh_io)

Open source projects from [mmesh](https://mmesh.io).

# mmesh-cli

[![Go Report Card](https://goreportcard.com/badge/mmesh.dev/m-cli)](https://goreportcard.com/report/mmesh.dev/m-cli)
[![Release](https://img.shields.io/github/v/release/mmesh/m-cli?display_name=tag&style=flat)](https://github.com/mmesh/m-cli/releases/latest)
[![GitHub](https://img.shields.io/github/license/mmesh/m-cli?style=flat)](/LICENSE)

This repository contains `mmeshctl`, a tool for managing the [mmesh](https://mmesh.io) platform from the command line.

`mmeshctl` is available for a variety of Linux platforms, macOS and Windows.

## Minimun Requirements

`mmeshctl` has the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements) as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

## Getting Started

See [Quick Start](https://mmesh.io/docs/platform/getting-started/quickstart/) to learn how to start building your mmesh.

## Documentation

For the complete mmesh platform documentation visit [mmesh.io/docs](https://mmesh.io/docs/).

## Installation

### Binary Downloads

Linux, macOS and Windows binary downloads are available from the [Releases](https://github.com/mmesh/m-cli/releases) page.

You can download the pre-compiled binaries and install them with the appropriate tools.

### Linux Installation

#### Linux binary installation with curl

1. Download the latest release.

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/linux/amd64/mmeshctl"
    ```

2. Validate the binary (optional).

    Download the mmeshctl checksum file:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/linux/amd64/mmeshctl_checksum.sha256"
    ```

    Validate the mmeshctl binary against the checksum file:

    ```bash
    sha256sum --check < mmeshctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    mmeshctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    mmeshctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install `mmeshctl`.

    ```shell
    sudo install -o root -g root -m 0755 mmeshctl /usr/local/bin/mmeshctl
    ```

    > **Note**:
    > If you do not have root access on the target system, you can still install mmeshctl to the `~/.local/bin` directory:
    >
    > ```shell
    > chmod +x mmeshctl
    > mkdir -p ~/.local/bin
    > mv ./mmeshctl ~/.local/bin/mmeshctl
    > # and then append (or prepend) ~/.local/bin to $PATH
    > ```

#### Package Repository

mmesh provides a package repository that contains both DEB and RPM downloads.

For DEB-based platforms (e.g. Ubuntu and Debian) run the following to set up a new APT sources.list entry and install `mmesh-cli`:

```shell
echo 'deb [trusted=yes] https://repo.mmesh.io/apt/ /' | sudo tee /etc/apt/sources.list.d/mmesh.list
sudo apt update
sudo apt install mmesh-cli
```

For RPM-based platforms (e.g. RHEL, CentOS) use the following to create a repo file and install `mmesh-cli`:

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/mmesh.repo
[mmesh]
name=mmesh Repository - Stable
baseurl=https://repo.mmesh.io/yum
enabled=1
gpgcheck=0
EOF
sudo yum install mmesh-cli
```

#### Homebrew installation on Linux

If you are on Linux and using [Homebrew](https://docs.brew.sh/Homebrew-on-Linux) package manager, you can install the mmesh CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install mmesh/tap/mmesh-cli
    ```

2. Test to ensure the version you installed is up-to-date:

    ```shell
    mmeshctl version show
    ```

### macOS Installation

#### macOS binary installation with curl

1. Download the latest release.

    **Intel**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/amd64/mmeshctl"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/arm64/mmeshctl"
    ```

2. Validate the binary (optional).

    Download the mmeshctl checksum file:

    **Intel**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/amd64/mmeshctl_checksum.sha256"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/arm64/mmeshctl_checksum.sha256"
    ```

    Validate the mmeshctl binary against the checksum file:

    ```console
    shasum --algorithm 256 --check mmeshctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    mmeshctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    mmeshctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Make the mmeshctl binary executable.

    ```shell
    chmod +x mmeshctl
    ```

4. Move the mmeshctl binary to a file location on your system `PATH`.

    ```shell
    sudo mkdir -p /usr/local/bin
    sudo mv mmeshctl /usr/local/bin/mmeshctl
    sudo chown root: /usr/local/bin/mmeshctl
    ```

    > **Note**: Make sure `/usr/local/bin` is in your `PATH` environment variable.

#### Homebrew installation on macOS

If you are on macOS and using [Homebrew](https://brew.sh/) package manager, you can install the mmesh CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install mmesh/tap/mmesh-cli
    ```

2. Test to ensure the version you installed is up-to-date:

    ```shell
    mmeshctl version show
    ```

### Windows Installation

#### Windows binary installation with curl

1. Open the Command Prompt as Administrator and create a folder for mmesh.

    ```shell
    mkdir 'C:\Program Files\mmesh'
    ```

2. Download the latest release into the mmesh folder.

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/windows/amd64/mmeshctl.exe"
    ```

3. Validate the binary (optional).

    Download the mmeshctl.exe checksum file:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/windows/amd64/mmeshctl.exe_checksum.sha256"
    ```

    Validate the mmeshctl.exe binary against the checksum file:

    - Using Command Prompt to manually compare CertUtil's output to the checksum file downloaded:

         ```shell
         CertUtil -hashfile mmeshctl.exe SHA256
         type mmeshctl.exe_checksum.sha256
         ```

    - Using PowerShell to automate the verification using the -eq operator to get a `True` or `False` result:

         ```powershell
         $($(CertUtil -hashfile mmeshctl.exe SHA256)[1] -replace " ", "") -eq $(type mmeshctl.exe_checksum.sha256).split(" ")[0]
         ```

4. Append or prepend the folder `C:\Program Files\mmesh` to your `PATH` environment variable.
5. Test to ensure the version of mmeshctl is the same as downloaded.

    ```shell
    mmeshctl version show
    ```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [Releases](https://github.com/mmesh/m-cli/releases) page:

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```console
COSIGN_EXPERIMENTAL=1 cosign verify mmeshdev/mmeshctl
```

## Configuration

The first time you run `mmeshctl`, you will be assisted to generate your `mmeshctl.yml`. This config file will be located by default at the `$HOME/.mmesh` directory.

See the [mmeshctl configuration reference](https://mmesh.io/docs/platform/reference/mmeshctl.yml/) to find all the configuration options.

## Usage

See usage with:

```shell
mmeshctl help
```

## Running with Docker

You can also run `mmeshctl` as a Docker container. See the example below.

Registries:

- `mmeshdev/mmeshctl`
- `ghcr.io/mmesh/mmeshctl`

Example usage:

```shell
docker run --rm -ti -v $HOME/.mmesh:/root/.mmesh:ro mmeshdev/mmeshctl help
```

## Community

Have questions, need support and or just want to talk?

Get in touch with the mmesh community!

[![Discord](https://img.shields.io/badge/Join_us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://mmesh.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/mmesh/discussions)
[![Twitter](https://img.shields.io/badge/Follow_on_Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white)](https://twitter.com/mmesh_io)

## Code of Conduct

Participation in the mmesh community is governed by the Contributor Covenant [Code of Conduct](https://github.com/mmesh/.github/blob/HEAD/CODE_OF_CONDUCT.md). Please make sure to read and observe this document.

Please make sure to read and observe this document. By participating, you are expected to uphold this code.

## License

The mmesh open source projects are licensed under the [Apache 2.0 License](/LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmmesh%2Fm-cli.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmmesh%2Fm-cli?ref=badge_large)
