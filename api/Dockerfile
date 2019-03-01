FROM golang:1.12 as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

# Build image
FROM scratch 
COPY --from=build-env /go/bin/app /app
COPY --from=build-env /app/config.yaml .
ENTRYPOINT ["/app"]