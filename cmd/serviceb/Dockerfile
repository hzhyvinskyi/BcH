FROM golang:1.15 as builder
ENV GO111MODULE=on
WORKDIR /build
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM scratch
WORKDIR /bin
COPY --from=builder /build/app .
#COPY .env .
EXPOSE 8084
ENTRYPOINT ["/bin/app"]
