# microservice_item
![https://ewanvalentine.io/microservices-in-golang-part-1/](https://img.shields.io/badge/tutorial-https%3A%2F%2Fewanvalentine.io%2Fmicroservices--in--golang--part--1%2F-brightgreen.svg)

# install

1. install protobuf
```
GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

2. install other
```
go get -u github.com/satori/go.uuid
go get -u google.golang.org/grpc
go get github.com/microservice_item
```

3. run
```
go run main.go
```

# LICENSE

mit
