FROM golang:alpine AS dev-env

WORKDIR /usr/local/go/src/airman.com/airmsExample
COPY . /usr/local/go/src/airman.com/airmsExample

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN go get ./...

RUN go build -o dist/airmsExample &&\
    cp -f dist/airmsExample /usr/local/bin/ &&\
    cp -f dist/airmsExample.ini /usr/local/etc/ &&\

RUN ls -l && ls -l dist

CMD ["/usr/local/bin/airmsExample", "-c", "/usr/local/etc/airmsExample.ini" ]