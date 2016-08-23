package main

import (
	"strconv"
)

// Convert 0...100 to 32..0
func dec2iosVolume(vol float32) float32 {
	return -(30.0 * float32(100-vol)) / 100.0
}

func ios2decVolume(vol float32) float32 {
	if vol < -30 {
		return 0
	}
	if vol > 0 {
		return 100
	}
	return 100 + (vol*100)/30
}

func toRaopVolume(cmd string) (float32, error) {
	v, err := strconv.ParseFloat(cmd, 32)
	if err != nil {
		return 0, err
	}
	return dec2iosVolume(float32(v)), nil
}
