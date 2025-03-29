# Base image für den Container
FROM mcr.microsoft.com/vscode/devcontainers/base:0-focal

# Setze den Arbeitsordner im Container
WORKDIR /workspace

# Installiere Node.js (für SvelteKit) und Go (für Gin)
RUN apt-get update && \
    apt-get install -y \
    nodejs \
    npm \
    golang \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Installiere SvelteKit und Go Module
RUN npm install -g svelte@latest  # Installiert Svelte global
RUN go get -u github.com/gin-gonic/gin  # Installiert Gin

# Setze den Standardbefehl beim Starten des Containers
CMD ["sleep", "infinity"]  # Container läuft weiter, um im Codespace aktiv zu bleiben
