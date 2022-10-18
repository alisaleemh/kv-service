FROM golang:latest

WORKDIR /opt/app

ADD bin /opt/app/bin

EXPOSE 8081
#CMD ["go", "run", "main.go"]
RUN chmod +x /opt/app/bin/main
CMD ["/opt/app/bin/main"]
