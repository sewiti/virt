package virt

import (
	"bufio"
	"os"
	"strings"
)

// IsContainer reports whetver application is running in a container.
func IsContainer() bool {
	const root = ""
	return openVZ(root) || lxc(root) || docker(root) || podman(root) || wsl(root)
}

// openVZ checks for OpenVZ / Virtuozzo container.
//
// 	"/proc/vz" // always exists if OpenVZ kernel is running (inside and outside container)
// 	"/proc/bc" // exists on node, but not inside container.
func openVZ(root string) bool {
	stat, err := os.Stat(root + "/proc/vz")
	if err != nil || !stat.IsDir() {
		return false
	}
	_, err = os.Stat(root + "/proc/bc")
	return os.IsNotExist(err)
}

// lxc checks for LXC container.
func lxc(root string) bool {
	f, err := os.Open(root + "/proc/1/environ")
	if err != nil {
		return false
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\000')
		if err != nil {
			return false
		}
		if strings.HasPrefix(line, "container=") {
			return true
		}
	}
}

// docker checks for Docker container.
func docker(root string) bool {
	stat, err := os.Stat(root + "/.dockerinit")
	if err == nil && !stat.IsDir() {
		return true
	}
	stat, err = os.Stat(root + "/.dockerenv")
	return err == nil && !stat.IsDir()
}

// podman checks for Podman container.
func podman(root string) bool {
	stat, err := os.Stat(root + "/run/.containerenv")
	return err == nil && !stat.IsDir()
}

// wsl checks for WSL container.
func wsl(root string) bool {
	f, err := os.Open(root + "/proc/sys/kernel/osrelease")
	if err != nil {
		return false
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')
	if err != nil {
		return false
	}
	// https://github.com/Microsoft/WSL/issues/423#issuecomment-221627364
	return strings.Contains(line, "Microsoft") || strings.Contains(line, "WSL")
}
