package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v6"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main(){
	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	// Search with a term query
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
			Index("pages").   // search in index "twitter"
			Query(termQuery).   // specify the query
			From(0).Size(10).   // take documents 0-9
			Pretty(true).       // pretty print request and response JSON
			Do(ctx)             // execute
	if err != nil {
			// Handle error
			panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over the process, see below.
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
			t := item.(Tweet)
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}

	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through the search results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
			fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

			// Iterate through results
			for _, hit := range searchResult.Hits.Hits {
					// hit.Index contains the name of the index

					// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
					var t Tweet
					err := json.Unmarshal(*hit.Source, &t)
					if err != nil {
							// Deserialization failed
					}

					// Work with tweet
					fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
			}
	} else {
			// No hits
			fmt.Print("Found no tweets\n")
	}
}
