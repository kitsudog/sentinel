FROM golang:1.22.2-alpine3.19
WORKDIR /src
COPY src/ /src/
RUN go build main.go && mkdir /app && cp main /app/sentinel
CMD [ "/app/sentinel", "80", "uname", "-a" ]