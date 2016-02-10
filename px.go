package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

// Get a printf() format truncating a string at terminal width
func getLineFormat() string {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		// stty failed
		return "%s\n"
	}

	width, err :=
		strconv.Atoi(strings.Split(strings.TrimSpace(string(out)), " ")[1])
	if err != nil {
		// Unexpected output from stty
		return "%s\n"
	}

	return fmt.Sprintf("%%.%ds\n", width)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func printProcesses(processes []*Process) {
	var pidW, userW, cpuTimeW, memPercentW, nameW int
	for _, proc := range processes {
		pidW = max(pidW, len(proc.PidString()))
		userW = max(userW, len(proc.User()))
		cpuTimeW = max(cpuTimeW, len(proc.CPUTimeString()))
		memPercentW = max(memPercentW, len(proc.MemoryPercentString()))
		nameW = max(nameW, len(proc.Name()))
	}

	procFormat := fmt.Sprintf("%%%ds %%-%ds %%%ds %%%ds %%-%ds %%s",
		pidW, userW, cpuTimeW, memPercentW, nameW)

	lineFormat := getLineFormat()
	for _, proc := range processes {
		line := fmt.Sprintf(procFormat,
			proc.PidString(),
			proc.User(),
			proc.CPUTimeString(),
			proc.MemoryPercentString(),
			proc.Name(),
			proc.Cmdline())
		fmt.Printf(lineFormat, line)
	}
}

// FIXME: This program is terribly slow (on OS X). Listing all processes takes
// over 3s. Doing the same thing with ps takes 0.05s on the same machine.
func main() {
	// FIXME: process.Pids doesn't list PID 0 (init)
	pids, _ := process.Pids()

	var processes []*Process
	for _, pid := range pids {
		psutilproc, err := process.NewProcess(pid)
		if err != nil {
			continue
		}
		processes = append(processes, NewProcess(psutilproc))
	}

	sort.Sort(ByRelevance(processes))
	printProcesses(processes)
}