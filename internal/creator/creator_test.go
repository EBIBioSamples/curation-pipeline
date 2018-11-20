package creator_test

import (
	"fmt"
	"github.com/EBIBioSamples/certification-pipeline/internal/creator"
	"github.com/EBIBioSamples/certification-pipeline/internal/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	sampleCreated = make(chan model.Sample)
	jsonSubmitted = make(chan string)
)

func TestCreateSample(t *testing.T) {
	tests := []struct {
		documentFile string
	}{
		{
			documentFile: "../../res/json/ncbi-SAMN03894263.json",
		},
	}
	for _, test := range tests {
		document, err := ioutil.ReadFile(test.documentFile)
		if err != nil {
			log.Fatal(errors.Wrap(err, fmt.Sprintf("read failed for: %s", test.documentFile)))
		}

		creator.NewCreator(
			log.New(os.Stdout, "TestCreateSample ", log.LstdFlags|log.Lshortfile),
			jsonSubmitted,
			sampleCreated)

		jsonSubmitted <- string(document)
		sample := <-sampleCreated

		assert.Regexp(t, `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`, sample.UUID)
	}
}
