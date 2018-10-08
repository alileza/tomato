package compare

import (
	"encoding/json"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

type jsondiff struct {
	deltas []gojsondiff.Delta
}

func (j *jsondiff) Deltas() []gojsondiff.Delta {
	return j.deltas
}

func (j *jsondiff) Modified() bool {
	return len(j.deltas) > 0
}

// JSON compares two json strings, processes them to handle wild cards,
func JSON(a []byte, b []byte) (string, error) {
	differ := gojsondiff.New()
	d, err := differ.Compare(a, b)
	if err != nil {
		return "", err
	}

	filteredDiffer := jsondiff{}
	for _, delta := range d.Deltas() {
		switch delta.(type) {
		case *gojsondiff.Modified:
			d := delta.(*gojsondiff.Modified)
			if d.NewValue.(string) == "*" {
				continue
			}
			filteredDiffer.deltas = append(filteredDiffer.deltas, delta)
		default:
			filteredDiffer.deltas = append(filteredDiffer.deltas, delta)
		}
	}

	var diffString string
	if filteredDiffer.Modified() {
		var aJSON map[string]interface{}
		json.Unmarshal(a, &aJSON)

		formatter := formatter.NewAsciiFormatter(aJSON, formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       false,
		})
		diffString, _ = formatter.Format(d)
	}
	return diffString, nil
}
