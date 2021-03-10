package graphql

import (
	"context"
	"net/http"

	gqV2 "github.com/machinebox/graphql"
	gqV1 "github.com/shurcooL/graphql"
)

// graphql client:
type Client struct {
	v1 *gqV1.Client // query: go struct
	v2 *gqV2.Client // query: raw json string
}

func NewClient(url string, httpClient *http.Client, opts ...gqV2.ClientOption) *Client {
	return &Client{
		v1: gqV1.NewClient(url, httpClient),
		v2: gqV2.NewClient(url, opts...),
	}
}

// HTTP GET: get
func (m *Client) Query(ctx context.Context, query interface{}, variables map[string]interface{}) error {
	return m.v1.Query(ctx, query, variables)
}

// HTTP POST: create
func (m *Client) Mutation(ctx context.Context, payload interface{}, variables map[string]interface{}) error {
	return m.v1.Mutate(ctx, payload, variables)
}

// HTTP GET: by raw query string
func (m *Client) QueryRaw(ctx context.Context, req *gqV2.Request, resp interface{}) (err error) {
	// run it and capture the response
	return m.v2.Run(ctx, req, &resp)
}
