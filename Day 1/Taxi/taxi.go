/*Exercise
The NYC Taxi data has URLs in the following format:

https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2019-01.csv
https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2019-02.csv
Using several HTTP HEAD requests, determine the size in megabytes of all the CSV files (both yellow and green) in 2019.

Didn't do entire assignment...

*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2019-01.parquet
	u := "https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2019-01.parquet"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)

	//This works as well
	/*req, err := http.NewRequest(http.MethodHead, u, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}*/

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	size := resp.Header.Get("Content-Length")

	fmt.Println("size: ", size)
}
