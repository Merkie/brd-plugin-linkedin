package brdlinkedin

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/merkie/brightdata-sdk-go/unblocker"
)

func FetchProfile(Unblocker *unblocker.UnblockerZone, LinkedinID string) (*Profile, error) {
	var cleanProfile *Profile

	for {
		// Fetch the HTML
		html, err := Unblocker.NewRequest("https://www.linkedin.com/in/" + LinkedinID + "/").Execute()
		if err != nil {
			return nil, err
		}

		// Check if the profile exists
		if !strings.Contains(html, `<script type="application/ld+json">`) {
			return nil, fmt.Errorf("Profile not found")
		}

		// Grab the internal json string
		jsonString := strings.Split(strings.Split(html, `<script type="application/ld+json">`)[1], `</script>`)[0]

		// Unmarshal the json `json:"knowledge"`string
		var profileJson linkedInProfileJSON
		err = json.Unmarshal([]byte(jsonString), &profileJson)
		if err != nil {
			return nil, err
		}

		// Create a new profile
		var profile Profile

		for _, graph := range profileJson.Graph {
			if graph.Type == "Person" {
				profile.Name = cleanString(graph.Name)
				profile.Description = cleanString(graph.Description)
				profile.ProfileImage = graph.Image.ContentURL
				profile.AddressLocality = cleanString(graph.Address.AddressLocality)
				profile.AddressCountry = cleanString(graph.Address.AddressCountry)
				profile.AlumniOf = []AlumniOf{}
				profile.JobTitles = []string{}
				profile.WorksFor = []WorksFor{}
				profile.Followers = 0
				profile.SharedArticles = []Article{}

				// NOTE: Start date and end date can be ints when it represents a year
				for _, alumniOf := range graph.AlumniOf {
					newAlumniOf := AlumniOf{
						Type:        alumniOf.Type,
						Name:        cleanString(alumniOf.Name),
						Url:         alumniOf.URL,
						Description: cleanString(alumniOf.Member.Description),
					}

					startDate, ok := alumniOf.Member.StartDate.(float64)
					if ok {
						newAlumniOf.StartDate = fmt.Sprintf("%d", int(startDate))
					} else {
						newAlumniOf.StartDate = fmt.Sprintf("%s", alumniOf.Member.StartDate)
					}

					endDate, ok := alumniOf.Member.EndDate.(float64)
					if ok {
						newAlumniOf.EndDate = fmt.Sprintf("%d", int(endDate))
					} else {
						newAlumniOf.EndDate = fmt.Sprintf("%s", alumniOf.Member.EndDate)
					}

					newAlumniOf.StartDate = cleanString(newAlumniOf.StartDate)
					newAlumniOf.EndDate = cleanString(newAlumniOf.EndDate)

					profile.AlumniOf = append(profile.AlumniOf, newAlumniOf)
				}

				profile.JobTitles = graph.JobTitle

				for _, worksFor := range graph.WorksFor {
					newWorksFor := WorksFor{
						Name:        cleanString(worksFor.Name),
						Url:         cleanString(worksFor.URL),
						Description: cleanString(worksFor.Member.Description),
					}

					startDate, ok := worksFor.Member.StartDate.(float64)
					if ok {
						newWorksFor.StartDate = fmt.Sprintf("%d", int(startDate))
					} else {
						newWorksFor.StartDate = fmt.Sprintf("%s", worksFor.Member.StartDate)
					}

					newWorksFor.StartDate = cleanString(newWorksFor.StartDate)

					profile.WorksFor = append(profile.WorksFor, newWorksFor)
				}

				if graph.InteractionStatistic.InteractionType == "https://schema.org/FollowAction" {
					profile.Followers = graph.InteractionStatistic.UserInteractionCount
				}
			} else if graph.Type == "Article" {
				profile.SharedArticles = append(profile.SharedArticles, Article{
					Name:        cleanString(graph.Name),
					ArticleUrl:  graph.URL,
					ArticleBody: cleanString(graph.ArticleBody),
					Author:      cleanString(graph.Author.Name),
					AuthorUrl:   graph.Author.URL,
				})
			}
		}

		censoredPattern := regexp.MustCompile(`^[ *-]+$`)
		isCensored := false

		// Check if the profile is censored
		for _, title := range profile.JobTitles {
			if censoredPattern.MatchString(title) {
				isCensored = true
				break
			}
		}

		if !isCensored {
			for _, worksFor := range profile.WorksFor {
				if censoredPattern.MatchString(worksFor.Name) {
					isCensored = true
					break
				}
			}
		}

		if !isCensored {
			for _, alumniOf := range profile.AlumniOf {
				if censoredPattern.MatchString(alumniOf.Name) {
					isCensored = true
					break
				}
			}
		}

		// If, after all checks, the profile is clean, return it
		if !isCensored {
			cleanProfile = &profile
			break
		}
	}

	return cleanProfile, nil
}
