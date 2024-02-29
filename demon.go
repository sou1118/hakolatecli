package main

import (
	"fmt"
	"log"
	"os/exec"
)

func installNerdctl() {
	fmt.Println("Installing nerdctl...")

	_, err := exec.Command("sudo", "apt-get", "update").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("sudo", "apt-get", "install", "-y", "curl").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("bash", "-c", "curl -sL https://github.com/containerd/nerdctl/releases/download/v0.11.2/nerdctl-0.11.2-linux-amd64.tar.gz | tar xzv -C /usr/local/bin").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("nerdctl installed successfully.")
}
