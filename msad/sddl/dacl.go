package sddl

// DACL a dacl type
type DACL struct {
	ACL
}

// NewDACL creates a new DACL
func NewDACL(descriptor []byte) (dacl *DACL, err error) {
	dacl = &DACL{}
	dacl.Parse(descriptor)
	return
}
