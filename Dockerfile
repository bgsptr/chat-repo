FROM golang

WORKDIR /userservice

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o cmd/main ./cmd

CMD [ "./cmd/main" ]
