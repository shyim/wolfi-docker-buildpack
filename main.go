package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"wolfi-docker-buildpack/buildpack"
)

var projectDir = flag.String("dir", "", "Project directory")
var imageName = flag.String("image-name", "", "If specified it builds an docker image with given name")

func main() {
	flag.Parse()

	if *projectDir == "" {
		log.Fatalf("Usage: %s <project>", os.Args[0])
	}

	generated, err := buildpack.GenerateImageFile(*projectDir)

	if err != nil {
		log.Fatalf("Error generating image: %s", err)
	}

	if err := os.WriteFile(path.Join(*projectDir, "Dockerfile"), []byte(generated.Dockerfile), 0644); err != nil {
		log.Fatalf("Error writing Dockerfile: %s", err)
	}

	if err := os.WriteFile(path.Join(*projectDir, ".dockerignore"), []byte(strings.Join(generated.Dockerignore, "\n")), 0644); err != nil {
		log.Fatalf("Error writing .dockerignore: %s", err)
	}

	if *imageName != "" {
		docker := exec.CommandContext(context.Background(), "docker", "build", "-t", *imageName, *projectDir)
		docker.Stdin = os.Stdin
		docker.Stderr = os.Stderr
		docker.Stdout = os.Stdout

		if err := docker.Run(); err != nil {
			log.Fatalf("Could not build image: %s", err)
		}
	}
}
