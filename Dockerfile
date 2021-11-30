FROM golang:1.17-alpine
WORKDIR /fafagans
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN source .env
RUN go build -o program
CMD ./program