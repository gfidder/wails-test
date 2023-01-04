package oid

type Oid struct {
	Name        string `json:"name"`
	OID         string `json:"oid"`
	Mib         string `json:"mib"`
	Syntax      string `json:"syntax"`
	Access      string `json:"access"`
	Status      string `json:"status"`
	DefVal      string `json:"defVal"`
	Indexes     string `json:"indexes"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func CreateNewOid() *Oid {
	return &Oid{}
}
