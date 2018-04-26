CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t goboilerplates/micro-rest .
docker tag  goboilerplates/micro-rest goboilerplates/micro-rest:0.0.0