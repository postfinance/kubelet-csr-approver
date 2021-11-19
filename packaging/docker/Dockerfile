FROM alpine:latest

ARG TZ=Europe/Zurich

ENV TZ=${TZ}
RUN apk add --no-cache tzdata

COPY approverd /usr/bin/approverd

COPY packaging/docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod 755 /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
