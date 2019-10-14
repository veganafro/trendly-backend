FROM golang:latest

LABEL maintainer="Jeremy Muhia <jmuhia80@gmail.com>"

ADD . /go/src/trendly-backend
WORKDIR /go/src/trendly-backend

ARG LOG_DIR=/go/src/trendly-backend/logs
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/trendly-backend.log

COPY . .

RUN go get trendly-backend
RUN go build -o main .

EXPOSE 8080
VOLUME [${LOG_DIR}]

CMD ["./main"]

