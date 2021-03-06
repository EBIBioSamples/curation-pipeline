package interrogator_test

import (
	"fmt"
	"github.com/EBIBioSamples/certification-pipeline/internal/interrogator"
	"github.com/EBIBioSamples/certification-pipeline/internal/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	logger     = log.New(os.Stdout, "TestInterrogate ", log.LstdFlags|log.Lshortfile)
	checklists = []model.Checklist{
		{Name: "NCBI Candidate Checklist", File: "../../res/schemas/ncbi-candidate-schema.json"},
		{Name: "BioSamples Checklist", File: "../../res/schemas/biosamples-schema.json"},
	}
)

func TestInterrogate(t *testing.T) {
	tests := []struct {
		documentFile           string
		expectedCandidateNames []string
	}{
		{
			documentFile:           "../../res/json/ncbi-SAMN03894263.json",
			expectedCandidateNames: []string{"NCBI Candidate Checklist"},
		},
		{
			documentFile:           "../../res/json/SAMEA3774859.json",
			expectedCandidateNames: []string{},
		},
	}
	for _, test := range tests {
		in := make(chan model.Sample)
		document, err := ioutil.ReadFile(test.documentFile)
		if err != nil {
			log.Fatal(errors.Wrap(err, fmt.Sprintf("read failed for: %s", test.documentFile)))
		}

		sampleInterrogated := interrogator.NewInterrogator(logger, in, checklists)

		sample := model.Sample{Accession: "test-uuid", Document: string(document)}

		in <- sample
		ir := <-sampleInterrogated

		candidateNames := make([]string, 0)
		for _, checklist := range ir.Checklists {
			candidateNames = append(candidateNames, checklist.Name)
		}
		assert.Equal(t, test.expectedCandidateNames, candidateNames)
	}
}
