### API

POST https://zybes.netlify.app/api/notify/mailgun

### Payload

```
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