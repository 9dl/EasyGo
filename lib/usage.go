package easygo

import (
	"context"
	"fmt"
	"time"

	helpers "github.com/9dl/EasyGo/helpers"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getUsage() (float64, float64, error) {
	info, err := host.Info()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get host information: %v", err)
	}

	var cpuUsage float64
	var memoryUsage float64

	switch info.OS {
	case "linux":
		cpuPercent, err := cpu.PercentWithContext(context.Background(), time.Second, false)
		if err != nil {
			helpers.HandleError(err)
		}
		cpuUsage = cpuPercent[0]

		memory, err := mem.VirtualMemory()
		if err != nil {
			helpers.HandleError(err)
		}
		memoryUsage = float64(memory.Used) / float64(memory.Total) * 100

	case "windows":
		cpuPercent, err := cpu.PercentWithContext(context.Background(), time.Second, true)
		if err != nil {
			helpers.HandleError(err)
		}
		cpuUsage = cpuPercent[0]

		memory, err := mem.VirtualMemory()
		if err != nil {
			helpers.HandleError(err)
		}
		memoryUsage = float64(memory.Used) / float64(memory.Total) * 100

	default:
		helpers.HandleError(fmt.Errorf("Unsupported OS"))
	}

	return cpuUsage, memoryUsage, nil
}

/*
func main() {
	cpuUsage, memoryUsage, err := getUsage()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Printf("Memory Usage: %.2f%%\n", memoryUsage)
}
*/
