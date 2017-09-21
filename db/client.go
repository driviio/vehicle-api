package db

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	// "google.golang.org/api/option"
)

func NewDataStoreClient(projectID string) (*datastore.Client, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
