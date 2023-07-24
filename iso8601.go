package clockogo

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const iso8601UTCFormat = "2006-01-02T15:04:05Z"

// mixed pointer/value receivers required for query.Encoder/json.Marshaler interfaces

type ISO8601UTC time.Time

func (t ISO8601UTC) String() string {
	return time.Time(t).Format(iso8601UTCFormat)
}

func (t ISO8601UTC) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *ISO8601UTC) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(iso8601UTCFormat, strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	*t = ISO8601UTC(tt)
	return nil
}

func (t ISO8601UTC) EncodeValues(key string, v *url.Values) error {
	v.Add(key, t.String())
	return nil
}

func (t ISO8601UTC) Time() time.Time {
	return time.Time(t)
}
