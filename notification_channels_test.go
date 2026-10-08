// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersCreateNotificationChannelRequest(t *testing.T) {
	t.Run("SetNotificationChannelID", func(t *testing.T) {
		obj := &CreateNotificationChannelRequest{}
		var fernTestValueNotificationChannelID string
		obj.SetNotificationChannelID(fernTestValueNotificationChannelID)
		assert.Equal(t, fernTestValueNotificationChannelID, obj.NotificationChannelID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateNotificationChannelRequest(t *testing.T) {
	t.Run("SetNotificationChannelID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateNotificationChannelRequest{}
		var fernTestValueNotificationChannelID string

		// Act
		obj.SetNotificationChannelID(fernTestValueNotificationChannelID)

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

func TestSettersGetNotificationChannelEventsRequest(t *testing.T) {
	t.Run("SetLastEventID", func(t *testing.T) {
		obj := &GetNotificationChannelEventsRequest{}
		var fernTestValueLastEventID *string
		obj.SetLastEventID(fernTestValueLastEventID)
		assert.Equal(t, fernTestValueLastEventID, obj.LastEventID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNotificationChannelID", func(t *testing.T) {
		obj := &GetNotificationChannelEventsRequest{}
		var fernTestValueNotificationChannelID string
		obj.SetNotificationChannelID(fernTestValueNotificationChannelID)
		assert.Equal(t, fernTestValueNotificationChannelID, obj.NotificationChannelID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetNotificationChannelEventsRequest(t *testing.T) {
	t.Run("SetLastEventID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetNotificationChannelEventsRequest{}
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

	t.Run("SetNotificationChannelID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetNotificationChannelEventsRequest{}
		var fernTestValueNotificationChannelID string

		// Act
		obj.SetNotificationChannelID(fernTestValueNotificationChannelID)

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

func TestSettersNotificationChannelCreated(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &NotificationChannelCreated{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotificationChannelCreated(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelCreated{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelCreated
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

}

func TestSettersMarkExplicitNotificationChannelCreated(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelCreated{}
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

func TestSettersNotificationChannelEvent(t *testing.T) {
	t.Run("SetType", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueType NotificationChannelEventType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversationID", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueConversationID string
		obj.SetConversationID(fernTestValueConversationID)
		assert.Equal(t, fernTestValueConversationID, obj.ConversationID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversationExternalID", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueConversationExternalID *string
		obj.SetConversationExternalID(fernTestValueConversationExternalID)
		assert.Equal(t, fernTestValueConversationExternalID, obj.ConversationExternalID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestID", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueRequestID string
		obj.SetRequestID(fernTestValueRequestID)
		assert.Equal(t, fernTestValueRequestID, obj.RequestID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOutput", func(t *testing.T) {
		obj := &NotificationChannelEvent{}
		var fernTestValueOutput *MessageEventData
		obj.SetOutput(fernTestValueOutput)
		assert.Equal(t, fernTestValueOutput, obj.Output)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotificationChannelEvent(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected NotificationChannelEventType
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetConversationID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected string
		obj.ConversationID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationID(), "getter should return the property value")
	})

	t.Run("GetConversationID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationID() // Should return zero value
	})

	t.Run("GetConversationExternalID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected *string
		obj.ConversationExternalID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationExternalID(), "getter should return the property value")
	})

	t.Run("GetConversationExternalID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		obj.ConversationExternalID = nil

		// Act & Assert
		assert.Nil(t, obj.GetConversationExternalID(), "getter should return nil when property is nil")
	})

	t.Run("GetConversationExternalID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationExternalID() // Should return zero value
	})

	t.Run("GetRequestID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected string
		obj.RequestID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestID(), "getter should return the property value")
	})

	t.Run("GetRequestID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestID() // Should return zero value
	})

	t.Run("GetOutput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var expected *MessageEventData
		obj.Output = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOutput(), "getter should return the property value")
	})

	t.Run("GetOutput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		obj.Output = nil

		// Act & Assert
		assert.Nil(t, obj.GetOutput(), "getter should return nil when property is nil")
	})

	t.Run("GetOutput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOutput() // Should return zero value
	})

}

func TestSettersMarkExplicitNotificationChannelEvent(t *testing.T) {
	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var fernTestValueType NotificationChannelEventType

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

	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
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

	t.Run("SetConversationID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
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

	t.Run("SetConversationExternalID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var fernTestValueConversationExternalID *string

		// Act
		obj.SetConversationExternalID(fernTestValueConversationExternalID)

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
		obj := &NotificationChannelEvent{}
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

	t.Run("SetOutput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}
		var fernTestValueOutput *MessageEventData

		// Act
		obj.SetOutput(fernTestValueOutput)

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

func TestGettersNotificationChannelStreamData(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelStreamData{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelStreamData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetNotification", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelStreamData{}
		var expected *NotificationChannelEvent
		obj.Notification = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNotification(), "getter should return the property value")
	})

	t.Run("GetNotification_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelStreamData{}
		obj.Notification = nil

		// Act & Assert
		assert.Nil(t, obj.GetNotification(), "getter should return nil when property is nil")
	})

	t.Run("GetNotification_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelStreamData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNotification() // Should return zero value
	})

	t.Run("GetServerError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelStreamData{}
		var expected *ServerErrorEventDetails
		obj.ServerError = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetServerError(), "getter should return the property value")
	})

	t.Run("GetServerError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelStreamData{}
		obj.ServerError = nil

		// Act & Assert
		assert.Nil(t, obj.GetServerError(), "getter should return nil when property is nil")
	})

	t.Run("GetServerError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelStreamData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetServerError() // Should return zero value
	})

}

func TestJSONMarshalingNotificationChannelCreated(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelCreated{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotificationChannelCreated
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelCreated
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelCreated
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingNotificationChannelEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotificationChannelEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringNotificationChannelCreated(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelCreated{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelCreated
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringNotificationChannelEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumNotificationChannelEventType(t *testing.T) {
	t.Run("NewFromString_notification", func(t *testing.T) {
		t.Parallel()
		val, err := NewNotificationChannelEventTypeFromString("notification")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, NotificationChannelEventType("notification"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewNotificationChannelEventTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewNotificationChannelEventTypeFromString("notification")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesNotificationChannelCreated(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelCreated{}
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
		var obj *NotificationChannelCreated
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesNotificationChannelEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelEvent{}
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
		var obj *NotificationChannelEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
