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

func TestSettersApprovalTool(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &ApprovalTool{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSource", func(t *testing.T) {
		obj := &ApprovalTool{}
		var fernTestValueSource *ApprovalToolSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMemory", func(t *testing.T) {
		obj := &ApprovalTool{}
		var fernTestValueMemory *ApprovalToolMemory
		obj.SetMemory(fernTestValueMemory)
		assert.Equal(t, fernTestValueMemory, obj.Memory)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersApprovalTool(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalTool
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		var expected *ApprovalToolSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalTool
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetMemory", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		var expected *ApprovalToolMemory
		obj.Memory = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMemory(), "getter should return the property value")
	})

	t.Run("GetMemory_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		obj.Memory = nil

		// Act & Assert
		assert.Nil(t, obj.GetMemory(), "getter should return nil when property is nil")
	})

	t.Run("GetMemory_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalTool
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMemory() // Should return zero value
	})

}

func TestSettersMarkExplicitApprovalTool(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
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

	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		var fernTestValueSource *ApprovalToolSource

		// Act
		obj.SetSource(fernTestValueSource)

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

	t.Run("SetMemory_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}
		var fernTestValueMemory *ApprovalToolMemory

		// Act
		obj.SetMemory(fernTestValueMemory)

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

func TestSettersApprovalToolMemory(t *testing.T) {
	t.Run("SetOperation", func(t *testing.T) {
		obj := &ApprovalToolMemory{}
		var fernTestValueOperation ApprovalToolMemoryOperation
		obj.SetOperation(fernTestValueOperation)
		assert.Equal(t, fernTestValueOperation, obj.Operation)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersApprovalToolMemory(t *testing.T) {
	t.Run("GetOperation", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolMemory{}
		var expected ApprovalToolMemoryOperation
		obj.Operation = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOperation(), "getter should return the property value")
	})

	t.Run("GetOperation_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolMemory
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOperation() // Should return zero value
	})

}

func TestSettersMarkExplicitApprovalToolMemory(t *testing.T) {
	t.Run("SetOperation_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolMemory{}
		var fernTestValueOperation ApprovalToolMemoryOperation

		// Act
		obj.SetOperation(fernTestValueOperation)

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

func TestSettersApprovalToolSource(t *testing.T) {
	t.Run("SetKind", func(t *testing.T) {
		obj := &ApprovalToolSource{}
		var fernTestValueKind ApprovalToolSourceKind
		obj.SetKind(fernTestValueKind)
		assert.Equal(t, fernTestValueKind, obj.Kind)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &ApprovalToolSource{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &ApprovalToolSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersApprovalToolSource(t *testing.T) {
	t.Run("GetKind", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		var expected ApprovalToolSourceKind
		obj.Kind = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetKind(), "getter should return the property value")
	})

	t.Run("GetKind_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetKind() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitApprovalToolSource(t *testing.T) {
	t.Run("SetKind_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
		var fernTestValueKind ApprovalToolSourceKind

		// Act
		obj.SetKind(fernTestValueKind)

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
		obj := &ApprovalToolSource{}
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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}
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

}

func TestSettersBoundaryProblem(t *testing.T) {
	t.Run("SetType", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueType string
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueTitle string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueStatus int
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDetail", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueDetail string
		obj.SetDetail(fernTestValueDetail)
		assert.Equal(t, fernTestValueDetail, obj.Detail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetParam", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueParam *string
		obj.SetParam(fernTestValueParam)
		assert.Equal(t, fernTestValueParam, obj.Param)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDecisionID", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueDecisionID string
		obj.SetDecisionID(fernTestValueDecisionID)
		assert.Equal(t, fernTestValueDecisionID, obj.DecisionID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetBoundary", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueBoundary *string
		obj.SetBoundary(fernTestValueBoundary)
		assert.Equal(t, fernTestValueBoundary, obj.Boundary)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRule", func(t *testing.T) {
		obj := &BoundaryProblem{}
		var fernTestValueRule *string
		obj.SetRule(fernTestValueRule)
		assert.Equal(t, fernTestValueRule, obj.Rule)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBoundaryProblem(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetTitle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected string
		obj.Title = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTitle(), "getter should return the property value")
	})

	t.Run("GetTitle_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTitle() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected int
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetDetail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected string
		obj.Detail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDetail(), "getter should return the property value")
	})

	t.Run("GetDetail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDetail() // Should return zero value
	})

	t.Run("GetParam", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected *string
		obj.Param = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetParam(), "getter should return the property value")
	})

	t.Run("GetParam_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		obj.Param = nil

		// Act & Assert
		assert.Nil(t, obj.GetParam(), "getter should return nil when property is nil")
	})

	t.Run("GetParam_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetParam() // Should return zero value
	})

	t.Run("GetDecisionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected string
		obj.DecisionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDecisionID(), "getter should return the property value")
	})

	t.Run("GetDecisionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDecisionID() // Should return zero value
	})

	t.Run("GetBoundary", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected *string
		obj.Boundary = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBoundary(), "getter should return the property value")
	})

	t.Run("GetBoundary_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		obj.Boundary = nil

		// Act & Assert
		assert.Nil(t, obj.GetBoundary(), "getter should return nil when property is nil")
	})

	t.Run("GetBoundary_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBoundary() // Should return zero value
	})

	t.Run("GetRule", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var expected *string
		obj.Rule = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRule(), "getter should return the property value")
	})

	t.Run("GetRule_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		obj.Rule = nil

		// Act & Assert
		assert.Nil(t, obj.GetRule(), "getter should return nil when property is nil")
	})

	t.Run("GetRule_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRule() // Should return zero value
	})

}

