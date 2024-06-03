# EcomMaster (In progress....)

EcomMaster is a comprehensive e-commerce platform built with a microservices architecture using Golang. This project is designed to practice and showcase advanced concepts such as microservices, distributed databases, RPC, logging, Gin, PostgreSQL, and RabbitMQ.

## Table of Contents

- [Overview](#overview)
- [Microservices](#microservices)
- [Technology Stack (expected)](#technology-stack)

## Overview

EcomMaster is an e-commerce platform designed to demonstrate the use of microservices and various technologies. It consists of multiple services handling different functionalities such as user management, product catalog, order management, payment processing, and notifications.

## Microservices

1. **User Service**

   - User registration, authentication, and profile management.
   - JWT-based authentication.

2. **Product Catalog Service**

   - CRUD operations for products.
   - Product searching and filtering.

3. **Order Service**

   - Order creation and management.
   - Integration with Payment and Inventory services.

4. **Payment Service**

   - Payment processing.
   - Integration with third-party payment gateways.

5. **Inventory Service**

   - Inventory management.
   - Stock level updates.

6. **Notification Service**
   - Sending email/SMS notifications.
   - Using RabbitMQ for message queuing.

## Technology Stack

- **Golang**: Main programming language for all microservices.
- **Gin**: Web framework for building RESTful APIs.
- **PostgreSQL**: Database for persistent storage.
- **gRPC**: Protocol for communication between microservices.
- **RabbitMQ**: Message broker for asynchronous communication.
- **Elasticsearch**: Full-text search engine (for Search Service).
- **Prometheus & Grafana**: Monitoring and alerting.
- **Redis**: Caching layer.
- **Kafka**: Event streaming platform for an event-driven architecture.
- **Jaeger/Zipkin**: Distributed tracing.

## Getting Started

### Prerequisites

- Golang (latest version)
- Docker
- PostgreSQL
- RabbitMQ
- Kafka
- Elasticsearch
- Redis
