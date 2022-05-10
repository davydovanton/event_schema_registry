package schemaregistry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoutingKey(t *testing.T) {
	var assert = assert.New(t)

	v := NewValidator("./test")

	data := []byte(`{
		"event_id": "a",
		"event_version": 1,
		"event_name": "AccountCreated",
		"event_time": "1",
		"producer": "pp",
		"data": {
			"public_id": "3ab651f8-a995-4293-ae3d-568a8b759916",
			"email": "a@a.a"
		}
	}`)

	assert.Equal(nil, v.Validate(data, "accounts.created", 1), "validate accounts schema")

	assert.NotNil(v.Validate(data, "accounts.create", 1), "bad schema name")
	assert.NotNil(v.Validate(data, "t.created", 1), "bad schema name")
}
