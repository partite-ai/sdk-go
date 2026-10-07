// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersAttachAssetsToConversationRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &AttachAssetsToConversationRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitAttachAssetsToConversationRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachAssetsToConversationRequest{}
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

func TestSettersCreateConversationRequestRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CreateConversationRequestRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIntentName", func(t *testing.T) {
		obj := &CreateConversationRequestRequest{}
		var fernTestValueIntentName string
		obj.SetIntentName(fernTestValueIntentName)
		assert.Equal(t, fernTestValueIntentName, obj.IntentName)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateConversationRequestRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationRequestRequest{}
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

	t.Run("SetIntentName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationRequestRequest{}
		var fernTestValueIntentName string

		// Act
		obj.SetIntentName(fernTestValueIntentName)

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

func TestSettersCreateConversationWithInitialRequestRequest(t *testing.T) {
	t.Run("SetIntentName", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueIntentName string
		obj.SetIntentName(fernTestValueIntentName)
		assert.Equal(t, fernTestValueIntentName, obj.IntentName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMeshID", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueMeshID string
		obj.SetMeshID(fernTestValueMeshID)
		assert.Equal(t, fernTestValueMeshID, obj.MeshID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionOrdinal", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueVersionOrdinal *int
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)
		assert.Equal(t, fernTestValueVersionOrdinal, obj.VersionOrdinal)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionLabel", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueVersionLabel *string
		obj.SetVersionLabel(fernTestValueVersionLabel)
		assert.Equal(t, fernTestValueVersionLabel, obj.VersionLabel)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateConversationWithInitialRequestRequest(t *testing.T) {
	t.Run("SetIntentName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueIntentName string

		// Act
		obj.SetIntentName(fernTestValueIntentName)

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

	t.Run("SetMeshID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueMeshID string

		// Act
		obj.SetMeshID(fernTestValueMeshID)

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

	t.Run("SetVersionOrdinal_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequestRequest{}
		var fernTestValueVersionOrdinal *int

		// Act
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)

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
		obj := &CreateConversationWithInitialRequestRequest{}
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

}

func TestSettersCreateEmptyConversationRequest(t *testing.T) {
	t.Run("SetMeshID", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueMeshID string
		obj.SetMeshID(fernTestValueMeshID)
		assert.Equal(t, fernTestValueMeshID, obj.MeshID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionOrdinal", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueVersionOrdinal *int
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)
		assert.Equal(t, fernTestValueVersionOrdinal, obj.VersionOrdinal)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionLabel", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueVersionLabel *string
		obj.SetVersionLabel(fernTestValueVersionLabel)
		assert.Equal(t, fernTestValueVersionLabel, obj.VersionLabel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNotificationChannel", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueNotificationChannel *string
		obj.SetNotificationChannel(fernTestValueNotificationChannel)
		assert.Equal(t, fernTestValueNotificationChannel, obj.NotificationChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracingEnabled", func(t *testing.T) {
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueTracingEnabled *bool
		obj.SetTracingEnabled(fernTestValueTracingEnabled)
		assert.Equal(t, fernTestValueTracingEnabled, obj.TracingEnabled)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateEmptyConversationRequest(t *testing.T) {
	t.Run("SetMeshID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueMeshID string

		// Act
		obj.SetMeshID(fernTestValueMeshID)

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

	t.Run("SetVersionOrdinal_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueVersionOrdinal *int

		// Act
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)

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
		obj := &CreateEmptyConversationRequest{}
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

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
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

	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
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

	t.Run("SetNotificationChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
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

	t.Run("SetTracingEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateEmptyConversationRequest{}
		var fernTestValueTracingEnabled *bool

		// Act
		obj.SetTracingEnabled(fernTestValueTracingEnabled)

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

func TestSettersGetConversationByExternalIDRequest(t *testing.T) {
	t.Run("SetExternalID", func(t *testing.T) {
		obj := &GetConversationByExternalIDRequest{}
		var fernTestValueExternalID string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetConversationByExternalIDRequest(t *testing.T) {
	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationByExternalIDRequest{}
		var fernTestValueExternalID string

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

}

func TestSettersGetConversationRequestsRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetConversationRequestsRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetConversationRequestsRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsRequest{}
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

func TestSettersGetIntentsForConversationRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetIntentsForConversationRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetIntentsForConversationRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIntentsForConversationRequest{}
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

func TestSettersGetIntentsForNewConversationRequest(t *testing.T) {
	t.Run("SetMeshID", func(t *testing.T) {
		obj := &GetIntentsForNewConversationRequest{}
		var fernTestValueMeshID string
		obj.SetMeshID(fernTestValueMeshID)
		assert.Equal(t, fernTestValueMeshID, obj.MeshID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionOrdinal", func(t *testing.T) {
		obj := &GetIntentsForNewConversationRequest{}
		var fernTestValueVersionOrdinal *int
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)
		assert.Equal(t, fernTestValueVersionOrdinal, obj.VersionOrdinal)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVersionLabel", func(t *testing.T) {
		obj := &GetIntentsForNewConversationRequest{}
		var fernTestValueVersionLabel *string
		obj.SetVersionLabel(fernTestValueVersionLabel)
		assert.Equal(t, fernTestValueVersionLabel, obj.VersionLabel)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetIntentsForNewConversationRequest(t *testing.T) {
	t.Run("SetMeshID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIntentsForNewConversationRequest{}
		var fernTestValueMeshID string

		// Act
		obj.SetMeshID(fernTestValueMeshID)

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

	t.Run("SetVersionOrdinal_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetIntentsForNewConversationRequest{}
		var fernTestValueVersionOrdinal *int

		// Act
		obj.SetVersionOrdinal(fernTestValueVersionOrdinal)

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
		obj := &GetIntentsForNewConversationRequest{}
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

}

func TestSettersGetRequestEventsRequest(t *testing.T) {
	t.Run("SetLastEventID", func(t *testing.T) {
		obj := &GetRequestEventsRequest{}
		var fernTestValueLastEventID *string
		obj.SetLastEventID(fernTestValueLastEventID)
		assert.Equal(t, fernTestValueLastEventID, obj.LastEventID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetRequestEventsRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &GetRequestEventsRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetRequestEventsRequest(t *testing.T) {
	t.Run("SetLastEventID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestEventsRequest{}
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
		obj := &GetRequestEventsRequest{}
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
		obj := &GetRequestEventsRequest{}
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

func TestSettersGetRequestResponseRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetRequestResponseRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &GetRequestResponseRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetRequestResponseRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestResponseRequest{}
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
		obj := &GetRequestResponseRequest{}
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

func TestSettersGetRequestTaskRequest(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &GetRequestTaskRequest{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &GetRequestTaskRequest{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetRequestTaskRequest(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetRequestTaskRequest{}
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
		obj := &GetRequestTaskRequest{}
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

func TestSettersAttachmentInfo(t *testing.T) {
	t.Run("SetReferenceURL", func(t *testing.T) {
		obj := &AttachmentInfo{}
		var fernTestValueReferenceURL string
		obj.SetReferenceURL(fernTestValueReferenceURL)
		assert.Equal(t, fernTestValueReferenceURL, obj.ReferenceURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFilename", func(t *testing.T) {
		obj := &AttachmentInfo{}
		var fernTestValueFilename string
		obj.SetFilename(fernTestValueFilename)
		assert.Equal(t, fernTestValueFilename, obj.Filename)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAttachmentInfo(t *testing.T) {
	t.Run("GetReferenceURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachmentInfo{}
		var expected string
		obj.ReferenceURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetReferenceURL(), "getter should return the property value")
	})

	t.Run("GetReferenceURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AttachmentInfo
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetReferenceURL() // Should return zero value
	})

	t.Run("GetFilename", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachmentInfo{}
		var expected string
		obj.Filename = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFilename(), "getter should return the property value")
	})

	t.Run("GetFilename_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AttachmentInfo
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFilename() // Should return zero value
	})

}

func TestSettersMarkExplicitAttachmentInfo(t *testing.T) {
	t.Run("SetReferenceURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachmentInfo{}
		var fernTestValueReferenceURL string

		// Act
		obj.SetReferenceURL(fernTestValueReferenceURL)

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

	t.Run("SetFilename_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachmentInfo{}
		var fernTestValueFilename string

		// Act
		obj.SetFilename(fernTestValueFilename)

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
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMeshID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueMeshID string
		obj.SetMeshID(fernTestValueMeshID)
		assert.Equal(t, fernTestValueMeshID, obj.MeshID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMeshVersionID", func(t *testing.T) {
		obj := &Conversation{}
		var fernTestValueMeshVersionID string
		obj.SetMeshVersionID(fernTestValueMeshVersionID)
		assert.Equal(t, fernTestValueMeshVersionID, obj.MeshVersionID)
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
	t.Run("GetConversationID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected string
		obj.ConversationID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationID(), "getter should return the property value")
	})

	t.Run("GetConversationID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationID() // Should return zero value
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

	t.Run("GetMeshID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected string
		obj.MeshID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMeshID(), "getter should return the property value")
	})

	t.Run("GetMeshID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMeshID() // Should return zero value
	})

	t.Run("GetMeshVersionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var expected string
		obj.MeshVersionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMeshVersionID(), "getter should return the property value")
	})

	t.Run("GetMeshVersionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Conversation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMeshVersionID() // Should return zero value
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
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
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

	t.Run("SetMeshID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueMeshID string

		// Act
		obj.SetMeshID(fernTestValueMeshID)

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

	t.Run("SetMeshVersionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Conversation{}
		var fernTestValueMeshVersionID string

		// Act
		obj.SetMeshVersionID(fernTestValueMeshVersionID)

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

func TestSettersConversationCreateOptions(t *testing.T) {
	t.Run("SetTags", func(t *testing.T) {
		obj := &ConversationCreateOptions{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExternalID", func(t *testing.T) {
		obj := &ConversationCreateOptions{}
		var fernTestValueExternalID *string
		obj.SetExternalID(fernTestValueExternalID)
		assert.Equal(t, fernTestValueExternalID, obj.ExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNotificationChannel", func(t *testing.T) {
		obj := &ConversationCreateOptions{}
		var fernTestValueNotificationChannel *string
		obj.SetNotificationChannel(fernTestValueNotificationChannel)
		assert.Equal(t, fernTestValueNotificationChannel, obj.NotificationChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracingEnabled", func(t *testing.T) {
		obj := &ConversationCreateOptions{}
		var fernTestValueTracingEnabled *bool
		obj.SetTracingEnabled(fernTestValueTracingEnabled)
		assert.Equal(t, fernTestValueTracingEnabled, obj.TracingEnabled)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersConversationCreateOptions(t *testing.T) {
	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationCreateOptions
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetExternalID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		var expected *string
		obj.ExternalID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetExternalID(), "getter should return the property value")
	})

	t.Run("GetExternalID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		obj.ExternalID = nil

		// Act & Assert
		assert.Nil(t, obj.GetExternalID(), "getter should return nil when property is nil")
	})

	t.Run("GetExternalID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationCreateOptions
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetExternalID() // Should return zero value
	})

	t.Run("GetNotificationChannel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		var expected *string
		obj.NotificationChannel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNotificationChannel(), "getter should return the property value")
	})

	t.Run("GetNotificationChannel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		obj.NotificationChannel = nil

		// Act & Assert
		assert.Nil(t, obj.GetNotificationChannel(), "getter should return nil when property is nil")
	})

	t.Run("GetNotificationChannel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationCreateOptions
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNotificationChannel() // Should return zero value
	})

	t.Run("GetTracingEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		var expected *bool
		obj.TracingEnabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTracingEnabled(), "getter should return the property value")
	})

	t.Run("GetTracingEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		obj.TracingEnabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetTracingEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetTracingEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationCreateOptions
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTracingEnabled() // Should return zero value
	})

}

func TestSettersMarkExplicitConversationCreateOptions(t *testing.T) {
	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
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

	t.Run("SetExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
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

	t.Run("SetNotificationChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
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

	t.Run("SetTracingEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}
		var fernTestValueTracingEnabled *bool

		// Act
		obj.SetTracingEnabled(fernTestValueTracingEnabled)

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

func TestSettersCreateConversationResponse(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CreateConversationResponse{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIntents", func(t *testing.T) {
		obj := &CreateConversationResponse{}
		var fernTestValueIntents IntentsList
		obj.SetIntents(fernTestValueIntents)
		assert.Equal(t, fernTestValueIntents, obj.Intents)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateConversationResponse(t *testing.T) {
	t.Run("GetConversationID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationResponse{}
		var expected string
		obj.ConversationID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationID(), "getter should return the property value")
	})

	t.Run("GetConversationID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationID() // Should return zero value
	})

	t.Run("GetIntents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationResponse{}
		var expected IntentsList
		obj.Intents = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIntents(), "getter should return the property value")
	})

	t.Run("GetIntents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIntents() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateConversationResponse(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationResponse{}
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

	t.Run("SetIntents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationResponse{}
		var fernTestValueIntents IntentsList

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

func TestSettersCreateRequestResponse(t *testing.T) {
	t.Run("SetConversationID", func(t *testing.T) {
		obj := &CreateRequestResponse{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &CreateRequestResponse{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateRequestResponse(t *testing.T) {
	t.Run("GetConversationID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestResponse{}
		var expected string
		obj.ConversationID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationID(), "getter should return the property value")
	})

	t.Run("GetConversationID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestResponse
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
		obj := &CreateRequestResponse{}
		var expected string
		obj.RequestID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestID(), "getter should return the property value")
	})

	t.Run("GetRequestID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestID() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateRequestResponse(t *testing.T) {
	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestResponse{}
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
		obj := &CreateRequestResponse{}
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

func TestSettersGetTaskResponse(t *testing.T) {
	t.Run("SetTaskID", func(t *testing.T) {
		obj := &GetTaskResponse{}
		var fernTestValueTaskID string
		obj.SetTaskID(fernTestValueTaskID)
		assert.Equal(t, fernTestValueTaskID, obj.TaskID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &GetTaskResponse{}
		var fernTestValueData *TaskData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetTaskResponse(t *testing.T) {
	t.Run("GetTaskID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}
		var expected string
		obj.TaskID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskID(), "getter should return the property value")
	})

	t.Run("GetTaskID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTaskResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTaskID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}
		var expected *TaskData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTaskResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitGetTaskResponse(t *testing.T) {
	t.Run("SetTaskID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}
		var fernTestValueData *TaskData

		// Act
		obj.SetData(fernTestValueData)

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

func TestSettersIntentsListItem(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &IntentsListItem{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &IntentsListItem{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &IntentsListItem{}
		var fernTestValueSchema map[string]any
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetResponses", func(t *testing.T) {
		obj := &IntentsListItem{}
		var fernTestValueResponses []*IntentsListItemResponsesItem
		obj.SetResponses(fernTestValueResponses)
		assert.Equal(t, fernTestValueResponses, obj.Responses)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersIntentsListItem(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItem
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
		obj := &IntentsListItem{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItem
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
		obj := &IntentsListItem{}
		var expected map[string]any
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItem
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
		obj := &IntentsListItem{}
		var expected []*IntentsListItemResponsesItem
		obj.Responses = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponses(), "getter should return the property value")
	})

	t.Run("GetResponses_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}
		obj.Responses = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponses(), "getter should return nil when property is nil")
	})

	t.Run("GetResponses_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetResponses() // Should return zero value
	})

}

func TestSettersMarkExplicitIntentsListItem(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}
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
		obj := &IntentsListItem{}
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
		obj := &IntentsListItem{}
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
		obj := &IntentsListItem{}
		var fernTestValueResponses []*IntentsListItemResponsesItem

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

func TestSettersIntentsListItemResponsesItem(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &IntentsListItemResponsesItem{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &IntentsListItemResponsesItem{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &IntentsListItemResponsesItem{}
		var fernTestValueSchema map[string]any
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersIntentsListItemResponsesItem(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItemResponsesItem
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
		obj := &IntentsListItemResponsesItem{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItemResponsesItem
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
		obj := &IntentsListItemResponsesItem{}
		var expected map[string]any
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItemResponsesItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitIntentsListItemResponsesItem(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}
		var fernTestValueName *string

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
		obj := &IntentsListItemResponsesItem{}
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
		obj := &IntentsListItemResponsesItem{}
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

func TestSettersListIntentsResponse(t *testing.T) {
	t.Run("SetIntents", func(t *testing.T) {
		obj := &ListIntentsResponse{}
		var fernTestValueIntents IntentsList
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
		var expected IntentsList
		obj.Intents = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIntents(), "getter should return the property value")
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
		var fernTestValueIntents IntentsList

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

func TestSettersRequestCreateOptions(t *testing.T) {
	t.Run("SetTags", func(t *testing.T) {
		obj := &RequestCreateOptions{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRequestCreateOptions(t *testing.T) {
	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestCreateOptions{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestCreateOptions{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestCreateOptions
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

}

func TestSettersMarkExplicitRequestCreateOptions(t *testing.T) {
	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestCreateOptions{}
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

func TestGettersResponseEventData(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetTask", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		var expected *TaskEventDetails
		obj.Task = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTask(), "getter should return the property value")
	})

	t.Run("GetTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Task = nil

		// Act & Assert
		assert.Nil(t, obj.GetTask(), "getter should return nil when property is nil")
	})

	t.Run("GetTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
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
		obj := &ResponseEventData{}
		var expected *TaskResponseEventDetails
		obj.TaskResponse = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskResponse(), "getter should return the property value")
	})

	t.Run("GetTaskResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.TaskResponse = nil

		// Act & Assert
		assert.Nil(t, obj.GetTaskResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetTaskResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
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
		obj := &ResponseEventData{}
		var expected *MessageEventDetails
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
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
		obj := &ResponseEventData{}
		var expected *ThinkingEventDetails
		obj.Thinking = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetThinking(), "getter should return the property value")
	})

	t.Run("GetThinking_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Thinking = nil

		// Act & Assert
		assert.Nil(t, obj.GetThinking(), "getter should return nil when property is nil")
	})

	t.Run("GetThinking_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
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
		obj := &ResponseEventData{}
		var expected *ResponseEventDetails
		obj.Response = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponse(), "getter should return the property value")
	})

	t.Run("GetResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Response = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
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
		obj := &ResponseEventData{}
		var expected *ErrorEventDetails
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

	t.Run("GetCancel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		var expected *CancelEventDetails
		obj.Cancel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCancel(), "getter should return the property value")
	})

	t.Run("GetCancel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.Cancel = nil

		// Act & Assert
		assert.Nil(t, obj.GetCancel(), "getter should return nil when property is nil")
	})

	t.Run("GetCancel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCancel() // Should return zero value
	})

	t.Run("GetServerError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		var expected *ServerErrorEventDetails
		obj.ServerError = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetServerError(), "getter should return the property value")
	})

	t.Run("GetServerError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventData{}
		obj.ServerError = nil

		// Act & Assert
		assert.Nil(t, obj.GetServerError(), "getter should return nil when property is nil")
	})

	t.Run("GetServerError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetServerError() // Should return zero value
	})

}

func TestSettersServerErrorEventDetails(t *testing.T) {
	t.Run("SetType", func(t *testing.T) {
		obj := &ServerErrorEventDetails{}
		var fernTestValueType ServerErrorEventDetailsType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetError", func(t *testing.T) {
		obj := &ServerErrorEventDetails{}
		var fernTestValueError string
		obj.SetError(fernTestValueError)
		assert.Equal(t, fernTestValueError, obj.Error)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersServerErrorEventDetails(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEventDetails{}
		var expected ServerErrorEventDetailsType
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ServerErrorEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEventDetails{}
		var expected string
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ServerErrorEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

}

func TestSettersMarkExplicitServerErrorEventDetails(t *testing.T) {
	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEventDetails{}
		var fernTestValueType ServerErrorEventDetailsType

		// Act
		obj.SetType(fernTestValueType)

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
		obj := &ServerErrorEventDetails{}
		var fernTestValueError string

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

func TestSettersAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("SetAttachments", func(t *testing.T) {
		obj := &AttachAssetsToConversationResponse{}
		var fernTestValueAttachments []*AttachmentInfo
		obj.SetAttachments(fernTestValueAttachments)
		assert.Equal(t, fernTestValueAttachments, obj.Attachments)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("GetAttachments", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachAssetsToConversationResponse{}
		var expected []*AttachmentInfo
		obj.Attachments = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAttachments(), "getter should return the property value")
	})

	t.Run("GetAttachments_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachAssetsToConversationResponse{}
		obj.Attachments = nil

		// Act & Assert
		assert.Nil(t, obj.GetAttachments(), "getter should return nil when property is nil")
	})

	t.Run("GetAttachments_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AttachAssetsToConversationResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAttachments() // Should return zero value
	})

}

func TestSettersMarkExplicitAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("SetAttachments_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachAssetsToConversationResponse{}
		var fernTestValueAttachments []*AttachmentInfo

		// Act
		obj.SetAttachments(fernTestValueAttachments)

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

func TestSettersGetConversationRequestsResponse(t *testing.T) {
	t.Run("SetRequests", func(t *testing.T) {
		obj := &GetConversationRequestsResponse{}
		var fernTestValueRequests []*GetConversationRequestsResponseRequestsItem
		obj.SetRequests(fernTestValueRequests)
		assert.Equal(t, fernTestValueRequests, obj.Requests)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetConversationRequestsResponse(t *testing.T) {
	t.Run("GetRequests", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponse{}
		var expected []*GetConversationRequestsResponseRequestsItem
		obj.Requests = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequests(), "getter should return the property value")
	})

	t.Run("GetRequests_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponse{}
		obj.Requests = nil

		// Act & Assert
		assert.Nil(t, obj.GetRequests(), "getter should return nil when property is nil")
	})

	t.Run("GetRequests_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequests() // Should return zero value
	})

}

func TestSettersMarkExplicitGetConversationRequestsResponse(t *testing.T) {
	t.Run("SetRequests_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponse{}
		var fernTestValueRequests []*GetConversationRequestsResponseRequestsItem

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

func TestSettersGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("SetRequestID", func(t *testing.T) {
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInputMessage", func(t *testing.T) {
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueInputMessage map[string]any
		obj.SetInputMessage(fernTestValueInputMessage)
		assert.Equal(t, fernTestValueInputMessage, obj.InputMessage)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueStatus GetConversationRequestsResponseRequestsItemStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueTags []*Tag
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEvents", func(t *testing.T) {
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueEvents []*MessageEventData
		obj.SetEvents(fernTestValueEvents)
		assert.Equal(t, fernTestValueEvents, obj.Events)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("GetRequestID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var expected string
		obj.RequestID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestID(), "getter should return the property value")
	})

	t.Run("GetRequestID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestID() // Should return zero value
	})

	t.Run("GetInputMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var expected map[string]any
		obj.InputMessage = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInputMessage(), "getter should return the property value")
	})

	t.Run("GetInputMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		obj.InputMessage = nil

		// Act & Assert
		assert.Nil(t, obj.GetInputMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetInputMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInputMessage() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var expected GetConversationRequestsResponseRequestsItemStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var expected []*Tag
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

	t.Run("GetEvents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var expected []*MessageEventData
		obj.Events = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvents(), "getter should return the property value")
	})

	t.Run("GetEvents_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		obj.Events = nil

		// Act & Assert
		assert.Nil(t, obj.GetEvents(), "getter should return nil when property is nil")
	})

	t.Run("GetEvents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvents() // Should return zero value
	})

}

func TestSettersMarkExplicitGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("SetRequestID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
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

	t.Run("SetInputMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueInputMessage map[string]any

		// Act
		obj.SetInputMessage(fernTestValueInputMessage)

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
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueStatus GetConversationRequestsResponseRequestsItemStatus

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

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
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

	t.Run("SetEvents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}
		var fernTestValueEvents []*MessageEventData

		// Act
		obj.SetEvents(fernTestValueEvents)

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

func TestJSONMarshalingAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachAssetsToConversationResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AttachAssetsToConversationResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AttachAssetsToConversationResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AttachAssetsToConversationResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAttachmentInfo(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AttachmentInfo{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AttachmentInfo
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AttachmentInfo
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AttachmentInfo
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
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

func TestJSONMarshalingConversationCreateOptions(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ConversationCreateOptions{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ConversationCreateOptions
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ConversationCreateOptions
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ConversationCreateOptions
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateConversationResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateConversationResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateRequestResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateRequestResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetConversationRequestsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetConversationRequestsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetConversationRequestsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetConversationRequestsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetConversationRequestsResponseRequestsItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetConversationRequestsResponseRequestsItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetConversationRequestsResponseRequestsItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetConversationRequestsResponseRequestsItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetTaskResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetTaskResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetTaskResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetTaskResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetTaskResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingIntentsListItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled IntentsListItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj IntentsListItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj IntentsListItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingIntentsListItemResponsesItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &IntentsListItemResponsesItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled IntentsListItemResponsesItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj IntentsListItemResponsesItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj IntentsListItemResponsesItem
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

func TestJSONMarshalingRequestCreateOptions(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestCreateOptions{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RequestCreateOptions
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RequestCreateOptions
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RequestCreateOptions
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingServerErrorEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ServerErrorEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ServerErrorEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ServerErrorEventDetails
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

func TestStringAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AttachAssetsToConversationResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AttachAssetsToConversationResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAttachmentInfo(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AttachmentInfo{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AttachmentInfo
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
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

func TestStringConversationCreateOptions(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationCreateOptions{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ConversationCreateOptions
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateConversationResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateRequestResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetConversationRequestsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetConversationRequestsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetConversationRequestsResponseRequestsItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetConversationRequestsResponseRequestsItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetTaskResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetTaskResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetTaskResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringIntentsListItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &IntentsListItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringIntentsListItemResponsesItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &IntentsListItemResponsesItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *IntentsListItemResponsesItem
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

func TestStringRequestCreateOptions(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RequestCreateOptions{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestCreateOptions
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringServerErrorEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ServerErrorEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ServerErrorEventDetails
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

func TestEnumGetConversationRequestsResponseRequestsItemStatus(t *testing.T) {
	t.Run("NewFromString_active", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetConversationRequestsResponseRequestsItemStatusFromString("active")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetConversationRequestsResponseRequestsItemStatus("active"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_cancelled", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetConversationRequestsResponseRequestsItemStatusFromString("cancelled")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetConversationRequestsResponseRequestsItemStatus("cancelled"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_complete", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetConversationRequestsResponseRequestsItemStatusFromString("complete")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetConversationRequestsResponseRequestsItemStatus("complete"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetConversationRequestsResponseRequestsItemStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetConversationRequestsResponseRequestsItemStatusFromString("active")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumServerErrorEventDetailsType(t *testing.T) {
	t.Run("NewFromString_server_error", func(t *testing.T) {
		t.Parallel()
		val, err := NewServerErrorEventDetailsTypeFromString("server_error")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ServerErrorEventDetailsType("server_error"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewServerErrorEventDetailsTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewServerErrorEventDetailsTypeFromString("server_error")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesAttachAssetsToConversationResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AttachAssetsToConversationResponse{}
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
		var obj *AttachAssetsToConversationResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAttachmentInfo(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AttachmentInfo{}
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
		var obj *AttachmentInfo
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
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

func TestExtraPropertiesConversationCreateOptions(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ConversationCreateOptions{}
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
		var obj *ConversationCreateOptions
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateConversationResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationResponse{}
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
		var obj *CreateConversationResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateRequestResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestResponse{}
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
		var obj *CreateRequestResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetConversationRequestsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetConversationRequestsResponse{}
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
		var obj *GetConversationRequestsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetConversationRequestsResponseRequestsItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetConversationRequestsResponseRequestsItem{}
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
		var obj *GetConversationRequestsResponseRequestsItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetTaskResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetTaskResponse{}
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
		var obj *GetTaskResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesIntentsListItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &IntentsListItem{}
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
		var obj *IntentsListItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesIntentsListItemResponsesItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &IntentsListItemResponsesItem{}
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
		var obj *IntentsListItemResponsesItem
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

func TestExtraPropertiesRequestCreateOptions(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RequestCreateOptions{}
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
		var obj *RequestCreateOptions
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesServerErrorEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ServerErrorEventDetails{}
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
		var obj *ServerErrorEventDetails
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
