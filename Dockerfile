# syntax=docker/dockerfile:1

# Build stage
ARG GO_VERSION=1.22.1
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Cache dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# Copy the application source
COPY . .

# Build the application
ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/app

# Final stage
FROM alpine:latest AS final

# Install dependencies
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        postgresql-client \
        && \
        update-ca-certificates

# Create a non-root user
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

USER appuser

# Copy the built binary from the build stage
COPY --from=build /bin/server /bin/server

ENTRYPOINT ["/bin/server"]
