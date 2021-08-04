package virt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTests() map[string]bool {
	return map[string]bool{
		"testdata/baremetal":               false,
		"testdata/docker":                  false,
		"testdata/kvm":                     false,
		"testdata/linux-vserver":           false,
		"testdata/lx86":                    false,
		"testdata/lxc":                     false,
		"testdata/openvz":                  false,
		"testdata/podman":                  false,
		"testdata/qemu":                    false,
		"testdata/rhel5-xen-dom0":          false,
		"testdata/rhel5-xen-domU-hvm-ia64": false,
		"testdata/rhel5-xen-domU-pv":       false,
		"testdata/uml":                     false,
		"testdata/wsl":                     false,
		"testdata/zvm":                     false,
	}
}

func run(t *testing.T, tests map[string]bool, f func(root string) bool) {
	for testdata, expected := range tests {
		t.Run(testdata, func(t *testing.T) {
			assert.Equal(t, expected, f(testdata))
		})
	}
}

func TestLinuxVServer(t *testing.T) {
	tests := getTests()
	tests["testdata/linux-vserver"] = true
	run(t, tests, linuxVServer)
}

func TestUML(t *testing.T) {
	tests := getTests()
	tests["testdata/uml"] = true
	run(t, tests, uml)
}

func TestPowerVMLx86(t *testing.T) {
	tests := getTests()
	tests["testdata/lx86"] = true
	run(t, tests, powerVMLx86)
}

func TestZVM(t *testing.T) {
	tests := getTests()
	tests["testdata/zvm"] = true
	run(t, tests, zvm)
}

func TestXen(t *testing.T) {
	tests := getTests()
	tests["testdata/rhel5-xen-dom0"] = true
	tests["testdata/rhel5-xen-domU-hvm-ia64"] = true
	tests["testdata/rhel5-xen-domU-pv"] = true
	run(t, tests, xen)
}

func TestQEMUKVM(t *testing.T) {
	tests := getTests()
	tests["testdata/kvm"] = true
	tests["testdata/qemu"] = true
	run(t, tests, qemuKVM)
}
