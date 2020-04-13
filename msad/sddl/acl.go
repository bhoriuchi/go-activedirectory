package sddl

import (
	"encoding/hex"

	"github.com/bhoriuchi/go-activedirectory/msad/common"
)

// ACL acl
type ACL struct {
	Revision int64  `json:"revision"`
	SBZ1     int64  `json:"sbz1"`
	SBZ2     int    `json:"sbz2"`
	ACECount int    `json:"aceCount"`
	ACEs     []*ACE `json:"aces"`
}

// NewACL creates a new ACL
func NewACL(descriptor []byte) (acl *ACL, err error) {
	acl = &ACL{}
	err = acl.Parse(descriptor)
	return
}

// AddACE adds an ACE
func (c *ACL) AddACE(aces ...*ACE) (err error) {
	for _, ace := range aces {
		if err = ace.Validate(); err != nil {
			return
		}
	}
	return
}

// RemoveACE removes an ace
func (c *ACL) RemoveACE(ace ...*ACE) (err error) {
	return
}

// Parse parses acl
func (c *ACL) Parse(descriptor []byte) (err error) {
	h := hex.EncodeToString(descriptor)
	c.ACEs = []*ACE{}

	if c.Revision, err = common.Hexdec(common.Substr(h, 0, 2)); err != nil {
		return
	} else if c.SBZ1, err = common.Hexdec(common.Substr(h, 2, 2)); err != nil {
		return
	} else if c.SBZ2, err = common.HexUShort16Le2int(common.Substr(h, 12, 4)); err != nil {
		return
	} else if c.ACECount, err = common.HexUShort16Le2int(common.Substr(h, 8, 4)); err != nil {
		return
	}

	pos := 16
	for i := 0; i < c.ACECount; i++ {
		var aceLength int
		var ace *ACE
		var aceb []byte

		aceLength, err = common.HexUShort16Le2int(common.Substr(h, pos+4, 4))
		if err != nil {
			return
		}

		aceLength = aceLength * 2
		if aceb, err = hex.DecodeString(common.Substr(h, pos, aceLength)); err != nil {
			return
		} else if ace, err = NewACE(aceb); err != nil {
			return
		} else if err = c.AddACE(ace); err != nil {
			return
		}

		pos += aceLength
	}

	return nil
}
