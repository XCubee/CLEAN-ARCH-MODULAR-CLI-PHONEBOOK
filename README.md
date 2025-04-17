# go-clean-phonebook

A **modular, clean-architecture-based CLI phonebook application** built in Go.  
Designed with the same software engineering principles used in top tech companies to ensure **scalability, maintainability, and testability**.

---

## Features

- Add new contacts
- View all saved contacts
- Search contacts by name
- Simple terminal-based interface
- Clean file structure and decoupled business logic

---

## Tech Stack & Design Principles

- **Language:** [Go](https://golang.org/)
- **Architecture Style:** Clean Architecture (modular & layered)
- **Coding Principles:**
  - Separation of concerns
  - Proper package structuring
  - Encapsulation using unexported variables and exported functions
  - In-memory storage abstraction (replaceable with DB/file later)
  - Pointers for memory efficiency
  - Reusable utilities

---

## 🗂️ Project Structure
go-clean-phonebook/
├── main.go                  
│
├── models/                  
│   └── contact.go          
│
├── storage/                
│   └── store.go             
│
├── handlers/               
│   └── contactHandler.go    
│
├── utils/                  
│   └── input.go             
│
└── go.mod                   


## 📦 How to Run

1. **Clone the repo:**

```bash
git clone https://github.com/yourusername/go-clean-phonebook.git
cd go-clean-phonebook