func TestSettersMarkExplicitBoundaryProblem(t *testing.T) {
	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueType string

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueTitle string

		// Act
		obj.SetTitle(fernTestValueTitle)

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
		obj := &BoundaryProblem{}
		var fernTestValueStatus int

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

	t.Run("SetDetail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueDetail string

		// Act
		obj.SetDetail(fernTestValueDetail)

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

	t.Run("SetParam_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueParam *string

		// Act
		obj.SetParam(fernTestValueParam)

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

	t.Run("SetDecisionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueDecisionID string

		// Act
		obj.SetDecisionID(fernTestValueDecisionID)

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

	t.Run("SetBoundary_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueBoundary *string

		// Act
		obj.SetBoundary(fernTestValueBoundary)

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

	t.Run("SetRule_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}
		var fernTestValueRule *string

		// Act
		obj.SetRule(fernTestValueRule)

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

func TestSettersCancelledEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &CancelledEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &CancelledEvent{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCancelledEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelledEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelledEvent
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
		obj := &CancelledEvent{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelledEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitCancelledEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelledEvent{}
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
		obj := &CancelledEvent{}
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

func TestSettersClientTask(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &ClientTask{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInput", func(t *testing.T) {
		obj := &ClientTask{}
		var fernTestValueInput map[string]any
		obj.SetInput(fernTestValueInput)
		assert.Equal(t, fernTestValueInput, obj.Input)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPresentations", func(t *testing.T) {
		obj := &ClientTask{}
		var fernTestValuePresentations []*Presentation
		obj.SetPresentations(fernTestValuePresentations)
		assert.Equal(t, fernTestValuePresentations, obj.Presentations)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersClientTask(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetInput", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
		var expected map[string]any
		obj.Input = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInput(), "getter should return the property value")
	})

	t.Run("GetInput_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
		obj.Input = nil

		// Act & Assert
		assert.Nil(t, obj.GetInput(), "getter should return nil when property is nil")
	})

	t.Run("GetInput_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInput() // Should return zero value
	})

	t.Run("GetPresentations", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
		var expected []*Presentation
		obj.Presentations = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPresentations(), "getter should return the property value")
	})

	t.Run("GetPresentations_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
		obj.Presentations = nil

		// Act & Assert
		assert.Nil(t, obj.GetPresentations(), "getter should return nil when property is nil")
	})

	t.Run("GetPresentations_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTask
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPresentations() // Should return zero value
	})

}

func TestSettersMarkExplicitClientTask(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
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

	t.Run("SetInput_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
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

	t.Run("SetPresentations_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}
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

func TestSettersClientTaskData(t *testing.T) {
	t.Run("SetResponseSchema", func(t *testing.T) {
		obj := &ClientTaskData{}
		var fernTestValueResponseSchema map[string]any
		obj.SetResponseSchema(fernTestValueResponseSchema)
		assert.Equal(t, fernTestValueResponseSchema, obj.ResponseSchema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetClientTask", func(t *testing.T) {
		obj := &ClientTaskData{}
		var fernTestValueClientTask *ClientTask
		obj.SetClientTask(fernTestValueClientTask)
		assert.Equal(t, fernTestValueClientTask, obj.ClientTask)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersClientTaskData(t *testing.T) {
	t.Run("GetResponseSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTaskData{}
		var expected map[string]any
		obj.ResponseSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponseSchema(), "getter should return the property value")
	})

	t.Run("GetResponseSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTaskData{}
		obj.ResponseSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponseSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetResponseSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTaskData
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
		obj := &ClientTaskData{}
		var expected *ClientTask
		obj.ClientTask = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClientTask(), "getter should return the property value")
	})

	t.Run("GetClientTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTaskData{}
		obj.ClientTask = nil

		// Act & Assert
		assert.Nil(t, obj.GetClientTask(), "getter should return nil when property is nil")
	})

	t.Run("GetClientTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClientTask() // Should return zero value
	})

}

func TestSettersMarkExplicitClientTaskData(t *testing.T) {
	t.Run("SetResponseSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTaskData{}
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
		obj := &ClientTaskData{}
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

}

func TestSettersErrorEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ErrorEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ErrorEvent{}
		var fernTestValueData string
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
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
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
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
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
		obj := &ErrorEvent{}
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
		var fernTestValueData string
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
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
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
		var fernTestValueData string
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
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
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

func TestSettersNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("SetData", func(t *testing.T) {
		obj := &NotificationChannelServerErrorEvent{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelServerErrorEvent{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelServerErrorEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("SetData_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelServerErrorEvent{}
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

	t.Run("GetServerError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		var expected *NotificationChannelServerErrorEvent
		obj.ServerError = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetServerError(), "getter should return the property value")
	})

	t.Run("GetServerError_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelSseEvent{}
		obj.ServerError = nil

		// Act & Assert
		assert.Nil(t, obj.GetServerError(), "getter should return nil when property is nil")
	})

	t.Run("GetServerError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelSseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetServerError() // Should return zero value
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

func TestSettersProblem(t *testing.T) {
	t.Run("SetType", func(t *testing.T) {
		obj := &Problem{}
		var fernTestValueType string
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &Problem{}
		var fernTestValueTitle string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &Problem{}
		var fernTestValueStatus int
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDetail", func(t *testing.T) {
		obj := &Problem{}
		var fernTestValueDetail string
		obj.SetDetail(fernTestValueDetail)
		assert.Equal(t, fernTestValueDetail, obj.Detail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetParam", func(t *testing.T) {
		obj := &Problem{}
		var fernTestValueParam *string
		obj.SetParam(fernTestValueParam)
		assert.Equal(t, fernTestValueParam, obj.Param)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersProblem(t *testing.T) {
	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var expected string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetTitle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var expected string
		obj.Title = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTitle(), "getter should return the property value")
	})

	t.Run("GetTitle_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTitle() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var expected int
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetDetail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var expected string
		obj.Detail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDetail(), "getter should return the property value")
	})

	t.Run("GetDetail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDetail() // Should return zero value
	})

	t.Run("GetParam", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var expected *string
		obj.Param = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetParam(), "getter should return the property value")
	})

	t.Run("GetParam_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		obj.Param = nil

		// Act & Assert
		assert.Nil(t, obj.GetParam(), "getter should return nil when property is nil")
	})

	t.Run("GetParam_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetParam() // Should return zero value
	})

}

