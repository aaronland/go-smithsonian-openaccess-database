package oembed

import (
	"context"
	oa_oembed "github.com/aaronland/go-smithsonian-openaccess/oembed"
)

type OEmbedDatabase interface {
	AddOEmbed(context.Context, *oa_oembed.OEmbedRecord) error
	GetRandomOEmbed(context.Context) (*oa_oembed.OEmbedRecord, error)
	GetOEmbedWithObjectURI(context.Context, string) ([]*oa_oembed.OEmbedRecord, error)
	Close() error
}
