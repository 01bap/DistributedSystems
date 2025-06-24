# Go-Image für devconatiner
FROM mcr.microsoft.com/devcontainers/go AS builder


# Kubernetes

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
RUN echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
RUN sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
RUN kubectl version --client


# Kind for kubernetes

#RUN go install sigs.k8s.io/kind@v0.29.0 && kind create cluster


# Node and svelte dependencies

RUN apt-get update && apt-get install -y nodejs npm
RUN apt-get install -y docker.io

RUN npm install -g svelte@latest
