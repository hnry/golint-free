package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func filter(line string, regs []*regexp.Regexp) bool {
	f := false
	for _, r := range regs {
		if r.MatchString(line) {
			f = true
		}
	}
	return f
}

func configRead(fp string) []*regexp.Regexp {
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

	for _, line := range byteLines(f) {
		l := line[:len(line)-1] // chomp \n
		if l != "" {
			regx, _ := regexp.Compile(l)
			r = append(r, regx)
		}
	}

	return r
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
	regs := configRead(configPath)

	golint := exec.Command("golint", os.Args[1:]...)
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
