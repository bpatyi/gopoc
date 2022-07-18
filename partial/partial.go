package partial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type Partial interface {
	Fields() []string
}

type partial struct {
	originalType  reflect.Type
	originalValue *reflect.Value
	fields        map[string]struct{}
}

func NewPartial(source interface{}) (Partial, error) {
	rt := reflect.TypeOf(source)
	if rt.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("cannot created partial model out of non-pointer %v", rt.String())
	}
	rtElem := rt.Elem()
	if rtElem.Kind() != reflect.Struct {
		return nil, fmt.Errorf("cannot created partial model out of non-struct %v", rtElem.String())
	}
	rv := reflect.ValueOf(source)
	return &partial{
		originalType:  rtElem,
		originalValue: &rv,
	}, nil
}

func (p *partial) UnmarshalJSON(data []byte) error {
	p.fields = make(map[string]struct{})
	var m map[string]*json.RawMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for key, value := range m {
		field, ok := p.originalType.FieldByNameFunc(func(name string) bool {
			return bytes.EqualFold([]byte(key), []byte(name))
		})
		if !ok {
			return fmt.Errorf("unknown field %s", key)
		}
		err = json.Unmarshal(*value, p.originalValue.Elem().FieldByIndex(field.Index).Addr().Interface())
		if err != nil {
			return fmt.Errorf("cannot unmarshal the %s field: %w", field.Name, err)
		}
		p.fields[field.Name] = struct{}{}
	}
	return nil
}

func (p *partial) Fields() []string {
	var fields []string
	for field := range p.fields {
		fields = append(fields, field)
	}
	return fields
}
