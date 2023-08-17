package apache_beam

import (
	"flag"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
)

var (
	input_text = flag.String("input-text", "sample input text", "Input text to print.")
)

func init() {
	// DoFns should be registered with Beam to be available in distributed runners.
	register.Function1x1(strings.Title)
	register.Emitter1[string]()
}

func myPipeline(scope beam.Scope, input_text string) beam.PCollection {
	elements := beam.Create(scope, "hello", "world!", input_text)
	return elements
}

func Hello() {
	flag.Parse()
	beam.Init()

	_, scope := beam.NewPipelineWithRoot()
	myPipeline(scope, *input_text)

}
