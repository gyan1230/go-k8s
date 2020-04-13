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
