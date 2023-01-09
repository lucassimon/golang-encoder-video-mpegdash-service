FROM golang:1.19-alpine as base
RUN apk --no-cache update && \
    apk add --no-cache git

FROM base as ci
WORKDIR /app/
COPY . .

FROM ci as builder
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o entrypoint .

FROM scratch
WORKDIR /app/
USER app
COPY --from=builder /app/entrypoint .

ENTRYPOINT [ "/app/entrypoint" ]
