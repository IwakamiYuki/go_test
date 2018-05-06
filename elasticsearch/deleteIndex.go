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

	// Delete the index again
	_, err = client.DeleteIndex("pages").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}
