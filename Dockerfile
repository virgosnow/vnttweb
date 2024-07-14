FROM golang:alpine

WORKDIR /home
COPY . /home

CMD ["go", "run", "main.go"]
