package main

import (
	"testing"
)

func TestNewFileInfo(t *testing.T) {

}

func TestParseExtension(t *testing.T) {
	{
		samplePath := "./basename.mp4"
		result := parseExtension(samplePath)

		if result != "mp4" {
			t.Logf("fail result = %s, must be mp4", result)
			t.Fail()
		}
	}
	{
		samplePath := "./basename.bmp"
		result := parseExtension(samplePath)

		if result != "bmp" {
			t.Logf("fail result = %s, must be bmp", result)
			t.Fail()
		}
	}

	{
		samplePath := "./basename.bMp"
		result := parseExtension(samplePath)

		if result != "bmp" {
			t.Logf("fail result = %s, must be bmp", result)
			t.Fail()
		}
	}
}

func TestApi01(t *testing.T) {
	t.Logf("name=[%s]\n", "TestApi01")
}
