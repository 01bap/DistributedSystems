cd frontend
echo "$1"
echo '{' > static/config.json
echo '  "VITE_PUBLIC_BACKEND_URL": "'$1'"' >> static/config.json
echo '}' >> static/config.json