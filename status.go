package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func printOSInfo() {
	color.Green("Operating System:")
	color.Green("---------------")
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
}

func printKernelVersion() {
	color.Green("Kernel Version:")
	color.Green("---------------")
	cmd := exec.Command("uname", "-r")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error retrieving kernel version:", err)
		return
	}
	fmt.Printf("Kernel Version: %s", output)
}

func printUptime() {
	color.Blue("Uptime:")
	color.Blue("---------------")
	uptime := time.Since(time.Now().Add(-1 * time.Second * time.Duration(runtimeUptime())))
	fmt.Printf("Uptime: %s\n", uptime)
}

func runtimeUptime() int64 {
	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		fmt.Println("Error retrieving uptime:", err)
		return 0
	}
	return info.Uptime
}

func printShell() {
	color.Cyan("Shell:")
	color.Cyan("---------------")
	fmt.Printf("Shell: %s\n", os.Getenv("SHELL"))
}

func printMemoryStatus() {
	color.Cyan("Memory Status:")
	color.Cyan("---------------")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Allocated memory: %d bytes\n", m.Alloc)
	fmt.Printf("Total memory allocated and not yet freed: %d bytes\n", m.TotalAlloc)
	fmt.Printf("System memory obtained from OS: %d bytes\n", m.Sys)
	fmt.Printf("Number of live objects: %d\n", m.Mallocs-m.Frees)
}

func printCPUStatus() {
	color.Yellow("CPU Status:")
	color.Yellow("---------------")
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)
}

func printOSDistro() {
	color.Green("OS Distribution:")
	color.Green("---------------")
	cmd := exec.Command("lsb_release", "-i", "-s")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error retrieving OS distribution:", err)
		return
	}
	distro := string(output)
	fmt.Printf("OS Distribution: %s", distro)
}

func printDiskStatus() {
	color.Magenta("Disk Status:")
	color.Magenta("---------------")
	var stat syscall.Statfs_t
	wd, _ := os.Getwd()
	syscall.Statfs(wd, &stat)
	blockSize := uint64(stat.Bsize)
	totalDiskSpace := stat.Blocks * blockSize
	freeDiskSpace := stat.Bfree * blockSize

	fmt.Printf("Total disk space: %s\n", humanReadableSize(totalDiskSpace))
	fmt.Printf("Free disk space: %s\n", humanReadableSize(freeDiskSpace))
}

func humanReadableSize(bytes uint64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
		tb = gb * 1024
	)

	switch {
	case bytes >= tb:
		return fmt.Sprintf("%.2f TB", float64(bytes)/tb)
	case bytes >= gb:
		return fmt.Sprintf("%.2f GB", float64(bytes)/gb)
	case bytes >= mb:
		return fmt.Sprintf("%.2f MB", float64(bytes)/mb)
	case bytes >= kb:
		return fmt.Sprintf("%.2f KB", float64(bytes)/kb)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}
