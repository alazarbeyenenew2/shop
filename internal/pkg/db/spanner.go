package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
)

func NewSpannerClient(ctx context.Context, projectID, instanceID, databaseID string) (*spanner.Client, error) {

	db := fmt.Sprintf(
		"projects/%s/instances/%s/databases/%s",
		projectID,
		instanceID,
		databaseID,
	)

	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return nil, err
	}

	return client, nil
}
