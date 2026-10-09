// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersA2UIPresentation(t *testing.T) {
	t.Run("SetContent", func(t *testing.T) {
		obj := &A2UIPresentation{}
		var fernTestValueContent *A2UISurfaces
		obj.SetContent(fernTestValueContent)
		assert.Equal(t, fernTestValueContent, obj.Content)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersA2UIPresentation(t *testing.T) {
	t.Run("GetContent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UIPresentation{}
		var expected *A2UISurfaces
		obj.Content = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContent(), "getter should return the property value")
	})

	t.Run("GetContent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UIPresentation{}
		obj.Content = nil

		// Act & Assert
		assert.Nil(t, obj.GetContent(), "getter should return nil when property is nil")
	})

	t.Run("GetContent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UIPresentation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContent() // Should return zero value
	})

}

func TestSettersMarkExplicitA2UIPresentation(t *testing.T) {
	t.Run("SetContent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UIPresentation{}
		var fernTestValueContent *A2UISurfaces

		// Act
		obj.SetContent(fernTestValueContent)

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

func TestSettersA2UISurfaceComponents(t *testing.T) {
	t.Run("SetComponents", func(t *testing.T) {
		obj := &A2UISurfaceComponents{}
		var fernTestValueComponents []map[string]any
		obj.SetComponents(fernTestValueComponents)
		assert.Equal(t, fernTestValueComponents, obj.Components)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersA2UISurfaceComponents(t *testing.T) {
	t.Run("GetComponents", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaceComponents{}
		var expected []map[string]any
		obj.Components = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetComponents(), "getter should return the property value")
	})

	t.Run("GetComponents_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaceComponents{}
		obj.Components = nil

		// Act & Assert
		assert.Nil(t, obj.GetComponents(), "getter should return nil when property is nil")
	})

	t.Run("GetComponents_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaceComponents
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetComponents() // Should return zero value
	})

}

func TestSettersMarkExplicitA2UISurfaceComponents(t *testing.T) {
	t.Run("SetComponents_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaceComponents{}
		var fernTestValueComponents []map[string]any

		// Act
		obj.SetComponents(fernTestValueComponents)

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

func TestSettersA2UISurfaces(t *testing.T) {
	t.Run("SetCatalogID", func(t *testing.T) {
		obj := &A2UISurfaces{}
		var fernTestValueCatalogID string
		obj.SetCatalogID(fernTestValueCatalogID)
		assert.Equal(t, fernTestValueCatalogID, obj.CatalogID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOpen", func(t *testing.T) {
		obj := &A2UISurfaces{}
		var fernTestValueOpen *A2UISurfaceComponents
		obj.SetOpen(fernTestValueOpen)
		assert.Equal(t, fernTestValueOpen, obj.Open)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCompleted", func(t *testing.T) {
		obj := &A2UISurfaces{}
		var fernTestValueCompleted *A2UISurfaceComponents
		obj.SetCompleted(fernTestValueCompleted)
		assert.Equal(t, fernTestValueCompleted, obj.Completed)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersA2UISurfaces(t *testing.T) {
	t.Run("GetCatalogID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var expected string
		obj.CatalogID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCatalogID(), "getter should return the property value")
	})

	t.Run("GetCatalogID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaces
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCatalogID() // Should return zero value
	})

	t.Run("GetOpen", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var expected *A2UISurfaceComponents
		obj.Open = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOpen(), "getter should return the property value")
	})

	t.Run("GetOpen_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		obj.Open = nil

		// Act & Assert
		assert.Nil(t, obj.GetOpen(), "getter should return nil when property is nil")
	})

	t.Run("GetOpen_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaces
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOpen() // Should return zero value
	})

	t.Run("GetCompleted", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var expected *A2UISurfaceComponents
		obj.Completed = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCompleted(), "getter should return the property value")
	})

	t.Run("GetCompleted_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		obj.Completed = nil

		// Act & Assert
		assert.Nil(t, obj.GetCompleted(), "getter should return nil when property is nil")
	})

	t.Run("GetCompleted_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaces
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCompleted() // Should return zero value
	})

}

