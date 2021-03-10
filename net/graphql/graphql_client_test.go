package graphql

import (
	"context"
	"github.com/shurcooL/graphql"
	"testing"
)

/*

ref:
	- https://github.com/hhstore/blog/issues/68

	- usage:
		- https://graphql.org/swapi-graphql/

		- ```

			query {
			  allStarships{

				totalCount,
				pageInfo{
				  hasNextPage,
				  hasPreviousPage
				}
				starships {
				  name,
				  model,
				  crew,
				  passengers,
				}
			  }
			}

			```

curl 'http://34.87.21.221:8000/subgraphs/name/SkyWalker'
	-H 'Accept-Encoding: gzip, deflate, br'
	-H 'Content-Type: application/json'
	-H 'Accept: application/json'
	-H 'Connection: keep-alive'
	-H 'DNT: 1'
	-H 'Origin: file://'
	--data-binary '{"query":"# Write your query or mutation here\n{\n  markets {\n    id\n    description\n  }\n  market(id: \"0x7f103852bac2de9e524ab5ade5757f1410fbd861\") {\n    id\n    description\n    marketType\n    status\n    balances\n    cliamBalance\n  }\n}\n"}' --compressed


*/
func TestClient_Query(t *testing.T) {

	ctx := context.Background()

	url := "https://graphql.org/swapi-graphql/"
	url2 := "http://34.87.21.221:8000/subgraphs/name/SkyWalker"

	client1 := NewClient(url, nil)
	client2 := NewClient(url2, nil)

	var query struct {
		Film struct {
			ID    graphql.String
			Title graphql.String
		} `graphql:"film(id:1, fileID:1)"`
	}

	err := client1.Query(ctx, &query, nil)
	t.Logf("query1 resp: %v, %+v", err, query)

	/*

	  markets{
	    id
	    description
	  }
	*/
	var query2 struct {
		// query many:
		Markets []struct {
			ID          string
			Description string
		}

		// query one:
		// market(id: ID!block: Block_height): Market
		Market struct {
			ID          string
			Description string
			Balances    string
			CliamBalance string
		} `graphql:"market(id: \"0x7f103852bac2de9e524ab5ade5757f1410fbd861\")"`
	}

	err = client2.Query(ctx, &query2, nil)
	t.Logf("query2 resp: %v, %+v", err, query2)

}

func TestClient_QueryRaw(t *testing.T) {
	qs := ""
	t.Log(qs)
}
