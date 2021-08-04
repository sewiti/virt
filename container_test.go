package virt

import "testing"

func TestOpenVZ(t *testing.T) {
	tests := getTests()
	tests["testdata/openvz"] = true
	run(t, tests, openVZ)
}

func TestLXC(t *testing.T) {
	tests := getTests()
	tests["testdata/lxc"] = true
	run(t, tests, lxc)
}

func TestDocker(t *testing.T) {
	tests := getTests()
	tests["testdata/docker"] = true
	run(t, tests, docker)
}

func TestPodman(t *testing.T) {
	tests := getTests()
	tests["testdata/podman"] = true
	run(t, tests, podman)
}

func TestWSL(t *testing.T) {
	tests := getTests()
	tests["testdata/wsl"] = true
	run(t, tests, wsl)
}
