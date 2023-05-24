package enginewidgets

import (
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/localize"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/types"
)

var (
	localizer *localize.Localizer

	widgets = types.NewUniqueObjectsWithInt64Key[globalThingsChangeListener]()
)

// naming is hard
type globalThingsChangeListener interface {
	LocalizationDidChange(struct{})
	// and other functions like ThemeDidChange(struct{})
}

func InitializeWidgets(
	aLocalizer *localize.Localizer,
	_ struct{},
) {
	localizer = aLocalizer
}

// ========
// Internal funcs
// ========

func AddWidget(widget globalThingsChangeListener, _ struct{}) int64 {
	return widgets.Add(widget)
}

func RemoveWidgetWithID(id int64, _ struct{}) {
	widgets.RemoveWithKey(id)
}

func LocalizationDidChange() {
	ch := make(chan globalThingsChangeListener)
	go widgets.AllObjectsEnumerated(ch)
	for widget := range ch {
		widget.LocalizationDidChange(struct{}{})
	}
}

func Localizer(_ struct{}) *localize.Localizer {
	return localizer
}
