package pfsenseapi

import "strings"

// TrueIfPresent is designed to unmarshal PFSense boolean values that can indicate
// truth by having an empty string as the value of the property
type StringArray []string

// UnmarshalJSON implements the json.Unmarshaler interface.
func (sa *StringArray) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*sa = make(StringArray, 0)
		return nil
	}

	*sa = strings.Split(string(data), ",")
	return nil
}
