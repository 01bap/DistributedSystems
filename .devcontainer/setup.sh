#!/bin/bash

go install sigs.k8s.io/kind@v0.29.0
kind create cluster

cd frontend
npm install
echo VITE_PUBLIC_BACKEND_URL=$VITE_PUBLIC_BACKEND_URL > .env
echo PUBLIC_BACKEND_URL=$VITE_PUBLIC_BACKEND_URL >> .env
echo '{' > public/config.json
echo '  "VITE_PUBLIC_BACKEND_URL": "$VITE_PUBLIC_BACKEND_URL"' >> public/config.json
echo '}' >> public/config.json

cd ../backend
go mod tidy

echo 'Post-Installationsprozess abgeschlossen'