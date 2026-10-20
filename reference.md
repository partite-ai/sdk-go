# Reference
## Conversations
<details><summary><code>client.Conversations.CreateConversation(request) -> *partiteai.CreatedConversation</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts a conversation on a mesh version or a route. With `input`, it also sends the first message, and answers like sending a message does: `wait` and a streaming `Accept` apply to that first request. An external ID already in use is a 409.
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
request := &partiteai.CreateConversationRequest{
    Params: &partiteai.CreateConversationParams{},
}
client.Conversations.CreateConversation(
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

**wait:** `*int` — Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
    
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

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partiteai.GetConversationRequest{
    ConversationID: "extid=thread-42",
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.ListIntents(ConversationID) -> *partiteai.ListIntentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The intents of the mesh version the conversation runs on, with the schema of the input each takes and of the responses each gives. An intent the caller's boundaries don't allow is left out.
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
request := &partiteai.ListIntentsRequest{
    ConversationID: "extid=thread-42",
}
client.Conversations.ListIntents(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.ListRequests(ConversationID) -> *partiteai.ListRequestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Every request of the conversation, oldest first.
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
request := &partiteai.ListRequestsRequest{
    ConversationID: "extid=thread-42",
}
client.Conversations.ListRequests(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CreateRequest(ConversationID, request) -> *partiteai.Request</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends a message to the conversation, which starts a request.
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
request := &partiteai.CreateRequestRequest{
    ConversationID: "extid=thread-42",
    Params: &partiteai.CreateRequestParams{
        Input: map[string]any{
            "key": "value",
        },
    },
}
client.Conversations.CreateRequest(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**wait:** `*int` — Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.GetRequest(ConversationID, RequestID) -> *partiteai.Request</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The request's state. With `?wait=N`, the answer waits until the request leaves `running`, for up to N seconds; a request still running then is reported as `running`.
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
request := &partiteai.GetRequestRequest{
    ConversationID: "extid=thread-42",
    RequestID: "request_id",
}
client.Conversations.GetRequest(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of a request of the conversation.
    
</dd>
</dl>

<dl>
<dd>

**wait:** `*int` — Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.StreamRequestEvents(ConversationID, RequestID) -> partiteai.RequestStreamEvent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The request's events: a `started` event, then every stored event from the first, or from after `Last-Event-ID`, live until the request ends. A stream does not end while the request waits on a task. In SSE, each event's `event:` line is its `type` (`request_error` for an `error`) and its `data:` is the event; in JSON Lines, each line is the event. The `started` event has no ID, so it never changes where a resumed stream starts.
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
request := &partiteai.StreamRequestEventsRequest{
    ConversationID: "extid=thread-42",
    RequestID: "request_id",
}
client.Conversations.StreamRequestEvents(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of a request of the conversation.
    
</dd>
</dl>

<dl>
<dd>

**lastEventID:** `*string` — The ID of the last event received. The stream starts after it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CancelRequest(ConversationID, RequestID) -> *partiteai.Request</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels the request at once: it is `cancelled` in the answer. Work already in flight is dropped at the mesh's next checkpoint, though a tool call already sent may still have its effect. Cancelling a request that has ended answers with its final state.
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
    ConversationID: "extid=thread-42",
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of a request of the conversation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CompleteTask(ConversationID, RequestID, TaskID, request) -> *partiteai.Request</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Answers an open task of the request. The body is the task's response and must match the task's `response_schema`; a response that does not is a 400 and leaves the task open. A `tool_call_approval` is answered with `{"approved": true}` to let the call go ahead, or with `approved` false to decline it, optionally with a `comment` for the agent.
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
    ConversationID: "extid=thread-42",
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of a request of the conversation.
    
</dd>
</dl>

<dl>
<dd>

**taskID:** `string` — The ID of an open task of the request.
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` — An object that matches the task's `response_schema`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.CancelTask(ConversationID, RequestID, TaskID) -> *partiteai.Request</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Declines an open task, and the request goes on. For a `client_task`, the tool call that started it returns an error to the agent. For a `tool_call_approval`, the call is declined, as if it was answered with `approved` false.
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
request := &partiteai.CancelTaskRequest{
    ConversationID: "extid=thread-42",
    RequestID: "request_id",
    TaskID: "task_id",
}
client.Conversations.CancelTask(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` — The ID of a request of the conversation.
    
</dd>
</dl>

<dl>
<dd>

**taskID:** `string` — The ID of an open task of the request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.ListResources(ConversationID) -> *partiteai.ResourceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Everything the conversation holds, in the order it entered, including what MCP servers and other activity added.
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
request := &partiteai.ListResourcesRequest{
    ConversationID: "extid=thread-42",
}
client.Conversations.ListResources(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Conversations.AddResources(ConversationID, request) -> *partiteai.ResourceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Each file part becomes a resource of the conversation, which later inputs can refer to by its `reference`.
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
request := &partiteai.AddResourcesRequest{
    ConversationID: "extid=thread-42",
}
client.Conversations.AddResources(
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

**conversationID:** `string` — A conversation ID, or `extid=` followed by the percent-encoded external ID.
    
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

Reads the event queue of a notification channel. Every request output of every conversation created with this channel appears here, across conversations, in a single resumable sequence. The Accept header selects the mode: `text/event-stream` streams as SSE, `application/jsonl` streams as JSON lines, and `application/json` returns one page and closes. With no Accept header, or `* /*`, the mode is `text/event-stream`. An Accept header that allows none of these returns 406. A channel that has not been created returns 404. A created channel with no events reads as empty; streams wait for the first event.
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

