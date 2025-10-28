# Reference
## Conversations
<details><summary><code>client.Conversations.CreateEmptyConversation() -> *partiteai.CreateConversationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new conversation, without specifying an initial request or assets. If neither the version_ordinal nor version_label query parameters are provided, the draft version of the mesh will be used.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.CreateEmptyConversationRequest{
        MeshId: "mesh_id",
        VersionOrdinal: partiteai.Int(
            1,
        ),
        VersionLabel: partiteai.String(
            "version_label",
        ),
    }
client.Conversations.CreateEmptyConversation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**meshId:** `string` — The ID of the Mesh for which to create a conversation
    
</dd>
</dl>

<dl>
<dd>

**versionOrdinal:** `*int` — The version number of a mesh version for which to create the conversation
    
</dd>
</dl>

<dl>
<dd>

**versionLabel:** `*string` — The version label of a mesh version for which to create the conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetIntentsForNewConversation() -> *partiteai.ListIntentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of intents defined on a particular mesh version for which a new conversation can be started.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.GetIntentsForNewConversationRequest{
        MeshId: "mesh_id",
        VersionOrdinal: partiteai.Int(
            1,
        ),
        VersionLabel: partiteai.String(
            "version_label",
        ),
    }
client.Conversations.GetIntentsForNewConversation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**meshId:** `string` — The ID of the Mesh for which to get intents
    
</dd>
</dl>

<dl>
<dd>

**versionOrdinal:** `*int` — The version number of a mesh version for which to get intents
    
</dd>
</dl>

<dl>
<dd>

**versionLabel:** `*string` — The version label of a mesh version for which to get intents
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CreateConversationWithInitialRequest(IntentName, request) -> *partiteai.CreateRequestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new conversation, specifying an initial request and attachments to start the conversation. If neither the version_ordinal nor version_label query parameters are provided, the draft version of the mesh will be used.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.CreateConversationWithInitialRequestRequest{
        MeshId: "mesh_id",
        VersionOrdinal: partiteai.Int(
            1,
        ),
        VersionLabel: partiteai.String(
            "version_label",
        ),
        InputMessage: map[string]any{
            "key": "value",
        },
    }
client.Conversations.CreateConversationWithInitialRequest(
        context.TODO(),
        "intent_name",
        nil,
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**intentName:** `string` — The name of the Intent defined on the target mesh which will be used to start the conversation
    
</dd>
</dl>

<dl>
<dd>

**meshId:** `string` — The ID of the Mesh for which to create a conversation
    
</dd>
</dl>

<dl>
<dd>

**versionOrdinal:** `*int` — The version number of a mesh version for which to create the conversation
    
</dd>
</dl>

<dl>
<dd>

**versionLabel:** `*string` — The version label of a mesh version for which to create the conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetIntentsForConversation(ConversationId) -> *partiteai.ListIntentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of intents associated with a particular conversation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.GetIntentsForConversation(
        context.TODO(),
        "conversation_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CreateConversationRequest(ConversationId, IntentName, request) -> *partiteai.CreateRequestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a request specified in JSON to an existing conversation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.CreateConversationRequestRequest{
        InputMessage: map[string]any{
            "key": "value",
        },
    }
client.Conversations.CreateConversationRequest(
        context.TODO(),
        "conversation_id",
        "intent_name",
        nil,
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**intentName:** `string` — The name of the Intent defined on the target mesh which will be used for this request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetConversationRequests(ConversationId) -> *partiteai.GetConversationRequestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of requests that have been made within a particular conversation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.GetConversationRequests(
        context.TODO(),
        "conversation_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CancelRequest(ConversationId, RequestId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels execution of a previously submitted request. Processing will be stopped at the next checkpoint, and the request will have a cancellation event added to its output.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.CancelRequest(
        context.TODO(),
        "conversation_id",
        "request_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestId:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CompleteTask(ConversationId, RequestId, TaskId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Used to submit the results of a task event that previously appeared as an output on the request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.CompleteTaskRequest{
        Body: map[string]any{
            "key": "value",
        },
    }
client.Conversations.CompleteTask(
        context.TODO(),
        "conversation_id",
        "request_id",
        "task_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestId:** `string` — The ID of an existing request
    
</dd>
</dl>

<dl>
<dd>

**taskId:** `string` — The ID of the task to complete
    
</dd>
</dl>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestEvents(ConversationId, RequestId) -> partiteai.ResponseEventData</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the results of a request, as a stream of events. This allows for real-time updates and progress tracking.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.GetRequestEventsRequest{
        LastEventId: partiteai.String(
            "Last-Event-ID",
        ),
    }
client.Conversations.GetRequestEvents(
        context.TODO(),
        "conversation_id",
        "request_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestId:** `string` — The ID of an existing request
    
</dd>
</dl>

<dl>
<dd>

**lastEventId:** `*string` — The ID of the last event received. If provided, the stream will start from the next event after this ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestResponse(ConversationId, RequestId) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the final response of a request. This will block until the request is complete.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.GetRequestResponse(
        context.TODO(),
        "conversation_id",
        "request_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestId:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestTask(ConversationId, RequestId) -> *partiteai.GetTaskResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

If the request is currently waiting for a task to be completed, this endpoint will return the details of the task. It will block until a task is available or the request completes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.GetRequestTask(
        context.TODO(),
        "conversation_id",
        "request_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestId:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.AttachAssetsToConversation(ConversationId, request) -> *partiteai.AttachAssetsToConversationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds assets to an existing conversation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Conversations.AttachAssetsToConversation(
        context.TODO(),
        "conversation_id",
        nil,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**conversationId:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
