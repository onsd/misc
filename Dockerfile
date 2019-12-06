

FROM golang:1.13 as server
WORKDIR /go/src/sample_docker_compose
COPY . .
RUN ls


RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o app

FROM alpine:3.9

COPY --from=server /go/src/sample_docker_compose/app /app

EXPOSE 8080
EXPOSE 80
EXPOSE 443

ENTRYPOINT /app





