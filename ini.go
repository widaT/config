package kasha

import (
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"
)

const middle = "========="
const dofualtSession = "default"

type Ini struct {
	mymap  map[string]string
	strcet string
}

func (c *Ini) Parse(path string) (Configer,error) {
	c.strcet = dofualtSession
	c.mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		return  nil,err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return  nil,err
		}
		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}
		if len(second) == 0 {
			continue
		}
		key := c.strcet + middle + frist
		c.mymap[key] = strings.TrimSpace(second)
	}
	return c ,nil
}

func (c *Ini) read(key string) string {
	keys:= strings.Split(key,".")
	if len(keys)  ==1 {
		key = dofualtSession + middle + keys[0]
	}else if len(keys) == 2 {
		key = keys[0] + middle + keys[1]
	}else {
		return ""
	}
	v, found := c.mymap[key]
	if !found {
		return ""
	}
	return v
}

func (c *Ini)String(key string) string{
	return c.read(key)
}

func (c *Ini)Strings(key string)[]string{
	ret := c.String(key)
	return strings.Split(ret,";")
}

func (c *Ini)Int(key string)(int,error) {
	ret :=  c.read(key)
	return strconv.Atoi(ret)
}

func (c *Ini)Int64(key string) (int64,error){
	ret :=  c.read(key)
	return strconv.ParseInt(ret,10,64)
}

func (c *Ini)Float(key string) (float32,error){
	ret :=  c.read(key)
	f,err:= strconv.ParseFloat(ret,32)
	if err != nil {
		return 0.0,err
	}
	return float32(f),nil
}

func (c *Ini)Float64(key string)(float64,error) {
	ret :=  c.read(key)
	return  strconv.ParseFloat(ret,64)
}

func (c *Ini)Bool(key string)(bool,error){
	ret := c.read(key)
	return strconv.ParseBool(ret)
}

func (c *Ini)StringEx(key ,de string) string{
	ret :=  c.read(key)
	if ret == "" {
		return de
	}
	return ret
}

func (c *Ini)StringsEx(key string,de []string) []string{
	ret := c.Strings(key)
	if ret == nil {
		return de
	}
	return ret
}

func (c *Ini)IntEx(key string,de int)int{
	ret,err := c.Int(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Ini)Int64Ex(key string,de int64 ) int64{
	ret,err := c.Int64(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Ini)FloatEx(key string,de float32 ) float32{
	ret,err := c.Float(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Ini)Float64Ex(key string,de float64 )float64{
	ret,err := c.Float64(key)
	if err != nil {
		return de
	}
	return ret
}

func (c *Ini)BoolEx(key string,de bool)bool{
	ret,err := c.Bool(key)
	if err != nil {
		return de
	}
	return ret
}
