ARG GO_VERSION=1.25
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE=server
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/bin/${SERVICE} ./cmd/${SERVICE}

FROM alpine:${ALPINE_VERSION} AS runtime

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/bin/${SERVICE} .

ARG SERVICE
ENV SERVICE=${SERVICE}

EXPOSE 8080 8081

CMD ["./${SERVICE}"]
