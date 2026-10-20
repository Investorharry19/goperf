package store

import (
	"strconv"
	"strings"
)

// Rebind rewrites positional placeholders for the target SQL dialect.
func Rebind(driver, query string) string {
	if driver != "postgres" {
		return query
	}
	var b strings.Builder
	inSingle := false
	idx := 0
	for i := 0; i < len(query); i++ {
		ch := query[i]
		if ch == '\'' {
			if inSingle && i+1 < len(query) && query[i+1] == '\'' {
				b.WriteByte(ch)
				b.WriteByte(query[i+1])
				i++
				continue
			}
			inSingle = !inSingle
			b.WriteByte(ch)
			continue
		}
		if ch == '?' && !inSingle {
			idx++
			b.WriteString("$")
			b.WriteString(strconv.Itoa(idx))
			continue
		}
		b.WriteByte(ch)
	}
	return b.String()
}
