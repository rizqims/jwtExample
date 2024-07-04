FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api-laundry
ENTRYPOINT ["/app/api-laundry"]
