#!/bin/bash

kubectl apply -f kubernetes/

# echo ""
# kubectl get pods -A
echo " ----- "
echo ""

# Subshell for frontend
(
    while [ "$(kubectl describe pod/frontend | grep 'Status:' | cut -d':' -f2 | xargs)" != "Running" ]; do
        sleep 5
        echo "Waiting for frontend..."
    done
    # kubectl describe pod/frontend
    kubectl port-forward service/frontend 3000:3000
) &

# Subshell for backend
(
    while [ "$(kubectl describe pod/backend | grep 'Status:' | cut -d':' -f2 | xargs)" != "Running" ]; do
        sleep 5
        echo "Waiting for backend..."
    done
    # kubectl describe pod/backend
    kubectl port-forward service/backend 8080:8080
) &

wait;