package localize

import (
	"errors"
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Localizer struct {
	assetsDirectory string

	bundle          *i18n.Bundle
	i18nLocalizer   *i18n.Localizer
	i18nLocalizerMx *sync.RWMutex

	currentLanguage string
}

func NewLocalizer(assetsDirectory string) (*Localizer, error) {
	if len(assetsDirectory) == 0 {
		return nil, errors.New("assets directory is nil")
	}

	return &Localizer{
		assetsDirectory: assetsDirectory,
		i18nLocalizerMx: &sync.RWMutex{},
	}, nil
}

func (l *Localizer) LoadLocalization() bool {
	if l == nil {
		return false
	}

	localizationDirectory := l.assetsDirectory + "/Localization/"

	l.bundle = i18n.NewBundle(language.English)
	l.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	enFilePath := localizationDirectory + "en.toml"
	_, err := l.bundle.LoadMessageFile(enFilePath)
	if err != nil {
		fmt.Println("Cant load English localization", map[string]interface{}{
			"error":    err,
			"filePath": enFilePath,
		})
		return false
	}

	ruFilePath := localizationDirectory + "ru.toml"
	_, err = l.bundle.LoadMessageFile(ruFilePath)
	if err != nil {
		fmt.Println("Cant load Russian localization", map[string]interface{}{
			"error":    err,
			"filePath": enFilePath,
		})
		return false
	}

	l.ChangeLanguage(language.English.String())

	return true
}

func (l *Localizer) ChangeLanguage(lang string) {
	if l == nil {
		return
	}

	l.i18nLocalizerMx.Lock()
	// always fallback to English
	l.i18nLocalizer = i18n.NewLocalizer(l.bundle, lang, language.English.String())
	l.currentLanguage = lang
	l.i18nLocalizerMx.Unlock()
}

func (l *Localizer) Localize(key string) string {
	if l == nil {
		return key
	}

	l.i18nLocalizerMx.RLock()
	defer l.i18nLocalizerMx.RUnlock()

	localization, err := l.i18nLocalizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: key,
		},
	})
	if err != nil {
		fmt.Println("Cant localize", map[string]interface{}{
			"error":           err,
			"key":             key,
			"currentLanguage": l.currentLanguage,
		})
		return key
	}

	return localization
}
