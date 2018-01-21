FROM golang
LABEL maintainer="michael habib"

COPY . /go/src/github.com/itsHabib/google-home-sms/
WORKDIR /go/src/github.com/itsHabib/google-home-sms/

# Have to copy this because package isn't in public repo yet
COPY twiliogo/ /go/src/github.com/itsHabib/twiliogo/

RUN go install github.com/itsHabib/twiliogo/
RUN go install
EXPOSE 8081
ENTRYPOINT [ "/go/bin/google-home-sms" ]