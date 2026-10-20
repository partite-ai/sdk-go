// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersAddResourcesRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &AddResourcesRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitAddResourcesRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AddResourcesRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCancelRequestRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CancelRequestRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &CancelRequestRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCancelRequestRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelRequestRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelRequestRequest{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCancelTaskRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CancelTaskRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &CancelTaskRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTaskID", func(t *testing.T) {
		obj := &CancelTaskRequest{}
		var fernTestValueTaskID string
		obj.SetTaskID(fernTestValueTaskID)
		assert.Equal(t, fernTestValueTaskID, obj.TaskID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCancelTaskRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelTaskRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelTaskRequest{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTaskID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelTaskRequest{}
		var fernTestValueTaskID string

		// Act
		obj.SetTaskID(fernTestValueTaskID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCompleteTaskRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CompleteTaskRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &CompleteTaskRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTaskID", func(t *testing.T) {
		obj := &CompleteTaskRequest{}
		var fernTestValueTaskID string
		obj.SetTaskID(fernTestValueTaskID)
		assert.Equal(t, fernTestValueTaskID, obj.TaskID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCompleteTaskRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CompleteTaskRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CompleteTaskRequest{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTaskID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CompleteTaskRequest{}
		var fernTestValueTaskID string

		// Act
		obj.SetTaskID(fernTestValueTaskID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateConversationRequest(t *testing.T) {
	t.Run("SetWait", func(t *testing.T) {
		obj := &CreateConversationRequest{}
		var fernTestValueWait *int
		obj.SetWait(fernTestValueWait)
		assert.Equal(t, fernTestValueWait, obj.Wait)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateConversationRequest(t *testing.T) {
	t.Run("SetWait_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationRequest{}
		var fernTestValueWait *int

		// Act
		obj.SetWait(fernTestValueWait)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateRequestRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CreateRequestRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWait", func(t *testing.T) {
		obj := &CreateRequestRequest{}
		var fernTestValueWait *int
		obj.SetWait(fernTestValueWait)
		assert.Equal(t, fernTestValueWait, obj.Wait)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateRequestRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetWait_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestRequest{}
		var fernTestValueWait *int

		// Act
		obj.SetWait(fernTestValueWait)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersGetConversationRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetConversationRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetConversationRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersGetRequestRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetRequestRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &GetRequestRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWait", func(t *testing.T) {
		obj := &GetRequestRequest{}
		var fernTestValueWait *int
		obj.SetWait(fernTestValueWait)
		assert.Equal(t, fernTestValueWait, obj.Wait)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetRequestRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestRequest{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetWait_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestRequest{}
		var fernTestValueWait *int

		// Act
		obj.SetWait(fernTestValueWait)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListIntentsRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &ListIntentsRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListIntentsRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIntentsRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListRequestsRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &ListRequestsRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListRequestsRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRequestsRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListResourcesRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &ListResourcesRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListResourcesRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListResourcesRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersStreamRequestEventsRequest(t *testing.T) {
	t.Run("SetLastEventID", func(t *testing.T) {
		obj := &StreamRequestEventsRequest{}
		var fernTestValueLastEventID *string
		obj.SetLastEventID(fernTestValueLastEventID)
		assert.Equal(t, fernTestValueLastEventID, obj.LastEventID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversationID", func(t *testing.T) {
		obj := &StreamRequestEventsRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &StreamRequestEventsRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitStreamRequestEventsRequest(t *testing.T) {
	t.Run("SetLastEventID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamRequestEventsRequest{}
		var fernTestValueLastEventID *string

		// Act
		obj.SetLastEventID(fernTestValueLastEventID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamRequestEventsRequest{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StreamRequestEventsRequest{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersConversation(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMesh", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueMesh *ConversationMeshVersion
		obj.SetMesh(fernTestValueMesh)
		assert.Equal(t, fernTestValueMesh, obj.Mesh)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueCreatedAt time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueUpdatedAt time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersConversation(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetExternalID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected *string
		obj.ExternalID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExternalID(), "getter should return the property value")
	})

	t.Run("GetExternalID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		obj.ExternalID = nil

		// Act & Assert
		assert.Nil(t, obj.GetExternalID(), "getter should return nil when property is nil")
	})

	t.Run("GetExternalID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExternalID() // Should return zero value
	})

	t.Run("GetMesh", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected *ConversationMeshVersion
		obj.Mesh = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMesh(), "getter should return the property value")
	})

	t.Run("GetMesh_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		obj.Mesh = nil

		// Act & Assert
		assert.Nil(t, obj.GetMesh(), "getter should return nil when property is nil")
	})

	t.Run("GetMesh_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMesh() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitConversation(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueExternalID *string

		// Act
		obj.SetExternalID(fernTestValueExternalID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMesh_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueMesh *ConversationMeshVersion

		// Act
		obj.SetMesh(fernTestValueMesh)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueTags []*Tag

		// Act
		obj.SetTags(fernTestValueTags)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueCreatedAt time.Time

		// Act
		obj.SetCreatedAt(fernTestValueCreatedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueUpdatedAt time.Time

		// Act
		obj.SetUpdatedAt(fernTestValueUpdatedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersConversationMesh(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ConversationMesh{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersion", func(t *testing.T) {
		obj := &ConversationMesh{}
		var fernTestValueVersion *int
		obj.SetVersion(fernTestValueVersion)
		assert.Equal(t, fernTestValueVersion, obj.Version)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionLabel", func(t *testing.T) {
		obj := &ConversationMesh{}
		var fernTestValueVersionLabel *string
		obj.SetVersionLabel(fernTestValueVersionLabel)
		assert.Equal(t, fernTestValueVersionLabel, obj.VersionLabel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDraft", func(t *testing.T) {
		obj := &ConversationMesh{}
		var fernTestValueDraft *bool
		obj.SetDraft(fernTestValueDraft)
		assert.Equal(t, fernTestValueDraft, obj.Draft)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersConversationMesh(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetVersion", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var expected *int
		obj.Version = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVersion(), "getter should return the property value")
	})

	t.Run("GetVersion_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		obj.Version = nil

		// Act & Assert
		assert.Nil(t, obj.GetVersion(), "getter should return nil when property is nil")
	})

	t.Run("GetVersion_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVersion() // Should return zero value
	})

	t.Run("GetVersionLabel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var expected *string
		obj.VersionLabel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVersionLabel(), "getter should return the property value")
	})

	t.Run("GetVersionLabel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		obj.VersionLabel = nil

		// Act & Assert
		assert.Nil(t, obj.GetVersionLabel(), "getter should return nil when property is nil")
	})

	t.Run("GetVersionLabel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVersionLabel() // Should return zero value
	})

	t.Run("GetDraft", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var expected *bool
		obj.Draft = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDraft(), "getter should return the property value")
	})

	t.Run("GetDraft_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		obj.Draft = nil

		// Act & Assert
		assert.Nil(t, obj.GetDraft(), "getter should return nil when property is nil")
	})

	t.Run("GetDraft_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDraft() // Should return zero value
	})

}

func TestSettersMarkExplicitConversationMesh(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVersion_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var fernTestValueVersion *int

		// Act
		obj.SetVersion(fernTestValueVersion)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVersionLabel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var fernTestValueVersionLabel *string

		// Act
		obj.SetVersionLabel(fernTestValueVersionLabel)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDraft_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}
		var fernTestValueDraft *bool

		// Act
		obj.SetDraft(fernTestValueDraft)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersConversationMeshVersion(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ConversationMeshVersion{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionID", func(t *testing.T) {
		obj := &ConversationMeshVersion{}
		var fernTestValueVersionID string
		obj.SetVersionID(fernTestValueVersionID)
		assert.Equal(t, fernTestValueVersionID, obj.VersionID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersConversationMeshVersion(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMeshVersion{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMeshVersion
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetVersionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMeshVersion{}
		var expected string
		obj.VersionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVersionID(), "getter should return the property value")
	})

	t.Run("GetVersionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMeshVersion
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVersionID() // Should return zero value
	})

}

func TestSettersMarkExplicitConversationMeshVersion(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMeshVersion{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVersionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMeshVersion{}
		var fernTestValueVersionID string

		// Act
		obj.SetVersionID(fernTestValueVersionID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateConversationParams(t *testing.T) {
	t.Run("SetMesh", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueMesh *ConversationMesh
		obj.SetMesh(fernTestValueMesh)
		assert.Equal(t, fernTestValueMesh, obj.Mesh)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRouteID", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueRouteID *string
		obj.SetRouteID(fernTestValueRouteID)
		assert.Equal(t, fernTestValueRouteID, obj.RouteID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNotificationChannel", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueNotificationChannel *string
		obj.SetNotificationChannel(fernTestValueNotificationChannel)
		assert.Equal(t, fernTestValueNotificationChannel, obj.NotificationChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracing", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueTracing *bool
		obj.SetTracing(fernTestValueTracing)
		assert.Equal(t, fernTestValueTracing, obj.Tracing)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIntent", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueIntent *string
		obj.SetIntent(fernTestValueIntent)
		assert.Equal(t, fernTestValueIntent, obj.Intent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInput", func(t *testing.T) {
		obj := &CreateConversationParams{}
		var fernTestValueInput map[string]any
		obj.SetInput(fernTestValueInput)
		assert.Equal(t, fernTestValueInput, obj.Input)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateConversationParams(t *testing.T) {
	t.Run("GetMesh", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *ConversationMesh
		obj.Mesh = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMesh(), "getter should return the property value")
	})

	t.Run("GetMesh_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.Mesh = nil

		// Act & Assert
		assert.Nil(t, obj.GetMesh(), "getter should return nil when property is nil")
	})

	t.Run("GetMesh_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMesh() // Should return zero value
	})

	t.Run("GetRouteID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *string
		obj.RouteID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRouteID(), "getter should return the property value")
	})

	t.Run("GetRouteID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.RouteID = nil

		// Act & Assert
		assert.Nil(t, obj.GetRouteID(), "getter should return nil when property is nil")
	})

	t.Run("GetRouteID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRouteID() // Should return zero value
	})

	t.Run("GetExternalID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *string
		obj.ExternalID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExternalID(), "getter should return the property value")
	})

	t.Run("GetExternalID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.ExternalID = nil

		// Act & Assert
		assert.Nil(t, obj.GetExternalID(), "getter should return nil when property is nil")
	})

	t.Run("GetExternalID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExternalID() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetNotificationChannel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *string
		obj.NotificationChannel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNotificationChannel(), "getter should return the property value")
	})

	t.Run("GetNotificationChannel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.NotificationChannel = nil

		// Act & Assert
		assert.Nil(t, obj.GetNotificationChannel(), "getter should return nil when property is nil")
	})

	t.Run("GetNotificationChannel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNotificationChannel() // Should return zero value
	})

	t.Run("GetTracing", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *bool
		obj.Tracing = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTracing(), "getter should return the property value")
	})

	t.Run("GetTracing_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.Tracing = nil

		// Act & Assert
		assert.Nil(t, obj.GetTracing(), "getter should return nil when property is nil")
	})

	t.Run("GetTracing_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTracing() // Should return zero value
	})

	t.Run("GetIntent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected *string
		obj.Intent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIntent(), "getter should return the property value")
	})

	t.Run("GetIntent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.Intent = nil

		// Act & Assert
		assert.Nil(t, obj.GetIntent(), "getter should return nil when property is nil")
	})

	t.Run("GetIntent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIntent() // Should return zero value
	})

	t.Run("GetInput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var expected map[string]any
		obj.Input = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInput(), "getter should return the property value")
	})

	t.Run("GetInput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		obj.Input = nil

		// Act & Assert
		assert.Nil(t, obj.GetInput(), "getter should return nil when property is nil")
	})

	t.Run("GetInput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInput() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateConversationParams(t *testing.T) {
	t.Run("SetMesh_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueMesh *ConversationMesh

		// Act
		obj.SetMesh(fernTestValueMesh)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRouteID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueRouteID *string

		// Act
		obj.SetRouteID(fernTestValueRouteID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueExternalID *string

		// Act
		obj.SetExternalID(fernTestValueExternalID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueTags []*Tag

		// Act
		obj.SetTags(fernTestValueTags)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetNotificationChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueNotificationChannel *string

		// Act
		obj.SetNotificationChannel(fernTestValueNotificationChannel)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTracing_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueTracing *bool

		// Act
		obj.SetTracing(fernTestValueTracing)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIntent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueIntent *string

		// Act
		obj.SetIntent(fernTestValueIntent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetInput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}
		var fernTestValueInput map[string]any

		// Act
		obj.SetInput(fernTestValueInput)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateRequestParams(t *testing.T) {
	t.Run("SetIntent", func(t *testing.T) {
		obj := &CreateRequestParams{}
		var fernTestValueIntent *string
		obj.SetIntent(fernTestValueIntent)
		assert.Equal(t, fernTestValueIntent, obj.Intent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInput", func(t *testing.T) {
		obj := &CreateRequestParams{}
		var fernTestValueInput map[string]any
		obj.SetInput(fernTestValueInput)
		assert.Equal(t, fernTestValueInput, obj.Input)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &CreateRequestParams{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateRequestParams(t *testing.T) {
	t.Run("GetIntent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var expected *string
		obj.Intent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIntent(), "getter should return the property value")
	})

	t.Run("GetIntent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		obj.Intent = nil

		// Act & Assert
		assert.Nil(t, obj.GetIntent(), "getter should return nil when property is nil")
	})

	t.Run("GetIntent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIntent() // Should return zero value
	})

	t.Run("GetInput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var expected map[string]any
		obj.Input = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInput(), "getter should return the property value")
	})

	t.Run("GetInput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		obj.Input = nil

		// Act & Assert
		assert.Nil(t, obj.GetInput(), "getter should return nil when property is nil")
	})

	t.Run("GetInput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInput() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestParams
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateRequestParams(t *testing.T) {
	t.Run("SetIntent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var fernTestValueIntent *string

		// Act
		obj.SetIntent(fernTestValueIntent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetInput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var fernTestValueInput map[string]any

		// Act
		obj.SetInput(fernTestValueInput)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}
		var fernTestValueTags []*Tag

		// Act
		obj.SetTags(fernTestValueTags)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreatedConversation(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMesh", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueMesh *ConversationMeshVersion
		obj.SetMesh(fernTestValueMesh)
		assert.Equal(t, fernTestValueMesh, obj.Mesh)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueCreatedAt time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueUpdatedAt time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequest", func(t *testing.T) {
		obj := &CreatedConversation{}
		var fernTestValueRequest *Request
		obj.SetRequest(fernTestValueRequest)
		assert.Equal(t, fernTestValueRequest, obj.Request)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreatedConversation(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetExternalID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected *string
		obj.ExternalID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExternalID(), "getter should return the property value")
	})

	t.Run("GetExternalID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		obj.ExternalID = nil

		// Act & Assert
		assert.Nil(t, obj.GetExternalID(), "getter should return nil when property is nil")
	})

	t.Run("GetExternalID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExternalID() // Should return zero value
	})

	t.Run("GetMesh", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected *ConversationMeshVersion
		obj.Mesh = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMesh(), "getter should return the property value")
	})

	t.Run("GetMesh_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		obj.Mesh = nil

		// Act & Assert
		assert.Nil(t, obj.GetMesh(), "getter should return nil when property is nil")
	})

	t.Run("GetMesh_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMesh() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

	t.Run("GetRequest", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var expected *Request
		obj.Request = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequest(), "getter should return the property value")
	})

	t.Run("GetRequest_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		obj.Request = nil

		// Act & Assert
		assert.Nil(t, obj.GetRequest(), "getter should return nil when property is nil")
	})

	t.Run("GetRequest_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequest() // Should return zero value
	})

}

