package http

type requestUser struct {
	Nik        string
	FullName   string
	Password   string
	LegalName  string
	BirthPlace string
	BirthDate  string
	Wages      int
	Photo      requestPhoto
}

type requestPhoto struct {
	selfie string
	ktp    string
}
