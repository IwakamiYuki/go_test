package main

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v6"
)

func main(){
	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	// Create an index
	_, err = client.CreateIndex("pages").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}
