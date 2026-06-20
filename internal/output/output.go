package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/tidwall/gjson"
)

type Format string

const (
	FormatAuto   Format = "auto"
	FormatJSON   Format = "json"
	FormatJSONL  Format = "jsonl"
	FormatYAML   Format = "yaml"
	FormatPretty Format = "pretty"
	FormatRaw    Format = "raw"
)

type Options struct {
	Format      Format
	FormatError Format
	Transform   string
	IsError     bool
	Writer      io.Writer
	ErrorWriter io.Writer
}

func DefaultOptions() Options {
	return Options{
		Format:      FormatAuto,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
}

func (o Options) writer() io.Writer {
	if o.IsError && o.ErrorWriter != nil {
		return o.ErrorWriter
	}
	if o.Writer != nil {
		return o.Writer
	}
	return os.Stdout
}

func Write(data []byte, opts Options) error {
	format := opts.Format
	if format == FormatAuto {
		format = FormatPretty
	}

	body := data
	if opts.Transform != "" {
		result := gjson.GetBytes(data, opts.Transform)
		if !result.Exists() {
			return fmt.Errorf("transform path not found: %s", opts.Transform)
		}
		body = []byte(result.Raw)
	}

	w := opts.writer()
	switch format {
	case FormatRaw:
		_, err := w.Write(body)
		return err
	case FormatJSON:
		return writeJSON(w, body, false)
	case FormatJSONL:
		return writeJSON(w, body, true)
	case FormatYAML:
		return writeYAML(w, body)
	case FormatPretty:
		return writePretty(w, body)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func WriteError(err error, opts Options) error {
	opts.IsError = true
	if opts.FormatError != "" {
		opts.Format = opts.FormatError
	} else if opts.Format == FormatAuto {
		opts.Format = FormatPretty
	}
	payload, marshalErr := json.Marshal(map[string]any{
		"error": err.Error(),
	})
	if marshalErr != nil {
		return err
	}
	return Write(payload, opts)
}

type OptionsWithErrorFormat struct {
	Options
	FormatError Format
}

func (o Options) WithFormatError(f Format) OptionsWithErrorFormat {
	return OptionsWithErrorFormat{Options: o, FormatError: f}
}

func writeJSON(w io.Writer, body []byte, compact bool) error {
	if !json.Valid(body) {
		_, err := w.Write(body)
		return err
	}
	if compact {
		_, err := w.Write(body)
		return err
	}
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		_, err2 := w.Write(body)
		return err2
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func writeYAML(w io.Writer, body []byte) error {
	if !json.Valid(body) {
		_, err := w.Write(body)
		return err
	}
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}
	out, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	return err
}

func writePretty(w io.Writer, body []byte) error {
	if !json.Valid(body) {
		_, err := w.Write(body)
		return err
	}
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		_, err2 := w.Write(body)
		return err2
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func ParseFormat(s string) (Format, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "auto", "":
		return FormatAuto, nil
	case "json":
		return FormatJSON, nil
	case "jsonl":
		return FormatJSONL, nil
	case "yaml":
		return FormatYAML, nil
	case "pretty":
		return FormatPretty, nil
	case "raw":
		return FormatRaw, nil
	default:
		return "", fmt.Errorf("unknown format %q", s)
	}
}
