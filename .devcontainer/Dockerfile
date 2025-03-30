# Go-Image für devconatiner
FROM mcr.microsoft.com/devcontainers/go AS builder

#RUN exit 1

WORKDIR /backend

# Installiere Air für Hot-Reloading
RUN go install github.com/cosmtrek/air@latest
# Installiere die Abhängigkeiten
RUN go mod download


# SvelteKit (not really needed)
# FROM cucb/sveltekit

# RUN npm install -g svelte@latest


# WORKDIR /workspace
