package compare

import "testing"

func TestJSON(t *testing.T) {
	for _, test := range []struct {
		description string
		a           []byte
		b           []byte
		expected    string
	}{
		{
			description: "DiffName",
			a:           []byte(` { "name":"alice", "occupation": "foo", "role": "bar" } `),
			b:           []byte(` { "name":"bob", "occupation": "foo", "role": "bar" } `),
			expected: `{
					-  "name": "alice",
					+  "name": "bob",
					   "occupation": "foo",
					   "role": "bar"
				  }`,
		},
	} {
		output, err := JSON(test.a, test.b)
		if err != nil {
			t.Errorf("%s - expecting no error, got %s", test.description, err)
		} else if output != test.expected {
			t.Errorf("%s - expecting %s, got %s", test.description, test.expected, output)
		}
	}
}

var cJSON = `
{ 
	"name":"*", 
	"occupation": "foo", 
	"role": "bar" 
}
`
