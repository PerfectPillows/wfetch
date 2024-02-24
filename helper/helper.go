package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"time"

	"github.com/StackExchange/wmi"
)

type Win32_ComputerSystem struct {
	TotalPhysicalMemory uint64
}

type Win32_OperatingSystem struct {
	FreePhysicalMemory     uint64
	TotalVisibleMemorySize uint64
}

type Win32_UserAccount struct {
	Name string
}

type Win32_ComputerSystem_HostInfo struct {
	Manufacturer string
	Model        string
}

type Win32_OperatingSystem_Uptime struct {
	LastBootUpTime string
}

func GetOSVersion() string {
	runtimeOS := runtime.GOOS
	return runtimeOS
}

// Win32_ComputerSystem
// TotalPhysicalMemory Units ("bytes")
// Total size of physical memory.
func GetTotalRAM() (float64, error) {
	var cs []Win32_ComputerSystem
	query := "SELECT TotalPhysicalMemory FROM Win32_ComputerSystem"

	err := wmi.Query(query, &cs)
	if err != nil {
		return 0, err
	}

	if len(cs) > 0 {
		totalMemoryBytes := cs[0].TotalPhysicalMemory
		totalMemoryGB := float64(totalMemoryBytes) / (1024 * 1024 * 1024)
		return totalMemoryGB, nil
	} else {
		return 0, err
	}
}

// Win32_OperatingSystem
// FreePhysicalMemory - Number, in kilobytes, of physical memory currently unused and available.
func GetUsedRAMAmount() (float64, error) {
	var osInfo []Win32_OperatingSystem
	query := "SELECT FreePhysicalMemory, TotalVisibleMemorySize FROM Win32_OperatingSystem"

	err := wmi.Query(query, &osInfo)
	if err != nil {
		return 0, err
	}

	if len(osInfo) > 0 {
		freePhysicalMemory := osInfo[0].FreePhysicalMemory
		totalPhysicalMemory := osInfo[0].TotalVisibleMemorySize

		usedMemory := totalPhysicalMemory - freePhysicalMemory
		usedMemoryGB := float64(usedMemory) / (1024 * 1024)

		return usedMemoryGB, nil
	} else {
		return 0, err
	}
}

func GetUserName() (string, error) {
	currentUser, err := user.Current()

	return currentUser.Username, err
}

func GetHostInfo() (string, string, error) {
	var computerSystems []Win32_ComputerSystem_HostInfo
	query := "SELECT Manufacturer, Model FROM Win32_ComputerSystem"

	err := wmi.Query(query, &computerSystems)
	if err != nil {
		return "", "", err
	}

	if len(computerSystems) > 0 {
		manufacturer := computerSystems[0].Manufacturer
		model := computerSystems[0].Model
		return manufacturer, model, nil
	} else {
		return "", "", err
	}
}

// ! - time.Parse function is not giving the expected result
func GetUptimeInfo() (string, error) {
	var bootTimeInfo []Win32_OperatingSystem_Uptime
	query := "SELECT LastBootUpTime FROM Win32_OperatingSystem"

	err := wmi.Query(query, &bootTimeInfo)
	if err != nil {
		return "", nil
	}
	// fmt.Println(bootTimeInfo[0].LastBootUpTime)
	if len(bootTimeInfo) > 0 {
		lastBootUpTime, err := time.Parse("20060102150405.999999", "20240215203430.500000")
		if err != nil {
			return "", nil
		}

		currentTime := time.Now()

		uptime := currentTime.Sub(lastBootUpTime.Local())
		return formatTime(uptime), nil
	} else {
		return "", err
	}
}

func formatTime(time time.Duration) string {
	hours := int64(time.Hours())
	minutes := int64(time.Minutes()) % 60
	seconds := int64(time.Seconds()) % 60

	res := strconv.FormatInt(hours, 10) + " hours, " + strconv.FormatInt(minutes, 10) + " minutes, " + strconv.FormatInt(seconds, 10) + " seconds"
	return res
}

func ReadOSAsciiArt() {
	wd, err := os.Getwd()
	teal := "\033[38;5;6m"
	reset := "\033[0m"

	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	file, err := os.Open(wd + "\\art\\" + "win_11_art.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(teal + scanner.Text() + reset)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}
}