func TestSettersMarkExplicitProblem(t *testing.T) {
	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var fernTestValueType string

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var fernTestValueTitle string

		// Act
		obj.SetTitle(fernTestValueTitle)

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
		obj := &Problem{}
		var fernTestValueStatus int

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

	t.Run("SetDetail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var fernTestValueDetail string

		// Act
		obj.SetDetail(fernTestValueDetail)

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

	t.Run("SetParam_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}
		var fernTestValueParam *string

		// Act
		obj.SetParam(fernTestValueParam)

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

func TestSettersRequestServerSentEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &RequestServerSentEvent{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEvent", func(t *testing.T) {
		obj := &RequestServerSentEvent{}
		var fernTestValueEvent RequestServerSentEventEvent
		obj.SetEvent(fernTestValueEvent)
		assert.Equal(t, fernTestValueEvent, obj.Event)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &RequestServerSentEvent{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersRequestServerSentEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestServerSentEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetEvent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
		var expected RequestServerSentEventEvent
		obj.Event = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEvent(), "getter should return the property value")
	})

	t.Run("GetEvent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestServerSentEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEvent() // Should return zero value
	})

	t.Run("GetData", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestServerSentEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitRequestServerSentEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
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

	t.Run("SetEvent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}
		var fernTestValueEvent RequestServerSentEventEvent

		// Act
		obj.SetEvent(fernTestValueEvent)

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
		obj := &RequestServerSentEvent{}
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

func TestSettersResponseEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ResponseEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ResponseEvent{}
		var fernTestValueData map[string]any
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersResponseEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEvent
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
		obj := &ResponseEvent{}
		var expected map[string]any
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitResponseEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEvent{}
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
		obj := &ResponseEvent{}
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

func TestSettersServerErrorEvent(t *testing.T) {
	t.Run("SetError", func(t *testing.T) {
		obj := &ServerErrorEvent{}
		var fernTestValueError string
		obj.SetError(fernTestValueError)
		assert.Equal(t, fernTestValueError, obj.Error)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersServerErrorEvent(t *testing.T) {
	t.Run("GetError", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEvent{}
		var expected string
		obj.Error = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetError(), "getter should return the property value")
	})

	t.Run("GetError_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ServerErrorEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetError() // Should return zero value
	})

}

func TestSettersMarkExplicitServerErrorEvent(t *testing.T) {
	t.Run("SetError_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEvent{}
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

func TestGettersTaskData(t *testing.T) {
	t.Run("GetKind", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		var expected string
		obj.Kind = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetKind(), "getter should return the property value")
	})

	t.Run("GetKind_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
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
		obj := &TaskData{}
		var expected *ClientTaskData
		obj.ClientTask = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClientTask(), "getter should return the property value")
	})

	t.Run("GetClientTask_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		obj.ClientTask = nil

		// Act & Assert
		assert.Nil(t, obj.GetClientTask(), "getter should return nil when property is nil")
	})

	t.Run("GetClientTask_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
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
		obj := &TaskData{}
		var expected *ToolCallApprovalData
		obj.ToolCallApproval = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetToolCallApproval(), "getter should return the property value")
	})

	t.Run("GetToolCallApproval_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskData{}
		obj.ToolCallApproval = nil

		// Act & Assert
		assert.Nil(t, obj.GetToolCallApproval(), "getter should return nil when property is nil")
	})

	t.Run("GetToolCallApproval_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetToolCallApproval() // Should return zero value
	})

}

func TestSettersTaskEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &TaskEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &TaskEvent{}
		var fernTestValueData *TaskData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTaskEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEvent
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
		obj := &TaskEvent{}
		var expected *TaskData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitTaskEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEvent{}
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
		obj := &TaskEvent{}
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

func TestSettersTaskResponseEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &TaskResponseEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &TaskResponseEvent{}
		var fernTestValueData *TaskResponseEventData
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInReplyTo", func(t *testing.T) {
		obj := &TaskResponseEvent{}
		var fernTestValueInReplyTo string
		obj.SetInReplyTo(fernTestValueInReplyTo)
		assert.Equal(t, fernTestValueInReplyTo, obj.InReplyTo)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersTaskResponseEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEvent
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
		obj := &TaskResponseEvent{}
		var expected *TaskResponseEventData
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEvent{}
		obj.Data = nil

		// Act & Assert
		assert.Nil(t, obj.GetData(), "getter should return nil when property is nil")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEvent
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
		obj := &TaskResponseEvent{}
		var expected string
		obj.InReplyTo = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInReplyTo(), "getter should return the property value")
	})

	t.Run("GetInReplyTo_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInReplyTo() // Should return zero value
	})

}

func TestSettersMarkExplicitTaskResponseEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEvent{}
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
		obj := &TaskResponseEvent{}
		var fernTestValueData *TaskResponseEventData

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
		obj := &TaskResponseEvent{}
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

