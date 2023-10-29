FROM golang:1.21.3-alpine3.18

ARG DB_HOSTNAME=localhost

WORKDIR /app

RUN rm -rf /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN mv .docker.env .env

# got it almost working, just need to pass the DB_HOSTNAME from docker-compose to here...
RUN printf '\n' >> .env
RUN echo DB_HOSTNAME=${DB_HOSTNAME} >> .env
RUN echo DB_CONNECTION=mongodb://\${DB_USERNAME}:\${DB_PASSWORD}@\${DB_HOSTNAME}:27017/ >> .env

RUN go build .

EXPOSE 8080

CMD [ "./xlearn" ]

# to restore vm disk space, resize it so it resets everything!