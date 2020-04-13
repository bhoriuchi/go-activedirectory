package sddl

import (
	"encoding/hex"
	"strings"

	"github.com/bhoriuchi/go-activedirectory/msad/common"
)

// ACETypes types of ace
var ACETypes = map[string]int64{
	"ACCESS_ALLOWED":                 0,
	"ACCESS_DENIED":                  1,
	"SYSTEM_AUDIT":                   2,
	"SYSTEM_ALARM":                   3,
	"ACCESS_ALLOWED_COMPOUND":        4,
	"ACCESS_ALLOWED_OBJECT":          5,
	"ACCESS_DENIED_OBJECT":           6,
	"SYSTEM_AUDIT_OBJECT":            7,
	"SYSTEM_ALARM_OBJECT":            8,
	"ACCESS_ALLOWED_CALLBACK":        9,
	"ACCESS_DENIED_CALLBACK":         10,
	"ACCESS_ALLOWED_CALLBACK_OBJECT": 11,
	"ACCESS_DENIED_CALLBACK_OBJECT":  12,
	"SYSTEM_AUDIT_CALLBACK":          13,
	"SYSTEM_ALARM_CALLBACK":          14,
	"SYSTEM_AUDIT_CALLBACK_OBJECT":   15,
	"SYSTEM_ALARM_CALLBACK_OBJECT":   16,
	"SYSTEM_MANDATORY_LABEL":         17,
	"SYSTEM_RESOURCE_ATTRIBUTE":      18,
	"SYSTEM_SCOPED_POLICY_ID":        19,
}

// ACEShortNames short names
var ACEShortNames = map[string]int64{
	"A":  0,
	"D":  1,
	"AU": 2,
	"AL": 3,
	"CA": 4,
	"OA": 5,
	"OD": 6,
	"OU": 7,
	"OL": 8,
	"XA": 9,
	"XD": 10,
	"ZA": 11,
	"ZD": 12,
	"XU": 13,
	"XL": 14,
	"ZU": 15,
	"ZL": 16,
	"ML": 17,
	"RA": 18,
	"SP": 19,
}

// ACE an ace
type ACE struct {
	RawType       int64  `rawType:"json"`
	Type          string `type:"json"`
	TypeShortName string `typeShortName:"json"`
}

// NewACE creates a new ACE
func NewACE(descriptor []byte) (ace *ACE, err error) {
	ace = &ACE{}
	err = ace.Parse(descriptor)
	return
}

// Parse parses the ace
func (c *ACE) Parse(descriptor []byte) (err error) {
	h := hex.EncodeToString(descriptor)

	if c.RawType, err = common.Hexdec(common.Substr(h, 0, 2)); err != nil {
		return
	}
	c.Type = GetACEType(c.RawType)
	c.TypeShortName = GetACEShortName(c.RawType)

	if strings.HasSuffix(c.Type, "_OBJECT") {

	}

	return
}

// Validate validates an ACE
func (c *ACE) Validate() (err error) {
	return
}

// GetACEType ace type
func GetACEType(i int64) (typeString string) {
	for name, id := range ACETypes {
		if id == i {
			typeString = name
			break
		}
	}
	return
}

// GetACEShortName ace type short name
func GetACEShortName(i int64) (typeString string) {
	for name, id := range ACEShortNames {
		if id == i {
			typeString = name
			break
		}
	}
	return
}
