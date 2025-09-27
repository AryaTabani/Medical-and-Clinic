# Clinic Information API

This project is a read-only REST API built with **Go** and **Gin**. It is designed to serve all the necessary content for a public-facing clinic or healthcare website.

## ✨ Features

-   **Read-Only Endpoints**: The API is optimized for fast data retrieval to supply a frontend application with information.
-   **Doctor Filtering**: The `/doctors` endpoint allows clients to filter the list of doctors by their medical specialty.
-   **Data Aggregation**: A dedicated `/about` endpoint provides summarized data, such as the total count of active doctors and clinics, perfect for an "About Us" page.
-   **Clean Architecture**: Built with a layered architecture (Controller, Service, Repository) for maintainability and clear separation of logic.
-   **Comprehensive Content**: Provides structured data for all key sections of a clinic website, including doctors, clinic locations, facilities, and patient testimonials.

## 🛠️ Tech Stack

-   **Language**: Go
-   **Web Framework**: Gin
-   **Database**: SQLite

## 📄 API Endpoints

All endpoints are `GET` requests and do not require authentication.

| Method | Path          | Description                                         |
| :----- | :------------ | :-------------------------------------------------- |
| `GET`  | `/doctors`    | Get a list of all doctors.                          |
| `GET`  | `/doctors?specialty=:spec` | Filter doctors by their specialty.     |
| `GET`  | `/clinics`    | Get a list of all clinic locations.                 |
| `GET`  | `/facilities` | Get a list of all clinic facilities and equipment.  |
| `GET`  | `/testimonials`| Get a list of all patient testimonials.            |
| `GET`  | `/about`      | Get aggregated data for the "About Us" page.        |

## 🚀 Getting Started

**Prerequisites:**
-   Go (version 1.18 or later)
-   Git

**Instructions:**
1.  Clone the repository:
    ```bash
    git clone https://github.com/AryaTabani/Medical-and-Clinic.git
    cd Medical-and-Clinic
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the application:
    ```bash
    go run main.go
    ```
    The API will be available on the default port.
