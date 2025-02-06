package tiger

import (
	java "github.com/wnxd/microdbg-java"
)

type SharedPreference map[string]any

func (p SharedPreference) Contains(key string) java.JBoolean {
	_, ok := p[key]
	return ok
}

func (p SharedPreference) GetBoolean(key string, defValue java.JBoolean) java.JBoolean {
	if v, ok := p[key]; ok {
		return v.(java.JBoolean)
	}
	return defValue
}

func (p SharedPreference) GetFloat(key string, defValue java.JFloat) java.JFloat {
	if v, ok := p[key]; ok {
		return v.(java.JFloat)
	}
	return defValue
}

func (p SharedPreference) GetInt(key string, defValue java.JInt) java.JInt {
	if v, ok := p[key]; ok {
		return v.(java.JInt)
	}
	return defValue
}

func (p SharedPreference) GetLong(key string, defValue java.JLong) java.JLong {
	if v, ok := p[key]; ok {
		return v.(java.JLong)
	}
	return defValue
}

func (p SharedPreference) GetString(key string, defValue java.IString) java.IString {
	if v, ok := p[key]; ok {
		return v.(java.IString)
	}
	return defValue
}

func (p SharedPreference) Clear() {
	clear(p)
}

func (p SharedPreference) SetBoolean(key string, value java.JBoolean) {
	p[key] = value
}

func (p SharedPreference) SetFloat(key string, value java.JFloat) {
	p[key] = value
}

func (p SharedPreference) SetInt(key string, value java.JInt) {
	p[key] = value
}

func (p SharedPreference) SetLong(key string, value java.JLong) {
	p[key] = value
}

func (p SharedPreference) SetString(key string, value java.IString) {
	p[key] = value
}

func (p SharedPreference) Remove(key string) {
	delete(p, key)
}