func TestSettersMarkExplicitCreatedConversation(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueExternalID *string

		// Act
		obj.SetExternalID(fernTestValueExternalID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMesh_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueMesh *ConversationMeshVersion

		// Act
		obj.SetMesh(fernTestValueMesh)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueTags []*Tag

		// Act
		obj.SetTags(fernTestValueTags)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueCreatedAt time.Time

		// Act
		obj.SetCreatedAt(fernTestValueCreatedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueUpdatedAt time.Time

		// Act
		obj.SetUpdatedAt(fernTestValueUpdatedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequest_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}
		var fernTestValueRequest *Request

		// Act
		obj.SetRequest(fernTestValueRequest)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersIntent(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &Intent{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &Intent{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &Intent{}
		var fernTestValueSchema map[string]any
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetResponses", func(t *testing.T) {
		obj := &Intent{}
		var fernTestValueResponses []*IntentResponsesItem
		obj.SetResponses(fernTestValueResponses)
		assert.Equal(t, fernTestValueResponses, obj.Responses)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersIntent(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var expected map[string]any
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSchema() // Should return zero value
	})

	t.Run("GetResponses", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var expected []*IntentResponsesItem
		obj.Responses = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponses(), "getter should return the property value")
	})

	t.Run("GetResponses_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		obj.Responses = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponses(), "getter should return nil when property is nil")
	})

	t.Run("GetResponses_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponses() // Should return zero value
	})

}

