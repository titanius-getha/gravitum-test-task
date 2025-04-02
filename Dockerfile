FROM golang:1.24-alpine AS deps
RUN go install github.com/air-verse/air@latest
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download

FROM deps AS dev
COPY .air.toml ./
COPY ./src ./
CMD ["air"]

FROM deps AS prod
COPY ./src .
RUN go build -o /app/main .
CMD ["/app/main"]