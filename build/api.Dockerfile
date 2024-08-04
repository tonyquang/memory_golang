FROM alpine:3

RUN apk --no-cache add tzdata ca-certificates

COPY ./api/binaries/serverd /
COPY ./api/app.env /

CMD ./serverd