func TestSettersMarkExplicitIntent(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var fernTestValueDescription *string

		// Act
		obj.SetDescription(fernTestValueDescription)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var fernTestValueSchema map[string]any

		// Act
		obj.SetSchema(fernTestValueSchema)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetResponses_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}
		var fernTestValueResponses []*IntentResponsesItem

		// Act
		obj.SetResponses(fernTestValueResponses)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersIntentResponsesItem(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &IntentResponsesItem{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &IntentResponsesItem{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &IntentResponsesItem{}
		var fernTestValueSchema map[string]any
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersIntentResponsesItem(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentResponsesItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentResponsesItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var expected map[string]any
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentResponsesItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitIntentResponsesItem(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var fernTestValueDescription *string

		// Act
		obj.SetDescription(fernTestValueDescription)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}
		var fernTestValueSchema map[string]any

		// Act
		obj.SetSchema(fernTestValueSchema)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersOpenClientTask(t *testing.T) {
	t.Run("SetResponseSchema", func(t *testing.T) {
		obj := &OpenClientTask{}
		var fernTestValueResponseSchema map[string]any
		obj.SetResponseSchema(fernTestValueResponseSchema)
		assert.Equal(t, fernTestValueResponseSchema, obj.ResponseSchema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetClientTask", func(t *testing.T) {
		obj := &OpenClientTask{}
		var fernTestValueClientTask *ClientTask
		obj.SetClientTask(fernTestValueClientTask)
		assert.Equal(t, fernTestValueClientTask, obj.ClientTask)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &OpenClientTask{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersOpenClientTask(t *testing.T) {
	t.Run("GetResponseSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var expected map[string]any
		obj.ResponseSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponseSchema(), "getter should return the property value")
	})

	t.Run("GetResponseSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		obj.ResponseSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponseSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetResponseSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponseSchema() // Should return zero value
	})

	t.Run("GetClientTask", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var expected *ClientTask
		obj.ClientTask = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClientTask(), "getter should return the property value")
	})

	t.Run("GetClientTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		obj.ClientTask = nil

		// Act & Assert
		assert.Nil(t, obj.GetClientTask(), "getter should return nil when property is nil")
	})

	t.Run("GetClientTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClientTask() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

}

func TestSettersMarkExplicitOpenClientTask(t *testing.T) {
	t.Run("SetResponseSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var fernTestValueResponseSchema map[string]any

		// Act
		obj.SetResponseSchema(fernTestValueResponseSchema)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetClientTask_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var fernTestValueClientTask *ClientTask

		// Act
		obj.SetClientTask(fernTestValueClientTask)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersOpenToolCallApproval(t *testing.T) {
	t.Run("SetResponseSchema", func(t *testing.T) {
		obj := &OpenToolCallApproval{}
		var fernTestValueResponseSchema map[string]any
		obj.SetResponseSchema(fernTestValueResponseSchema)
		assert.Equal(t, fernTestValueResponseSchema, obj.ResponseSchema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetToolCallApproval", func(t *testing.T) {
		obj := &OpenToolCallApproval{}
		var fernTestValueToolCallApproval *ToolCallApproval
		obj.SetToolCallApproval(fernTestValueToolCallApproval)
		assert.Equal(t, fernTestValueToolCallApproval, obj.ToolCallApproval)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &OpenToolCallApproval{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersOpenToolCallApproval(t *testing.T) {
	t.Run("GetResponseSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var expected map[string]any
		obj.ResponseSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponseSchema(), "getter should return the property value")
	})

	t.Run("GetResponseSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		obj.ResponseSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponseSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetResponseSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponseSchema() // Should return zero value
	})

	t.Run("GetToolCallApproval", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var expected *ToolCallApproval
		obj.ToolCallApproval = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetToolCallApproval(), "getter should return the property value")
	})

	t.Run("GetToolCallApproval_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		obj.ToolCallApproval = nil

		// Act & Assert
		assert.Nil(t, obj.GetToolCallApproval(), "getter should return nil when property is nil")
	})

	t.Run("GetToolCallApproval_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetToolCallApproval() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

}

func TestSettersMarkExplicitOpenToolCallApproval(t *testing.T) {
	t.Run("SetResponseSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var fernTestValueResponseSchema map[string]any

		// Act
		obj.SetResponseSchema(fernTestValueResponseSchema)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetToolCallApproval_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var fernTestValueToolCallApproval *ToolCallApproval

		// Act
		obj.SetToolCallApproval(fernTestValueToolCallApproval)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersRequest(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueStatus RequestStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInput", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueInput map[string]any
		obj.SetInput(fernTestValueInput)
		assert.Equal(t, fernTestValueInput, obj.Input)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueCreatedAt time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetResponse", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueResponse map[string]any
		obj.SetResponse(fernTestValueResponse)
		assert.Equal(t, fernTestValueResponse, obj.Response)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTasks", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueTasks []*Task
		obj.SetTasks(fernTestValueTasks)
		assert.Equal(t, fernTestValueTasks, obj.Tasks)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetError", func(t *testing.T) {
		obj := &Request{}
		var fernTestValueError *string
		obj.SetError(fernTestValueError)
		assert.Equal(t, fernTestValueError, obj.Error)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRequest(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected RequestStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetInput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected map[string]any
		obj.Input = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInput(), "getter should return the property value")
	})

	t.Run("GetInput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		obj.Input = nil

		// Act & Assert
		assert.Nil(t, obj.GetInput(), "getter should return nil when property is nil")
	})

	t.Run("GetInput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInput() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetResponse", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected map[string]any
		obj.Response = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponse(), "getter should return the property value")
	})

	t.Run("GetResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		obj.Response = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponse() // Should return zero value
	})

	t.Run("GetTasks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected []*Task
		obj.Tasks = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTasks(), "getter should return the property value")
	})

	t.Run("GetTasks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		obj.Tasks = nil

		// Act & Assert
		assert.Nil(t, obj.GetTasks(), "getter should return nil when property is nil")
	})

	t.Run("GetTasks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTasks() // Should return zero value
	})

	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var expected *string
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

}

