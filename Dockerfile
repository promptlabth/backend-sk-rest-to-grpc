FROM golang:1.21-alpine as build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY . .

# Download the dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOARCH=amd64 go build -mod=readonly -v -o main


FROM gcr.io/distroless/static:nonroot as runtime
COPY --chown=nonroot:nonroot --from=build /app/main /main

EXPOSE 8080

ENTRYPOINT [ "/main", "-alsologtostderr" ]