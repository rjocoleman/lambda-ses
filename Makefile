all: build

build: ses lambda.zip

release: ses release.zip

clean:
	rm -f send-mail lambda.zip

ses: *.go
	GOOS=linux GOARCH=amd64 go build

lambda.zip: lambda.js ses
	zip lambda.zip lambda.js lambda-ses config.json

release.zip: lambda.js ses
	zip lambda.zip lambda.js lambda-ses

vendor:
	godep save -r -copy=true ./...
