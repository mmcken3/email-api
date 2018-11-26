# Email API

This is a tool to serve as an email and text message API for personal websites. If you 
want to use a contact form but not set up a third party tool or write your own backend then this could be used.
You could download the binary of this repo or download and build yourself to start the API. It will
take simple 3 field forms for the post bodies and then send those forms onto the set up recievers of
the forms in the Email API.

*Still under development*

Right now it is hard coded to work for only a few personal websites and the email endpoint is not set up
at all. I am currently working on setting it up to make the emails it sends too/from and the Twilio numbers
it sends to/from configuable so that this API can be plug and played for anyones personal websites.

The send text endpoint is pretty much ready to go and works for the hard coded numbers. The email one 
is the main endpoint still being worked on.

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
