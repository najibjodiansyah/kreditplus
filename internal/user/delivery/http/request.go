package http

type CreateUser struct {
	Nik          string `json:"nik" validate:"required"`
	FullName     string `json:"full_name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	LegalName    string `json:"legal_name"`
	BirthPlace   string `json:"birth_place"`
	BirthDate    string `json:"birth_date"`
	Wages        int    `json:"wages" validate:"required"`
	Photo_ktp    string `json:"photo_ktp"`
	Photo_selfie string `json:"photo_selfie"`
}

type UpdateUser struct {
	FullName     string
	LegalName    string
	BirthPlace   string
	BirthDate    string
	Wages        int
	Photo_ktp    string
	Photo_selfie string
}

type LoginUser struct {
	Nik      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string
}
