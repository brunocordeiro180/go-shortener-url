# Go URL Shortener

Go URL Shortener is a web application that allows users to input a long URL and receive a shortened version. It is built using Go's `html/template` package for rendering HTML pages and `SQLite3` for data storage.
Was built just for learning golang concepts and it works as a MVP.

## Features

- Shorten long URLs into concise, easy-to-share links.
- Redirect users from shortened URLs to the original URLs.
- Store URL mappings in an SQLite3 database.
- Render HTML pages using Go templates.

## Prerequisites

- Go 1.16 or higher installed on your machine.
- SQLite3 installed for database management.

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/brunocordeiro180/go-shortener-url.git
cd go-shortener-url
```

2. **Install dependencies::**

Ensure that the necessary Go packages are installed. You can use go mod tidy to install any missing dependencies.

```bash
go mod tidy
```

3. **Run the application:**
```bash
go run main.go
```

## Usage
- Navigate to http://localhost:8080 in your web browser.
- Enter a long URL into the input field and submit the form.
- The page will display the shortened URL.
- Accessing the shortened URL will redirect you to the original URL.