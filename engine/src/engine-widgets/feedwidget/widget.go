package feedwidget

import (
	"fmt"

	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/runtime"
)

type Widget struct {
	id         int64
	interactor *interactor
}

func Create(display Display) *Widget {
	if display == nil {
		fmt.Println("widget is nil")
		return nil
	}

	result := &Widget{}
	widgetID := enginewidgets.AddWidget(result, struct{}{})
	result.id = widgetID
	result.interactor = newInteractor(widgetID, display)
	return result
}

func (w *Widget) Free() {
	enginewidgets.RemoveWidgetWithID(w.id, struct{}{})
	w.interactor.free()
	runtime.FreeOSMemory()
}

func (w *Widget) CancelTask(taskID int64) {
	w.interactor.cancelTask(taskID)
}

func (w *Widget) InitialData() *InitialDataViewModel {
	return w.interactor.handle_initialData_Request()
}

func (w *Widget) FetchFeedAsync() int64 {
	return w.interactor.handle_fetchFeed_Request()
}

func (w *Widget) FeedItemsCount() int {
	return w.interactor.handle_feedItemsCount_Request()
}

func (w *Widget) FeedItemAtIndex(index int) []byte { // *protos.FeedItemViewModel
	return w.interactor.handle_feedItemAtIndex_Request(index)
}

// ========
// Internal methods
// ========

func (w *Widget) LocalizationDidChange(_ struct{}) {
	w.interactor.handle_localizationDidChange()
}