func TestGettersTaskResponseEventData(t *testing.T) {
	t.Run("GetStringUnknownMap", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventData{}
		var expected map[string]any
		obj.StringUnknownMap = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStringUnknownMap(), "getter should return the property value")
	})

	t.Run("GetStringUnknownMap_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventData{}
		obj.StringUnknownMap = nil

		// Act & Assert
		assert.Nil(t, obj.GetStringUnknownMap(), "getter should return nil when property is nil")
	})

	t.Run("GetStringUnknownMap_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStringUnknownMap() // Should return zero value
	})

	t.Run("GetTaskResponseEventDataOne", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEventData{}
		var expected TaskResponseEventDataOne
		obj.TaskResponseEventDataOne = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTaskResponseEventDataOne(), "getter should return the property value")
	})

	t.Run("GetTaskResponseEventDataOne_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEventData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTaskResponseEventDataOne() // Should return zero value
	})

}

func TestSettersThinkingEvent(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &ThinkingEvent{}
		var fernTestValueID string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetData", func(t *testing.T) {
		obj := &ThinkingEvent{}
		var fernTestValueData string
		obj.SetData(fernTestValueData)
		assert.Equal(t, fernTestValueData, obj.Data)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersThinkingEvent(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEvent{}
		var expected string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEvent
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
		obj := &ThinkingEvent{}
		var expected string
		obj.Data = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetData(), "getter should return the property value")
	})

	t.Run("GetData_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEvent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetData() // Should return zero value
	})

}

