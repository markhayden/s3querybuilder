// Package s3querybuilder is used for generating the full encoded url to allow an anonymous user to access a single file in a AWS S3 bucket
package s3querybuilder

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

// Cfg containts the necessary params to generate a secure, self destructing s3 download url
type Cfg struct {
	File        string
	Bucket      string
	AccessKey   string
	SecretKey   string
	HTTPVerb    string // Defaults to GET if not specified
	ContentMD5  string // Optional
	ContentType string // Optional, S3 will default to octet/stream if not specified
	Expire      int    // Seconds for the URL to expire
}

// Link creates the necessary headers / signatures for anonymous user to download/upload assets from s3 bucket
func (d *Cfg) Link() string {
	// convert date to string so it doesnt explode
	expiration := time.Now().Unix() + int64(d.Expire)

	// generate base url from the bucket name
	s3BaseURL := "https://" + d.Bucket + ".s3.amazonaws.com"

	// prepare the uri for the selected file
	uri := "/" + d.Bucket + "/" + d.File

	// defaults to GET to not break compatibility
	if d.HTTPVerb == "" {
		d.HTTPVerb = "GET"
	}

	// generate the aws message in preparation for signature generation
	message := fmt.Sprintf("%s\n%s\n%s\n%d\n%s", d.HTTPVerb, d.ContentMD5, d.ContentType, expiration, uri)

	// generate a signature for aws handshake
	sha256 := sha1.New
	hash := hmac.New(sha256, []byte(d.SecretKey))
	hash.Write([]byte(message))
	sha := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// encode special charactes on all the query string parameters
	toEncode := url.Values{}
	toEncode.Add("AWSAccessKeyId", d.AccessKey)
	toEncode.Add("Expires", fmt.Sprintf("%d", expiration))
	toEncode.Add("Signature", sha)
	query := toEncode.Encode()

	// return the fully formatted url
	return fmt.Sprintf("%s/%s?%s", s3BaseURL, d.File, query)
}
