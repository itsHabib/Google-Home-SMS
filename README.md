# Google Home SMS

This project was a solution to the fact that the Google Home products do
not support texting yet. The Repository contains the server code responsible
for handling requests and sending SMS. It contains the code that is used for the 
Google Action that is responsible for gathering the data as in who the text
should be sent to and the message. 

## Current Progress
As of 1 / 08 / 2018 the server for sending texts is complete. 
The plan is to complete the Google Action by 1 / 14/ 2018. 

## Usage
In order to use the code I've written so far you need to first set up a
Twilio account and get an account SID and auth token. Once you obtain those
you can clone / fork the repository and just run the main go file, providing your creds. 
I also created a docker file so that the program can be run in a container. This is how 
I have the server running, in a container on a Google Compute Engine instance.
In order for a text to be sent using the handler a POST request needs to be sent
to the server containing a body of:
```json
{
    "To": "NUMBER_TO",
    "From": "NUMBER_FROM",
    "Body": "Message"
}
```

### Notes
It is completely possible to just have the Google Action be responsible for gathering the
data and sending the SMS. I decided not to do this several reasons. I really enjoy writing
code in Go and have never used Twilio with it before. I also have never used Google's Compute
Engine as a host to a container. Lastly, I plan on refactoring out the Twilio code into its own
go package for others to use.


### Future Plans As of 1 / 08 / 2018
- [] Set up database that holds contacts
- [] Create Google Action
- [] Refactor Twilio code into its own go package 
