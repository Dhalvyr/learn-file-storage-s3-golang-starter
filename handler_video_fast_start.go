package main

import (
	"fmt"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	processPath := filePath + ".processing"
	cmd := exec.Command(
		"ffmpeg", "-i",
		filePath, "-c",
		"copy", "-movflags",
		"faststart", "-f",
		"mp4", processPath,
	)

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Unable to run ffmpeg: %w", err)
	}

	return processPath, nil
}
