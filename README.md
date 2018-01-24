# Google Home SMS

This project was a solution to the fact that the Google Home products do
not support texting yet. The Repository contains the server code responsible
for handling requests and sending SMS. It contains the code that is used for the 
Google Action that is responsible for gathering the data as in who the text
should be sent to and the message. 

## Usage
In order to use the code I've written so far you need to first set up a
Twilio account and get an account SID and auth token. Once you obtain those
you can clone / fork the repository and build and run the file. If you run the file without 
docker keep in mind you will need a TWILIO_SID and TWILIO_TOKEN environment variable set on your machine. The Twilio handler uses these environment variables to autheticate when making requests. I also created a docker file so that the program can be run in a container and a start up script for a Google Compute Engine instance. This is how  I have the server running, in a container on a Google Compute Engine instance. To get a more in depth explanation of how to implement my solution check out my blog post [How I get my Google Home To Send Texts](https://medium.com/@itsHabib). In order for a text to be sent using the handler a POST request needs to be sent to the server containing a body of:
```json
{
    "To": "NUMBER_TO",
    "From": "NUMBER_FROM",
    "Body": "Message"
}
```

## Notes
It is completely possible to just have the Google Action be responsible for gathering the
data and sending the SMS.