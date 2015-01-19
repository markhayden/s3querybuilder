# Self Destructing S3 Content Link Generator

[![GoDoc](https://godoc.org/github.com/markhayden/s3querybuilder?status.png)](https://godoc.org/github.com/markhayden/s3querybuilder)

## License
Contributors: Mark Hayden  
Tags: go, golang, s3, aws, amazon web services, s3querybuilder  
License: GPLv2 or later  
License URI: http://www.gnu.org/licenses/gpl-2.0.html

## Overview

This is a simple package used for generating the full encoded url to allow an anonymous user to access a single file in a AWS S3 bucket. Generated link includes a expiration setting to ensure bucket security. Expiration is set by the second.

You can specify an HTTP verb such as `PUT` to allow users to upload an specific file, if no verb is specified it defaults to `GET`, allowing the file to be downloaded.

`go get github.com/markhayden/s3querybuilder`

## Usage

### Generate a URL which expires in 5 minutes to download a file from s3


```Go
	q := s3querybuilder.Cfg{
		File:      "file.txt",
		Bucket:    "my-s3-bucket",
		AccessKey: "<AWS_KEY>",
		SecretKey: "<AWS_KEY_SECRET>",
		Expire:    300,
	}
	fmt.Println(q.Link())
```

### Generate a URL to upload a file to s3

```Go
	q := s3querybuilder.Cfg{
		File:      "file.txt",
		Bucket:    "my-s3-bucket",
		AccessKey: "<AWS_KEY>",
		SecretKey: "<AWS_KEY_SECRET>",
		HTTPVerb:  "PUT",
		Expire:    300,
	}
	fmt.Println(q.Link())
```

This will output the URL, you can use it with curl for example:

```curl -v -X PUT --upload-file "file.txt" "<URL>"```

Notice that you must put the URL inside double quotes or curl will not send all the query parameters.