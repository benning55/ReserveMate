FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

EXPOSE 8080

CMD [ "./target/server" ]