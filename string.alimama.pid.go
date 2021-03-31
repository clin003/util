package util

import (
	// "bytes"
	"strings"
)

func TbkPidParse(pid string) (siteID, adzoneID string) {
	pid = strings.ReplaceAll(pid, "mm_", "")
	siteID = BetweenStr(pid, "_", "_")
	pid = strings.ReplaceAll(pid, "_"+siteID, "")
	adzoneID = BetweenStr(pid+"_", "_", "_")
	return
}
