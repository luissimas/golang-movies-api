FROM golang:1.20 AS build

WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Copy the build to a lean image
FROM gcr.io/distroless/base-debian11
COPY --from=build /server /server
COPY config-docker.toml /config.toml

EXPOSE 3000

CMD ["/server"]
