package feedwidget

import (
	"context"

	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/localize"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/types"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/workers"
)

type interactor struct {
	widgetID   int64
	tasks      types.UniqueObjects[int64, *enginewidgets.Task]
	state      *state
	localizer  *localize.Localizer
	feedWorker *workers.FeedWorker
}

func newInteractor(
	widgetID int64,
	display Display,
) *interactor {
	return &interactor{
		widgetID:   widgetID,
		tasks:      types.NewUniqueObjectsWithInt64Key[*enginewidgets.Task](),
		state:      newState(widgetID, display),
		localizer:  enginewidgets.Localizer(struct{}{}),
		feedWorker: workers.NewFeedWorker(),
	}
}

func (i *interactor) free() {
	i.doCancelAllTasks()
	i.state.free()
}

func (i *interactor) cancelTask(taskID int64) {
	if task := enginewidgets.TaskWithID(i.tasks, taskID, struct{}{}); task != nil {
		task.Cancel(struct{}{})
	}
	i.tasks.RemoveWithKey(taskID)
}

// ========
// User actions parts
// ========

func (i *interactor) handle_initialData_Request() *InitialDataViewModel {
	return i.state.initialData()
}

func (i *interactor) handle_fetchFeed_Request() int64 {
	ctx, cancel := context.WithCancel(context.Background())

	task := enginewidgets.NewTask(ctx, cancel, struct{}{})
	taskID := i.tasks.Add(task)
	task.SetID(taskID, struct{}{})

	go func(ctx context.Context, task *enginewidgets.Task) {
		defer func() {
			i.tasks.RemoveWithKey(task.ID(struct{}{}))
		}()

		select {
		case <-ctx.Done():
			return
		default:
			if task.IsCancelled(struct{}{}) {
				return
			}

			i.doHandle_fetchFeed_Request(ctx)
		}
	}(ctx, task)

	return taskID
}

func (i *interactor) handle_feedItemsCount_Request() int {
	ch := make(chan int, 1)
	i.state.feedItemsCount(ch)
	result := <-ch
	return result
}

func (i *interactor) handle_feedItemAtIndex_Request(index int) []byte {
	ch := make(chan []byte, 1)
	i.state.feedItemAtIndex(ch, index)
	result := <-ch
	return result
}

func (i *interactor) handle_localizationDidChange() {
	i.state.localizationDidChange()
}

// ========
// Private funcs
// ========

func (i *interactor) doCancelAllTasks() {
	i.tasks.RemoveAllWithFn(
		func(task *enginewidgets.Task) {
			task.Cancel(struct{}{})
		},
	)
}

func (i *interactor) doHandle_fetchFeed_Request(ctx context.Context) {
	feed := i.feedWorker.FetchFeed(ctx)
	i.state.feedDidFetch(feed)
}
