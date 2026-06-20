package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/isul/tossinvest-cli/internal/output"
)

func TestWrite_PrettyJSON(t *testing.T) {
	var buf bytes.Buffer
	opts := output.DefaultOptions()
	opts.Format = output.FormatPretty
	opts.Writer = &buf

	payload := []byte(`{"symbol":"005930","price":"70000"}`)
	if err := output.Write(payload, opts); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "005930") {
		t.Fatalf("output: %s", buf.String())
	}
}

func TestWrite_Transform(t *testing.T) {
	var buf bytes.Buffer
	opts := output.DefaultOptions()
	opts.Format = output.FormatJSON
	opts.Transform = "price"
	opts.Writer = &buf

	payload := []byte(`{"symbol":"005930","price":"70000"}`)
	if err := output.Write(payload, opts); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "70000") {
		t.Fatalf("output: %s", buf.String())
	}
}

func TestParseFormat(t *testing.T) {
	f, err := output.ParseFormat("json")
	if err != nil || f != output.FormatJSON {
		t.Fatalf("got %v %v", f, err)
	}
}
