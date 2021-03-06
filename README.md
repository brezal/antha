# Antha
[![GoDoc](http://godoc.org/github.com/antha-lang/antha?status.svg)](http://godoc.org/github.com/antha-lang/antha)
[![Build Status](https://travis-ci.org/antha-lang/antha.svg?branch=master)](https://travis-ci.org/antha-lang/antha)

Antha v0.5

Contents:
- Installation Instructions
  - OSX (Native)
  - Linux (Native)
  - Windows (Native)
- Checking Your Installation
- Making and Running Antha Elements
- Adding Custom Equipment Drivers
  - List of Supported Interfaces
  - Connecting the Driver to Antha
- Demo

## Installation Instructions

Antha is divided into the core language and tools in this repo and protocols
and elements which are in
[antha-lang/elements](https://github.com/antha-lang/elements). Instructions
for using both are stored here.

### OSX (Native)

First step is to install or upgrade to go 1.7. Follow the instructions at the
[Golang](http://golang.org/doc/install) site. 

Update the GOPATH:

```bash
cat<<EOF>>$HOME/.bash_profile
export GOPATH=$HOME/go
export PATH=\$PATH:$HOME/go/bin
EOF
source ~/.bash_profile
```

After you install go, if you don't have [Homebrew](http://brew.sh/), please
install it. Then, follow these steps to setup a working antha development
environment:
```bash
# Install the xcode developer tools
xcode-select --install

# Install some external dependencies
brew update
brew install pkg-config homebrew/science/glpk sqlite3 opencv

# Install antha
go get github.com/antha-lang/elements/cmd/...
```

### Linux (Native)

Depending on your Linux distribution, you may not have the most recent version
of go available from your distribution's package repository. We recommend you
[download](https://golang.org/) go directly. 

For Debian-based distributions like Ubuntu on x86_64 machines, the installation
instructions follow. If you do not use a Debian based system or if you are not
using an x86_64 machine, you will have to modify these instructions by
replacing the go binary with one that corresponds to your platform and
replacing ``apt-get`` with your package manager.
```bash
# Install go
curl -O https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz

# Add /usr/local/go/bin to the path
export PATH=$PATH:/usr/local/go/bin

# Install antha external dependencies
sudo apt-get install -y libglpk-dev libopencv-dev libsqlite3-dev git

# Now, we are ready to get antha
go get github.com/antha-lang/elements/cmd/...

# add the local go bin to the path
export PATH=$PATH:$HOME/go/bin
```

### Windows (Native)

Installing antha on Windows is significantly more involved than for OSX or
Linux. The basic steps are:

  - Setup a go development environment:
    - Install the source code manager [git](https://git-scm.com/download/win)
    - Install [go](https://golang.org/dl/)
    - Install the compiler [mingw](http://sourceforge.net/projects/mingw/files/Installer/mingw-get-setup.exe/download).
      Depending on whether you installed the 386 (32-bit) or amd64 (64-bit) version
      of go, you need to install the corresponding version of mingw.
  - Download antha external dependencies
    - Install [glpk](http://sourceforge.net/projects/winglpk/) development library and make sure that
      mingw can find it.

If this procedure sounds daunting, you can try using some scripts we developed
to automate the installation procedure on Windows.
[Download](scripts/windows/windows-install.zip), unzip them and run
``install.bat``. This will try to automatically apply the Windows installation
procedure with the default options. Caveat emptor.

## Checking Your Installation

After following the installation instructions for your machine. You can check
if Antha is working properly by running a test protocol
```bash
cd $HOME/go/src/github.com/antha-lang/elements/workflows/examples/Aliquot
antha run --workflow workflow.json --parameters parameters.yml
```

## Making and Running Antha Elements

The easiest way to start developing your own antha elements is to place them
in the ``$HOME/go/src/github.com/antha-lang/elements/an`` directory and follow the structure of the
existing elements there. Afterwards, you can compile and use your elements
with the following commands:
```bash
make -C $HOME/go/src/github.com/antha-lang/elements
antha run --workflow myworkflowdefinition.json --parameters myparameters.yml
```

## Adding Custom Equipment Drivers

In order to write a custom driver for a piece of equipment and use it with Antha, you would need:

1. Find out what device interfaces does Antha currently support (see the list below)
2. Implementing the driver against that gRPC interface
3. Connect the driver to Antha.

### List of Supported Interfaces

This list could be interpreted as a list of device functions that this version of Antha can automate.

- **LIQUID HANDLING**

  - Interface: https://github.com/antha-lang/manualLiquidHandler#implementation
  - Dummy Driver (example code): https://github.com/antha-lang/manualLiquidHandler

### Connecting the Driver to Antha

Connecting the driver is as simple as running antharun with a flag --driver [driver_tcp_port]. It could look like this:

```
antha run --workflow wf.json --parameters params.json --driver localhost:50051
```

More instructions can be found here:

- https://github.com/antha-lang/elements/blob/master/starter/AnthaAcademy/Lesson1_Sample/B_parallelruns/readme_drivers.txt
- https://github.com/antha-lang/manualLiquidHandler#antharun-as-client
