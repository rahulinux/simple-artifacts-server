# Start by building the application.
FROM golang:1.8 as build
COPY serve.go /
WORKDIR /
RUN go get github.com/gorilla/handlers
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o serve .



# Now copy it into our base image.
FROM gcr.io/distroless/base
ADD artifacts /opt/artifacts
COPY --from=build /serve /
CMD ["/serve"]