func TestSettersMarkExplicitThinkingEvent(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEvent{}
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
		obj := &ThinkingEvent{}
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

func TestSettersToolCallApproval(t *testing.T) {
	t.Run("SetTool", func(t *testing.T) {
		obj := &ToolCallApproval{}
		var fernTestValueTool *ApprovalTool
		obj.SetTool(fernTestValueTool)
		assert.Equal(t, fernTestValueTool, obj.Tool)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetArguments", func(t *testing.T) {
		obj := &ToolCallApproval{}
		var fernTestValueArguments map[string]any
		obj.SetArguments(fernTestValueArguments)
		assert.Equal(t, fernTestValueArguments, obj.Arguments)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMessage", func(t *testing.T) {
		obj := &ToolCallApproval{}
		var fernTestValueMessage *string
		obj.SetMessage(fernTestValueMessage)
		assert.Equal(t, fernTestValueMessage, obj.Message)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDecisionID", func(t *testing.T) {
		obj := &ToolCallApproval{}
		var fernTestValueDecisionID string
		obj.SetDecisionID(fernTestValueDecisionID)
		assert.Equal(t, fernTestValueDecisionID, obj.DecisionID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersToolCallApproval(t *testing.T) {
	t.Run("GetTool", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var expected *ApprovalTool
		obj.Tool = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTool(), "getter should return the property value")
	})

	t.Run("GetTool_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		obj.Tool = nil

		// Act & Assert
		assert.Nil(t, obj.GetTool(), "getter should return nil when property is nil")
	})

	t.Run("GetTool_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTool() // Should return zero value
	})

	t.Run("GetArguments", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var expected map[string]any
		obj.Arguments = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetArguments(), "getter should return the property value")
	})

	t.Run("GetArguments_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		obj.Arguments = nil

		// Act & Assert
		assert.Nil(t, obj.GetArguments(), "getter should return nil when property is nil")
	})

	t.Run("GetArguments_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetArguments() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var expected *string
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

	t.Run("GetDecisionID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var expected string
		obj.DecisionID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDecisionID(), "getter should return the property value")
	})

	t.Run("GetDecisionID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApproval
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDecisionID() // Should return zero value
	})

}

func TestSettersMarkExplicitToolCallApproval(t *testing.T) {
	t.Run("SetTool_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var fernTestValueTool *ApprovalTool

		// Act
		obj.SetTool(fernTestValueTool)

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

	t.Run("SetArguments_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var fernTestValueArguments map[string]any

		// Act
		obj.SetArguments(fernTestValueArguments)

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

	t.Run("SetMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var fernTestValueMessage *string

		// Act
		obj.SetMessage(fernTestValueMessage)

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

	t.Run("SetDecisionID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}
		var fernTestValueDecisionID string

		// Act
		obj.SetDecisionID(fernTestValueDecisionID)

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

func TestSettersToolCallApprovalData(t *testing.T) {
	t.Run("SetResponseSchema", func(t *testing.T) {
		obj := &ToolCallApprovalData{}
		var fernTestValueResponseSchema map[string]any
		obj.SetResponseSchema(fernTestValueResponseSchema)
		assert.Equal(t, fernTestValueResponseSchema, obj.ResponseSchema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetToolCallApproval", func(t *testing.T) {
		obj := &ToolCallApprovalData{}
		var fernTestValueToolCallApproval *ToolCallApproval
		obj.SetToolCallApproval(fernTestValueToolCallApproval)
		assert.Equal(t, fernTestValueToolCallApproval, obj.ToolCallApproval)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersToolCallApprovalData(t *testing.T) {
	t.Run("GetResponseSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalData{}
		var expected map[string]any
		obj.ResponseSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetResponseSchema(), "getter should return the property value")
	})

	t.Run("GetResponseSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalData{}
		obj.ResponseSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetResponseSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetResponseSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalData
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
		obj := &ToolCallApprovalData{}
		var expected *ToolCallApproval
		obj.ToolCallApproval = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetToolCallApproval(), "getter should return the property value")
	})

	t.Run("GetToolCallApproval_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalData{}
		obj.ToolCallApproval = nil

		// Act & Assert
		assert.Nil(t, obj.GetToolCallApproval(), "getter should return nil when property is nil")
	})

	t.Run("GetToolCallApproval_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalData
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetToolCallApproval() // Should return zero value
	})

}

func TestSettersMarkExplicitToolCallApprovalData(t *testing.T) {
	t.Run("SetResponseSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalData{}
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
		obj := &ToolCallApprovalData{}
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

}

func TestSettersToolCallApprovalResponse(t *testing.T) {
	t.Run("SetApproved", func(t *testing.T) {
		obj := &ToolCallApprovalResponse{}
		var fernTestValueApproved bool
		obj.SetApproved(fernTestValueApproved)
		assert.Equal(t, fernTestValueApproved, obj.Approved)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetComment", func(t *testing.T) {
		obj := &ToolCallApprovalResponse{}
		var fernTestValueComment *string
		obj.SetComment(fernTestValueComment)
		assert.Equal(t, fernTestValueComment, obj.Comment)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersToolCallApprovalResponse(t *testing.T) {
	t.Run("GetApproved", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}
		var expected bool
		obj.Approved = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetApproved(), "getter should return the property value")
	})

	t.Run("GetApproved_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetApproved() // Should return zero value
	})

	t.Run("GetComment", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}
		var expected *string
		obj.Comment = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetComment(), "getter should return the property value")
	})

	t.Run("GetComment_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}
		obj.Comment = nil

		// Act & Assert
		assert.Nil(t, obj.GetComment(), "getter should return nil when property is nil")
	})

	t.Run("GetComment_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetComment() // Should return zero value
	})

}

