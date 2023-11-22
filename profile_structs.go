package brdlinkedin

type AlumniOf struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type WorksFor struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
}

type Article struct {
	Name        string `json:"name"`
	ArticleBody string `json:"articleBody"`
	ArticleUrl  string `json:"articleUrl"`
	Author      string `json:"author"`
	AuthorUrl   string `json:"authorUrl"`
}

type Profile struct {
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	ProfileImage    string     `json:"profileImage"`
	AddressLocality string     `json:"addressLocality"`
	AddressCountry  string     `json:"addressCountry"`
	AlumniOf        []AlumniOf `json:"alumniOf"`
	JobTitles       []string   `json:"jobTitle"`
	WorksFor        []WorksFor `json:"worksFor"`
	Followers       int        `json:"followers"`
	SharedArticles  []Article  `json:"sharedArticles"`
}

type linkedInProfileJSON struct {
	Context string `json:"@context"`
	Graph   []struct {
		Type   string `json:"@type"`
		Name   string `json:"name,omitempty"`
		Author struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"author,omitempty"`
		ArticleBody string `json:"articleBody,omitempty"`
		URL         string `json:"url"`
		Address     struct {
			Type            string `json:"@type"`
			AddressLocality string `json:"addressLocality"`
			AddressCountry  string `json:"addressCountry"`
		} `json:"address,omitempty"`
		AlumniOf []struct {
			Type   string `json:"@type"`
			Name   string `json:"name"`
			URL    string `json:"url"`
			Member struct {
				Type        string      `json:"@type"`
				Description string      `json:"description"`
				StartDate   interface{} `json:"startDate"`
				EndDate     interface{} `json:"endDate"`
			} `json:"member"`
		} `json:"alumniOf,omitempty"`
		Awards []any `json:"awards,omitempty"`
		Image  struct {
			Type       string `json:"@type"`
			ContentURL string `json:"contentUrl"`
		} `json:"image,omitempty"`
		JobTitle []string `json:"jobTitle,omitempty"`
		SameAs   string   `json:"sameAs,omitempty"`
		MemberOf []any    `json:"memberOf,omitempty"`
		WorksFor []struct {
			Type   string `json:"@type"`
			Name   string `json:"name"`
			URL    string `json:"url,omitempty"`
			Member struct {
				Type        string      `json:"@type"`
				StartDate   interface{} `json:"startDate"`
				Description string      `json:"description"`
			} `json:"member"`
		} `json:"worksFor,omitempty"`
		KnowsLanguage             []any  `json:"knowsLanguage,omitempty"`
		DisambiguatingDescription string `json:"disambiguatingDescription,omitempty"`
		InteractionStatistic      struct {
			Type                 string `json:"@type"`
			InteractionType      string `json:"interactionType"`
			Name                 string `json:"name"`
			UserInteractionCount int    `json:"userInteractionCount"`
		} `json:"interactionStatistic,omitempty"`
		Description string `json:"description,omitempty"`
		ReviewedBy  struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"reviewedBy,omitempty"`
	} `json:"@graph"`
}
