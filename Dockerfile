FROM golang:latest AS builder
WORKDIR /builder
COPY . /builder
RUN make build

FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add tzdata ca-certificates libc6-compat
COPY --from=builder /builder/server /app
COPY --from=builder /builder/resource /app/resource
EXPOSE 8000
ENTRYPOINT ["/app/server"]
