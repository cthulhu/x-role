package xrole

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FileCacheProvider struct {
	Creds *credentials.Credentials
}

func (f *FileCacheProvider) Retrieve() (credentials.Value, error) {
	// TODO check file credential cache before looking at nested credentials.
	// Fall back to underlying credentials, and repopulate the cache.
	return f.Creds.Get()
}
func (f *FileCacheProvider) IsExpired() bool {
	// TODO check file cache is expired? Fall back to underlying credentials
	return f.Creds.IsExpired()
}

func Run() error {

}
