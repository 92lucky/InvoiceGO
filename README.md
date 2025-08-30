# 🧾 InvoiceGO - PDF Invoice Generator in Go

**InvoiceGO**“This internal tool enhances PT PERTAMINA’s main operations by making invoice management faster and more efficient. It simplifies the process of generating invoices and creating detailed Operational Reports, helping teams save time, stay organized, and maintain accuracy.”. [🎥 Demo Aplikasi](
https://youtu.be/DrMlzEFmyto?si=L2B3OkGbR56zVQP5)

---

##  Features
- ✅ Invoice generator
- ✅ Operational Report generator

## Technical Implementation
- ✅ Google Outh Login
- ✅ Generate invoice PDF from form input
- ✅ Preview/Download PDF
- ✅ Extract/parsing raw excel LO & Add vendor name heading & total cell
- ✅ Clean/scalable folder structure
- ✅ CI/CD Pipeline
- ✅ Docker,DockerHub
- ✅ Production Ready / Deployment

---

## ⚙️ Installation

```bash
git clone https://github.com/92lucky/InvoiceGO.git
cd InvoiceGO
go mod tidy
go run cmd/main.go
go test ./test

## 🧑‍💻 Usage

1. Visit `http://localhost:8080/setup`
2. Fill in the invoice form
3. Click "Preview" to see the invoice
4. Click "Download" to get PDF

## 📡 API Endpoints

| Method | Endpoint         | Description                            | Auth Required |
|--------|------------------|----------------------------------------|----------------|
| GET    | `/`              | Show login or landing page             | ❌             |
| GET    | `/index`         | Dashboard / Home (after login)         | ✅             |
| GET    | `/setup`         | Setup invoice form                     | ✅             |
| POST   | `/generate`      | Generate and preview invoice           | ✅             |
| POST   | `/generate-pdf`  | Generate and download invoice as PDF   | ✅             |
| GET    | `/lo`            | Show LO (Letter of Offer) form         | ✅             |
| POST   | `/previewLo`     | Preview LO PDF                         | ✅             |
| POST   | `/downloadLo`    | Download LO PDF                        | ✅             |
| GET    | `/static/...`    | Serve static files (CSS, JS, etc.)     | ❌             |


## monolith app go + htmx








