## First example

```bash
syntax = "proto3"

message Greeting {
    string first_name = 1;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse{
    string result = 1;
}

service GreetService{
    rpc Greet(GreetRequest) return GreetResponse {}
}
```

## Eficiency vs JSON

_JSON: 52 bytes(compressed)_

```js
{
    "age": 26,
    "firstName": "Manuel",
    "lastName": "Escobedo"
}
```

_Same in Protocol Buffers: 17 bytes_

```bash
syntax = "proto3"

messsage Person {
    uint31 age = 1;
    string firstName = 2;
    string lastName = 3;
}
```

- Less CPU intensive
- Less Memmory
- Faster
- Mobile and Microcontrollers
- Support to many languajes

## HTTP/2

**How HTTP 1 works**

- TCP Connection per request
- Plaintext headers
- Request / Response

![HTTP Image](/01%20gRPC%20internals%20deep%20dive/http1.png "asdf")

**How HTTP 1 works**

- one TCP Connection
- server push
- Multiplexing (server can push multiple messages)
- data is compressed to binary
- SSL
- Less chatter
- Less bandwidth

![HTTP Image](/01%20gRPC%20internals%20deep%20dive/http2.png "asdf")

## Types of gRPC api

![HTTP Image](/01%20gRPC%20internals%20deep%20dive/grpc%20types.png "asdf")
