package yql

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

func TestBuildQuery(t *testing.T){
    fields := []string{"x","y","z"}
    tables := []string{"table1", "table2"}
    where := []string{"x == 1", "y > 0", "z != 0"}
    andOr := []bool{true, false}

    expected := "select x,y,z from table1,table2 where x == 1 AND y > 0 OR z != 0"
    result := BuildQuery(fields, tables, where, andOr)
    if result != expected{
        t.Errorf("Query string is different than expected.\nExpected:%s\nGot:%s\n", expected, result)
    }
}
