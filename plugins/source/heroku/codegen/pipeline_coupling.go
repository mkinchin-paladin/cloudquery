// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func PipelineCouplings() *schema.Table {
	return &schema.Table{
		Name:        "heroku_pipeline_couplings",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#pipeline-coupling-attributes",
		Resolver:    fetchPipelineCouplings,
		Columns: []schema.Column{
			{
				Name:     "app",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("App"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "pipeline",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Pipeline"),
			},
			{
				Name:     "stage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Stage"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchPipelineCouplings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.PipelineCouplingList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
