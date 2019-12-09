package forms

type CreateUserCommand struct {
	FirstName   string  `json:"firstname" binding:"required"`
	LastName    string  `json:"lastname" binding:"required"`
	Company     string  `json:"company" binding:"required"`
	Email       string  `json:"email"`
	Mobile      string  `json:"mobile"`
	Title       string  `json:"title"`
	Extend1     string  `json:"extend1"`
	Extend2     string  `json:"extend2"`
}

type ListUserCommand struct {
	Keyword  string    `json:"keyword"`
}

type UpdateUserImageCommand struct {
	ID              string   `json:"id" binding:"required"`
	Image           string   `json:"image" binding:"required"`
	PersonID        string   `json:"personid""`
	FaceRegistered  bool     `json:"face_registered"`
}

type VerifyImageCommand struct {
  Threshold   float64   `json:"threshold" binding:"required"`
	Max         int       `json:"max" binding:"required"`
	Image       string    `json:"image" binding:"required"`
}
type FindUserCommand struct {
	ID   string  `json:"id" binding:"required"`
}

type DeleteUserCommand struct {
	ID   string  `json:"id" binding:"required"`
}

type UpdateCheckCommand struct {
	ID        string  `json:"id" binding:"required"`
	BoothName string  `json:"boothName" binding:"required"`
	Checked   bool    `json:"checked"`
}


type UpdateUserCommand struct {
	ID          string  `json:"id" binding:"required"`
	Email       string  `json:"email"`
	Mobile      string  `json:"mobile"`
	Title       string   `json:"title"`
	Extend1     string  `json:"extend1"`
	Extend2     string  `json:"extend2"`
	Registered         *bool   `json:"registered" validate:"exists"`
	CounterRegistered  *bool   `json:"counterRegistered" validate:"exists"`
}
