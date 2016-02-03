package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

type config struct {
	Golint string   `json:"golint"`
	Ignore []string `json:"ignore"`
}

func filter(line string, regs []*regexp.Regexp) bool {
	f := false
	for _, r := range regs {
		if r.MatchString(line) {
			f = true
		}
	}
	return f
}

func configRead(fp string) (string, []*regexp.Regexp) {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal("There was an issue reading your config file, expected it to be at ", fp)
	}
	defer file.Close()

	var r []*regexp.Regexp

	f, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	cfg := new(config)
	err = json.Unmarshal(f, &cfg)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range cfg.Ignore {
		regx, _ := regexp.Compile(line)
		r = append(r, regx)
	}

	return cfg.Golint, r
}

func byteLines(data []byte) []string {
	var strs []string

	b := bytes.NewBuffer(data)
	for line, err := b.ReadBytes(byte('\n')); err == nil; line, err = b.ReadBytes(byte('\n')) {
		strs = append(strs, string(line))
	}

	return strs
}

func main() {
	configPath := os.Getenv("HOME") + "/.golint-free"
	golintPath, regs := configRead(configPath)

	golint := exec.Command(golintPath, os.Args[1:]...)
	output, err := golint.Output()
	if err != nil {
		log.Fatal(err)
	}

	out := byteLines(output)
	for _, outputLine := range out {
		if !filter(outputLine, regs) {
			os.Stdout.Write([]byte(outputLine))
		}
	}
}
