FROM alpine
RUN apk update && apk add --update --no-cache ca-certificates
COPY mmeshctl /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/mmeshctl"]
