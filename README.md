# Product & User Microservice Suite

A modular microservice architecture built with **Go (Golang)**, **GORM**, and **MySQL**, utilizing clean architecture patterns and connection pooling.

---

## 🏛️ Architecture Overview

This repository contains two core microservices that communicate independently with dedicated SQL schemas:

* **Product Service:** Handles product creation, updates, pricing, and catalog search.
* **User Service:** Manages user registration, profiles, and identity using GUID/UUID primary keys.
