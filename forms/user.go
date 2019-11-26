package forms

type CreateUserCommand struct {
	FirstName   string  `json:"firstname" binding:"required"`
	LastName    string  `json:"lastname" binding:"required"`
	Company     string  `json:"company" binding:"required"`
	Email       string  `json:"email"`
	Mobile      string  `json:"mobile"`
	Extend1     string  `json:"extend1"`
	Extend2     string  `json:"extend2"`
}


type UpdateUserImageCommand struct {
	ID          string   `json:"id" binding:"required"`
	Image       string   `json:"image" binding:"required"`
	PersonID      string   `json:"personid""`
	FaceRegistered  bool     `json:"face_registerd"`
}

type VerifyImageCommand struct {
  Threshold   float64   `json:"threshold" binding:"required"`
	Max         int       `json:"max" binding:"required"`
	Image       string   `json:"image" binding:"required"`
}


type DeleteUserCommand struct {
	ID   string  `json:"id" binding:"required"`
}

type UpdateUserCommand struct {
	Name   string  `json:"name" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}
