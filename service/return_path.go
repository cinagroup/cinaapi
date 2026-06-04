package service

import (
	"strings"

	"github.com/cinagroup/cinatoken/common"
	"github.com/cinagroup/cinatoken/setting/system_setting"
)

func PaymentReturnURL(suffix string) string {
	base := strings.TrimRight(system_setting.ServerAddress, "/")
	return base + common.ThemeAwarePath(suffix)
}
