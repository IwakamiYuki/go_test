package main

import (
	"context"

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

	// Add a document to the index
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("pages").
		Type("doc").
		BodyJson(tweet).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}
