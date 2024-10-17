FROM golang:1.22.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base-debian12:debug
WORKDIR /app
ENV APP_ENV_FILE=.env
COPY --from=builder /app/.env .
COPY --from=builder /app/constants/locales ./constants/locales
COPY --from=builder /app/main .
ENTRYPOINT [ "/app/main" ]
