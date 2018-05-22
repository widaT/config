package config

import (
	"testing"
	"os"
	"fmt"
)

func TestYaml(t *testing.T) {
	var (
		yamlcontext = `
"appname": test
"httpport": 8080
"mysqlport": 3600
"PI": 3.1415976
"runmode": dev
"autorender": false
"copyrequestbody": true
"PATH": GOPATH
"path1": GOPATH
"path2": GOPATH
"emptystrings":
- aaa
- dddd
`

		keyValue = map[string]interface{}{
			"appname":         "test",
			"httpport":        8080,
			"mysqlport":       int64(3600),
			"PI":              3.1415976,
			"runmode":         "dev",
			"autorender":      false,
			"copyrequestbody": true,
			"PATH":            "GOPATH",
			"path1":          "GOPATH",
			"path2":           "GOPATH",
			"error":           "",
			"emptystrings":    []string{"aaa","dddd"},
		}
	)
	f, err := os.Create("testyaml.yaml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(yamlcontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer func() {
		os.Remove("testyaml.yaml")
	}()


	config ,err:= NewConfig("yaml","testyaml.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}

	if config.String("appname") != "test" {
		t.Fatal("appname not equal to beeapi")
	}

	for k, v := range keyValue {
		var (
			value interface{}
			err   error
		)

		switch v.(type) {
		case int:
			value, err = config.Int(k)
		case int64:
			value, err = config.Int64(k)
		case float64:
			value, err = config.Float64(k)
		case bool:
			value, err = config.Bool(k)
		case []string:
			value = config.Strings(k)
		case string:
			value = config.String(k)
		}
		if err != nil {
			t.Errorf("get key %q value fatal,%v err %s", k, v, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", value) {
			t.Errorf("get key %q value, want %v got %v .", k, v, value)
		}

	}
}