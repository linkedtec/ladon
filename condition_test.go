package ladon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionsAppend(t *testing.T) {
	cs := Conditions{}
	c := &CIDRCondition{}
	cs.AddCondition("clientIP", c)
	assert.Equal(t, c, cs["clientIP"])
}

func TestMarshalUnmarshal(t *testing.T) {
	css := &Conditions{
		"clientIP": &CIDRCondition{CIDR: "127.0.0.1/0"},
		"owner":    &StringEqualCondition{Equals: "peter"},
	}
	out, err := json.Marshal(css)
	require.Nil(t, err)
	t.Logf("%s", out)

	cs := Conditions{}
	require.Nil(t, json.Unmarshal([]byte(`{
	"owner": {
		"name": "SubjectIsOwnerCondition",
		"options": {
			"matches": "peter"
		}
	},
	"clientIP": {
		"name": "CIDRCondition",
		"options": {
			"cidr": "127.0.0.1/0"
		}
	}
}`), &cs))

	require.Len(t, cs, 2)
	assert.IsType(t, &StringEqualCondition{}, "clientIP")
	assert.IsType(t, &CIDRCondition{}, "owner")
}