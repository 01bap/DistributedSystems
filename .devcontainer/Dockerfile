# Go-Image für devconatiner
FROM mcr.microsoft.com/devcontainers/go AS builder

RUN apt-get update && apt-get install -y nodejs npm
RUN apt-get install -y docker.io

RUN npm install -g svelte@latest