func TestSettersMarkExplicitRequest(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueID string

		// Act
		obj.SetID(fernTestValueID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueStatus RequestStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetInput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueInput map[string]any

		// Act
		obj.SetInput(fernTestValueInput)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueTags []*Tag

		// Act
		obj.SetTags(fernTestValueTags)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueCreatedAt time.Time

		// Act
		obj.SetCreatedAt(fernTestValueCreatedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetResponse_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueResponse map[string]any

		// Act
		obj.SetResponse(fernTestValueResponse)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTasks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueTasks []*Task

		// Act
		obj.SetTasks(fernTestValueTasks)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetError_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}
		var fernTestValueError *string

		// Act
		obj.SetError(fernTestValueError)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestGettersRequestStreamEvent(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetStarted", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *StartedEvent
		obj.Started = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStarted(), "getter should return the property value")
	})

	t.Run("GetStarted_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Started = nil

		// Act & Assert
		assert.Nil(t, obj.GetStarted(), "getter should return nil when property is nil")
	})

	t.Run("GetStarted_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStarted() // Should return zero value
	})

	t.Run("GetTask", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *TaskEvent
		obj.Task = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTask(), "getter should return the property value")
	})

	t.Run("GetTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Task = nil

		// Act & Assert
		assert.Nil(t, obj.GetTask(), "getter should return nil when property is nil")
	})

	t.Run("GetTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTask() // Should return zero value
	})

	t.Run("GetTaskResponse", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *TaskResponseEvent
		obj.TaskResponse = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskResponse(), "getter should return the property value")
	})

	t.Run("GetTaskResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.TaskResponse = nil

		// Act & Assert
		assert.Nil(t, obj.GetTaskResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetTaskResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTaskResponse() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *MessageEvent
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

	t.Run("GetThinking", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *ThinkingEvent
		obj.Thinking = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetThinking(), "getter should return the property value")
	})

	t.Run("GetThinking_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Thinking = nil

		// Act & Assert
		assert.Nil(t, obj.GetThinking(), "getter should return nil when property is nil")
	})

	t.Run("GetThinking_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetThinking() // Should return zero value
	})

	t.Run("GetResponse", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *ResponseEvent
		obj.Response = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponse(), "getter should return the property value")
	})

	t.Run("GetResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Response = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponse() // Should return zero value
	})

	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *ErrorEvent
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

	t.Run("GetCancelled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *CancelledEvent
		obj.Cancelled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCancelled(), "getter should return the property value")
	})

	t.Run("GetCancelled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.Cancelled = nil

		// Act & Assert
		assert.Nil(t, obj.GetCancelled(), "getter should return nil when property is nil")
	})

	t.Run("GetCancelled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCancelled() // Should return zero value
	})

	t.Run("GetServerError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		var expected *ServerErrorEvent
		obj.ServerError = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetServerError(), "getter should return the property value")
	})

	t.Run("GetServerError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestStreamEvent{}
		obj.ServerError = nil

		// Act & Assert
		assert.Nil(t, obj.GetServerError(), "getter should return nil when property is nil")
	})

	t.Run("GetServerError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestStreamEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetServerError() // Should return zero value
	})

}

