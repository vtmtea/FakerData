package base

import "github.com/vtmtea/fakerdata/util"

var operators = []string{
	"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "170", "178", "182", "183", "184", "187", "188", // China Mobile
	"130", "131", "132", "145", "155", "156", "171", "176", "185", "186", // China Unicom
	"133", "153", "177", "180", "181", "189", // China Telecom
}

func PhoneNumber(locale string) string {
	operatorsLen := len(operators)
	return operators[util.RandomDigit(operatorsLen)] + util.RandomLenDigit(8)
}
