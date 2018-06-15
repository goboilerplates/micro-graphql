echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/micro-graphql
docker push goboilerplates/micro-graphql:0.0.0