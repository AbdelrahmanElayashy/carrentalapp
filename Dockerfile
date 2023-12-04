# ==================
# BUILD STAGE
# ==================
FROM golang:1.21.0-alpine3.18 AS build

WORKDIR /app
COPY /src/go.mod /src/go.sum ./
RUN go mod download
COPY /src/ ./
RUN go build -o main .

# ==================
# TEST STAGE
# ==================
# Use the previous build stage as a base for test stage
FROM build AS test

RUN go test -v ./...

# ==================
# OPTIMIZATION STAGE
# ==================
FROM build AS optimize

RUN apk --no-cache add binutils
RUN strip main

# ==================
# PRODUCTION STAGE
# ==================
FROM alpine:3.18 AS production

WORKDIR /root/
COPY --from=optimize /app/main .
EXPOSE 80
CMD ["./main"]