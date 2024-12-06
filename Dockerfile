FROM alpine:3.20

RUN mkdir -p /etc/env
RUN apk update && \
    apk upgrade && \
    apk add bash envsubst && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

COPY migrations/*.sql migrations/
COPY migrate.sh .
COPY ./.env /etc/env/local.env


RUN chmod +x migrate.sh

ENTRYPOINT ["bash", "migrate.sh"]