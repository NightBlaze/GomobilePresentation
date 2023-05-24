package feedwidget

import (
	"fmt"

	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/helpers/array"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/localize"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/models"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/protos"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/types"
	"google.golang.org/protobuf/proto"
)

type Display interface {
	LocalizationDidChange(data []byte) // protos.FeedLocalizationDidChangeViewModel
	FeedDidFetch()
}

type state struct {
	widgetID  int64
	display   Display
	localizer *localize.Localizer

	closeCh                 chan struct{}
	localizationDidChangeCh chan struct{}
	feedDidFetchCh          chan []models.FeedItem
	feedItemsCountCh        chan chan int
	feedItemAtIndexCh       chan types.Pair[chan []byte, int]

	feedItems []models.FeedItem
}

func newState(
	widgetID int64,
	display Display,
) *state {
	result := &state{
		widgetID:                widgetID,
		display:                 display,
		localizer:               enginewidgets.Localizer(struct{}{}),
		closeCh:                 make(chan struct{}),
		localizationDidChangeCh: make(chan struct{}),
		feedDidFetchCh:          make(chan []models.FeedItem),
		feedItemsCountCh:        make(chan chan int),
		feedItemAtIndexCh:       make(chan types.Pair[chan []byte, int]),
	}

	go result.startInbox()

	return result
}

func (s *state) free() {
	close(s.closeCh)
	s.display = nil
}

func (s *state) startInbox() {
	for {
		select {
		case <-s.closeCh:
			return
		case <-s.localizationDidChangeCh:
			s.doLocalizationDidChange()
		case feed := <-s.feedDidFetchCh:
			s.doUpdateFeedItems(feed)
		case ch := <-s.feedItemsCountCh:
			ch <- len(s.feedItems)
		case pair := <-s.feedItemAtIndexCh:
			s.doFeedItemAtIndex(pair.First, pair.Second)
		}
	}
}

func (s *state) initialData() *InitialDataViewModel {
	return &InitialDataViewModel{}
}

func (s *state) localizationDidChange() {
	s.localizationDidChangeCh <- struct{}{}
}

func (s *state) feedDidFetch(feed []models.FeedItem) {
	s.feedDidFetchCh <- feed
}

func (s *state) feedItemsCount(ch chan int) {
	s.feedItemsCountCh <- ch
}

func (s *state) feedItemAtIndex(ch chan []byte, index int) {
	s.feedItemAtIndexCh <- types.NewPair(ch, index)
}

// ========
// Private funcs
// ========

func (s *state) doLocalizationDidChange() {
	viewModel := &protos.FeedLocalizationDidChangeViewModel{
		FeedItems: array.Map(s.feedItems, func(feedItem models.FeedItem) *protos.FeedItemLocalizableViewModel {
			return &protos.FeedItemLocalizableViewModel{
				Id:      feedItem.ID,
				Caption: s.localizer.Localize("feed.feed_item_caption"),
			}
		}),
	}
	data, err := proto.Marshal(viewModel)
	if err != nil {
		fmt.Println("Cant serialize FeedLocalizationDidChangeViewModel", err)
	}
	s.display.LocalizationDidChange(data)
}

func (s *state) doUpdateFeedItems(feed []models.FeedItem) {
	s.feedItems = feed
	s.display.FeedDidFetch()
}

func (s *state) doFeedItemAtIndex(ch chan []byte, index int) {
	if index < 0 || index >= len(s.feedItems) {
		ch <- nil
		return
	}

	viewModel := s.feedItemToViewModel(s.feedItems[index])
	data, err := proto.Marshal(viewModel)
	if err != nil {
		fmt.Println("Cant serialize FeedItemViewModel", err)
	}
	ch <- data
}

func (s *state) feedItemToViewModel(feedItem models.FeedItem) *protos.FeedItemViewModel {
	return &protos.FeedItemViewModel{
		Id:       feedItem.ID,
		Caption:  s.localizer.Localize("feed.feed_item_caption"),
		Title:    feedItem.Title,
		ImageUrl: "https://hsto.org/getpro/habr/avatars/6c7/234/254/6c723425499b09b2554dbed3e5a246ad.jpg",
	}
}