func TestSettersResource(t *testing.T) {
	t.Run("SetReference", func(t *testing.T) {
		obj := &Resource{}
		var fernTestValueReference Reference
		obj.SetReference(fernTestValueReference)
		assert.Equal(t, fernTestValueReference, obj.Reference)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &Resource{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &Resource{}
		var fernTestValueDescription string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMimeType", func(t *testing.T) {
		obj := &Resource{}
		var fernTestValueMimeType string
		obj.SetMimeType(fernTestValueMimeType)
		assert.Equal(t, fernTestValueMimeType, obj.MimeType)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersResource(t *testing.T) {
	t.Run("GetReference", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var expected Reference
		obj.Reference = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetReference(), "getter should return the property value")
	})

	t.Run("GetReference_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetReference() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var expected string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetMimeType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var expected string
		obj.MimeType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMimeType(), "getter should return the property value")
	})

	t.Run("GetMimeType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMimeType() // Should return zero value
	})

}

func TestSettersMarkExplicitResource(t *testing.T) {
	t.Run("SetReference_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var fernTestValueReference Reference

		// Act
		obj.SetReference(fernTestValueReference)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var fernTestValueDescription string

		// Act
		obj.SetDescription(fernTestValueDescription)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMimeType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}
		var fernTestValueMimeType string

		// Act
		obj.SetMimeType(fernTestValueMimeType)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersResourceList(t *testing.T) {
	t.Run("SetResources", func(t *testing.T) {
		obj := &ResourceList{}
		var fernTestValueResources []*Resource
		obj.SetResources(fernTestValueResources)
		assert.Equal(t, fernTestValueResources, obj.Resources)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersResourceList(t *testing.T) {
	t.Run("GetResources", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResourceList{}
		var expected []*Resource
		obj.Resources = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResources(), "getter should return the property value")
	})

	t.Run("GetResources_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResourceList{}
		obj.Resources = nil

		// Act & Assert
		assert.Nil(t, obj.GetResources(), "getter should return nil when property is nil")
	})

	t.Run("GetResources_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResourceList
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResources() // Should return zero value
	})

}

