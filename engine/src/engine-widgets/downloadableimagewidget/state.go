package downloadableimagewidget

type Display interface {
	DisplayImage(data []byte)
}

type state struct {
	widgetID int64
	display  Display

	closeCh                 chan struct{}
	localizationDidChangeCh chan struct{}
	imageDidDownloadCh      chan []byte
}

func newState(
	widgetID int64,
	display Display,
) *state {
	result := &state{
		widgetID:                widgetID,
		display:                 display,
		closeCh:                 make(chan struct{}),
		localizationDidChangeCh: make(chan struct{}),
		imageDidDownloadCh:      make(chan []byte),
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
		case data := <-s.imageDidDownloadCh:
			s.display.DisplayImage(data)
		}
	}
}

func (s *state) initialData() *InitialDataViewModel {
	return &InitialDataViewModel{}
}

func (s *state) localizationDidChange() {
	s.localizationDidChangeCh <- struct{}{}
}

func (s *state) imageDidDownload(data []byte) {
	s.imageDidDownloadCh <- data
}

// ========
// Private funcs
// ========

func (s *state) doLocalizationDidChange() {
}
