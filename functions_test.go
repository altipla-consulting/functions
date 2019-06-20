package functions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeFirestoreEvent(t *testing.T) {
	input := `
		{
			"value": {
				"fields": {
					"groups": {
						"stringValue": "foo"
					}
				}
			}
		}`
	event := new(FirestoreEvent)
	require.NoError(t, json.Unmarshal([]byte(input), &event))

	require.Equal(t, event.Value.Fields["groups"].GetStringValue(), "foo")
}
