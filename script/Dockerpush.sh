echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/micro-rest
docker push goboilerplates/micro-rest:0.0.0