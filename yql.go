package stockfeed

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "bytes"
    "encoding/json"
)

/**
    Represents an instance of the YQL Api
 */
type YQL struct {
    Url, Env, Fmt string
}

/**
    The data response received from a YQL Query
 */
type Response struct{
    Created, Lang string
    Results map[string]interface{}
}

/**
    Send a YQL query to the Yahoo! servers specified in the YQL object.  Returns Response struct with
    response from server.
 */
func (y *YQL) Query(q string) (Response, error){
    var r Response
    queryUrl := y.buildURL(q)
    resp, err := http.Get(queryUrl)
    
    // HTTP error
    if err != nil {
        return r, err
    }
    
    resp_str, err := ioutil.ReadAll(resp.Body)
    // Read error
    if err != nil {
        return r, err
    }
    defer resp.Body.Close()
    
    // Build the Response
    var jsonArray map[string]interface{}
    err = json.Unmarshal(resp_str, &jsonArray)
    // JSON error
    if err != nil {
        return r, err
    }

    // TODO: Revisit once understand of type system is better.  Lacks safety (i think?)
    jsonArray = (jsonArray["query"]).(map[string]interface{})
    r = Response{ jsonArray["created"].(string), jsonArray["lang"].(string), jsonArray["results"].(map[string]interface{}) }

    return r, err
}

/**
    Build the URL for a YQL query
 */
func (y *YQL) buildURL(query string) (string){
    return y.Url + "?q=" + url.QueryEscape(query) + "&format=" + y.Fmt + "&env=" + url.QueryEscape(y.Env)
}

/**
    Helper function to create the actual YQL query string
 */
func BuildQuery(fields []string, tables []string, where []string, andOr bool) (string) {
    // Validate
    if len(fields) == 0 || len(tables) == 0 || len(where) == 0{
        return ""
    }
    
    // Setup Buffer
    var query_buffer bytes.Buffer
    
    // Select
    query_buffer.WriteString("select ")
    for key, value := range fields{
        if key > 0 {
            query_buffer.WriteString(",")
        }
        query_buffer.WriteString(value) 
    }
    
    // From
    query_buffer.WriteString(" from ")
    for key, value := range tables{
        if key > 0 {
            query_buffer.WriteString(",")
        }
        query_buffer.WriteString(value) 
    }
    
    // Where
    query_buffer.WriteString(" where ")
    for key, value := range where{
        if key > 0 {
            if andOr == true{
                query_buffer.WriteString(" AND ")
            } else {
                query_buffer.WriteString(" OR ")
            }
        }
        query_buffer.WriteString(value) 
    }
    
    return query_buffer.String()
}

