package main

import (
	"fmt"
	"os"
	"strconv"

	aero "github.com/aerospike/aerospike-client-go"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  fmt.Println("Starting app...")

  host := os.Getenv("host")
  if (len(host) == 0) {
	host = "127.0.0.1"
  }

  portEnv := os.Getenv("port")
  var port int64
  if (len(portEnv) == 0) {
	port = 3000
  } else {
	port,_ = strconv.ParseInt(portEnv, 0, 8)
  }

  // define a client to connect to
  client, err := aero.NewClient(host, int(port))
  panicOnError(err)

  key, err := aero.NewKey("hotels", "aerospike", "key")
  panicOnError(err)

  // define some bins with data
  bins := aero.BinMap{
    "bin1": 42,
    "bin2": "An elephant is a mouse with an operating system",
    "bin3": []interface{}{"Go", 2009},
  }

  // write the bins
  err = client.Put(nil, key, bins)
  panicOnError(err)

  // read it back!
  rec, err := client.Get(nil, key)
  panicOnError(err)

  // print values
  fmt.Println("BINS:")
  fmt.Println(rec.Bins)
  fmt.Println("RECORD:")
  fmt.Println(rec)
  fmt.Println("\n")

  // delete the key, and check if key exists
  existed, err := client.Delete(nil, key)
  panicOnError(err)
  fmt.Printf("Record existed before delete? %v\n", existed)
}
