kubectl apply -f kubernetes/

sleep 10000

kubectl get pods -A

# Subshell for frontend
(
    kubectl port-forward service/frontend 3000:3000
)

# Subshell for backend
(
    kubectl port-forward service/backend 8080:8080
)

while true:
done