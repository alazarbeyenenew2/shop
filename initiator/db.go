package initiator

import (
	"context"
	"os"

	"cloud.google.com/go/spanner"
	"github.com/alazarbeyenenew2/shop/internal/pkg/db"
	"go.uber.org/zap"
)

func InitDB(logger *zap.Logger, ctx context.Context) *spanner.Client {
	spannerProject := os.Getenv("SPANNER_PROJECT")
	spannerInstance := os.Getenv("SPANNER_INSTANCE")
	spannerDatabase := os.Getenv("SPANNER_DATABASE")

	client, err := db.NewSpannerClient(ctx, spannerProject, spannerInstance, spannerDatabase)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return client
}