func TestSettersMarkExplicitToolCallApprovalResponse(t *testing.T) {
	t.Run("SetApproved_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}
		var fernTestValueApproved bool

		// Act
		obj.SetApproved(fernTestValueApproved)

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

	t.Run("SetComment_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}
		var fernTestValueComment *string

		// Act
		obj.SetComment(fernTestValueComment)

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

func TestJSONMarshalingApprovalTool(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalTool{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ApprovalTool
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalTool
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalTool
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingApprovalToolMemory(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolMemory{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ApprovalToolMemory
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalToolMemory
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalToolMemory
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingApprovalToolSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ApprovalToolSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ApprovalToolSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalToolSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ApprovalToolSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingBoundaryProblem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BoundaryProblem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BoundaryProblem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BoundaryProblem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BoundaryProblem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCancelledEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CancelledEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CancelledEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CancelledEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CancelledEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingClientTask(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTask{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ClientTask
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ClientTask
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ClientTask
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingClientTaskData(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ClientTaskData{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ClientTaskData
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ClientTaskData
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ClientTaskData
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

func TestJSONMarshalingNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &NotificationChannelServerErrorEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled NotificationChannelServerErrorEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelServerErrorEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj NotificationChannelServerErrorEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingProblem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &Problem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled Problem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj Problem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj Problem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingRequestServerSentEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &RequestServerSentEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled RequestServerSentEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj RequestServerSentEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj RequestServerSentEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingResponseEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ResponseEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ResponseEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ResponseEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ResponseEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingServerErrorEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ServerErrorEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ServerErrorEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ServerErrorEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ServerErrorEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTaskEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TaskEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TaskEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TaskEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingTaskResponseEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &TaskResponseEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled TaskResponseEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj TaskResponseEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj TaskResponseEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingThinkingEvent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ThinkingEvent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ThinkingEvent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ThinkingEvent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ThinkingEvent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingToolCallApproval(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApproval{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ToolCallApproval
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApproval
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApproval
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingToolCallApprovalData(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalData{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ToolCallApprovalData
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApprovalData
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApprovalData
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingToolCallApprovalResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ToolCallApprovalResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ToolCallApprovalResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApprovalResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ToolCallApprovalResponse
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

func TestStringApprovalTool(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalTool{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalTool
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringApprovalToolMemory(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalToolMemory{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolMemory
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringApprovalToolSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalToolSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ApprovalToolSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringBoundaryProblem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BoundaryProblem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BoundaryProblem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCancelledEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CancelledEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CancelledEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringClientTask(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ClientTask{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTask
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringClientTaskData(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ClientTaskData{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ClientTaskData
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

func TestStringNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelServerErrorEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *NotificationChannelServerErrorEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringProblem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &Problem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *Problem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringRequestServerSentEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &RequestServerSentEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *RequestServerSentEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringResponseEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ResponseEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ResponseEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringServerErrorEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ServerErrorEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ServerErrorEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTaskEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TaskEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringTaskResponseEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &TaskResponseEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *TaskResponseEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringThinkingEvent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ThinkingEvent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ThinkingEvent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringToolCallApproval(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApproval{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApproval
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringToolCallApprovalData(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApprovalData{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalData
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringToolCallApprovalResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApprovalResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ToolCallApprovalResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumApprovalToolMemoryOperation(t *testing.T) {
	t.Run("NewFromString_query", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolMemoryOperationFromString("query")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolMemoryOperation("query"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_write", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolMemoryOperationFromString("write")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolMemoryOperation("write"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_expand", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolMemoryOperationFromString("expand")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolMemoryOperation("expand"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewApprovalToolMemoryOperationFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewApprovalToolMemoryOperationFromString("query")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumApprovalToolSourceKind(t *testing.T) {
	t.Run("NewFromString_mcp", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("mcp")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("mcp"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_particle", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("particle")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("particle"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_client_task", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("client_task")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("client_task"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_resource_reader", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("resource_reader")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("resource_reader"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_skill", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("skill")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("skill"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_memory", func(t *testing.T) {
		t.Parallel()
		val, err := NewApprovalToolSourceKindFromString("memory")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ApprovalToolSourceKind("memory"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewApprovalToolSourceKindFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewApprovalToolSourceKindFromString("mcp")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumRequestServerSentEventEvent(t *testing.T) {
	t.Run("NewFromString_started", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("started")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("started"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_task", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("task")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("task"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_task_response", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("task_response")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("task_response"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_message", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("message")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("message"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_thinking", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("thinking")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("thinking"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_response", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("response")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("response"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_request_error", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("request_error")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("request_error"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_cancelled", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("cancelled")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("cancelled"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_server_error", func(t *testing.T) {
		t.Parallel()
		val, err := NewRequestServerSentEventEventFromString("server_error")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, RequestServerSentEventEvent("server_error"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewRequestServerSentEventEventFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewRequestServerSentEventEventFromString("started")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumTaskResponseEventDataOne(t *testing.T) {
	t.Run("NewFromString_cancelled", func(t *testing.T) {
		t.Parallel()
		val, err := NewTaskResponseEventDataOneFromString("cancelled")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, TaskResponseEventDataOne("cancelled"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewTaskResponseEventDataOneFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewTaskResponseEventDataOneFromString("cancelled")
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

func TestExtraPropertiesApprovalTool(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalTool{}
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
		var obj *ApprovalTool
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesApprovalToolMemory(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalToolMemory{}
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
		var obj *ApprovalToolMemory
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesApprovalToolSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ApprovalToolSource{}
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
		var obj *ApprovalToolSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesBoundaryProblem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BoundaryProblem{}
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
		var obj *BoundaryProblem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCancelledEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CancelledEvent{}
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
		var obj *CancelledEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesClientTask(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ClientTask{}
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
		var obj *ClientTask
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesClientTaskData(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ClientTaskData{}
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
		var obj *ClientTaskData
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

func TestExtraPropertiesNotificationChannelServerErrorEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &NotificationChannelServerErrorEvent{}
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
		var obj *NotificationChannelServerErrorEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesProblem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &Problem{}
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
		var obj *Problem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesRequestServerSentEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &RequestServerSentEvent{}
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
		var obj *RequestServerSentEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesResponseEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ResponseEvent{}
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
		var obj *ResponseEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesServerErrorEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ServerErrorEvent{}
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
		var obj *ServerErrorEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTaskEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TaskEvent{}
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
		var obj *TaskEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesTaskResponseEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &TaskResponseEvent{}
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
		var obj *TaskResponseEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesThinkingEvent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ThinkingEvent{}
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
		var obj *ThinkingEvent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesToolCallApproval(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApproval{}
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
		var obj *ToolCallApproval
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesToolCallApprovalData(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApprovalData{}
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
		var obj *ToolCallApprovalData
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesToolCallApprovalResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ToolCallApprovalResponse{}
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
		var obj *ToolCallApprovalResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
