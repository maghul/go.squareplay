package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolumeDec2Ios(t *testing.T) {
	assert.Equal(t, float32(-30.0), dec2iosVolume(0))
	assert.Equal(t, float32(-22.5), dec2iosVolume(25))
	assert.Equal(t, float32(-15.0), dec2iosVolume(50))
	assert.Equal(t, float32(-7.5), dec2iosVolume(75))
	assert.Equal(t, float32(0.0), dec2iosVolume(100))
}

func TestVolumeIos2Dec(t *testing.T) {
	assert.Equal(t, float32(0), ios2decVolume(-40))
	assert.Equal(t, float32(0), ios2decVolume(-30))
	assert.Equal(t, float32(25), ios2decVolume(-22.5))
	assert.Equal(t, float32(50), ios2decVolume(-15))
	assert.Equal(t, float32(75), ios2decVolume(-7.5))
	assert.Equal(t, float32(100), ios2decVolume(0))
	assert.Equal(t, float32(100), ios2decVolume(10))
}
