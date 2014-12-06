package stockfeed

import (
    "testing"
    // "fmt"
)

func TestYQLQuery(t *testing.T) {
    y := YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    r, err := y.Query("select * from yahoo.finance.quote where symbol in (\"YHOO\",\"AAPL\",\"GOOG\",\"MSFT\")")
    
    if err != nil || r == nil {
        t.Errorf("Query Error: %s", err)
    }
}

