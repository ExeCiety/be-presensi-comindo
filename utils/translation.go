package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var Lang *string
var Localizer *i18n.Localizer
var bundle *i18n.Bundle

func localization() {
	var defaultLang language.Tag
	var locale string
	path := *RootPath

	if Lang == nil {
		locale = "id"
	} else {
		locale = *Lang
	}

	if locale == "id" {
		defaultLang = language.Indonesian
	} else if locale == "en" {
		defaultLang = language.English
	}

	bundle = i18n.NewBundle(defaultLang)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if locale == "id" {
		_, err := bundle.LoadMessageFile(path + "utils/languages/id.json")
		if err != nil {
			log.Error("Error when reading file id.json")
			return
		}
	} else if locale == "en" {
		_, err := bundle.LoadMessageFile(path + "utils/languages/en.json")
		if err != nil {
			log.Error("Error when reading file en.json")
			return
		}
	}

	Localizer = i18n.NewLocalizer(bundle, defaultLang.String())
}

func SetLocale(locale string) {
	Lang = &locale
	localization()
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