func TestSettersMarkExplicitA2UISurfaces(t *testing.T) {
	t.Run("SetCatalogID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var fernTestValueCatalogID string

		// Act
		obj.SetCatalogID(fernTestValueCatalogID)

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

	t.Run("SetOpen_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var fernTestValueOpen *A2UISurfaceComponents

		// Act
		obj.SetOpen(fernTestValueOpen)

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

	t.Run("SetCompleted_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}
		var fernTestValueCompleted *A2UISurfaceComponents

		// Act
		obj.SetCompleted(fernTestValueCompleted)

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

func TestSettersBadRequestErrorResponse(t *testing.T) {
	t.Run("SetError", func(t *testing.T) {
		obj := &BadRequestErrorResponse{}
		var fernTestValueError string
		obj.SetError(fernTestValueError)
		assert.Equal(t, fernTestValueError, obj.Error)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetParameter", func(t *testing.T) {
		obj := &BadRequestErrorResponse{}
		var fernTestValueParameter *string
		obj.SetParameter(fernTestValueParameter)
		assert.Equal(t, fernTestValueParameter, obj.Parameter)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBadRequestErrorResponse(t *testing.T) {
	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}
		var expected string
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BadRequestErrorResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

	t.Run("GetParameter", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}
		var expected *string
		obj.Parameter = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetParameter(), "getter should return the property value")
	})

	t.Run("GetParameter_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}
		obj.Parameter = nil

		// Act & Assert
		assert.Nil(t, obj.GetParameter(), "getter should return nil when property is nil")
	})

	t.Run("GetParameter_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BadRequestErrorResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetParameter() // Should return zero value
	})

}

func TestSettersMarkExplicitBadRequestErrorResponse(t *testing.T) {
	t.Run("SetError_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}
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

	t.Run("SetParameter_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}
		var fernTestValueParameter *string

		// Act
		obj.SetParameter(fernTestValueParameter)

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

func TestSettersCancelEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &CancelEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &CancelEventDetails{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCancelEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelEventDetails{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitCancelEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelEventDetails{}
		var fernTestValueData string

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

func TestSettersCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("SetInputMessage", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueInputMessage map[string]any
		obj.SetInputMessage(fernTestValueInputMessage)
		assert.Equal(t, fernTestValueInputMessage, obj.InputMessage)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversationOptions", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueConversationOptions *ConversationCreateOptions
		obj.SetConversationOptions(fernTestValueConversationOptions)
		assert.Equal(t, fernTestValueConversationOptions, obj.ConversationOptions)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestOptions", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueRequestOptions *RequestCreateOptions
		obj.SetRequestOptions(fernTestValueRequestOptions)
		assert.Equal(t, fernTestValueRequestOptions, obj.RequestOptions)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAttachments", func(t *testing.T) {
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueAttachments []string
		obj.SetAttachments(fernTestValueAttachments)
		assert.Equal(t, fernTestValueAttachments, obj.Attachments)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("GetInputMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var expected map[string]any
		obj.InputMessage = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInputMessage(), "getter should return the property value")
	})

	t.Run("GetInputMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		obj.InputMessage = nil

		// Act & Assert
		assert.Nil(t, obj.GetInputMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetInputMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationWithInitialRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInputMessage() // Should return zero value
	})

	t.Run("GetConversationOptions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var expected *ConversationCreateOptions
		obj.ConversationOptions = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversationOptions(), "getter should return the property value")
	})

	t.Run("GetConversationOptions_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		obj.ConversationOptions = nil

		// Act & Assert
		assert.Nil(t, obj.GetConversationOptions(), "getter should return nil when property is nil")
	})

	t.Run("GetConversationOptions_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationWithInitialRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversationOptions() // Should return zero value
	})

	t.Run("GetRequestOptions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var expected *RequestCreateOptions
		obj.RequestOptions = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestOptions(), "getter should return the property value")
	})

	t.Run("GetRequestOptions_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		obj.RequestOptions = nil

		// Act & Assert
		assert.Nil(t, obj.GetRequestOptions(), "getter should return nil when property is nil")
	})

	t.Run("GetRequestOptions_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationWithInitialRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestOptions() // Should return zero value
	})

	t.Run("GetAttachments", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var expected []string
		obj.Attachments = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAttachments(), "getter should return the property value")
	})

	t.Run("GetAttachments_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		obj.Attachments = nil

		// Act & Assert
		assert.Nil(t, obj.GetAttachments(), "getter should return nil when property is nil")
	})

	t.Run("GetAttachments_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationWithInitialRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAttachments() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("SetInputMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
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

	t.Run("SetConversationOptions_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueConversationOptions *ConversationCreateOptions

		// Act
		obj.SetConversationOptions(fernTestValueConversationOptions)

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

	t.Run("SetRequestOptions_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueRequestOptions *RequestCreateOptions

		// Act
		obj.SetRequestOptions(fernTestValueRequestOptions)

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

	t.Run("SetAttachments_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}
		var fernTestValueAttachments []string

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

func TestSettersCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("SetInputMessage", func(t *testing.T) {
		obj := &CreateRequestWithAttachmentsRequest{}
		var fernTestValueInputMessage map[string]any
		obj.SetInputMessage(fernTestValueInputMessage)
		assert.Equal(t, fernTestValueInputMessage, obj.InputMessage)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRequestOptions", func(t *testing.T) {
		obj := &CreateRequestWithAttachmentsRequest{}
		var fernTestValueRequestOptions *RequestCreateOptions
		obj.SetRequestOptions(fernTestValueRequestOptions)
		assert.Equal(t, fernTestValueRequestOptions, obj.RequestOptions)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAttachments", func(t *testing.T) {
		obj := &CreateRequestWithAttachmentsRequest{}
		var fernTestValueAttachments []string
		obj.SetAttachments(fernTestValueAttachments)
		assert.Equal(t, fernTestValueAttachments, obj.Attachments)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("GetInputMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		var expected map[string]any
		obj.InputMessage = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInputMessage(), "getter should return the property value")
	})

	t.Run("GetInputMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		obj.InputMessage = nil

		// Act & Assert
		assert.Nil(t, obj.GetInputMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetInputMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestWithAttachmentsRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInputMessage() // Should return zero value
	})

	t.Run("GetRequestOptions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		var expected *RequestCreateOptions
		obj.RequestOptions = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRequestOptions(), "getter should return the property value")
	})

	t.Run("GetRequestOptions_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		obj.RequestOptions = nil

		// Act & Assert
		assert.Nil(t, obj.GetRequestOptions(), "getter should return nil when property is nil")
	})

	t.Run("GetRequestOptions_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestWithAttachmentsRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRequestOptions() // Should return zero value
	})

	t.Run("GetAttachments", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		var expected []string
		obj.Attachments = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAttachments(), "getter should return the property value")
	})

	t.Run("GetAttachments_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		obj.Attachments = nil

		// Act & Assert
		assert.Nil(t, obj.GetAttachments(), "getter should return nil when property is nil")
	})

	t.Run("GetAttachments_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestWithAttachmentsRequest
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAttachments() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("SetInputMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
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

	t.Run("SetRequestOptions_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		var fernTestValueRequestOptions *RequestCreateOptions

		// Act
		obj.SetRequestOptions(fernTestValueRequestOptions)

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

	t.Run("SetAttachments_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}
		var fernTestValueAttachments []string

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

func TestSettersErrorEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ErrorEvent{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ErrorEvent{}
		var fernTestValueData *ServerErrorEventDetails
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersErrorEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		var expected *ServerErrorEventDetails
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitErrorEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		var fernTestValueID *string

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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}
		var fernTestValueData *ServerErrorEventDetails

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

func TestSettersErrorEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ErrorEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ErrorEventDetails{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersErrorEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEventDetails{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitErrorEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEventDetails{}
		var fernTestValueData string

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

func TestSettersGenericErrorResponse(t *testing.T) {
	t.Run("SetError", func(t *testing.T) {
		obj := &GenericErrorResponse{}
		var fernTestValueError string
		obj.SetError(fernTestValueError)
		assert.Equal(t, fernTestValueError, obj.Error)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGenericErrorResponse(t *testing.T) {
	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GenericErrorResponse{}
		var expected string
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GenericErrorResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

}

func TestSettersMarkExplicitGenericErrorResponse(t *testing.T) {
	t.Run("SetError_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GenericErrorResponse{}
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

func TestSettersMessageEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &MessageEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &MessageEvent{}
		var fernTestValueData *MessageEventData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersMessageEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}
		var expected *MessageEventData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitMessageEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}
		var fernTestValueData *MessageEventData

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

func TestGettersMessageEventData(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *TaskEventDetails
		obj.Task = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTask(), "getter should return the property value")
	})

	t.Run("GetTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Task = nil

		// Act & Assert
		assert.Nil(t, obj.GetTask(), "getter should return nil when property is nil")
	})

	t.Run("GetTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *TaskResponseEventDetails
		obj.TaskResponse = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskResponse(), "getter should return the property value")
	})

	t.Run("GetTaskResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.TaskResponse = nil

		// Act & Assert
		assert.Nil(t, obj.GetTaskResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetTaskResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *MessageEventDetails
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *ThinkingEventDetails
		obj.Thinking = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetThinking(), "getter should return the property value")
	})

	t.Run("GetThinking_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Thinking = nil

		// Act & Assert
		assert.Nil(t, obj.GetThinking(), "getter should return nil when property is nil")
	})

	t.Run("GetThinking_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *ResponseEventDetails
		obj.Response = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponse(), "getter should return the property value")
	})

	t.Run("GetResponse_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Response = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponse(), "getter should return nil when property is nil")
	})

	t.Run("GetResponse_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *ErrorEventDetails
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
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
		obj := &MessageEventData{}
		var expected *CancelEventDetails
		obj.Cancelled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCancelled(), "getter should return the property value")
	})

	t.Run("GetCancelled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventData{}
		obj.Cancelled = nil

		// Act & Assert
		assert.Nil(t, obj.GetCancelled(), "getter should return nil when property is nil")
	})

	t.Run("GetCancelled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCancelled() // Should return zero value
	})

}

func TestSettersMessageEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &MessageEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &MessageEventDetails{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersMessageEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventDetails{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitMessageEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventDetails{}
		var fernTestValueData string

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

func TestSettersNotificationChannelMessageEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &NotificationChannelMessageEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &NotificationChannelMessageEvent{}
		var fernTestValueData *NotificationChannelEvent
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotificationChannelMessageEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelMessageEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}
		var expected *NotificationChannelEvent
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelMessageEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitNotificationChannelMessageEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}
		var fernTestValueData *NotificationChannelEvent

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

func TestGettersNotificationChannelSseEvent(t *testing.T) {
	t.Run("GetEvent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		var expected string
		obj.Event = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvent(), "getter should return the property value")
	})

	t.Run("GetEvent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelSseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvent() // Should return zero value
	})

	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		var expected *ErrorEvent
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelSseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		var expected *NotificationChannelMessageEvent
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelSseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

}

func TestGettersPresentation(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Presentation{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Presentation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetA2UI", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Presentation{}
		var expected *A2UIPresentation
		obj.A2UI = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetA2UI(), "getter should return the property value")
	})

	t.Run("GetA2UI_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Presentation{}
		obj.A2UI = nil

		// Act & Assert
		assert.Nil(t, obj.GetA2UI(), "getter should return nil when property is nil")
	})

	t.Run("GetA2UI_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Presentation
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetA2UI() // Should return zero value
	})

}

func TestSettersResponseEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ResponseEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ResponseEventDetails{}
		var fernTestValueData map[string]any
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersResponseEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}
		var expected map[string]any
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitResponseEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}
		var fernTestValueData map[string]any

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

func TestGettersSseEvent(t *testing.T) {
	t.Run("GetEvent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SseEvent{}
		var expected string
		obj.Event = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvent(), "getter should return the property value")
	})

	t.Run("GetEvent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvent() // Should return zero value
	})

	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SseEvent{}
		var expected *ErrorEvent
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SseEvent{}
		obj.Error = nil

		// Act & Assert
		assert.Nil(t, obj.GetError(), "getter should return nil when property is nil")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SseEvent{}
		var expected *MessageEvent
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &SseEvent{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *SseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

}

func TestSettersTaskData(t *testing.T) {
	t.Run("SetInput", func(t *testing.T) {
		obj := &TaskData{}
		var fernTestValueInput map[string]any
		obj.SetInput(fernTestValueInput)
		assert.Equal(t, fernTestValueInput, obj.Input)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTaskType", func(t *testing.T) {
		obj := &TaskData{}
		var fernTestValueTaskType string
		obj.SetTaskType(fernTestValueTaskType)
		assert.Equal(t, fernTestValueTaskType, obj.TaskType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOutputSchema", func(t *testing.T) {
		obj := &TaskData{}
		var fernTestValueOutputSchema map[string]any
		obj.SetOutputSchema(fernTestValueOutputSchema)
		assert.Equal(t, fernTestValueOutputSchema, obj.OutputSchema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPresentations", func(t *testing.T) {
		obj := &TaskData{}
		var fernTestValuePresentations []*Presentation
		obj.SetPresentations(fernTestValuePresentations)
		assert.Equal(t, fernTestValuePresentations, obj.Presentations)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTaskData(t *testing.T) {
	t.Run("GetInput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var expected map[string]any
		obj.Input = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInput(), "getter should return the property value")
	})

	t.Run("GetInput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		obj.Input = nil

		// Act & Assert
		assert.Nil(t, obj.GetInput(), "getter should return nil when property is nil")
	})

	t.Run("GetInput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInput() // Should return zero value
	})

	t.Run("GetTaskType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var expected string
		obj.TaskType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskType(), "getter should return the property value")
	})

	t.Run("GetTaskType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTaskType() // Should return zero value
	})

	t.Run("GetOutputSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var expected map[string]any
		obj.OutputSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOutputSchema(), "getter should return the property value")
	})

	t.Run("GetOutputSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		obj.OutputSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetOutputSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetOutputSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOutputSchema() // Should return zero value
	})

	t.Run("GetPresentations", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var expected []*Presentation
		obj.Presentations = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPresentations(), "getter should return the property value")
	})

	t.Run("GetPresentations_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		obj.Presentations = nil

		// Act & Assert
		assert.Nil(t, obj.GetPresentations(), "getter should return nil when property is nil")
	})

	t.Run("GetPresentations_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPresentations() // Should return zero value
	})

}

