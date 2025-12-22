# 📧 Email Campaign Dispatcher (Go)

A **lightweight, concurrent email campaign dispatcher** built in **Go** using the **Producer–Consumer pattern** and **Goroutines**.  
This project is an **MVP / Proof of Concept** designed to send bulk emails efficiently at low cost by leveraging Go’s powerful concurrency model.

---

## Why This Project Exists

Most email marketing tools become **very expensive** as your email list grows.

Instead of relying on third-party tools:
- We build **our own email dispatcher**
- We control **cost, performance, and scaling**
- We use **Go**, because concurrency is its strongest feature

This project focuses on **core backend functionality only** — no unnecessary complexity.

---

## Core Concept

**Problem**  
Sending emails sequentially is slow and wastes system resources.

**Solution**  
Send emails **concurrently** using worker goroutines.

### Key Components
- **CSV file** → Source of recipient data  
- **Producer** → Reads CSV and creates jobs  
- **Consumers (Workers)** → Send emails concurrently  
- **Channels** → Communication between goroutines  
- **SMTP Server** → External email delivery service  

---

## Architecture Overview

```text
CSV File
   |
   v
Producer (CSV Reader)
   |
   v
Go Channel (Job Queue)
   |
   v
Multiple Consumer Workers
   |
   v
SMTP Server
   |
   v
End Users
```
## Project Structure
```text
email-dispatcher/
│
├── main.go          # Application entry point
├── emails.csv       # Sample CSV file with recipients
├── go.mod           # Go module configuration
└── README.md
```

## 🛠️ Setup & Run

### 1. Check Go Installation
  ```
    go version
  ```
### 2. Initialize Go Module
  ```
    go mod init github.com/your-username/email-dispatcher
  ```
### 3. Run the Application
 ```
    go run main.go
  ```

## MVP Goals

- Concurrent email sending
- Efficient CPU and memory usage
- Simple and readable code

### Not Included
- Retry logic
- Scheduling
- UI or dashboard