func TestSettersMarkExplicitResourceList(t *testing.T) {
	t.Run("SetResources_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResourceList{}
		var fernTestValueResources []*Resource

		// Act
		obj.SetResources(fernTestValueResources)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersStartedEvent(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &StartedEvent{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &StartedEvent{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersStartedEvent(t *testing.T) {
	t.Run("GetConversationID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StartedEvent{}
		var expected string
		obj.ConversationID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationID(), "getter should return the property value")
	})

	t.Run("GetConversationID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StartedEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationID() // Should return zero value
	})

	t.Run("GetRequestID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StartedEvent{}
		var expected string
		obj.RequestID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestID(), "getter should return the property value")
	})

	t.Run("GetRequestID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StartedEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestID() // Should return zero value
	})

}

func TestSettersMarkExplicitStartedEvent(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StartedEvent{}
		var fernTestValueConversationID string

		// Act
		obj.SetConversationID(fernTestValueConversationID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StartedEvent{}
		var fernTestValueRequestID string

		// Act
		obj.SetRequestID(fernTestValueRequestID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersTag(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &Tag{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetValue", func(t *testing.T) {
		obj := &Tag{}
		var fernTestValueValue string
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTag(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Tag{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Tag
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Tag{}
		var expected string
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Tag
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitTag(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Tag{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Tag{}
		var fernTestValueValue string

		// Act
		obj.SetValue(fernTestValueValue)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestGettersTask(t *testing.T) {
	t.Run("GetKind", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Task{}
		var expected string
		obj.Kind = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetKind(), "getter should return the property value")
	})

	t.Run("GetKind_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Task
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetKind() // Should return zero value
	})

	t.Run("GetClientTask", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Task{}
		var expected *OpenClientTask
		obj.ClientTask = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClientTask(), "getter should return the property value")
	})

	t.Run("GetClientTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Task{}
		obj.ClientTask = nil

		// Act & Assert
		assert.Nil(t, obj.GetClientTask(), "getter should return nil when property is nil")
	})

	t.Run("GetClientTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Task
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClientTask() // Should return zero value
	})

	t.Run("GetToolCallApproval", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Task{}
		var expected *OpenToolCallApproval
		obj.ToolCallApproval = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetToolCallApproval(), "getter should return the property value")
	})

	t.Run("GetToolCallApproval_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Task{}
		obj.ToolCallApproval = nil

		// Act & Assert
		assert.Nil(t, obj.GetToolCallApproval(), "getter should return nil when property is nil")
	})

	t.Run("GetToolCallApproval_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Task
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetToolCallApproval() // Should return zero value
	})

}

