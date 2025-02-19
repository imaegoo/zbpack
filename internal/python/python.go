package python

import (
	"github.com/zeabur/zbpack/pkg/types"
)

// GenerateDockerfile generates the Dockerfile for Python projects.
func GenerateDockerfile(meta types.PlanMeta) (string, error) {
	installCmd := meta["install"]
	startCmd := meta["start"]
	aptDeps := meta["apt-deps"]

	dockerfile := "FROM docker.io/library/python:3.8.2-slim-buster\n"

	dockerfile += `WORKDIR /app
RUN apt-get update
RUN apt-get install ` + aptDeps + ` gcc -y
RUN rm -rf /var/lib/apt/lists/*
COPY . .
RUN ` + installCmd + `
EXPOSE 8080
CMD ` + startCmd

	return dockerfile, nil
}
