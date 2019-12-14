package goliveconfig

import (
	"reflect"
	"testing"
)

func TestLiveConfig_Values(t *testing.T) {
	type fields struct {
		activeKeys     map[string]LiveValue
		concreteConfig interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		"test",
		fields{
			make(map[string]LiveValue),
			map[string]interface{}{"hello": 1, "world": []string{"yes", "no"}, "bool": true},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LiveConfig{
				activeKeys:     tt.fields.activeKeys,
				concreteConfig: tt.fields.concreteConfig,
			}
			lv := lc.LiveValue("$.hello")
			v := lv.Value()
			if !reflect.DeepEqual(v, 1) {
				t.Errorf("LiveConfig.Values() = %v, want %v", v, 1)
			}
		})
	}
}
