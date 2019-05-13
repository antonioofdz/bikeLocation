FROM golang:1.12

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download && \
    go install ./cmd/dra

EXPOSE 5002:5002

CMD ["dra"]