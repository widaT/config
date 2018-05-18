package kasha

import (
	"strings"
	"os"
	"gopkg.in/yaml.v2"
	"errors"
)

type Yaml struct {
	mymap  map[string]interface{}
}

func (c *Yaml) Parse(path string) (Configer,error) {
	f, err := os.Open(path)
	if err != nil {
		return  nil,err
	}
	dec := yaml.NewDecoder(f)
	if err := dec.Decode(&c.mymap);err != nil {
		return nil,err
	}
	return c ,nil
}

func (c *Yaml) read(key string) (value interface{}) {
	path := strings.Split(key,".")
	var ok bool
	for i,v:= range path {
		if i ==0 {
			value ,ok = c.mymap[v]
			if !ok {
				return nil
			}
		}else {
			switch value.(type) {
			case map[string]interface{}:
				value = value.(map[string]interface{})[v]
			default:
				return
			}
		}
	}
	return
}

func (c *Yaml)String(key string) string{
	ret,ok:= c.read(key).(string)
	if !ok {
		return ""
	}
	return ret
}

func (c *Yaml)Strings(key string)[]string{
	ret := c.String(key)
	return strings.Split(ret,";")
}

func (c *Yaml)Int(key string)(int,error) {
	v := c.read(key)
	if v == nil {
		return 0,errors.New("not found")
	}
	if ret,ok:= v.(int) ;ok{
		return ret,nil
	}else if ret,ok:= v.(int64);ok {
		return int(ret),nil
	}
	return 0, errors.New("not int value")
}

func (c *Yaml)Int64(key string) (int64,error){
	v := c.read(key)
	if v == nil {
		return 0,errors.New("not found")
	}
	if ret,ok:= v.(int) ;ok{
		return int64(ret),nil
	}else if ret,ok:= v.(int64);ok {
		return ret,nil
	}
	return 0, errors.New("not int value")
}

func (c *Yaml)Float(key string) (float32,error){
	v := c.read(key)
	if v == nil {
		return 0.0,errors.New("not found")
	}
	if ret,ok:= v.(float32) ;ok{
		return ret,nil
	}else if ret,ok:= v.(float64);ok {
		return float32(ret),nil
	}else if ret, ok := v.(int); ok {
		return float32(ret), nil
	} else if ret, ok := v.(int64); ok {
		return float32(ret), nil
	}
	return 0.0, errors.New("not float value")
}

func (c *Yaml)Float64(key string)(float64,error) {
	v := c.read(key)
	if v == nil {
		return 0.0,errors.New("not found")
	}
	if ret,ok:= v.(float32) ;ok{
		return float64(ret),nil
	}else if ret,ok:= v.(float64);ok {
		return ret,nil
	}else if ret, ok := v.(int); ok {
		return float64(ret), nil
	} else if ret, ok := v.(int64); ok {
		return float64(ret), nil
	}
	return 0.0, errors.New("not float value")
}

func (c *Yaml)Bool(key string)(bool,error){
	v := c.read(key)
	if v == nil {
		return false,errors.New("not found")
	}
	if ret,ok:= v.(bool) ;ok{
		return ret,nil
	}
	return false, errors.New("not bool value")
}

func (c *Yaml)StringEx(key ,de string) string{
	ret :=  c.String(key)
	if ret == "" {
		return de
	}
	return ret
}

func (c *Yaml)StringsEx(key string,de []string) []string{
	ret := c.Strings(key)
	if ret == nil {
		return de
	}
	return ret
}

func (c *Yaml)IntEx(key string,de int)int{
	ret,err := c.Int(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Yaml)Int64Ex(key string,de int64 ) int64{
	ret,err := c.Int64(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Yaml)FloatEx(key string,de float32 ) float32{
	ret,err := c.Float(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Yaml)Float64Ex(key string,de float64 )float64{
	ret,err := c.Float64(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Yaml)BoolEx(key string,de bool)bool{
	ret,err := c.Bool(key)
	if err != nil {
		return de
	}
	return ret
}


// Exists returns true if key exists.
func (c *Yaml) Exists(key string) bool {
	if _, ok := c.mymap[key]; ok {
		return true
	}
	return false
}