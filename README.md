# Message

Base library to transport transactional message through the process flow. By default, the message converts the data structure to JSON and as well the message body, any other formate needs to be implemented on the input/output plugin, never in this module.

## Usage

Use this library is simple, you just need to create a header and body, then create a message using it. follow there is an example of usage. 

```go
header := Header{}
header.Add("hello", "world")

str := "{\"hello\":\"world\"}"
body := NewBody(str)

msg := New(header, body)
```

## License

Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements.  See the NOTICE file distributed with this work for additional information regarding copyright ownership.  The ASF licenses this file to you under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.