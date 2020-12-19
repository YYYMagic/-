package main

import (
	ui "github.com/gizak/termui/v3"
	"log"
	"testing"
)

func TestGetMusicList(t *testing.T) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	key := "徐佳莹"
	result, _ := GetMusicList(key)
	PrintMusicList(result)
}