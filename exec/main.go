// Copyright 2020 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	errCh := make(chan error)
	go LaunchServer(errCh)
	go LaunchCaddy(errCh)
	e := <-errCh
	log.Fatalf("%v\n", e)
}

func LaunchServer(e chan<- error) {
	cmd := exec.Command("/server")
	stdout, stderr, err := pipe(cmd)
	if err != nil {
		e <- err
		return
	}
	if err := cmd.Start(); err != nil {
		e <- err
		return
	}
	go forwardOutput(stdout)
	go forwardOutput(stderr)
	if err := cmd.Wait(); err != nil {
		e <- err
		return
	}
}

func LaunchCaddy(e chan<- error) {
	cmd := exec.Command("/caddy", "run", "--config", "/etc/Caddyfile", "--adapter", "caddyfile")
	stdout, stderr, err := pipe(cmd)
	if err != nil {
		e <- err
		return
	}
	if err := cmd.Start(); err != nil {
		e <- err
		return
	}
	go forwardOutput(stdout)
	go forwardOutput(stderr)
	if err := cmd.Wait(); err != nil {
		e <- err
		return
	}
}

func pipe(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	return stdout, stderr, nil
}

func forwardOutput(rc io.ReadCloser) {
	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Println(l)
	}
}
