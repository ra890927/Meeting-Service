package clients

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	gcsInstance *GcsInstance
	gcsOnce     sync.Once
)

type GcsInstance struct {
	client     *storage.Client
	bucketName string
}

func GetGCSInstance() *GcsInstance {
	if gcsInstance == nil {
		gcsOnce.Do(func() {
			ctx := context.Background()

			client, err := storage.NewClient(ctx, option.WithCredentialsFile("./docs/meeting-center-0f48b48e9bd5.json"))
			if err != nil {
				panic(err)
			}

			gcsInstance = &GcsInstance{
				client:     client,
				bucketName: "meeting-center-storage-test",
			}
		})
	}

	return gcsInstance
}

func (gcs GcsInstance) Bucket() *storage.BucketHandle {
	return gcs.client.Bucket(gcs.bucketName)
}

func (gcs GcsInstance) Close() error {
	return gcs.client.Close()
}
