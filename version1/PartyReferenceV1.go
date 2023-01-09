package version1

type PartyReferenceV1 struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewPartyReferenceV1(id, name, email string) *PartyReferenceV1 {
	return &PartyReferenceV1{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
