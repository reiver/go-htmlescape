package htmlescape

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInternalError = erorr.Error("htmlescape: internal error")
	errNotUTF8       = erorr.Error("htmlespcape: not UTF-8")
)
