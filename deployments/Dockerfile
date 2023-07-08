FROM golang:1.20.5-alpine as base
WORKDIR /builder
# install upx and make
RUN apk add upx make
ENV GO111MODULE=on CGO_ENABLED=0
COPY . .
RUN go mod tidy
RUN make build

FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/bin/asynqmon-multi-service /app/bin/asynqmon-multi-service
CMD ["/app/bin/asynqmon-multi-service", "server"]
