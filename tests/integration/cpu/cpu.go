package cpu

import "golang.org/x/sys/unix"

const CPUForTests = 0 // CPU to pin test processes to

// SetCPUs pins the current process to a specific CPU
func SetCPUs(id ...int) {
	if len(id) == 0 {
		id = append(id, CPUForTests)
	}

	cpuMask := unix.CPUSet{}
	for _, i := range id {
		cpuMask.Set(i)
	}
	_ = unix.SchedSetaffinity(0, &cpuMask)
}
