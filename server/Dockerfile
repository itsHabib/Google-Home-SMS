FROM golang
LABEL maintainer="michael habib"
RUN useradd google-home-sms
USER google-home-sms

COPY . /go/src/github.com/itsHabib/google-home-sms/
RUN go install github.com/itsHabib/google-home-sms/
EXPOSE 8081
CMD ["/go/bin/google-home-sms"]