// Package virt is a library that helps detect if application is running in a
// virtual machine, container or bare-metal.
package virt

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

// IsVM reports whetver application is running in a virtual machine.
// Uses cpuid as primary source, filesystem as fallback.
func IsVM() bool {
	const root = ""
	// Check cpuid first
	switch cpuid.CPU.VendorID {
	case cpuid.Bhyve, cpuid.KVM, cpuid.MSVM, cpuid.VMware, cpuid.XenHVM:
		return true
	}
	return cpuid.CPU.VM() || linuxVServer(root) || uml(root) ||
		powerVMLx86(root) || zvm(root) || xen(root) || qemuKVM(root)
}

// getVendorID returns `vendor_id` value from `/proc/cpuinfo`.
func getVendorID(root string) string {
	f, err := os.Open(root + "/proc/cpuinfo")
	if err != nil {
		return ""
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "vendor_id\t:") {
			continue
		}
		line = strings.TrimPrefix(line, "vendor_id\t:")
		return strings.TrimSpace(line)
	}
	return ""
}

// linuxVServer checks for Linux-VServer guest.
func linuxVServer(root string) bool {
	f, err := os.Open(root + "/proc/self/status")
	if err != nil {
		return false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "VxID:") {
			continue
		}
		line = strings.TrimPrefix(line, "VxID:")
		line = strings.TrimSpace(line)
		id, err := strconv.Atoi(line)
		return err == nil && id != 0
	}
	return false
}

// uml checks for UML.
func uml(root string) bool {
	return strings.HasPrefix(getVendorID(root), "User Mode Linux")
}

// powerVMLx86 checks for IBM PowerVM Lx86 Linux/x86 emulator.
func powerVMLx86(root string) bool {
	return getVendorID(root) == "PowerVM Lx86"
}

// zvm checks for z/VM.
func zvm(root string) bool {
	if getVendorID(root) == "IBM/S390" {
		return true
	}

	f, err := os.Open(root + "/proc/sysinfo")
	if err != nil {
		return false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "VM00 Control Program:") {
			continue
		}
		line = strings.TrimPrefix(line, "VM00 Control Program:")
		line = strings.TrimSpace(line)
		return strings.HasPrefix(line, "z/VM")
	}
	return false
}

// xen checks for Xen.
func xen(root string) bool {
	stat, err := os.Stat(root + "/proc/xen")
	if err == nil && stat.IsDir() {
		return true
	}

	stat, err = os.Stat(root + "/sys/hypervisor/type")
	if err == nil && !stat.IsDir() {
		return true
	}

	stat, err = os.Stat(root + "/sys/bus/xen")
	if err != nil || !stat.IsDir() {
		return false
	}

	_, err = os.Stat(root + "/sys/bus/xen-backend")
	return os.IsNotExist(err)
}

// qemuKVM checks for QEMU/KVM.
func qemuKVM(root string) bool {
	f, err := os.Open(root + "/proc/cpuinfo")
	if err != nil {
		return false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "QEMU") {
			return true
		}
	}
	return false
}
