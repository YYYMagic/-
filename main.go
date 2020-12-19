package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"log"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("食用方法: ./term-music <搜索内容>")
		os.Exit(-1)
	}
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	key := os.Args[1]
	result, _ := GetMusicList(key)
	PrintMusicList(result)
}