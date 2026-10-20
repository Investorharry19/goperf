package metrics

import "net/http"

// RouteLabel extracts the route label from an HTTP request pattern.
func RouteLabel(r *http.Request) string {
	if r == nil {
		return "unmatched"
	}
	if r.Pattern != "" {
		pattern := r.Pattern
		if pattern != "" && pattern[0] != '/' {
			for i, ch := range pattern {
				if ch == '/' {
					pattern = pattern[i:]
					break
				}
			}
		}
		if pattern != "" {
			return pattern
		}
	}
	if r.URL != nil && r.URL.Path != "" {
		return r.URL.Path
	}
	return "unmatched"
}
