package rpi

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const cpuTempFile = "/sys/class/thermal/thermal_zone0/temp"

func ReadCPUTemp() (float64, error) {
    bytes, err := os.ReadFile(cpuTempFile)
    if err != nil {
        return 0, fmt.Errorf("could not read temp file: %w", err)
    }

    // The temperature is represented as a string, in millidegrees Celsius.
    tempStr := string(bytes[:len(bytes)-1]) // Strip trailing newline
    tempInt, err := strconv.ParseInt(tempStr, 10, 64)
    if err != nil {
        return 0, fmt.Errorf("could not parse temp: %w", err)
    }

    // Convert millidegrees to degrees.
    temp := float64(tempInt) / 1000

    return temp, nil
}

func GetDeviceID() (string, error) {
    out, err := exec.Command("sh", "-c", "cat /proc/cpuinfo | grep Serial | cut -d ' ' -f 2").Output()

    if err != nil {
        return "", fmt.Errorf("error retrieving Device ID: %w", err)
    }

    deviceId := strings.TrimSpace(string(out))
    return deviceId, nil
}