func TestSettersMarkExplicitTaskData(t *testing.T) {
	t.Run("SetInput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
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

	t.Run("SetTaskType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var fernTestValueTaskType string

		// Act
		obj.SetTaskType(fernTestValueTaskType)

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

	t.Run("SetOutputSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var fernTestValueOutputSchema map[string]any

		// Act
		obj.SetOutputSchema(fernTestValueOutputSchema)

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

	t.Run("SetPresentations_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var fernTestValuePresentations []*Presentation

		// Act
		obj.SetPresentations(fernTestValuePresentations)

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

func TestSettersTaskEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &TaskEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &TaskEventDetails{}
		var fernTestValueData *TaskData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTaskEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}
		var expected *TaskData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitTaskEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}
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

func TestSettersTaskResponseEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &TaskResponseEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &TaskResponseEventDetails{}
		var fernTestValueData *TaskResponseEventDetailsData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInReplyTo", func(t *testing.T) {
		obj := &TaskResponseEventDetails{}
		var fernTestValueInReplyTo string
		obj.SetInReplyTo(fernTestValueInReplyTo)
		assert.Equal(t, fernTestValueInReplyTo, obj.InReplyTo)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTaskResponseEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		var expected *TaskResponseEventDetailsData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

	t.Run("GetInReplyTo", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		var expected string
		obj.InReplyTo = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInReplyTo(), "getter should return the property value")
	})

	t.Run("GetInReplyTo_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInReplyTo() // Should return zero value
	})

}

func TestSettersMarkExplicitTaskResponseEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		var fernTestValueData *TaskResponseEventDetailsData

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

	t.Run("SetInReplyTo_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}
		var fernTestValueInReplyTo string

		// Act
		obj.SetInReplyTo(fernTestValueInReplyTo)

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

func TestGettersTaskResponseEventDetailsData(t *testing.T) {
	t.Run("GetStringUnknownMap", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetailsData{}
		var expected map[string]any
		obj.StringUnknownMap = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStringUnknownMap(), "getter should return the property value")
	})

	t.Run("GetStringUnknownMap_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetailsData{}
		obj.StringUnknownMap = nil

		// Act & Assert
		assert.Nil(t, obj.GetStringUnknownMap(), "getter should return nil when property is nil")
	})

	t.Run("GetStringUnknownMap_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetailsData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStringUnknownMap() // Should return zero value
	})

	t.Run("GetTaskResponseEventDetailsDataOne", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetailsData{}
		var expected TaskResponseEventDetailsDataOne
		obj.TaskResponseEventDetailsDataOne = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskResponseEventDetailsDataOne(), "getter should return the property value")
	})

	t.Run("GetTaskResponseEventDetailsDataOne_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetailsData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTaskResponseEventDetailsDataOne() // Should return zero value
	})

}

