Deploy, manage, and scale a simple Go web app on Kubernetes.

Dockerizing the Go application

Create a file named Dockerfile inside the project’s folder and add configurations in the Dockerfile.

 Build and push the docker image of our Go app on docker hub so that we can later use this image while deploying the app on Kubernetes 

 # Build the docker image
$ docker build -t go-k8s .

# Tag the image
$ docker tag go-kubernetes gyanchand1230/go-webapp:1.0.0

# Login to docker with your docker Id
$ docker login

# Push the image to docker hub
$ docker push gyanchand1230/go-webapp:1.0.0


Creating a Kubernetes deployment

It’s a simple yaml file with a bunch of configurations containing the desired state of 
application.

start a Kubernetes cluster:
minikube start

Deploy our app to the minikube cluster by applying the deployment manifest using kubectl.
kubectl apply -f k8s-deployment.yml

The deployment is created. You can get the deployments like this:

$ kubectl get deployments

to get the pods in the cluster:

$ kubectl get pods

Pods are allocated a private IP address by default and cannot be reached outside of the cluster. You can use the kubectl port-forward command to map a local port to a port inside the pod like this:

$ kubectl port-forward go-webapp-69b45499fb-7fh87 8080:8080

now interact with the Pod on the forwarded port:

$ curl localhost:8080/home