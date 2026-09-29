package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Unable to run ffprobe: %w", err)
	}

	type StreamType struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	type output struct {
		Streams []StreamType `json:"streams"`
	}

	var filedata output
	if err = json.Unmarshal(buffer.Bytes(), &filedata); err != nil {
		return "", fmt.Errorf("Unable to unmarshal data: %w", err)
	}

	tolerance := 0.01
	actualRatio := float64(filedata.Streams[0].Width) / float64(filedata.Streams[0].Height)
	if filedata.Streams[0].Width > filedata.Streams[0].Height {
		expectedRatio := 16.0 / 9.0
		difference := expectedRatio - actualRatio
		if difference <= tolerance && difference >= -tolerance {
			return "16:9", nil
		}
	} else {
		expectedRatio := 9.0 / 16.0
		difference := expectedRatio - actualRatio
		if difference <= tolerance && difference >= -tolerance {
			return "9:16", nil
		}
	}

	return "other", nil
}
