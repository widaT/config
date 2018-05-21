package config

type Configer interface {
	String(key string) string
	Strings(key string) []string
	Int(key string)(int,error)
	Int64(key string) (int64,error)
	Float(key string) (float32,error)
	Float64(key string)(float64,error)
	Bool(key string)(bool,error)
	StringEx(key ,de string) string
	StringsEx(key string,de []string) []string
	IntEx(key string,de int)int
	Int64Ex(key string,de int64 ) int64
	FloatEx(key string,de float32 ) float32
	Float64Ex(key string,de float64 )float64
	BoolEx(key string,de bool)bool
}

type  ConfigProvider interface {
	Parse(path string) (Configer,error)
}