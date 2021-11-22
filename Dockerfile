FROM golang:1.17

ENV SERVER_HOST=localhost
ENV SERVER_PORT=8080

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=test
ENV DB_PASS=test
ENV DB_NAME=test
ENV JWT_SECRET=test

EXPOSE 8080

COPY src/ /
WORKDIR /

RUN go build -o gwi/assignment

CMD ["./gwi/assignment"]