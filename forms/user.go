package forms

type CreateUserCommand struct {
	FirstName   string  `json:"firstname" binding:"required"`
	LastName    string  `json:"lastname" binding:"required"`
	Company     string  `json:"company" binding:"optional"`
	Email       string  `json:"email" binding:"optional"`
	Mobile      string  `json:"mobile" binding:"optional"`
}

type UpdateUserCommand struct {
	Name   string  `json:"name" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}
