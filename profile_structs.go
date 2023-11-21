package brdlinkedin

type Profile struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	JobTitles   string       `json:"jobTitles"`
	WorksFor    []Work       `json:"worksFor"`
	Experience  []Experience `json:"experience"`
	Education   []Education  `json:"education"`
}

type Work struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	Description string `json:"description"`
}

type Experience struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
}

type Education struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
}
