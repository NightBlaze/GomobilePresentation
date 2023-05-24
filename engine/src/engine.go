package engine

import (
	"log"

	enginewidgets "github.com/NightBlaze/GomobilePresentation/engine/src/engine-widgets"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/localize"
	"github.com/NightBlaze/GomobilePresentation/engine/src/internal/runtime"
	_ "golang.org/x/mobile/bind"
)

var (
	localizer *localize.Localizer
)

type InitializeConfig struct {
	EngineAssetsDirectory string
	LanguageCode          string
}

func Initialize(cfg *InitializeConfig) {
	localizer = doInitializeLocalizer(cfg.EngineAssetsDirectory, cfg.LanguageCode)

	doStartBackgroundJobs()

	enginewidgets.InitializeWidgets(localizer, struct{}{})
}

// ========
// Private funcs
// ========

func doInitializeLocalizer(assetsDirectory, languageCode string) *localize.Localizer {
	localizer, err := localize.NewLocalizer(assetsDirectory)
	if err != nil {
		log.Fatalln(err)
	}
	ok := localizer.LoadLocalization()
	if !ok {
		log.Fatalln("cant initialize Localize module")
	}
	localizer.ChangeLanguage(languageCode)
	return localizer
}

func doStartBackgroundJobs() {
	runtime.StartFreeOSMemoryLoop()
}
