CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t goboilerplates/micro-graphql .
docker tag  goboilerplates/micro-graphql goboilerplates/micro-graphql:0.0.0