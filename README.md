# Parallized and Distributed Systems

A repository to practice in class on distributed systems.

# Run the application

## Backend
Execute the following commands:
    cd backend/
    go run .

## Frontend
Execute the following commands:
    cd frontend/
    npm run dev

Make sure to create an environment file .env in the frontend directory:

```
PUBLIC_BACKEND_URL="${BACKEND_URL}"
```

## Working in a dev container

When opening a github work space the project should provide the neccassary dependancies. Ports for the frontend and backend will also be forwarded but are still in private and have to be changed to public manually for security reasons.
For other vs-code extansion on setup add does in the .devcontainer/devcontainer.json.