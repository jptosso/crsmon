package crsmon

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type CrsMode int

const CRS_VERSION = "3.4"

const (
	MODE_SCORING        CrsMode = 1
	MODE_SELF_CONTAINED CrsMode = 2
)

type auditEngine struct {
	Engine  string
	Args    string
	pattern string
}

type Policy struct {
	mode       CrsMode
	directives map[string]string
	mimes      []string
	audit      []auditEngine
	auditParts *AuditLogParts

	cachepath string
}

/*
 * LIBRARY SPECIFICS
 */

/*
 * CRS SPECIFIC
 */

func (c *Policy) SetMode(mode CrsMode) {
	c.mode = mode
}

func (c *Policy) DisableRules(ids []int) {

}

/*
 * CORAZA SPECIFIC
 */

func (c *Policy) SetTmp(dir string) {

}

func (c *Policy) SetDataDir(dir string) {
}

func (c *Policy) AllowRequestBodyAccess(rba bool) {
	c.directives["SecRequestBodyAccess"] = onoff(rba)
}

func (c *Policy) AllowResponseBodyAccess(rba bool) {
	c.directives["SecResponseBodyAccess"] = onoff(rba)
}

func (c *Policy) AddResponseMime(mime string) {
	c.mimes = append(c.mimes, mime)
}

func (c *Policy) AllowInterruption(it bool) {
	if it {
		c.directives["SecRuleEngine"] = "On"
	} else {
		c.directives["SecRuleEngine"] = "DetectionOnly"
	}
}

func (c *Policy) AddAuditEngine(engine string, args string, pattern string) {
	c.audit = append(c.audit, auditEngine{engine, args, pattern})
}

func (c *Policy) SetAuditLogParts(alp *AuditLogParts) {
	c.auditParts = alp
}

func (c *Policy) OverwriteDirective(directive string, arguments string) {
	c.directives[directive] = arguments
}

func (c *Policy) AuditLogParts() *AuditLogParts {
	return c.auditParts
}

func (c *Policy) Build() error {
	bf := strings.Builder{}
	for key, value := range c.directives {
		bf.WriteString(fmt.Sprintf("%s %s\n", key, value))
	}
	lastminor, err := getLastMinor()
	if err != nil {
		return err
	}
	err = downloadCrs(lastminor, c.cachepath)
	if err != nil {
		return err
	}
	file := path.Join(c.cachepath, "crs.conf")
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	os.Remove(file)
	bf.WriteByte('\n')
	if c.mode == MODE_SCORING {
		bf.WriteString("SecDefaultAction \"phase:1,log,auditlog,pass\"\nSecDefaultAction \"phase:2,log,auditlog,pass\"")
	} else if c.mode == MODE_SELF_CONTAINED {
		bf.WriteString("SecDefaultAction \"phase:1,log,auditlog,deny,status:403\"\n SecDefaultAction \"phase:2,log,auditlog,deny,status:403\"")
	}
	bf.WriteString("SecAction \"id:900990,phase:1,nolog,pass,t:none,setvar:tx.crs_setup_version=340\"\n")
	//TODO add more default actions
	data = append([]byte(bf.String()), data...)
	return ioutil.WriteFile(file, data, 0644)
}

func NewPolicy(cachepath string) *Policy {
	defaults := map[string]string{
		"SecUnicodeMap":        "20127",
		"SecCookieFormat":      "0",
		"SecArgumentSeparator": "&",
	}
	ap := NewAuditLogParts()
	ap.EnableAll()
	return &Policy{
		directives: defaults,
		auditParts: ap,
		cachepath:  cachepath,
	}
}
