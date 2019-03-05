package google

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

const (
	bucket = "purviata"
)

// GCS is a struct that enables interaction
// with a Google Cloud storage object
type GCS struct {
	client *storage.Client
}

// NewGCS creates a new storage Client to upload to GCS
func NewGCS() *GCS {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return &GCS{
			client: client,
		}
	}
	return nil
}

func ReadFromBucket(object string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(context.Background())
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
