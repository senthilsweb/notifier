
# Introduction

A multipurpose notification library to send mail, text, slack, push, telegram etc. developed in `Go`

## Pre-requisites
* Should have Go installed in your local development environment
* Optional [netlify](https://app.netlify.com/) account to deploy the golang api's as serverless functions
* Optional [netlify](https://app.netlify.com/) CLI
* Optional if you want to send HTML formatted email using mailgun template. I used [Maizzle](https://maizzle.com/) to build HTML emails with
[Tailwind CSS](https://tailwindcss.com/)
https://maizzle.com/
* [mailgun](https://app.mailgun.com/) account as the `notify` (sendmail) api is built using `mailgun`
* VSCode Editor or [Gitpod](https://gitpod.io/) online VSCode editor 50 hours per month free plan

## Local Development

> Clone repository.

```bash
git clone https://github.com/senthilsweb/notifier.git
```

> Local build (Mac OS)

```bash
go build
```

> Local Run

```bash
./notifier -p "3000"
```

##  API endpoints

`https://zybes.netlify.app/api`

Request |       Endpoints                                                |       Functionality
--------|----------------------------------------------------------------|--------------------------------
POST    |  https://zybes.netlify.app/api/notify/mailgun                  |   Send email  
GET     |  https://zybes.netlify.app/api/ping                            |   Health check


Payload for `https://zybes.netlify.app/api/notify/mailgun`

```Json
{
    "message": {
        "subject": "This is subject",
        "body": "This is body",
        "template": "welcome_email",
        "recipient": "name <your email@gmail.com>",
        "payload": {"name":"John Smith"}
    },
    "MAILGUN_DOMAIN": "your domain",
    "MAILGUN_KEY": "your key",
    "EMAIL_SENDER": "Mailgun Sandbox <your sender>"
}
```

## Netlify Deployment

Refer the following documentaion and blog post to host the server (and optional Single Page Application) at [Netlify](https://docs.netlify.com/)

* [Netlify](https://docs.netlify.com/) documentaion 
* [blog post](https://blog.carlmjohnson.net/post/2020/how-to-host-golang-on-netlify-for-free/) by [Carl M. Johnson](https://carlmjohnson.net/)

## Key Frameworks and Libraries used

- [x] github.com/apex/gateway
- [x] github.com/gin-gonic/gin
- [x] github.com/mailgun/mailgun-go/v4
- [x] github.com/sirupsen/logrus
- [x] github.com/spf13/viper
- [x] github.com/tidwall/gjson
