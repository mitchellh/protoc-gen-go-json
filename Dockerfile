#--------------------------------------------------------------------
# builder builds the binaries
#--------------------------------------------------------------------

FROM golang:1.17-alpine AS builder

# Ensure we're doing the right thing for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Copy our app
RUN mkdir -p /src
WORKDIR /src
COPY . .

# Build
RUN go mod download
RUN go build -o /protoc-gen-go-json .

#--------------------------------------------------------------------
# copy the built static binary to a scratch image
#--------------------------------------------------------------------

FROM scratch

COPY --from=builder /protoc-gen-go-json /protoc-gen-go-json

ENTRYPOINT ["/protoc-gen-go-json"]