func TestSettersListIntentsResponse(t *testing.T) {
	t.Run("SetIntents", func(t *testing.T) {
		obj := &ListIntentsResponse{}
		var fernTestValueIntents []*Intent
		obj.SetIntents(fernTestValueIntents)
		assert.Equal(t, fernTestValueIntents, obj.Intents)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListIntentsResponse(t *testing.T) {
	t.Run("GetIntents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIntentsResponse{}
		var expected []*Intent
		obj.Intents = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIntents(), "getter should return the property value")
	})

	t.Run("GetIntents_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIntentsResponse{}
		obj.Intents = nil

		// Act & Assert
		assert.Nil(t, obj.GetIntents(), "getter should return nil when property is nil")
	})

	t.Run("GetIntents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIntentsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIntents() // Should return zero value
	})

}

func TestSettersMarkExplicitListIntentsResponse(t *testing.T) {
	t.Run("SetIntents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIntentsResponse{}
		var fernTestValueIntents []*Intent

		// Act
		obj.SetIntents(fernTestValueIntents)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListRequestsResponse(t *testing.T) {
	t.Run("SetRequests", func(t *testing.T) {
		obj := &ListRequestsResponse{}
		var fernTestValueRequests []*Request
		obj.SetRequests(fernTestValueRequests)
		assert.Equal(t, fernTestValueRequests, obj.Requests)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListRequestsResponse(t *testing.T) {
	t.Run("GetRequests", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRequestsResponse{}
		var expected []*Request
		obj.Requests = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequests(), "getter should return the property value")
	})

	t.Run("GetRequests_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRequestsResponse{}
		obj.Requests = nil

		// Act & Assert
		assert.Nil(t, obj.GetRequests(), "getter should return nil when property is nil")
	})

	t.Run("GetRequests_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRequestsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequests() // Should return zero value
	})

}

func TestSettersMarkExplicitListRequestsResponse(t *testing.T) {
	t.Run("SetRequests_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRequestsResponse{}
		var fernTestValueRequests []*Request

		// Act
		obj.SetRequests(fernTestValueRequests)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestJSONMarshalingConversation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Conversation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Conversation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Conversation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingConversationMesh(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMesh{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ConversationMesh
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ConversationMesh
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ConversationMesh
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingConversationMeshVersion(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationMeshVersion{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ConversationMeshVersion
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ConversationMeshVersion
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ConversationMeshVersion
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateConversationParams(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationParams{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateConversationParams
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationParams
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationParams
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateRequestParams(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestParams{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateRequestParams
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestParams
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestParams
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreatedConversation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreatedConversation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreatedConversation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreatedConversation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreatedConversation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingIntent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Intent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Intent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Intent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Intent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingIntentResponsesItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentResponsesItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled IntentResponsesItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj IntentResponsesItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj IntentResponsesItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListIntentsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListIntentsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListIntentsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListIntentsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListIntentsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListRequestsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRequestsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListRequestsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListRequestsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListRequestsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingOpenClientTask(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenClientTask{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled OpenClientTask
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj OpenClientTask
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj OpenClientTask
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingOpenToolCallApproval(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &OpenToolCallApproval{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled OpenToolCallApproval
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj OpenToolCallApproval
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj OpenToolCallApproval
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRequest(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Request{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Request
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Request
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Request
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingResource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Resource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Resource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Resource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Resource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingResourceList(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResourceList{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ResourceList
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ResourceList
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ResourceList
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingStartedEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &StartedEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled StartedEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj StartedEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj StartedEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTag(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Tag{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Tag
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Tag
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Tag
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringConversation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Conversation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringConversationMesh(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationMesh{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringConversationMeshVersion(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationMeshVersion{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMeshVersion
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateConversationParams(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationParams{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateRequestParams(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestParams{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestParams
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreatedConversation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreatedConversation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringIntent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Intent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringIntentResponsesItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &IntentResponsesItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentResponsesItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListIntentsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListIntentsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIntentsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListRequestsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListRequestsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRequestsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringOpenClientTask(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &OpenClientTask{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenClientTask
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringOpenToolCallApproval(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &OpenToolCallApproval{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenToolCallApproval
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRequest(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Request{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringResource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Resource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringResourceList(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ResourceList{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResourceList
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringStartedEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &StartedEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StartedEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTag(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Tag{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Tag
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumRequestStatus(t *testing.T) {
	t.Run("NewFromString_running", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestStatusFromString("running")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestStatus("running"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_waiting", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestStatusFromString("waiting")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestStatus("waiting"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_completed", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestStatusFromString("completed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestStatus("completed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_failed", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestStatusFromString("failed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestStatus("failed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_cancelled", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestStatusFromString("cancelled")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestStatus("cancelled"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewRequestStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewRequestStatusFromString("running")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesConversation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Conversation{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesConversationMesh(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationMesh{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMesh
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesConversationMeshVersion(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationMeshVersion{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationMeshVersion
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateConversationParams(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationParams{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationParams
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateRequestParams(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestParams{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestParams
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreatedConversation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreatedConversation{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreatedConversation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesIntent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Intent{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Intent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesIntentResponsesItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &IntentResponsesItem{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentResponsesItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListIntentsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListIntentsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListIntentsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListRequestsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListRequestsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRequestsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesOpenClientTask(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &OpenClientTask{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenClientTask
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesOpenToolCallApproval(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &OpenToolCallApproval{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *OpenToolCallApproval
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRequest(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Request{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Request
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesResource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Resource{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Resource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesResourceList(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ResourceList{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResourceList
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesStartedEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &StartedEvent{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *StartedEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTag(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Tag{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Tag
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
