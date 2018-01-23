#! /bin/bash
docker container run -d --name google-home-sms -p 8081:8081 \
-e "TWILIO_SID=XXXX" \
-e "TWILIO_TOKEN=XXXX" \
itshabib/google-home-sms
