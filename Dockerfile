FROM golang:1.24-alpine AS deps
RUN go install github.com/air-verse/air@latest
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download

FROM deps AS dev
COPY .air.toml ./
COPY ./src ./
CMD ["air"]

FROM deps AS build_prod
COPY ./src .
RUN go build -ldflags="-s -w" -trimpath -o backend .

FROM alpine:latest AS prod
WORKDIR /app
COPY --from=build_prod /app/backend ./
CMD ["./backend"]