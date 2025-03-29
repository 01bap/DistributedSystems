# Basis-Image für den Container
FROM mcr.microsoft.com/vscode/devcontainers/base:0-focal

# Setze den Arbeitsordner im Container
WORKDIR /workspace

# Update Node.js auf v18 und installiere es
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# Installiere Go
RUN apt-get update && apt-get install -y golang

# Installiere SvelteKit und Go-Modul für Gin
RUN npm install -g svelte@latest
RUN go get -u github.com/gin-gonic/gin

# Setze den Standardbefehl beim Starten des Containers
CMD ["sleep", "infinity"]  # Container bleibt offen
