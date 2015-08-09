# Lambda SES

An AWS Lambda function to send an email via SES.
Sends email to a set recipient with the content of the JSON input.

[Blog Post](http://www.rjocoleman.io/2015/08/08/contact-form-overkill.html)


## Features:

* Shells out to a Go binary to send the email.
* Policies to allow only one sender and receiver.
* Low line count.


## Usage

For full setup and usage instructions, including an example with AWS API Gateway check out my [Blog post](http://www.rjocoleman.io/2015/08/08/contact-form-overkill.html)

Input JSON:

```json
{
  "FromAddress": "from_address@example.com",
  "ToAddress": "to_address@example.com",
  "Subject": "Test Message",
  "Message": "Test Body"
}
```


## Setup

You need:

* An AWS account
* A Lambda function with the basic execution policy.
* SES set up with an email address and policy (permissons); an example set of permissions is in the [policies/ses.json](policies/ses.json) file.
* Your SES address configured in `config.json` (see [config.example.json](config.example.json))
* Build your own, or download the latest release from the releases section above.
* Add your config.json, zip and upload the zip to your Lambda.


### Build your own

You'll need Golang installed, then:

```shell
$ make build
```


## Standalone usage

There's nothing at all limiting the executable `lambda-ses` to being used only in lambda. You can pass in the below JSON and AWS credentials come from ENV.
`AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`.

Standalone Input JSON:

```json
{
  "SenderAddress": "sender_address@example.com",
  "FromAddress": "from_address@example.com",
  "ToAddress": "to_address@example.com",
  "Subject": "Test Message",
  "Message": "Test Body"
}
