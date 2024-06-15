# IPCheck
Small utility to list the current external/public IP address, implemented in Go

The code opens a simple connection to [https://ident.me](https://ident.me) and displays the returned text, which happens to be the IP address of the caller.

This may not give the expected answer when running within a VM. 

## Building
Build using standard Go tooling.
```shell
go build .
```
Alternately, tweak the link step to reduce the executable size.

```shell
go build -ldflags="-s -w" .
```
There is no side effect with using these switches in a utility that is as simple as this one. 

## Installing
For ease of use, the recommendation is to copy the executable to somewhere that is on the path.

## Running
Run the built executable or using `go run`. In both cases, the utility require no arguments.

```bash
# This assumes the executable is on the path
ipcheck
```

```shell
# This assumes we are in the directory containing the project source code
go run .
```