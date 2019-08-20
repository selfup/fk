package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/selfup/gosh"
)

func main() {
	switch os.Args[1] {
	case "mkdir":
		checkCmd(gosh.MkDir(os.Args[2]))
		break
	case "ls", "dir":
		dir, err := gosh.Ls(os.Args[2])
		if err != nil {
			panic(err)
		}

		lsOutput(dir)
		break
	case "wr", "write", "echo":
		mode, err := strconv.Atoi(os.Args[4])
		check(err)

		checkCmd(gosh.Wr(os.Args[2], []byte(os.Args[3]), os.FileMode(mode)))
		break
	case "cat", "rd", "read":
		checkBytes(gosh.Rd(os.Args[2]))
		break
	case "fex", "exists":
		checkBool(gosh.Fex(os.Args[2]))
		break
	default:
		panic("invalid command")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkCmd(err error) {
	if err != nil {
		panic(err)
	}
}

func checkBool(result bool) {
	if result == false {
		os.Exit(1)
	}
}

func checkBytes(bytes []byte) {
	if len(bytes) == 0 {
		os.Exit(1)
	} else {
		fmt.Print(string(bytes))
	}
}

func lsOutput(dir []os.FileInfo) {
	var paths []string

	for _, path := range dir {
		var output string

		if path.IsDir() {
			output = "d: " + path.Name()
		} else {
			output = "f: " + path.Name()
		}

		paths = append(paths, output)
	}

	fmt.Println(strings.Join(paths, "\n"))
}
