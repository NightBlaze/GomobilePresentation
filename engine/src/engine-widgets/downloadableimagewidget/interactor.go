package downloadableimagewidget

import (
	"context"
	"fmt"

	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/network"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/types"
)

type interactor struct {
	widgetID        int64
	tasks           types.UniqueObjects[int64, *enginewidgets.Task]
	state           *state
	imageDownloader *network.ImageDownloader
}

func newInteractor(
	widgetID int64,
	display Display,
) *interactor {
	return &interactor{
		widgetID:        widgetID,
		tasks:           types.NewUniqueObjectsWithInt64Key[*enginewidgets.Task](),
		state:           newState(widgetID, display),
		imageDownloader: network.NewImageDownloader(),
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

func (i *interactor) handle_downloadImage_Request(url string) int64 {
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

			i.doHandle_downloadImage_Request(ctx, url)
		}
	}(ctx, task)

	return taskID
}

// ========
// User actions parts
// ========

func (i *interactor) handle_initialData_Request() *InitialDataViewModel {
	return i.state.initialData()
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

func (i *interactor) doHandle_downloadImage_Request(ctx context.Context, url string) {
	data, err := i.imageDownloader.Download(ctx, url)
	if err != nil {
		fmt.Println("Cant download image", map[string]interface{}{
			"error": err,
			"url":   url,
		})
	}
	i.state.imageDidDownload(data)
}
