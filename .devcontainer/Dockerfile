# Basis-Image für den Container
FROM mcr.microsoft.com/vscode/devcontainers/base:0-focal

# Setze den Arbeitsordner im Container
WORKDIR /workspace

# Update Node.js auf v18 und installiere es
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# Installiere Go (neueste stabile Version)
RUN apt-get update && apt-get install -y golang-go

# Bereinige den Go-Modul-Cache und installiere Go-Module
RUN go clean -modcache
RUN go mod tidy
RUN go get -u github.com/gin-gonic/gin

# Installiere SvelteKit global
RUN npm install -g svelte@latest

# Setze den Standardbefehl beim Starten des Containers
CMD ["sleep", "infinity"]
