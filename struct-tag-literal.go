//
// Package thestruct provides helpers
// to parse struct tag literal and list of actual struct fields w/ types
//
package thestruct

//
// ParseStructTagLiteral is used to extract definitions from struct tag string
//
func ParseStructTagLiteral(s string) (*Tags, error) {

	o := Tags{}

	arr, err := explodeStructTag(s)
	if err != nil {
		return nil, err
	}

	if err := o.index(arr); err != nil {
		return nil, err
	}

	return &o, nil
}
