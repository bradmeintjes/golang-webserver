package routing

import (
	"path"
	"strings"
)

// Routing through the use of this package is accomplished by iteratively
// processing each segment of the url, and passing the shifted path down to
// child handlers to do the same until the message is ultimately
// handled by the appropriate conmponent

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}