FROM golang:1.17-alpine
WORKDIR /fafagans
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o program
RUN source .env
CMD ./program