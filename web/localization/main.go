package localization

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"slices"
	"strings"
)

const (
	DefaultLanguageCode = "en_US"
)

type Localization struct {
	Name    string
	Code    string
	Strings map[string]string
}

type Item struct {
	Name string
	Code string
}

var (
	localizations map[string]Localization
	items         []Item
)

func init() {
	localizations = make(map[string]Localization)

	files, err := os.ReadDir("assets/l10n")
	if err != nil {
		log.Fatalf("unable to read l10n files: %s", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		code := strings.TrimSuffix(file.Name(), ".json")

		contents, err := os.ReadFile(fmt.Sprintf("assets/l10n/%s", file.Name()))
		if err != nil {
			log.Errorf("unable to read l10n file %s: %s", file.Name(), err)
			continue
		}

		m := make(map[string]string)
		if err := json.Unmarshal(contents, &m); err != nil {
			log.Errorf("unable to unmarshal l10n file %s: %s", file.Name(), err)
			continue
		}

		localizations[code] = Localization{
			Name:    m["__name__"],
			Code:    code,
			Strings: m,
		}
	}

	for _, loc := range localizations {
		items = append(items, Item{
			Code: loc.Code,
			Name: loc.Name,
		})
	}

	slices.SortFunc(items, func(a, b Item) int {
		return strings.Compare(a.Code, b.Code)
	})
}

func List() []Item {
	return items
}

func Exists(name string) bool {
	_, ok := localizations[name]
	return ok
}

func Code(c echo.Context) string {
	langCookie, err := c.Cookie("playtime-l10n")
	if err == nil && Exists(langCookie.Value) {
		return langCookie.Value
	} else {
		return DefaultLanguageCode
	}
}

func LocalizeCtx(c echo.Context, s string, args ...any) string {
	return Localize(Code(c), s, args)
}

func Localize(lang, s string, args []any) string {
	loc, ok := localizations[lang]
	if !ok {
		return fmt.Sprintf("unknown l10n language: %s", lang)
	}

	line, ok := loc.Strings[s]
	if !ok {
		if lang != DefaultLanguageCode {
			return Localize(DefaultLanguageCode, s, args)
		} else {
			return fmt.Sprintf("unknown l10n line: %s.%s", lang, s)
		}
	}

	return fmt.Sprintf(line, args...)
}
