package model

import (
	"strconv"
	"strings"
)

// CSVStringList is a custom type that encodes []string as a single comma-separated string.
type CSVStringList []string

func (s CSVStringList) Strings() []string {
	return []string(s)
}

// MarshalText implements encoding.TextMarshaler.
func (s CSVStringList) MarshalText() ([]byte, error) {
	return []byte(strings.Join(s, ",")), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CSVStringList) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*s = nil
		return nil
	}
	*s = strings.Split(string(text), ",")
	return nil
}

// CSVFloatList is a custom type that encodes []string as a single comma-separated string.
type CSVFloatList []float64

func (f CSVFloatList) Float64s() []float64 {
	return []float64(f)
}

// MarshalText implements encoding.TextMarshaler.
func (f CSVFloatList) MarshalText() ([]byte, error) {
	strs := make([]string, len(f))
	for i, v := range f {
		strs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return []byte(strings.Join(strs, ",")), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *CSVFloatList) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))
	if s == "" {
		*f = nil
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]float64, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			return err
		}
		result = append(result, v)
	}
	*f = result
	return nil
}
