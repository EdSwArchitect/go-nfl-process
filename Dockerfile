FROM ubuntu:20.10

RUN apt update

RUN mkdir /data /usr/local/ekb

WORKDIR /usr/local/ekb

COPY go-nfl-process go-nfl-process

VOLUME [ "/data" ]

EXPOSE 18080

ENTRYPOINT ["/usr/local/ekb/go-nfl-process", "-bootstrap-server", "kafka-svc:9092", "-dir", "/data", "-verbose=true", "-weeks-topic", "nfl.weeks", "-plays-topic", "nfl.plays", "-games-topic", "nfl.games" ]
 