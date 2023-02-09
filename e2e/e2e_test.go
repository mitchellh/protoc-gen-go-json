package e2e

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {
	type basicWrapper struct{ Basic }
	var optionalPresent = "present"
	var optionalEmpty = ""
	var cases = []struct {
		Name string

		// Value and Expected MUST be pointers to structs. If Expected is
		// nil, then it is expected to be identical to Value.
		Value    interface{}
		Expected interface{}
	}{
		{
			"basic",
			&Basic{
				A: "hello",
				B: &Basic_Int{
					Int: 42,
				},
			},
			nil,
		},

		{
			"basic wrapped in Go struct",
			&basicWrapper{
				Basic: Basic{
					A: "hello",
					B: &Basic_Int{
						Int: 42,
					},
				},
			},
			nil,
		},

		{
			"nested",
			&Nested_Message{
				Basic: &Basic{
					A: "hello",
					B: &Basic_Int{
						Int: 42,
					},
				},
			},
			nil,
		},

		{
			"optional present",
			&Basic{
				A: "hello",
				B: &Basic_Int{
					Int: 42,
				},
				O: &optionalPresent,
			},
			nil,
		},

		{
			"optional empty",
			&Basic{
				A: "hello",
				B: &Basic_Int{
					Int: 42,
				},
				O: &optionalEmpty,
			},
			nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			require := require.New(t)

			// Verify marshaling doesn't error
			bs, err := json.Marshal(tt.Value)
			require.NoError(err)
			require.NotEmpty(bs)

			// Determine what we expect the result to be
			expected := tt.Expected
			if expected == nil {
				expected = tt.Value
			}

			// Unmarshal. We want to do this into a concrete type so we
			// use reflection here (you can't just decode into interface{})
			// and have that work.
			val := reflect.New(reflect.ValueOf(expected).Elem().Type())
			require.NoError(json.Unmarshal(bs, val.Interface()))
			require.Equal(val.Interface(), expected)
		})
	}
}
