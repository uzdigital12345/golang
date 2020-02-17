# container for contact list
FROM golang:1.13.4

RUN mkdir -p $GOPATH/src/github.com/uzdigital12345/golang/models_grpc_contact_list_docker
WORKDIR $GOPATH/src/github.com/uzdigital12345/golang/models_grpc_contact_list_docker

COPY . .
RUN go build server/contact_service.go


ENTRYPOINT ["./contact_service"]
