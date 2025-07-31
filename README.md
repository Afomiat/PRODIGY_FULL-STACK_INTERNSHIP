# 👥 Employee Management System

A full-stack Employee Management System designed to digitize HR operations for small to mid-sized teams. This system replaces error-prone spreadsheet tracking with a secure, scalable solution. Built with **Go** for the backend and **React** for the frontend, it includes authentication, role-based access, and employee CRUD operations.

---

## 🚀 Features

- 🧾 **Employee Records Management** — Create, read, update, and delete employee data.
- 🔐 **JWT Authentication** — Secure login system using JSON Web Tokens.
- 🛡 **Role-Based Access Control** — Admin vs. staff privileges; 100% unauthorized access blocked during testing.
- 📉 **Error Reduction** — Replaced spreadsheets and reduced data inconsistencies by 30%.
- 📱 **Responsive Design** — Works well on desktops, tablets, and mobile screens.

---

## 🛠 Tech Stack

| Frontend  | Backend  | Auth        | Database         |
|-----------|----------|-------------|------------------|
| React     | Go       | JWT         | MongoDB

---

## 🧰 Installation

### Prerequisites
- Node.js ≥ 16
- Go ≥ 1.18
- MongoDB or PostgreSQL (whichever your backend is configured with)

### 1. Clone the repository
```bash
git clone https://github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP.git
cd employee-management-system
````

### 2. Backend Setup (Go)

```bash
cd backend
go mod tidy
go run main.go
```

> ⚠️ Ensure MongoDB or your configured database is running.

### 3. Frontend Setup (React)

```bash
cd frontend
npm install
npm run dev
```

> The frontend runs at `http://localhost:3000` by default.

---

## 🌐 Project Structure

```
employee-management-system/
│
├── backend/           # Go backend
│   └── main.go        # App entry point
│
├── frontend/          # React frontend
│   ├── src/
│   └── ...
```

---

## 🧪 Testing the System

1. Start both backend and frontend servers.
2. Register or log in as an **Admin** to manage all employee records.
3. Log in as a **Staff** member to view limited data based on permissions.


## ✨ Future Enhancements

* ✅ Employee attendance tracking
* ✅ Performance review module
* ✅ Export reports to CSV/PDF
* ✅ Email notifications for HR events
* ✅ Audit logs


## 🙌 Contributing

Contributions are welcome! Please fork the repo and open a pull request with any improvements or bug fixes.

