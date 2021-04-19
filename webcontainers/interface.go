package webcontainers

import (
	"blocks/web"
)

type IWebContainer interface {
	SetContent(content []string)

	web.IWeb
}
