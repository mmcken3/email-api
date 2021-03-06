# Email API

This is a tool to serve as an email and text message API for personal websites. If you
want to use a contact form but not set up a third party tool or write your own backend then this could be used.
You could download the binary of this repo or download and build yourself to start the API. It will
take simple 3 field forms for the post bodies and then send those forms onto the set up receivers of
the forms in the Email API.

The email endpoint is set up so that this API could be stood up by anyone with a Gmail account and configured
to send messages to any email they like through the configured host email. This auth and email address
set up is configured through environment variables.

The text endpoint is set up so that this API could be stood up by anyone with a Twilio account and configured
to send messages to any number they like through a phone number in that account. This auth and phone number
set up is configured through environment variables.

*Still under development*

## Development

Download the source:

    go get github.com/mmcken3/email-api

Go Install:

    go install ./...

Run the Server:

    emailapi

The APi can then be hit locally at http://localhost:3000.

## Email API Endpoints

HTTP request | Description
------------ | -------------
**GET** /health    | returns a health check for the server |
**POST** /v1/send/email    | submits a request to send an email with data from the body |
**POST** /v1/send/text    | submits a request to send an text with data from the body |

The send text endpoint uses a Twilio integration. A Gist where you can see an example of
this type of integration can be found [here](https://gist.github.com/mmcken3/d2a485cb713b9f68ebeb28cc73c0c2af).

#### Request Body

For both of the post endpoints they will take the same JSON body to send the message. This
JSON body is like this:

    {
        "name": "Example Name",
        "email_address": "test@test.com",
        "message": "This is an example message"
    }

## Configuration

Environment variables that need to be set in order to run the Email API. This is your Twilio account
auth, the Twilio phone numbers text come from in your account and then number you want text to be
sent too. These phone numbers must be in the format with a plus like: "+11234567890".

    TWILIO_ACCOUNT_SID
    TWILIO_AUTH_TOKEN
    FROM_TWILIO_NUMBER
    TO_PHONE_NUMBER

These env variables are for your email endpoint. EMAIL_USER is your email account to end from and
its password. Then for gmail these are the server and port that you would want to use. Anyone
can configure this to work with their email and send to a specific email as well.

    EMAIL_USER
    EMAIL_PASSWORD
    EMAIL_SERVER='smtp.gmail.com'
    EMAIL_PORT='587'
    SEND_TO
    EMAIL_FROM

## Deploying

For me I am running this on a linux server. You could just pull down the code into a linux box and build
it there like above if you like. However, you could also run:

    GOOS=linux go build -o emailapi ./cmd/emailapi

And then move this binary over to your server to run.

## External Libraries Used

[go-chi router](https://github.com/go-chi/chi)

[go-chi render](https://github.com/go-chi/render)

[godotenv](https://github.com/joho/godotenv)

[envconfig](https://github.com/kelseyhightower/envconfig)

[errors pkg](https://github.com/pkg/errors)
