package workers

import (
	"context"

	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/models"
	id "github.com/satori/go.uuid"
)

type FeedWorker struct {
}

func NewFeedWorker() *FeedWorker {
	return &FeedWorker{}
}

func (w *FeedWorker) FetchFeed(ctx context.Context) []models.FeedItem {
	result := make([]models.FeedItem, 0)
	for i := 0; i < 100; i++ {
		result = append(result, models.FeedItem{
			ID:    id.NewV4().String(),
			Title: id.NewV4().String(),
		})
	}

	return result
}
