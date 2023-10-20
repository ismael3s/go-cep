# Go-CEP


## What is?

This project is a simple CEP tracker where the application can call N CEP gateways and return the response from the fastest one, if no  gateway responds within specified time the request will fail with timeout error.

## What you will find in the source code:

1. Redis integration
2. Decorator Pattern
3. Concurrent requests with an timeout after 2 seconds
4. Decoupled layers (Basics of `Ports and Adapters` and `Clean Architeture` concepts)
5. Basic testing (Unit and integration)
6. Docker with multistage build
7. Auto update image in hub.docker with github actions

## Diagram of the aplication

```mermaid
sequenceDiagram
    participant User
    participant API
    User->>API: GET /api/v1/:cep
    alt invalid cep
        API->>User: Bad Request
    end
    par API to ViaCEP
        API->>ViaCEP: Get address data from this cep
    and API to BrasilAPI
        API->>BrasilAPI: Get address data from this cep
    and API create timeout
        API->>API: Create a timeout of two seconds
    end
    break when some gateway respond or timeout
        ViaCEP->>API: Return data
        BrasilAPI->>API: Return data
        API->>API: Timeout error
    end
```
