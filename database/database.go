package database

import (
	"context"
	"github.com/aaronland/go-smithsonian-openaccess/oembed"
)

type OEmbedDatabase interface {
	AddOEmbed(context.Context, *oembed.OEmbedRecord) error
	GetRandomOEmbed(context.Context) (*oembed.OEmbedRecord, error)
	Close() error
}
