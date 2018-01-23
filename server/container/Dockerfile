FROM golang
LABEL maintainer="michael habib"
RUN useradd google-home-sms
USER google-home-sms

RUN go get github.com/itsHabib/google-home-sms
EXPOSE 8081
CMD ["/go/bin/google-home-sms"]