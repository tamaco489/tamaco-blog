FROM golang:1.25-alpine AS builder

WORKDIR /go/src/api

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -tags appsec \
    -o /go/bin/app \
    -ldflags="-w -s -extldflags '-static'" \
    ./cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot

USER nonroot:nonroot

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.9.1 --chown=nonroot:nonroot /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=builder --chown=nonroot:nonroot /go/bin/app /app

CMD ["/app"]
