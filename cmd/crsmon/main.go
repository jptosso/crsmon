package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/jptosso/crsmon"
)

func main() {
	dir := flag.String("path", "/opt/coreruleset/", "Path to store the configurations")
	flag.Parse()
	policy := crsmon.NewPolicy(*dir)
	if err := policy.Build(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Policy was written to %s\n", path.Join(*dir, "crs.conf"))
}
