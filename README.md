Welcome to crsmon

CRSMON is a test project developed to help people using [Coraza WAF](https://github.com/jptosso/coraza-waf) to automatically add [OWASP Core Ruleset](https://github.com/coreruleset/coreruleset) policies to heir project.

## Requirements

* Go 1.16+
* Access to github.com

## Important

If you are using Coraza you are required to use CGO_ENABLED=1 and install libinjection and libpcre, see [this tutorial](https://coraza.io/docs/tutorials/dependencies/)

## Example

```go
import(
    //...
    coraza"github.com/jptosso/coraza-waf"
    "github.com/jptosso/coraza-waf/seclang"
    corazagin"github.com/jptosso/coraza-gin"
    "github.com/jptosso/crsmon"
)
func main() {
    // Creates a router without any middleware by default
    r := gin.New()
    waf := coraza.NewWaf()
    // path to CRS
    path := "/opt/coreruleset/"
    parser := seclang.NewParser(waf)
    policy := crsmon.NewPolicy(path)
    policy.Build()
    parser.FromFile(path+"crs.conf")
    r.Use(corazagin.Coraza(waf))

    // Per route middleware, you can add as many as you desire.
    r.GET("/mypath", MyFunction(), Endpoint)

    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```
