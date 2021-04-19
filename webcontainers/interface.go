package webcontainers

import (
	"blocks/web"
)

type IWebContainer interface {
	SetMarkdown(markdown []string)

	web.IWeb
}
