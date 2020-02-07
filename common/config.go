package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	file2 "myTool/file"
	"os"
	"path/filepath"
	"strings"
)

var VideoWaterCon *Config

type Config struct {
	VideoPath string `toml:"videoPath"`
	System int `toml:"system"`
	Task int `toml:"task"`
	Output string `toml:"output"`
	AppId string `toml:"appId"`
	CutSection CutSection
	Snip Snip
	Format Format
	FrameRate FrameRate
	BitRate BitRate
	CutFront CutFront	`toml:"cutFront"`
	CutBack CutBack	`toml:"cutBack"`
	Crop   Crop
	Crop1   Crop1 `toml:"crop1"`
	ClearWater ClearWater
	ClearWater1 ClearWater `toml:"clearWater1"`
	Mirror Mirror
	Resolution Resolution
	Compress Compress
	WaterText WaterText
	RunWaterText RunWaterText `toml:"RunWaterText"`
	WaterImage WaterImage
	Speed Speed
	FilmTitle FilmTitle
	FilmEnd FilmEnd

}


type CutSection struct {
	SectionPath string `toml:"sectionPath"`
	Switch bool
	Duration int
}


type Snip struct {
	Switch bool
	T int
	R int
}

type Format struct {
	Switch bool
	Form string
}

type FrameRate struct {
	Switch bool
	Value string
}

type BitRate struct {
	Switch bool
	Value string
}

type CutFront struct {
	Switch bool
	Value int
}

type CutBack struct {
	Switch bool
	Value int
}

type Crop struct {
	Switch bool
	Start int64
	Duration int64
	X int
	Y int
	W int
	H int
}

type Crop1 struct {
	Switch bool
	Start int64
	Duration int64
	Left int
	Right int
	Top int
	Bottom int
}

type ClearWater struct {
	Switch bool
	X int
	Y int
	W int
	H int
}

type Mirror struct {
	Switch bool
	Direction string
}

type Resolution struct {
	Switch bool
	W int
	H int
}

type Compress struct {
	Switch bool
	Preset string
	Crf int
}

type WaterText struct {
	Switch bool
	Content string
	Path string
	Size int
	Color string
	Alpha float32
	Style int
	Sp1 int
	Sp2 int
}

type RunWaterText struct {
	Switch  bool
	Content string
	Path    string
	Size    int
	Color   string
	IsTop   bool `toml:"isTop"`
	LeftToRight   bool `toml:"leftToRight"`
	Sp      int
}

type WaterImage struct {
	Switch bool
	Path string
	Style int
	Sp1 int
	Sp2 int
}

type Speed struct {
	Switch bool
	V string
}

type FilmTitle struct {
	Switch bool
	Path   string
}

type FilmEnd struct {
	Switch bool
	Path   string
}


func ReadConfig(file string) *Config  {
	if VideoWaterCon != nil {
		return VideoWaterCon
	}

	conPath := ""
	cur, _ := os.Getwd()
	if file == "" {
		conPath = filepath.Join(cur, "config.toml")

		if file2.PathExist(conPath) == false {
			conPath = os.ExpandEnv("$HOME") + "/Desktop/vm/config.toml"
		}

		file = conPath

	}

	_, err := toml.DecodeFile(file, &VideoWaterCon)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if strings.HasPrefix(VideoWaterCon.VideoPath,"./") {

	} else {
		VideoWaterCon.VideoPath = filepath.Join(cur,VideoWaterCon.VideoPath)
	}

	return VideoWaterCon

}