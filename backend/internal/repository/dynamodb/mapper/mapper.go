package mapper

import (
	"strings"
)

func extractKey(prefix string, key string) string {
	return strings.TrimPrefix(key, prefix)
}
