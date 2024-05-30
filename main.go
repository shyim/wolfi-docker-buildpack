package main

import (
	"log"
	"os"
	"path"
	"strings"
	"wolfi-docker-buildpack/buildpack"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <project>", os.Args[0])
	}

	projectPath := os.Args[1]

	generated, err := buildpack.GenerateImageFile(projectPath)

	if err != nil {
		log.Fatalf("Error generating image: %s", err)
	}

	if err := os.WriteFile(path.Join(projectPath, "Dockerfile"), []byte(generated.Dockerfile), 0644); err != nil {
		log.Fatalf("Error writing Dockerfile: %s", err)
	}

	if err := os.WriteFile(path.Join(projectPath, ".dockerignore"), []byte(strings.Join(generated.Dockerignore, "\n")), 0644); err != nil {
		log.Fatalf("Error writing .dockerignore: %s", err)
	}
}
