#!/bin/bash

# Updating config files
sed -i '/PUBLIC_BACKEND_URL/c\      PUBLIC_BACKEND_URL: '$VITE_PUBLIC_BACKEND_URL'' docker-compose.yml
sed -i "/PUBLIC_BACKEND_URL/{n;s/.*/      value: ${VITE_PUBLIC_BACKEND_URL//\//\\/}/;}" kubernetes/pod-frontend.yaml

# Initalize cluster with kind
go install sigs.k8s.io/kind@v0.29.0
kind create cluster

cd frontend
npm install
# ENV -- Not used anymore
echo VITE_PUBLIC_BACKEND_URL=$VITE_PUBLIC_BACKEND_URL > .env
echo PUBLIC_BACKEND_URL=$VITE_PUBLIC_BACKEND_URL >> .env
# Workaround with config file
echo '{' > static/config.json
echo '  "VITE_PUBLIC_BACKEND_URL": "'$VITE_PUBLIC_BACKEND_URL'"' >> static/config.json
echo '}' >> static/config.json

cd ../backend
go mod tidy

echo 'Post-Installationsprozess abgeschlossen'