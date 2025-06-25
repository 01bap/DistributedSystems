# Parallized and Distributed Systems


This repository contains a full-stack cloud-native application developed for a parallized and distributed systems course . It is based on [Matthias Häussler's and Cagri Tasci's Repository](https://github.com/maeddes/hse-25-summer). It follows modern DevOps priciples and implements the [12-Factor App](https://12factor.net/) methodology for scalable, maintainable and production-ready software.

The application consists of:
- A **SvelteKit frontend** for interacting with the system via web UI
- A **Go backend using Gin** web framework for a RESTful API
- A **PostgreSQL database** for item storage

This project demonstrates containerization, REST architecture and cloud deployment readiness.

# Table of Contents
- Project Overview
- [Setup Instructions](#setup-instructions)
    - GitHub Codespaces
    - Frontend
    - Backend
    - Database
- [12-Factor App Principles](#12-factor-app-principles)
- [Deployment](#deployment)
    - Containerize (with Docker)
    - Kubernetes (with kind)
- [Difficulties](#difficulties)

# Project Overview
This project aims to develop a modular, cloud-ready microservice architecture that enables users to perform CRUD operations on items. The frontend interacts with the backend through a RESTful API. Each component is containerized using Docker and orchestrated with docker-compose. Configuration values are externalized via environment variables inside the `devcontainer.json`, so no code changes are required to customize the enviroment variables.The codebase is designed to be clean, maintainable, and easily extendable.

# Setup Instructions

These are instructions to primarily run the application in a development environment but also setup up some configurations for eventual deployments.

## GitHub Codespaces

1. To run the entire stack seamlessly without local setup, use GitHub Codespaces:
2. Click the `<> Code` button in your repository.
3. Select `Codespaces` > `Create codespace on main`.
4. Wait for the container to build (it uses devcontainer.json).
5. On the bottom is the panel, please switch to Ports panel view, copy the backend forwarded address paste it inside the `devcontainer.json` in the variabel `"VITE_PUBLIC_BACKEND_URL"` and delete the last `/`.
6. You are now asked if you want to rebuild the container. Please do that.
    - This will rerun the setup script and prepare all configuration.
7. After rebuilding you can now start the back- and frontend with the terminal by the following commands:
    
    - Backend:

            cd backend/
            go run .

    - Frontend: 

            cd frontend/
            npm run dev

8. Last but not least again switch to the ports panel view and change both Back- and Frontends visibility to `Public`.
9. You are now able to acces both Back- and Frontend by for example clicking the link in the ports panel view

When opening a github workspace the project should provide the neccassary dependancies. Ports for the frontend and backend will also be forwarded but are still in private and have to be changed to public manually for security reasons.
For other vs-code extensions on setup add those in the `devcontainer.json`.

## Backend
- Located in `/backend`
- Build with Gin framework
- Configuration is read from environment variables (via `config.LoadConfig()`, which are defined in `devcontainer.json`)

Can be started with the following commands:

    cd backend/
    go run .

## Database
- PostgreSQL via Docker
- Credentials and database name are defined in `devcontainer.json`
- Runs as a service inside the Docker environment (`db`)

## Frontend
- Located in /frontend
- Built with SvelteKit
- .env is dynamically created via devcontainer.json

Can be started with the following commands:

    cd frontend/
    npm run dev

# 12-Factor App Principles
This project aligns with the 12-Factor App methodology, which promotes best practices for building modern, scalable, and maintainable web applications. Below is an overview of how each principle is implemented **in this project**:

## I. Codebase
**Definition**: One codebase tracked in revision control, many deploys

**In this Project**: All code (frontend, backend, config) lives in a single Git repository.

## II. Dependencies

**Definition**: Explicitly declare and isolate dependencies

**In this project**: Go uses go.mod, and the frontend uses package.json to manage dependencies.

## III. Config

**Definition**: Store config in the environment

**In this project**: Environment variables are defined in .env, devcontainer.json, or docker-compose.yml depending on the environment. Go accesses them via os.Getenv.

## IV. Backing Services

**Definition**: Treat backing services as attached resources

**In this project**: PostgreSQL is treated as an external, replaceable service. It can be swapped without code changes.

## V. Build, Release, Run

**Definition**: Strictly separate build and run stages

**In this project**: npm run build and go build are used to build separately from runtime (npm run dev, go run .).

## VI. Processes

**Definition**: Execute the app as one or more stateless processes

**In this project**: The backend runs as a stateless service; no session or file-based state is kept.

## VII. Port Binding

**Definition**: Export services via port binding

**In this project**: The backend listens on port :8080, and the Svelte frontend on :5173.

## VIII. Concurrency

**Definition**: Scale out via the process model

**In this project**: Both frontend and backend services can be replicated to scale horizontally.

## IX. Disposability

**Definition**: Maximize robustness with fast startup and graceful shutdown

**In this project**: The backend starts quickly with Gin and handles shutdowns gracefully.

## X. Dev/prod Parity

**Definition**: Keep development, staging, and production as similar as possible

**In this project**: Docker and Codespaces mirror the production environment closely.

## XI. Logs

**Definition**: Treat logs as event streams

**In this project**: Logs are printed to stdout and are easily capturable by logging systems.

## XII. Admin Processes

**Definition**: Run admin/management tasks as one-off processes

**In this project**: Administrative actions (e.g., database tasks) can be triggered via CLI or dedicated REST endpoints.

We aim to deploy the system on Kubernetes using Helm charts and manifests. This will include:
- Deployment & Service YAMLs
- ConfigMaps & Secrets for env variables
- Persistent Volume setup for PostgreSQL#

# Deployment

We have defined our deployment strategy using Docker and Kubernetes (with Kind). While Kubernetes should only be executed in the web codespace Docker can be used everywhere.

**Before deploying make sure to configure the right files**.</br>
*If you are using GitHub Codespaces follow the instructions [here](#github-codespaces).*

## Containerize

Our whole application is containerized using Docker. This includes separate Dockerfiles for the backend and frontend, as well as a docker-compose.yml file for local development.

If you are ready for [deployment](#deployment) run the `docker-compose.yml` file.
``` bash
    docker-compose up --build
```

## Kubernetes

Kubernetes is an open-source container orchestration platform that automates the deployment, scaling and management of containerized applications. In this project we use Kubernetes (with kind) to demonstrate the orchestration of our services. Configure more options within the `kubernetes` directory and the corresponding files. You can change for example the images (by default our images of this repository are used) or `replicas: <count of replicas>`.

If you are ready for [deployment](#deployment) run the `start-kubernetes.sh` file.
``` bash
    bash ./start-kubernetes.sh
```

# Difficulties

- **Configuration of deployment**
</br>
In GitHub Codespaces everything should work fine after following the instruction for [GitHub Codespaces](#github-codespaces).
</br>
The problem arises when using VSC for codespaces. Somehow the forward functionality was not working as expected. When running a service the port should be forwarded locally and be accessible but sometimes it could not be reached. Therefore it needs to be set public but then it is only accessible through the GitHub url. That means while following the instructions for codespaces but using VSC do not use localhost as a backend url because of flakey behaviour. After that the dockerization should work properly but kubernetes should still not be working. That is because when trying to access the cluster it is searching locally but kubernetes and kind are installed on the VM of the codespace. For more advanced, try to automate or remotely control the cluster but this is not done in this repository.