package main

import (
	"fmt"
	"log"
	"os/exec"
)

func installFirecracker() {
	fmt.Println("Installing firecracker...")

	_, err := exec.Command("sudo", "apt-get", "update").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("sudo", "apt-get", "install", "-y", "git", "curl", "wget", "libseccomp-dev", "gcc", "g++", "make").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("git", "clone", "https://github.com/firecracker-microvm/firecracker.git").Output()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("bash", "-c", "cd firecracker && tools/devtool build")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Firecracker installed successfully.")
}

func installRunc() {
	fmt.Println("Installing runc...")

	_, err := exec.Command("sudo", "apt-get", "update").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("sudo", "apt-get", "install", "-y", "git", "golang-go", "make").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("git", "clone", "https://github.com/opencontainers/runc.git").Output()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("bash", "-c", "cd runc && make runc")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("runc installed successfully.")
}

func installYouki() {
	fmt.Println("Installing youki...")

	// Install necessary dependencies
	_, err := exec.Command("sudo", "apt-get", "update").Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("sudo", "apt-get", "install", "-y", "git", "curl", "build-essential", "libseccomp-dev", "pkg-config", "libtool", "autoconf", "automake").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Clone youki repository
	_, err = exec.Command("git", "clone", "https://github.com/containers/youki.git").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Build youki
	cmd := exec.Command("bash", "-c", "cd youki && make")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("youki installed successfully.")
}
