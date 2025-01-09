FROM alpine:3.21.1

RUN     mkdir /cortex
WORKDIR /cortex
ADD     ./cortex ./
