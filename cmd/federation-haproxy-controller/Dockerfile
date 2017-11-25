FROM golang:1.8.5

ENV fedApiServerAddr 127.0.0.1
ENV fedApiServerPort 6443

WORKDIR /usr/src/app

COPY . .

RUN go build main.go

CMD ./main 
