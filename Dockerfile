FROM golang:1.21 as builder
ENV DEBIAN_FRONTEND=noninteractive
ADD . /build
RUN apt-get update &&\
    apt-get install -y curl git gpg wget unzip &&\
    curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key | gpg --dearmor >> /nodesource-key.gpg &&\
    echo "deb [signed-by=/nodesource-key.gpg] https://deb.nodesource.com/node_20.x bookworm main" >> /etc/apt/sources.list.d/nodesource.list &&\
    echo "deb-src [signed-by=/nodesource-key.gpg] https://deb.nodesource.com/node_20.x bookworm main" >> /etc/apt/sources.list.d/nodesource.list &&\
    apt-get install -y nodejs npm &&\
    cd /build &&\
    CGO_ENABLED=0 GOOS=linux go build -a -o app . &&\
    ./install.sh

###############################################################################

FROM ubuntu:22.04

COPY --from=builder /build/assets /app/assets/
COPY --from=builder /build/templates /app/templates/
COPY --from=builder /build/app /app/
ADD docker/run.sh /app/

ENV DEBIAN_FRONTEND=noninteractive

RUN useradd --user-group --system playtime &&\
    cd /app &&\
    mkdir -m 0777 data uploads &&\
    chmod 0777 /app &&\
    chown -R playtime:playtime /app

USER playtime
WORKDIR /app
EXPOSE 3000

VOLUME ["/app/data", "/app/uploads"]

CMD ["./run.sh"]
