package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type cmdInfo struct {
	pathToEnv string
	cmd       []string
}

func parseArgs() (*cmdInfo, error) {
	if len(os.Args) < 3 {
		return nil, fmt.Errorf("too few args")
	}

	pathToEnv := os.Args[1]
	cmd := os.Args[2:]

	return &cmdInfo{pathToEnv, cmd}, nil
}

func readDir(dir string) (map[string]string, error) {
	envMap := make(map[string]string)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "=") || file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		line := strings.Split(string(content), "\n")[0]

		envMap[file.Name()] = strings.TrimSpace(line)
	}

	return envMap, nil
}

func runCmd(cmd []string, env map[string]string) int {
	for k, v := range env {
		os.Setenv(k, v)
	}

	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Start(); err != nil {
		return 111
	}

	err := c.Wait()

	if e, ok := err.(*exec.ExitError); ok {
		return e.ExitCode()
	}

	if err != nil {
		return 111
	}

	return 0
}

func main() {
	info, err := parseArgs()

	if err != nil {
		log.Fatal(err)
	}

	envMap, err := readDir(info.pathToEnv)

	if err != nil {
		log.Fatal(err)
	}

	if exitCode := runCmd(info.cmd, envMap); exitCode != 0 {
		if exitCode == 111 {
			fmt.Fprintln(os.Stderr, "can not run command")
		}
		os.Exit(exitCode)
	}
}
