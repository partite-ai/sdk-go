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

	t.Run("SetLimit", func(t *testing.T) {
		obj := &GetNotificationChannelEventsRequest{}
		var fernTestValueLimit *int
		obj.SetLimit(fernTestValueLimit)
		assert.Equal(t, fernTestValueLimit, obj.Limit)
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

	t.Run("SetLimit_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetNotificationChannelEventsRequest{}
		var fernTestValueLimit *int

		// Act
		obj.SetLimit(fernTestValueLimit)

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

func TestSettersNotificationChannelEventList(t *testing.T) {
	t.Run("SetEvents", func(t *testing.T) {
		obj := &NotificationChannelEventList{}
		var fernTestValueEvents []*NotificationChannelEvent
		obj.SetEvents(fernTestValueEvents)
		assert.Equal(t, fernTestValueEvents, obj.Events)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNext", func(t *testing.T) {
		obj := &NotificationChannelEventList{}
		var fernTestValueNext string
		obj.SetNext(fernTestValueNext)
		assert.Equal(t, fernTestValueNext, obj.Next)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotificationChannelEventList(t *testing.T) {
	t.Run("GetEvents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}
		var expected []*NotificationChannelEvent
		obj.Events = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvents(), "getter should return the property value")
	})

	t.Run("GetEvents_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}
		obj.Events = nil

		// Act & Assert
		assert.Nil(t, obj.GetEvents(), "getter should return nil when property is nil")
	})

	t.Run("GetEvents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEventList
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvents() // Should return zero value
	})

	t.Run("GetNext", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}
		var expected string
		obj.Next = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNext(), "getter should return the property value")
	})

	t.Run("GetNext_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEventList
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNext() // Should return zero value
	})

}

func TestSettersMarkExplicitNotificationChannelEventList(t *testing.T) {
	t.Run("SetEvents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}
		var fernTestValueEvents []*NotificationChannelEvent

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

	t.Run("SetNext_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}
		var fernTestValueNext string

		// Act
		obj.SetNext(fernTestValueNext)

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

func TestJSONMarshalingNotificationChannelEventList(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelEventList{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotificationChannelEventList
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelEventList
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelEventList
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

func TestStringNotificationChannelEventList(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelEventList{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelEventList
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
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

func TestExtraPropertiesNotificationChannelEventList(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelEventList{}
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
		var obj *NotificationChannelEventList
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
