package models

import (
	"fmt"
)

type company struct {
	// Add fields as needed
	Name        string `json:"name"`
	Ranking		int `json:"ranking"`
	// This is split into 3 basically, europe, america and the rest of the world
	Region		string `json:"region"`
}

type companies struct {
	companies []company
}

func DefaultCompaniesList() []company {
	return []company{
		{Name: "blackrock", Ranking: 1, Region: "usa"},
		{Name: "amundi", Ranking: 2, Region: "europe"},
		{Name: "ubs", Ranking: 3, Region: "europe"},
	}
}

func (c *company) String() string {
	return fmt.Sprintf("Company Name: %s, Ranking: %d, Region: %s", c.Name, c.Ranking, c.Region)
}

// Since this is a static list we dont need to implement setters
func (c *company) GetName() string {
	return c.Name
}

func (c *company) GetRanking() int {
	return c.Ranking
}

func (c *company) GetRegion() string {
	return c.Region
}
func (c *companies) DefineDefaultCompaniesList() {
	c.companies = []company{
		{Name: "blackrock", Ranking: 1, Region: "usa"},
		{Name: "amundi", Ranking: 2, Region: "europe"},
		{Name: "ubs", Ranking: 3, Region: "europe"},
	}
}
