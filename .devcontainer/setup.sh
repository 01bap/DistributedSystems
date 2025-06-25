#!/bin/bash

if [ "$VITE_PUBLIC_BACKEND_URL" != "$1" ];then
    echo "Warning: Global variable 'VITE_PUBLIC_BACKEND_URL' doesnt match passed parameter $1"
fi

# Updating config files
sed -i '/PUBLIC_BACKEND_URL/c\      PUBLIC_BACKEND_URL: '$1'' docker-compose.yml
sed -i "/PUBLIC_BACKEND_URL/{n;s/.*/      value: ${1//\//\\/}/;}" kubernetes/pod-frontend.yaml

# Initalize cluster with kind
go install sigs.k8s.io/kind@v0.29.0
kind create cluster

cd frontend
npm install
# ENV
echo VITE_PUBLIC_BACKEND_URL=$1 > .env
echo PUBLIC_BACKEND_URL=$1 >> .env
# Workaround with config file
echo '{' > static/config.json
echo '  "VITE_PUBLIC_BACKEND_URL": "'$1'"' >> static/config.json
echo '}' >> static/config.json

cd ../backend
go mod tidy

echo 'Post-Installationsprozess abgeschlossen'