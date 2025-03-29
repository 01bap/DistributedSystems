# Verwende das Go-Image
FROM golang:1.19 AS builder

# Setze Arbeitsverzeichnis
WORKDIR /workspace/backend

# Kopiere Go-Moduldateien
COPY backend/go.mod backend/go.sum ./

# Installiere die Abhängigkeiten
RUN go mod download

# Kopiere den Rest des Codes
COPY backend/ .

# Baue die Anwendung
RUN go build -o /workspace/backend/app

# Verwende ein Node.js-Image für das Frontend
FROM node:latest

# Setze Arbeitsverzeichnis für das Frontend
WORKDIR /workspace/frontend

# Kopiere und installiere Node.js-Abhängigkeiten
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Kopiere den Rest des Frontend-Codes
COPY frontend/ .

# Setze das Arbeitsverzeichnis zurück
WORKDIR /workspace

# Führe die Backend-App aus (optional)
CMD ["/workspace/backend/app"]
