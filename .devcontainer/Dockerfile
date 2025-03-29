# Verwende ein grundlegendes Go-Image
FROM golang:1.19 AS go-builder

# Installiere Node.js
FROM node:latest

# Setze Arbeitsverzeichnis
WORKDIR /workspace

# Kopiere Go-Tools vom go-builder
COPY --from=go-builder /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# Installiere Svelte CLI
RUN npm install -g degit

# Installiere Gin (optional, du kannst auch go module verwenden)
RUN go get -u github.com/gin-gonic/gin

# Installiere Abhängigkeiten für das Frontend
WORKDIR /workspace/src
COPY src/package.json src/package-lock.json ./
RUN npm install

# Installiere Abhängigkeiten für das Backend
WORKDIR /workspace/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Setze das Arbeitsverzeichnis wieder auf das Projektverzeichnis
WORKDIR /workspace