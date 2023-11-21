package brdlinkedin

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	brightdatasdk "github.com/merkie/brightdata-sdk-go"
)

func FetchProfile(BrdClient *brightdatasdk.BrightDataClient, LinkedinID string) (*Profile, error) {
	// Fetch the HTML
	html, err := BrdClient.Unblocker("https://www.linkedin.com/in/" + LinkedinID + "/").Execute()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	description := strings.TrimSpace(doc.Find("h2.top-card-layout__headline").Text())
	location := strings.TrimSpace(doc.Find("div.top-card__subline-item").Text())

	var personBlob map[string]interface{}
	doc.Find("script[type='application/ld+json']").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := s.Text()
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(text), &data); err == nil {
			if graph, ok := data["@graph"].([]interface{}); ok {
				if len(graph) > 0 {
					if person, ok := graph[0].(map[string]interface{}); ok {
						if person["@type"] == "Person" {
							personBlob = person
							return false // Break the loop
						}
					}
				}
			}
		}
		return true
	})

	if personBlob == nil {
		return nil, nil
	}

	profile := parseProfile(personBlob, description, location)
	return &profile, nil
}

func parseProfile(blob map[string]interface{}, description, location string) Profile {
	profile := Profile{
		Name:        getString(blob["name"]),
		Description: description,
		Location:    location,
		JobTitles:   getString(blob["jobTitle"]),
		WorksFor:    parseWorks(blob["worksFor"]),
		Experience:  parseExperience(blob["alumniOf"], "Organization"),
		Education:   parseEducation(blob["alumniOf"], "EducationalOrganization"),
	}

	if profile.Location == "" {
		if address, ok := blob["address"].(map[string]interface{}); ok {
			profile.Location = getString(address["addressLocality"])
		}
	}

	return profile
}

func parseWorks(data interface{}) []Work {
	var works []Work
	if worksFor, ok := data.([]interface{}); ok {
		for _, item := range worksFor {
			if workMap, ok := item.(map[string]interface{}); ok {
				work := Work{
					Name:        getString(workMap["name"]),
					Location:    getString(workMap["location"]),
					StartDate:   getString(workMap["member"].(map[string]interface{})["startDate"]),
					Description: getString(workMap["member"].(map[string]interface{})["description"]),
				}
				works = append(works, work)
			}
		}
	}
	return works
}

func parseExperience(data interface{}, typeFilter string) []Experience {
	var experiences []Experience
	if alumni, ok := data.([]interface{}); ok {
		for _, item := range alumni {
			if alumniMap, ok := item.(map[string]interface{}); ok {
				if alumniMap["@type"] == typeFilter {
					experience := Experience{
						Name:        getString(alumniMap["name"]),
						Location:    getString(alumniMap["location"]),
						StartDate:   getString(alumniMap["member"].(map[string]interface{})["startDate"]),
						EndDate:     getString(alumniMap["member"].(map[string]interface{})["endDate"]),
						Description: getString(alumniMap["member"].(map[string]interface{})["description"]),
					}
					experiences = append(experiences, experience)
				}
			}
		}
	}
	return experiences
}

func parseEducation(data interface{}, typeFilter string) []Education {
	var educations []Education
	if alumni, ok := data.([]interface{}); ok {
		for _, item := range alumni {
			if alumniMap, ok := item.(map[string]interface{}); ok {
				if alumniMap["@type"] == typeFilter {
					education := Education{
						Name:        getString(alumniMap["name"]),
						Location:    getString(alumniMap["location"]),
						StartDate:   getString(alumniMap["member"].(map[string]interface{})["startDate"]),
						EndDate:     getString(alumniMap["member"].(map[string]interface{})["endDate"]),
						Description: getString(alumniMap["member"].(map[string]interface{})["description"]),
					}
					educations = append(educations, education)
				}
			}
		}
	}
	return educations
}

func getString(v interface{}) string {
	if str, ok := v.(string); ok {
		return strings.TrimSpace(str)
	}
	return ""
}
