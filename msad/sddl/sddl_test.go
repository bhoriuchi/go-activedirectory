package sddl

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"gopkg.in/ldap.v3"
)

func TestSDDL(t *testing.T) {
	user := os.Getenv("LDAP_USER")
	pass := os.Getenv("LDAP_PASS")
	host := os.Getenv("LDAP_HOST")

	conn, err := ldap.Dial("tcp", host)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	err = conn.Bind(user, pass)
	if err != nil {
		log.Fatalln(err)
	}

	userSearch := ldap.NewSearchRequest(
		"DC=cjp,DC=blackline,DC=corp",
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		"(&(objectClass=user)(sAMAccountName=USC1SQL3VC1$))",
		[]string{
			"nTSecurityDescriptor",
		},
		[]ldap.Control{
			ldap.NewControlString(
				"1.2.840.113556.1.4.801",
				true,
				fmt.Sprintf("%c%c%c%c%c", 48, 3, 2, 1, 7),
			),
		},
	)

	results, err := conn.Search(userSearch)
	if err != nil {
		log.Fatalln(err)
	}
	descriptor := results.Entries[0].Attributes[0].ByteValues[0]
	sddl, err := NewSDDL(descriptor)
	if err != nil {
		log.Fatalln(err)
	}

	j, _ := json.MarshalIndent(sddl, "", "  ")
	fmt.Printf("%s\n", j)
}
