FROM golang:1.22

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# baixando a lib de comunicação do kafka
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]