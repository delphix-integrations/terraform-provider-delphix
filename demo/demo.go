package main

import (
	"context"
	"dct-go/swagger"
	"fmt"
	"os"
)

func main() {
	var apiKey swagger.APIKey

	apiKey.Key = "2.a8lvIpB2PFfkOfJezQsxlxaNSzuVZZGmOedk6Hz0NYM82bqun1Lp8zNRDDOAs2Xq"
	apiKey.Prefix = "apk"

	ctx := context.WithValue(context.Background(), swagger.ContextAPIKey, apiKey)

	cfg := swagger.NewConfiguration()
	cfg.Host = "localhost"
	cfg.Scheme = "HTTPS"

	bookmarkRes, httpRes, err := swagger.NewAPIClient(cfg).BookmarksApi.GetBookmarks(ctx)

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print("\n__________RES__________\n")
	fmt.Print(bookmarkRes)
	fmt.Print("\n__________HTTP-RESP__________\n")
	fmt.Print(httpRes)
}