func TestSettersThinkingEventDetails(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ThinkingEventDetails{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ThinkingEventDetails{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersThinkingEventDetails(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEventDetails{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEventDetails{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEventDetails
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitThinkingEventDetails(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEventDetails{}
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

	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEventDetails{}
		var fernTestValueData string

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

func TestJSONMarshalingA2UIPresentation(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UIPresentation{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled A2UIPresentation
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj A2UIPresentation
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj A2UIPresentation
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingA2UISurfaceComponents(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaceComponents{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled A2UISurfaceComponents
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj A2UISurfaceComponents
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj A2UISurfaceComponents
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingA2UISurfaces(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &A2UISurfaces{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled A2UISurfaces
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj A2UISurfaces
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj A2UISurfaces
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingBadRequestErrorResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BadRequestErrorResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BadRequestErrorResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BadRequestErrorResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BadRequestErrorResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCancelEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CancelEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CancelEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CancelEventDetails
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateConversationWithInitialRequest{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateConversationWithInitialRequest
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationWithInitialRequest
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateConversationWithInitialRequest
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateRequestWithAttachmentsRequest{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateRequestWithAttachmentsRequest
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestWithAttachmentsRequest
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateRequestWithAttachmentsRequest
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingErrorEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ErrorEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ErrorEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ErrorEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingErrorEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ErrorEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ErrorEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ErrorEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ErrorEventDetails
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGenericErrorResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GenericErrorResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GenericErrorResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GenericErrorResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GenericErrorResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingMessageEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled MessageEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj MessageEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj MessageEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingMessageEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &MessageEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled MessageEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj MessageEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj MessageEventDetails
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

func TestJSONMarshalingNotificationChannelMessageEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelMessageEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotificationChannelMessageEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelMessageEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelMessageEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingResponseEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ResponseEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ResponseEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ResponseEventDetails
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

func TestJSONMarshalingTaskData(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TaskData
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TaskData
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TaskData
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTaskEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TaskEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TaskEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TaskEventDetails
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTaskResponseEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TaskResponseEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TaskResponseEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TaskResponseEventDetails
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingThinkingEventDetails(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEventDetails{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ThinkingEventDetails
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ThinkingEventDetails
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ThinkingEventDetails
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringA2UIPresentation(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &A2UIPresentation{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UIPresentation
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringA2UISurfaceComponents(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &A2UISurfaceComponents{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaceComponents
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringA2UISurfaces(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &A2UISurfaces{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *A2UISurfaces
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringBadRequestErrorResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BadRequestErrorResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BadRequestErrorResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCancelEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CancelEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelEventDetails
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationWithInitialRequest{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateConversationWithInitialRequest
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestWithAttachmentsRequest{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateRequestWithAttachmentsRequest
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringErrorEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ErrorEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringErrorEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ErrorEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ErrorEventDetails
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGenericErrorResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GenericErrorResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GenericErrorResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringMessageEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &MessageEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringMessageEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &MessageEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *MessageEventDetails
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

func TestStringNotificationChannelMessageEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelMessageEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelMessageEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringResponseEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ResponseEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEventDetails
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

func TestStringTaskData(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TaskData{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTaskEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TaskEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEventDetails
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTaskResponseEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TaskResponseEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventDetails
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringThinkingEventDetails(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ThinkingEventDetails{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEventDetails
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
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

func TestEnumTaskResponseEventDetailsDataOne(t *testing.T) {
	t.Run("NewFromString_cancelled", func(t *testing.T) {
		t.Parallel()
		val, err := NewTaskResponseEventDetailsDataOneFromString("cancelled")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, TaskResponseEventDetailsDataOne("cancelled"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewTaskResponseEventDetailsDataOneFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewTaskResponseEventDetailsDataOneFromString("cancelled")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesA2UIPresentation(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &A2UIPresentation{}
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
		var obj *A2UIPresentation
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesA2UISurfaceComponents(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &A2UISurfaceComponents{}
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
		var obj *A2UISurfaceComponents
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesA2UISurfaces(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &A2UISurfaces{}
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
		var obj *A2UISurfaces
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesBadRequestErrorResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BadRequestErrorResponse{}
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
		var obj *BadRequestErrorResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCancelEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CancelEventDetails{}
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
		var obj *CancelEventDetails
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateConversationWithInitialRequest(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateConversationWithInitialRequest{}
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
		var obj *CreateConversationWithInitialRequest
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateRequestWithAttachmentsRequest(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateRequestWithAttachmentsRequest{}
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
		var obj *CreateRequestWithAttachmentsRequest
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesErrorEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ErrorEvent{}
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
		var obj *ErrorEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesErrorEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ErrorEventDetails{}
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
		var obj *ErrorEventDetails
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGenericErrorResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GenericErrorResponse{}
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
		var obj *GenericErrorResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesMessageEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &MessageEvent{}
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
		var obj *MessageEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesMessageEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &MessageEventDetails{}
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
		var obj *MessageEventDetails
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

func TestExtraPropertiesNotificationChannelMessageEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelMessageEvent{}
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
		var obj *NotificationChannelMessageEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesResponseEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ResponseEventDetails{}
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
		var obj *ResponseEventDetails
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

func TestExtraPropertiesTaskData(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TaskData{}
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
		var obj *TaskData
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTaskEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TaskEventDetails{}
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
		var obj *TaskEventDetails
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTaskResponseEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TaskResponseEventDetails{}
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
		var obj *TaskResponseEventDetails
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesThinkingEventDetails(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ThinkingEventDetails{}
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
		var obj *ThinkingEventDetails
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
