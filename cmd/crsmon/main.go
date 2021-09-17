package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jptosso/crsmon"
)

type variables []string

func (i *variables) String() string {
	return "..."
}

func (i *variables) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var _ flag.Value = &variables{}

func main() {
	var vars variables
	dir := flag.String("path", "/opt/coreruleset/", "Path to store the configurations")
	prepend := flag.String("prepend", "/opt/coreruleset/custom.conf", "Prepend directives from file [optional]")
	flag.Var(&vars, "v", "Transaction variables")
	flag.Parse()
	policy := crsmon.NewPolicy(*dir)
	policy.Prepend(*prepend)
	for _, v := range vars {
		spl := strings.SplitN(v, "=", 2)
		policy.AddVar(spl[0], spl[1])
	}
	if err := policy.Build(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Policy was written to %s\n", path.Join(*dir, "crs.conf"))
}
