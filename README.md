# ERP System - Invoice Management

This is a Go-based ERP system for managing invoices, clients, and products. It uses **SQLite** as the database and a frontend built with **HTML, CSS, and JavaScript**.

## Features

* Create, read, and manage invoices (buying, selling, receipts, delivery notes)
* Store and manage client and product information
* Generate PDF invoices from HTML templates
* Interactive frontend for invoice creation and data management

## Project Structure

```
├── cmd/app                  # Main application entry point
├── internal                 # Backend logic, models, services, and utilities
├── assets                   # HTML templates and PDF templates
├── static                   # Frontend scripts, styles, images
├── go.mod                   # Go module file
├── go.sum                   # Go module checksums
├── replacements.txt         # File for sensitive data replacement
```

## Requirements

* Go >= 1.25
* SQLite3
* `.env` file containing environment variables (database path, API keys, etc.)

## Setup and Run

1. **Clone the repository**

```bash
git clone <repo-url>
cd <repo-folder>
```

2. **Add `.env` file**

* Place your environment variables in the `.env` file at the root of the project.

3. **Export environment variables**

```bash
cd cmd/app
. ./exportenv.sh
```

4. **Run the application**

```bash
go run .
```

* The backend server will start and interact with the SQLite database.
* Access the frontend through the static HTML/CSS/JS files in your browser.

## Usage

* Use the frontend pages in `assets/templates` and `static/scripts` to create invoices, manage clients, and products.
* PDFs are generated using the templates in `assets/pdftemplates`.

## Notes

* Make sure the SQLite database file path is correctly set in `.env`.
* All sensitive data should be kept out of the repository when making it public; use `replacements.txt` if needed to scrub secrets before sharing.

---

**Enjoy managing your invoices efficiently with this ERP system!**
