FROM  golang:alpine as builder
WORKDIR /code
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -mod vendor -o audio ./cmd/

FROM alpine:latest
WORKDIR /code
COPY --from=builder /code/audio .
COPY --from=builder /code/configs/* ./configs/
CMD ["./audio",  "--conf",  "./configs"]



