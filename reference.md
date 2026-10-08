# Reference
## Conversations
<details><summary><code>client.Conversations.CreateEmptyConversation(request) -> *partiteai.CreateConversationResponse</code></summary>
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
    MeshID: "mesh_id",
}
client.Conversations.CreateEmptyConversation(
    context.TODO(),
    request,
)
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

**meshID:** `string` — The ID of the Mesh for which to create a conversation
    
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

<dl>
<dd>

**tags:** `[]*partiteai.Tag` — Tags to apply to the conversation.
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` — Optional client-chosen external ID for the new conversation: 1 to 255 bytes. Immutable and unique (per organization). Use `GET /conversations/lookup/external-id/{external_id}` to find the conversation again.
    
</dd>
</dl>

<dl>
<dd>

**notificationChannel:** `*string` — Optional notification channel for the new conversation: 1 to 255 bytes. Every request output the conversation writes is also appended to this channel, readable via `GET /notification-channels/{notification_channel_id}/events`. The channel is created implicitly on first use and cannot be changed later.
    
</dd>
</dl>

<dl>
<dd>

**tracingEnabled:** `*bool` — Optional. `false` turns off tracing for the new conversation, so the mesh records no trace for any of its requests. Defaults to `true`. Cannot be changed later.
    
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
    MeshID: "mesh_id",
}
client.Conversations.GetIntentsForNewConversation(
    context.TODO(),
    request,
)
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

**meshID:** `string` — The ID of the Mesh for which to get intents
    
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
    IntentName: "intent_name",
    MeshID: "mesh_id",
    InputMessage: map[string]any{
        "key": "value",
    },
}
client.Conversations.CreateConversationWithInitialRequest(
    context.TODO(),
    request,
)
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

**meshID:** `string` — The ID of the Mesh for which to create a conversation
    
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

<details><summary><code>client.Conversations.GetConversation(ConversationID) -> *partiteai.Conversation</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the conversation.
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
request := &partiteai.GetConversationRequest{
    ConversationID: "conversation_id",
}
client.Conversations.GetConversation(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetIntentsForConversation(ConversationID) -> *partiteai.ListIntentsResponse</code></summary>
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
request := &partiteai.GetIntentsForConversationRequest{
    ConversationID: "conversation_id",
}
client.Conversations.GetIntentsForConversation(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CreateConversationRequest(ConversationID, IntentName, request) -> *partiteai.CreateRequestResponse</code></summary>
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
    ConversationID: "conversation_id",
    IntentName: "intent_name",
    InputMessage: map[string]any{
        "key": "value",
    },
}
client.Conversations.CreateConversationRequest(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
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

<details><summary><code>client.Conversations.GetConversationRequests(ConversationID) -> *partiteai.GetConversationRequestsResponse</code></summary>
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
request := &partiteai.GetConversationRequestsRequest{
    ConversationID: "conversation_id",
}
client.Conversations.GetConversationRequests(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CancelRequest(ConversationID, RequestID) -> error</code></summary>
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
request := &partiteai.CancelRequestRequest{
    ConversationID: "conversation_id",
    RequestID: "request_id",
}
client.Conversations.CancelRequest(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CompleteTask(ConversationID, RequestID, TaskID, request) -> error</code></summary>
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
    ConversationID: "conversation_id",
    RequestID: "request_id",
    TaskID: "task_id",
    Body: map[string]any{
        "key": "value",
    },
}
client.Conversations.CompleteTask(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of an existing request
    
</dd>
</dl>

<dl>
<dd>

**taskID:** `string` — The ID of the task to complete
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` — An object conforming to the task's `output_schema`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestEvents(ConversationID, RequestID) -> partiteai.ResponseEventData</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the results of a request, as a stream of events. This allows for real-time updates and progress tracking. The Accept header selects the format: `text/event-stream` streams as SSE, and `application/x-ndjson` or `application/jsonl` streams as JSON lines. With no Accept header, or `* /*`, the format is `text/event-stream`. An Accept header that allows none of these returns 406.
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
    ConversationID: "conversation_id",
    RequestID: "request_id",
}
client.Conversations.GetRequestEvents(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of an existing request
    
</dd>
</dl>

<dl>
<dd>

**lastEventID:** `*string` — The ID of the last event received. If provided, the stream will start from the next event after this ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestResponse(ConversationID, RequestID) -> map[string]any</code></summary>
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
request := &partiteai.GetRequestResponseRequest{
    ConversationID: "conversation_id",
    RequestID: "request_id",
}
client.Conversations.GetRequestResponse(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequestTask(ConversationID, RequestID) -> *partiteai.GetTaskResponse</code></summary>
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
request := &partiteai.GetRequestTaskRequest{
    ConversationID: "conversation_id",
    RequestID: "request_id",
}
client.Conversations.GetRequestTask(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of an existing request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.AttachAssetsToConversation(ConversationID, request) -> *partiteai.AttachAssetsToConversationResponse</code></summary>
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
request := &partiteai.AttachAssetsToConversationRequest{
    ConversationID: "conversation_id",
}
client.Conversations.AttachAssetsToConversation(
    context.TODO(),
    request,
)
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

**conversationID:** `string` — The ID of an existing conversation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetConversationByExternalID(ExternalID) -> *partiteai.Conversation</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the conversation that was created with the given `external_id`.
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
request := &partiteai.GetConversationByExternalIDRequest{
    ExternalID: "external_id",
}
client.Conversations.GetConversationByExternalID(
    context.TODO(),
    request,
)
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

**externalID:** `string` — The external ID supplied when the conversation was created.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notification Channels
<details><summary><code>client.NotificationChannels.GetNotificationChannelEvents(NotificationChannelID) -> partiteai.NotificationChannelStreamData</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reads the event queue of a notification channel. Every request output of every conversation created with this channel appears here, across conversations, in a single resumable sequence. The Accept header selects the mode: `text/event-stream` streams as SSE, `application/x-ndjson` or `application/jsonl` streams as JSON lines, and `application/json` returns one page and closes. With no Accept header, or `* /*`, the mode is `text/event-stream`. An Accept header that allows none of these returns 406. A channel that has not been created returns 404. A created channel with no events reads as empty; streams wait for the first event.
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
request := &partiteai.GetNotificationChannelEventsRequest{
    NotificationChannelID: "notification_channel_id",
}
client.NotificationChannels.GetNotificationChannelEvents(
    context.TODO(),
    request,
)
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

**notificationChannelID:** `string` — The notification channel id supplied at conversation creation.
    
</dd>
</dl>

<dl>
<dd>

**lastEventID:** `*string` — An event id (`rnco_…`) from a previous event on this channel. If provided, reading starts after that event. Event ids are ordered and unique per channel, not globally; an id is only meaningful for the channel it came from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NotificationChannels.CreateNotificationChannel(NotificationChannelID) -> *partiteai.NotificationChannelCreated</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.CreateNotificationChannelRequest{
    NotificationChannelID: "notification_channel_id",
}
client.NotificationChannels.CreateNotificationChannel(
    context.TODO(),
    request,
)
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

**notificationChannelID:** `string` — The client-supplied id of the channel to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

