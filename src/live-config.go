package goliveconfig

import "github.com/yalp/jsonpath"

type LiveConfig struct {
	activeKeys     map[string]LiveValue
	concreteConfig interface{}
}

func (lc *LiveConfig) Values() interface{} {
	return lc.concreteConfig
}

func (lc *LiveConfig) Update() {

}

func (lc *LiveConfig) LiveValue(key string) LiveValue {
	v, ok := lc.activeKeys[key]
	if ok {
		return v
	}

	val, err := jsonpath.Read(lc.Values(), key)
	if err != nil {
		lv := LiveValue{key, val, true}
		lc.activeKeys[key] = lv
		return lv
	}
	lv := LiveValue{key, nil, false}
	lc.activeKeys[key] = lv
	return lv
}

type LiveValue struct {
	key   string
	value interface{}
	isSet bool
}

func (lv *LiveValue) Value() (interface{}, bool) {
	return lv.value, true
}

func (lv *LiveValue) ValueAsString() (string, bool) {
	v, ok := lv.value.(string)
	return v, ok
}

func (lv *LiveValue) ValueAsInt() (int, bool) {
	v, ok := lv.value.(int)
	return v, ok
}
