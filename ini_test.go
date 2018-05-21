package config

import (
	"fmt"
	"os"
	"testing"
)

func TestIni(t *testing.T) {

	var (
		inicontext = `
appname = aaaaa
httpport = 8080
mysqlport = 3600
pi = 3.141597
runmode = dev
autorender = false
copyrequestbody = true
[demo]
key1=wida
key2 = xie
CaseInsensitive = true
peers = one;two;three
`

		keyValue = map[string]interface{}{
			"appname":               "aaaaa",
			"httpport":              8080,
			"mysqlport":             int64(3600),
			"pi":                    3.141597,
			"runmode":               "dev",
			"autorender":            false,
			"copyrequestbody":       true,
			"demo.key1":            "wida",
			"demo.key2":            "xie",
			"demo.CaseInsensitive": true,
			"demo.peers":           []string{"one", "two", "three"},
			"null":                  "",
			"demo2.key1":           "",
			"error":                 "",
			"emptystrings":          []string{},
		}
	)

	f, err := os.Create("testini.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(inicontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("testini.conf")
	ini := Ini{}
	iniconf, err := ini.Parse( "testini.conf")
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range keyValue {
		var err error
		var value interface{}
		switch v.(type) {
		case int:
			value, err = iniconf.Int(k)
		case int64:
			value, err = iniconf.Int64(k)
		case float64:
			value, err = iniconf.Float(k)
		case bool:
			value, err = iniconf.Bool(k)
		case []string:
			value = iniconf.Strings(k)
		case string:
			value = iniconf.String(k)
		}
		if err != nil {
			t.Fatalf("get key %q value fail,err %s", k, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", value) {
			t.Fatalf("get key %q value, want %v got %v .", k, v, value)
		}

	}
}
