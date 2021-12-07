# RESTful API

## 标准方法
| 标准方法 | HTTP 映射                   | HTTP 请求正文 | HTTP 响应正文           |
| -------- | --------------------------- | ------------- | ----------------------- |
| List     | GET <collection URL>        | 无            | 资源*列表               |
| Get      | GET <resource URL>          | 无            | 资源*                   |
| Create   | POST <collection URL>       | 资源          | 资源*                   |
| Update   | PUT or PATCH <resource URL> | 资源          | 资源*                   |
| Delete   | DELETE <resource URL>       | 不适用        | google.protobuf.Empty** |

### 列表
List 方法将一个集合名称和零个或多个参数作为输入，并返回与输入匹配的资源列表。

List 通常用于搜索资源。List 适用于来自单个集合的数据，该集合的大小有限且不进行缓存。

适用的常见模式：分页、结果排序。

HTTP 映射：
- List 方法 必须使用 HTTP GET 动词。
- 接收其资源正在列出的集合名称的请求消息字段应该映射到网址路径。如果集合名称映射到网址路径，则网址模板的最后一段（集合 ID）必须是字面量。
- 所有剩余的请求消息字段应该映射到网址查询参数。
- 响应正文应该包含资源列表以及可选元数据。

示例：
```protobuf
// Lists books in a shelf.
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse) {
  // List method maps to HTTP GET.
  option (google.api.http) = {
    // The `parent` captures the parent resource name, such as "shelves/shelf1".
    get: "/v1/{parent=shelves/*}/books"
  };
}

message ListBooksRequest {
  // The parent resource name, for example, "shelves/shelf1".
  string parent = 1;

  // The maximum number of items to return.
  int32 page_size = 2;

  // The next_page_token value returned from a previous List request, if any.
  string page_token = 3;
}

message ListBooksResponse {
  // The field name should match the noun "books" in the method name.  There
  // will be a maximum number of items returned based on the page_size field
  // in the request.
  repeated Book books = 1;

  // Token to retrieve the next page of results, or empty if there are no
  // more results in the list.
  string next_page_token = 2;
}
```

### 获取
Get 方法需要一个资源名称和零个或多个参数作为输入，并返回指定的资源。

HTTP 映射：
- Get 方法 必须使用 HTTP GET 动词。
- 接收资源名称的请求消息字段应该映射到网址路径。
- 所有剩余的请求消息字段应该映射到网址查询参数。
- 没有请求正文，API 配置不得声明 body 子句。
- 返回的资源应该映射到整个响应正文。

示例：
```protobuf
// Gets a book.
rpc GetBook(GetBookRequest) returns (Book) {
  // Get maps to HTTP GET. Resource name is mapped to the URL. No body.
  option (google.api.http) = {
    // Note the URL template variable which captures the multi-segment resource
    // name of the requested book, such as "shelves/shelf1/books/book2"
    get: "/v1/{name=shelves/*/books/*}"
  };
}

message GetBookRequest {
  // The field will contain name of the resource requested, for example:
  // "shelves/shelf1/books/book2"
  string name = 1;
}
```

### 创建
Create 方法需要一个父资源名称、一个资源以及零个或多个参数作为输入。它在指定的父资源下创建新资源，并返回新建的资源。

如果 API 支持创建资源，则应该为每一个可以创建的资源类型设置 Create 方法。

HTTP 映射：
- Create 方法 必须使用 HTTP POST 动词。
- 请求消息应该具有字段 parent，以指定要在其中创建资源的父资源名称。
- 包含资源的请求消息字段必须映射到请求正文。如果将 google.api.http 注释用于 Create 方法，则必须使用 body: "<resource_field>" 表单。
- 该请求可以包含名为 <resource>_id 的字段，以允许调用者选择客户端分配的 ID。该字段可以在资源内。
- 所有剩余的请求消息字段应该映射到网址查询参数。
- 返回的资源应该映射到整个 HTTP 响应正文。

如果 Create 方法支持客户端分配的资源名称并且资源已存在，则请求应该失败并显示错误代码 ALREADY_EXISTS 或使用服务器分配的不同的资源名称，并且文档应该清楚地记录创建的资源名称可能与传入的不同。

Create 方法必须使用输入资源，以便在资源架构更改时，无需同时更新请求架构和资源架构。对于客户端无法设置的资源字段，必须将它们记录为“仅限输出”字段。

