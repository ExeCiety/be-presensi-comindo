package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var DefaultLang = language.Indonesian
var Locale *string
var Lang language.Tag
var Bundle *i18n.Bundle
var Localizer *i18n.Localizer

func SetAllI18nBundles() {
	Bundle = i18n.NewBundle(DefaultLang)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	path := *RootPath

	_, err := Bundle.LoadMessageFile(path + "utils/languages/id.json")
	if err != nil {
		log.Error("Error when reading file id.json")
		return
	}

	_, err = Bundle.LoadMessageFile(path + "utils/languages/en.json")
	if err != nil {
		log.Error("Error when reading file en.json")
		return
	}
}

func SetLocale(locale string) {
	Locale = &locale

	if *Locale == "id" {
		Lang = language.Indonesian
	} else if *Locale == "en" {
		Lang = language.English
	}

	Localizer = i18n.NewLocalizer(Bundle, Lang.String())
}

func Translate(messageId string, data map[string]interface{}) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: data,
	}

	localized, err := Localizer.Localize(&localizeConfig)
	if err != nil {
		log.Error("Error when localizing message")
		return "error on " + messageId
	}

	return localized
}
