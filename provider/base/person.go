package base

import (
	"github.com/vtmtea/fakerdata/provider/en_US"
	"github.com/vtmtea/fakerdata/provider/zh_CN"
)

func Name(locale string) string {
	switch locale {
	case "zh_CN":
		return zh_CN.Name()
	case "en_US":
		return en_US.Name()
	default:
		return ""
	}
}
