package model

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/tiktok-business/util"
)

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var i uint64
	if bytes.Contains(b, []byte{'.'}) {
		f, _ := strconv.ParseFloat(string(b), 64)
		i = uint64(f)
	} else {
		i, _ = strconv.ParseUint(string(b), 10, 64)
	}
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Value() uint64 {
	return uint64(u64)
}

func (u64 Uint64) String() string {
	return strconv.FormatUint(uint64(u64), 10)
}

// Int64 support string quoted number in json
type Int64 int64

// UnmarshalJSON implement json Unmarshal interface
func (i64 *Int64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var i int64
	if bytes.Contains(b, []byte{'.'}) {
		f, _ := strconv.ParseFloat(string(b), 64)
		i = int64(f)
	} else {
		i, _ = strconv.ParseInt(string(b), 10, 64)
	}
	*i64 = Int64(i)
	return
}

func (i64 Int64) Value() int64 {
	return int64(i64)
}

func (i64 Int64) String() string {
	return strconv.FormatInt(int64(i64), 10)
}

// Int support string quoted number in json
type Int int

// UnmarshalJSON implement json Unmarshal interface
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var v int
	if bytes.Contains(b, []byte{'.'}) {
		f, _ := strconv.ParseFloat(string(b), 64)
		v = int(f)
	} else {
		v, _ = strconv.Atoi(string(b))
	}
	*i = Int(v)
	return
}

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// Float64 support string quoted number in json
type Float64 float64

// UnmarshalJSON implement json Unmarshal interface
func (f64 *Float64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseFloat(string(b), 64)
	*f64 = Float64(i)
	return
}

func (f64 Float64) Value() float64 {
	return float64(f64)
}

func (f64 Float64) String(prec int) string {
	return strconv.FormatFloat(float64(64), 'f', prec, 64)
}

// Bool support number/string in json
type Bool bool

// UnmarshalJSON implement json Unmarshal interface
func (bl *Bool) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var ret bool
	str := string(b)
	if str == "true" {
		ret = true
	} else if str == "false" {
		ret = false
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		return err
	} else if i == 0 {
		ret = false
	} else {
		ret = true
	}
	*bl = Bool(ret)
	return
}

func (b Bool) Value() bool {
	return bool(b)
}

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

const dateFormat = "2006-01-02 15:04:04"

type DateTime time.Time

func (d DateTime) IsZero() bool {
	return d.Time().IsZero()
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) String() string {
	return time.Time(d).UTC().Format(dateFormat)
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(util.StringsJoin(`"`, d.String(), `"`)), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	s := strings.TrimSpace(string(data))

	// Try to parse as an integer (timestamp)
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' { // Quick check if it could be a number
		if timestamp, err := strconv.ParseInt(s, 10, 64); err == nil {
			// Check for common timestamp lengths (seconds vs milliseconds)
			// If it's a very large number, it's likely milliseconds
			if timestamp > 1000000000000 { // Roughly > year 2001 in milliseconds
				*d = DateTime(time.UnixMilli(timestamp)) // Treat as milliseconds
			} else {
				*d = DateTime(time.Unix(timestamp, 0)) // Treat as seconds
			}
			return nil
		}
	}
	// Try to parse as a string (YYYY-MM-DD HH:MM:SS)
	// Remove quotes if present
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		s = strings.Trim(s, `"`)
	}

	parsed, err := time.ParseInLocation(dateFormat, s, time.UTC)
	if err != nil {
		return err
	}

	*d = DateTime(parsed)
	return nil
}

// TimeLocation is a custom type to handle JSON marshaling/unmarshaling of time.Location.
type TimeLocation time.Location

func (l *TimeLocation) Location() *time.Location {
	return (*time.Location)(l)
}

func (l *TimeLocation) String() string {
	return l.Location().String()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *TimeLocation) UnmarshalJSON(b []byte) error {
	// Remove the surrounding quotes from the JSON string.
	var tzName string
	if err := json.Unmarshal(b, &tzName); err != nil {
		return err
	}

	// Load the time zone based on the IANA name.
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return err
	}
	*l = TimeLocation(*loc)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (l TimeLocation) MarshalJSON() ([]byte, error) {
	loc := (*time.Location)(&l)
	if loc == nil {
		return []byte("null"), nil
	}
	// Marshal the location's name back to a JSON string.
	return json.Marshal(loc.String())
}
