package settingswidget

import (
	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/localize"
)

type Display interface {
	LocalizationDidChange(viewModel *SettingLocalizationDidChangeViewModel)
}

type state struct {
	widgetID  int64
	display   Display
	localizer *localize.Localizer

	closeCh                 chan struct{}
	localizationDidChangeCh chan struct{}
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
		}
	}
}

func (s *state) initialData() *InitialDataViewModel {
	return &InitialDataViewModel{
		RuLocalizationTitle: s.localizer.Localize("russian"),
		EnLocalizationTitle: s.localizer.Localize("english"),
	}
}

func (s *state) localizationDidChange() {
	s.localizationDidChangeCh <- struct{}{}
}

// ========
// Private funcs
// ========

func (s *state) doLocalizationDidChange() {
	viewModel := &SettingLocalizationDidChangeViewModel{
		RuLocalizationTitle: s.localizer.Localize("russian"),
		EnLocalizationTitle: s.localizer.Localize("english"),
	}
	s.display.LocalizationDidChange(viewModel)
}
