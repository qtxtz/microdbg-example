package tiger

import (
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type SharedPreference struct {
	CookieID string
}

func (p *SharedPreference) Contains(key string) java.JBoolean {
	return false
}

func (p *SharedPreference) GetBoolean(key string, defValue java.JBoolean) java.JBoolean {
	return defValue
}

func (p *SharedPreference) GetFloat(key string, defValue java.JFloat) java.JFloat {
	return defValue
}

func (p *SharedPreference) GetInt(key string, defValue java.JInt) java.JInt {
	return defValue
}

func (p *SharedPreference) GetLong(key string, defValue java.JLong) java.JLong {
	return defValue
}

func (p *SharedPreference) GetString(key string, defValue java.IString) java.IString {
	switch key {
	case "TT_COOKIEID", "TT_COOKIEID_NEW":
		return gava.FakeString(p.CookieID)
	}
	return defValue
}
