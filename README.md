
Kalibrium
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)


Kalibrium is the reference full node Kalibrium implementation written in Go (golang).

## What is Kalibrium

Kalibrium represents a groundbreaking advancement in the realm of cryptocurrency, aiming to redefine transaction speeds and costs. Built upon the PHANTOM protocol—a sophisticated evolution of Nakamoto consensus—Kalibrium introduces near-instantaneous confirmations and unparalleled block times of merely one second. 

Kalibrium is evolving with the upcoming integration of smart contracts and support for Web3 technologies. This development is a step towards enhancing the platform's capabilities, allowing for more complex and programmable transactions. It's about laying the groundwork for a more interactive and automated blockchain ecosystem, where users can engage with decentralized applications and services directly on the Kalibrium network. This update is part of our commitment to advancing blockchain functionality, providing our community with the tools they need to explore new possibilities in a decentralized world.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install Kalibrium including all dependencies:

```bash
$ git clone https://github.com/kalibriumnet/kalibrium
$ cd kalibrium
$ go install . ./cmd/...
```

- Kalibrium (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

Kalibrium has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ Kalibrium
```

## Discord
Join our discord server using the following link: https://discord.gg/qmmRQ9Tz

## Issue Tracker

The [integrated github issue tracker](https://github.com/kalibriumnet/kalibrium/issues)
is used for this project.


## Documentation

The [documentation](https://github.com/kalibriumnet/docs) is a work-in-progress

## License

Kalibrium is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
