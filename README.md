# virt

[![Go Reference](https://pkg.go.dev/badge/github.com/sewiti/virt.svg)](https://pkg.go.dev/github.com/sewiti/virt)

Package `virt` is a library that helps detect if application is running in a
virtual machine, container or bare-metal.

It depends on [klauspost/cpuid](https://github.com/klauspost/cpuid) library for
getting low-level CPU flags.

Detection mechanism is based on the checks in
[chuckleb/virt-what](https://github.com/chuckleb/virt-what).

## Install

```sh
go get -u github.com/sewiti/virt
```

## Usage Example

```go
package main

import (
	"fmt"

	"github.com/sewiti/virt"
)

func main() {
	fmt.Println("Virtual machine: ", virt.IsVM())
	fmt.Println("Container:       ", virt.IsContainer())
}
```
