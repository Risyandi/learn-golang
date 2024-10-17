package config

import (
	"encoding/json"
	"os"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func I18n() *fiberi18n.Config {
	langLists := []language.Tag{}

	for _, lang := range conf.Lang.AcceptedLanguages {

		langLists = append(langLists, language.Make(lang))
	}

	return &fiberi18n.Config{
		RootPath:         conf.Lang.PathLanguage,
		AcceptLanguages:  langLists,
		DefaultLanguage:  language.Make(conf.Lang.DefaultLanguage),
		FormatBundleFile: "json",
		UnmarshalFunc:    json.Unmarshal,
		Loader:           fiberi18n.LoaderFunc(os.ReadFile),
		LangHandler: func(ctx *fiber.Ctx, defaultLang string) string {
			return ctx.Get("Accept-Language")
		},
	}

}
