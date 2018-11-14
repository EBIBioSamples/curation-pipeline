package model

import "fmt"

//ValidationResult contains the results of a validation
type ValidationResult struct {
	Valid   bool     `json:"valid"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

//Checklist contains the name and file of a checklist
type Checklist struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	File    string `json:"file"`
}

func (c *Checklist) ID() string {
	return fmt.Sprintf("%s-%s", c.Name, c.Version)
}

//Sample tracks samples JSON in the pipeline
type Sample struct {
	UUID     string
	Document string
}

//InterrogationResult contains the checklists sample is a candidate for
type InterrogationResult struct {
	Sample              Sample
	CandidateChecklists []Checklist
}

//Curation is a transformation of a sample document content
type Curation struct {
	Characteristic string `json:"characteristic"`
	Value          string `json:"value"`
}

//PlanResult is the result of executing a curation plan
type PlanResult struct {
	Plan   Plan
	Sample Sample
}

//Certificate is certification given to a Sample against a Checklist
type Certificate struct {
	Sample        Sample
	SampleHash    string
	Checklist     Checklist
	ChecklistHash string
}

func (c *Certificate) Badge() string {
	return fmt.Sprintf("https://img.shields.io/badge/%s_%s-%s-%s.svg", c.Checklist.Name, c.Checklist.Version, "valid", "green")
}
