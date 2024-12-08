# Chat Backend General
==========================

**Table of Contents**
-----------------

1. [Overview](#overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Project Structure](#project-structure)
5. [Getting Started](#getting-started)
6. [Configuration](#configuration)
9. [Contributing](#contributing)

## Overview
--------

Chat Backend General is a robust, scalable backend infrastructure designed to support chat applications, incorporating advanced language models (LLMs) and Retrieval-Augmented Generation (RAG) systems. This project adheres to a clean architecture principle, ensuring maintainability, testability, and flexibility.

## Features
--------

* **Websocket Integration**: Establishes a persistent, low-latency connection between clients and the Azure Open AI API.
* **Real-time API Support**: Handles real-time requests from clients, forwarding them to the Azure Open AI API.
* **Kong API Gateway Compatible**: Configured for seamless integration with Kong API Gateway.

## Architecture
------------

* **Clean Architecture Pattern**: Separates concerns into distinct layers for domain logic, use cases, interface adapters, and infrastructure.
* **Layered Structure**:
    + **Domain**: Defines business logic and entities (e.g., LLMs, RAG).
    + **Use Cases**: Specifies interactions with the domain layer (e.g., querying LLMs).
    + **Interface Adapters**: Translates data between the outside world and the use cases (e.g., Websocket, Azure Open AI API client).
    + **Infrastructure**: Provides the underlying technology for the interface adapters (e.g., FastHTTP, PostgreSQL).

## Project Structure
---------------------

The project is organized into the following key directories:
```
chat.backend.general/
├── README.md
├── cmd
│   └── main.go
├── config
│   └── config.go
├── go.mod
├── go.sum
└── internal
    ├── adaptors
    │   ├── http
    │   │   └── file_handlers.go
    │   ├── mq
    │   │   ├── azure_service_bus_adapter.go
    │   │   └── mq_handlers.go
    │   ├── storage
    │   │   └── azure_blob_storage.go
    │   ├── validation
    │   │   ├── file_size_validator.go
    │   │   ├── file_size_validator_test.go
    │   │   ├── file_type_validator.go
    │   │   └── file_type_validator_test.go
    │   └── websocket
    ├── domain
    │   ├── celery_message.go
    │   ├── file
    │   ├── file.go
    │   ├── file_repository.go
    │   ├── file_validator.go
    │   ├── llm
    │   ├── message_queue.go
    │   ├── rag
    │   ├── storage
    │   └── usecase.go
    ├── infra
    │   ├── database
    │   │   └── postgre.go
    │   ├── http
    │   │   └── gin_server.go
    │   ├── storage
    │   │   ├── azure_blob.go
    │   │   └── s3.go
    │   └── websocket
    │       └── wss.go
    ├── llm
    │   └── llm_usecases.go
    └── usecases
        ├── file_upload.go
        ├── file_upload_impl.go
        └── mq
            └── message_queue.go
```

```
+----------------+
|  HTTP Handler  |
+----------------+
         |
         | (UploadedFile)
         v
+-------------------+
| FileUploadUseCase |
|  (Business Logic) |
+-------------------+
         |
         | (FileRepository)
         v
+------------------+
|  FileRepository  |
|  (Interface)     |
+------------------+
         |
         | (Implementation)
         v
+---------------------+
| BlobStorageAdapter  |
|  (Concrete Storage) |
+---------------------+
         |
         | (Validation)
         v
+-----------------+
|  FileValidator  |
|  (Interface)    |
+-----------------+
         |
         | (Implementation)
         v
+------------------------+
| FileSizeValidator      |
| FileExtensionValidator |
+------------------------+
		
```

### Overview
This structure outlines the organization of a Go-based backend project for managing file uploads, message queues, and other services.

### Top-Level Directories and Files
- **`README.md`**: Documentation for the project setup, usage, and other details.
- **`go.mod` / `go.sum`**: Go module configuration and dependency files.
- **`cmd`**:
  - `main.go`: Entry point for the application.
- **`config`**:
  - `config.go`: Configuration handling (e.g., environment variables, app settings).

### Internal Directory: Core Project Implementation
1. **Adaptors**
Bridges between external services or protocols and the application's domain layer.

- **`http`**:
    - `file_handlers.go`: Handlers for HTTP endpoints related to file operations.
- **`mq`**:
    - `azure_service_bus_adapter.go`: Adapter for Azure Service Bus integration.
    - `mq_handlers.go`: Handlers for processing messages from the queue.
- **`storage`**:
    - `azure_blob_storage.go`: Integration with Azure Blob Storage for file storage.
- **`validation`**:
    - `file_size_validator.go`: Validates file sizes.
    - `file_type_validator.go`: Validates file types.
    - `websocket`: Placeholder for WebSocket-related logic.
2. **Domain**
Contains core business logic, models, and interfaces.

- **`celery_message.go`**: Represents a message for Celery (Python task queue).
- **`file.go`**: Data structure representing file-related information.
- **`file_repository.go`**: Interface for file storage/repository operations.
- **`file_validator.go`**: Interface for file validation logic.
- **`message_queue.go`**: Interface for message queue interactions.

3. **Infra**
Handles infrastructure-level concerns (database, HTTP server, storage, WebSocket).

- **`database`**: Placeholder for database-related implementation.
- **`http`**:
  - `gin_server.go`: HTTP server implementation using Gin framework.
- **`storage`**: Placeholder for storage-related infrastructure.
- **`websocket`**: Placeholder for WebSocket-related infrastructure.

4. **LLM**
Manages use cases or logic related to large language models.

- **`llm_usecases.go`**: Business logic or use cases involving LLM functionality.

5. **Usecases**
Implements application-specific business use cases.

- **`File Upload`**:
  - `file_upload.go`: Defines the file upload use case.
  - `file_upload_impl.go`: Implementation of the file upload use case.
- **`Message Queue`**:
  - `message_queue.go`: Use case for handling message queues.

**Key Highlights**
- **`Separation of`** Concerns: Layers (adaptors, domain, infra, use cases) isolate responsibilities, ensuring maintainable and scalable code.
- **`Extensibility:`** Modular design facilitates integration with various services (e.g., Azure Service Bus, Blob Storage).
- **`Domain-Driven`** Design: Focuses on core business logic and its interface with external systems.


## Getting Started
---------------

1. **Clone the Repository**:
   ```bash
   git clone [repository-url]
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Run the Application**:
   ```bash
   go run cmd/main.go
   ```

## Configuration
---------------
Environment Variables:
- AZURE_OPENAI_ENDPOINT: Azure Open AI API endpoint
- OPENAI_ENDPOINT: Open AI API endpoint
- CLAUDE_ENDPOINT: Claude API endpoint
- LLAMA31_ENDPOINT: Llama 3.1 API endpoint
- PERPLEXITY_ENDPOINT: Perplexity API endpoint

## Contributing
---------------

- Fork the Project
- Create your Feature Branch (git checkout -b feature/AmazingFeature)
- Commit your Changes (git commit -m 'Add some AmazingFeature')
- Push to the Branch (git push origin feature/AmazingFeature)
- Open a Pull Request