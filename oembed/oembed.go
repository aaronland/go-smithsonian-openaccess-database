package oembed

import (
	"context"
	oa_oembed "github.com/aaronland/go-smithsonian-openaccess/oembed"
)

type OEmbedDatabase interface {
	AddOEmbed(context.Context, *oa_oembed.OEmbedRecord) error
	GetRandomOEmbed(context.Context) (*oa_oembed.OEmbedRecord, error)
	Close() error
}
