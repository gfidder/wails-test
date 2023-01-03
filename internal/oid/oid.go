package oid

type Oid struct {
	Name   string
	OID    string
	Mib    string
	Syntax string
	Access string
	Status string
}

func CreateNewOid() *Oid {

	return &Oid{}
}
