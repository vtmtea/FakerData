package faker

import (
	"github.com/vtmtea/fakerdata/provider/base"
	"github.com/vtmtea/fakerdata/util"
)

const defaultLocale = "zh_CN"

var validLocale = []string{"en_US", "zh_CN"}

type Engine struct {
	locale string
}

func New() *Engine {
	engine := &Engine{
		locale: defaultLocale,
	}
	return engine
}

func (engine *Engine) setLocale(locale string) *Engine {
	if util.Contain(validLocale, locale) {
		engine.locale = locale
	}
	return engine
}

func (engine *Engine) Name() string {
	return base.Name(engine.locale)
}

func PhoneNumber() string {
	return ""
}
