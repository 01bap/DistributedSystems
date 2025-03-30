# Verwende das Go-Image
FROM mcr.microsoft.com/devcontainers/go AS builder


# Installiere Air für Hot-Reloading
RUN go install github.com/cosmtrek/air@latest

COPY ./backend/go.mod ./backend/go.sum ./backend/
# Installiere die Abhängigkeiten
RUN go mod download

COPY ./backend ./backend

# Standardbefehl: Starte die App mit Air für Live-Reload
# CMD ["air"]

# SvelteKit (not really needed)
# FROM cucb/sveltekit

# RUN npm install -g svelte@latest


WORKDIR /workspace