package sddl

import (
	"encoding/hex"

	"github.com/bhoriuchi/go-activedirectory/msad/common"
)

// SDDL an sddl
type SDDL struct {
	Revision    int64  `json:"revision"`
	OffsetOwner int    `json:"offsetOwner"`
	OffsetGroup int    `json:"offsetGroup"`
	OffsetSacl  int    `json:"offsetSacl"`
	OffsetDacl  int    `json:"offsetDacl"`
	Owner       string `json:"owner"`
	Group       string `json:"group"`
	DACL        *DACL  `json:"dacl"`
}

// NewSDDL creates a new sddl
func NewSDDL(descriptor []byte) (sddl *SDDL, err error) {
	sddl = &SDDL{}
	err = sddl.Parse(descriptor)
	return
}

// Parse parses the descriptor
func (c *SDDL) Parse(descriptor []byte) (err error) {
	var daclb []byte

	h := hex.EncodeToString(descriptor)
	if c.Revision, err = common.Hexdec(common.Substr(h, 0, 2)); err != nil {
		return
	}

	if c.OffsetOwner, err = common.HexULong32Le2int(common.Substr(h, 8, 8)); err != nil {
		return
	} else if c.OffsetGroup, err = common.HexULong32Le2int(common.Substr(h, 16, 8)); err != nil {
		return
	} else if c.OffsetSacl, err = common.HexULong32Le2int(common.Substr(h, 24, 8)); err != nil {
		return
	} else if c.OffsetDacl, err = common.HexULong32Le2int(common.Substr(h, 32, 8)); err != nil {
		return
	}

	c.OffsetOwner = c.OffsetOwner * 2
	c.OffsetGroup = c.OffsetGroup * 2
	c.OffsetSacl = c.OffsetSacl * 2
	c.OffsetDacl = c.OffsetDacl * 2

	if c.OffsetOwner != 0 {
		if c.Owner, err = common.Hex2sid(h[c.OffsetOwner:]); err != nil {
			return
		}
	}
	if c.OffsetGroup != 0 {
		if c.Group, err = common.Hex2sid(h[c.OffsetGroup:]); err != nil {
			return
		}
	}

	if daclb, err = hex.DecodeString(h[c.OffsetDacl:]); err != nil {
		return
	} else if c.DACL, err = NewDACL(daclb); err != nil {
		return
	}

	return
}
