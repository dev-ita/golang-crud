package entitie

type Usuario struct {
	Id    uint32 `json:"id" example:"1"`
	Nome  string `json:"nome" example:"teste"`
	Email string `json:"email" example:"teste@email.com"`
}
