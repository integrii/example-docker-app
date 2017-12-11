FROM golang
ENV VERSION=DEFAULT

RUN mkdir /app
ADD main.go /app/main.go
WORKDIR /app
RUN go build

EXPOSE 80
ENTRYPOINT /app/app
