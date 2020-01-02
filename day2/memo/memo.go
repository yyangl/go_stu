package memo

type Memo struct {
	f Func
	cache map[string]result
}

type result struct {
	value interface{}
	err error
}
type Func func(key string) (interface{}, error)

func New(f Func) *Memo{
	return &Memo{f:f,cache:make(map[string]result)}
}

func (m *Memo) Get(key string)(interface{},error)  {
	res , ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}