示例：
```protobuf
// Creates a book in a shelf.
rpc CreateBook(CreateBookRequest) returns (Book) {
  // Create maps to HTTP POST. URL path as the collection name.
  // HTTP request body contains the resource.
  option (google.api.http) = {
    // The `parent` captures the parent resource name, such as "shelves/1".
    post: "/v1/{parent=shelves/*}/books"
    body: "book"
  };
}

message CreateBookRequest {
  // The parent resource name where the book is to be created.
  string parent = 1;

  // The book id to use for this book.
  string book_id = 3;

  // The book resource to create.
  // The field name should match the Noun in the method name.
  Book book = 2;
}

rpc CreateShelf(CreateShelfRequest) returns (Shelf) {
  option (google.api.http) = {
    post: "/v1/shelves"
    body: "shelf"
  };
}

message CreateShelfRequest {
  Shelf shelf = 1;
}
```

### 更新
Update 方法需要一条包含一个资源的请求消息和零个或多个参数作为输入。它更新指定的资源及其属性，并返回更新后的资源。

除了包含资源名称或父资源的属性之外，Update 方法应该可以改变可变资源的属性。Update 方法不得包含任何“重命名”或“移动”资源的功能，这些功能应该由自定义方法来处理。

HTTP 映射：
- 标准 Update 方法应该支持部分资源更新，并将 HTTP 动词 PATCH 与名为 update_mask 的 FieldMask字段一起使用。应忽略客户端提供的作为输入的输出字段。
- 需要更高级修补语义的 Update 方法（例如附加到重复字段）应该由自定义方法提供。
- 如果 Update 方法仅支持完整资源更新，则必须使用 HTTP 动词 PUT。但是，强烈建议不要进行完整更新，因为在添加新资源字段时会出现向后兼容性问题。
- 接收资源名称的消息字段必须映射到网址路径。该字段可以位于资源消息本身中。
- 包含资源的请求消息字段必须映射到请求正文。
- 所有剩余的请求消息字段必须映射到网址查询参数。
- 响应消息必须是更新的资源本身。

如果 API 接受客户端分配的资源名称，则服务器可以允许客户端指定不存在的资源名称并创建新资源。 否则，使用不存在的资源名称的 Update 方法应该失败。 如果这是唯一的错误条件，则应该使用错误代码 NOT_FOUND。

具有支持资源创建的 Update 方法的 API 还应该提供 Create 方法。原因是，如果 Update 方法是唯一的方法，则它将不知道如何创建资源。

示例：
```protobuf
// Updates a book.
rpc UpdateBook(UpdateBookRequest) returns (Book) {
  // Update maps to HTTP PATCH. Resource name is mapped to a URL path.
  // Resource is contained in the HTTP request body.
  option (google.api.http) = {
    // Note the URL template variable which captures the resource name of the
    // book to update.
    patch: "/v1/{book.name=shelves/*/books/*}"
    body: "book"
  };
}

message UpdateBookRequest {
  // The book resource which replaces the resource on the server.
  Book book = 1;

  // The update mask applies to the resource. For the `FieldMask` definition,
  // see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
  FieldMask update_mask = 2;
}
```

### 删除
Delete 方法需要一个资源名称和零个或多个参数作为输入，并删除或计划删除指定的资源。Delete 方法应该返回 google.protobuf.Empty。

API 不应该依赖于 Delete 方法返回的任何信息，因为它不能重复调用。

HTTP 映射：
- Delete 方法 必须使用 HTTP DELETE 动词。
- 接收资源名称的请求消息字段应该映射到网址路径。
- 所有剩余的请求消息字段应该映射到网址查询参数。
- 没有请求正文，API 配置不得声明 body 子句。
- 如果 Delete 方法立即移除资源，则应该返回空响应。
- 如果 Delete 方法启动长时间运行的操作，则应该返回长时间运行的操作。
- 如果 Delete 方法仅将资源标记为已删除，则应该返回更新后的资源。

对 Delete 方法的调用在效果上应该是幂等的，但不需要产生相同的响应。任意数量的 Delete 请求都应该导致资源（最终）被删除，但只有第一个请求会产生成功代码。后续请求应生成 google.rpc.Code.NOT_FOUND。

示例：
```protobuf
// Deletes a book.
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  // Delete maps to HTTP DELETE. Resource name maps to the URL path.
  // There is no request body.
  option (google.api.http) = {
    // Note the URL template variable capturing the multi-segment name of the
    // book resource to be deleted, such as "shelves/shelf1/books/book2"
    delete: "/v1/{name=shelves/*/books/*}"
  };
}

message DeleteBookRequest {
  // The resource name of the book to be deleted, for example:
  // "shelves/shelf1/books/book2"
  string name = 1;
}
```