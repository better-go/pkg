package orm

import (
	"testing"
)

func TestDeletedAt_QueryClauses(t *testing.T) {
	t.Log(zeroTime())

	var at DeletedAt
	t.Logf("query: %+v", at.QueryClauses())
}
