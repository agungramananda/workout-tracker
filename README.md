# Workout Tracker

## Latar Belakang dan Tujuan

Aplikasi Workout Tracker dibuat untuk membantu pengguna dalam mencatat dan memantau aktivitas olahraga mereka. Dengan aplikasi ini, pengguna dapat mencatat jenis latihan, durasi, dan intensitas latihan mereka, serta melihat perkembangan mereka dari waktu ke waktu. Tujuan utama dari aplikasi ini adalah untuk memotivasi pengguna agar tetap konsisten dalam berolahraga dan mencapai tujuan kebugaran mereka.

## Stack yang Digunakan

- **Frontend:**

  - React.js

- **Backend:**

  - Go (Golang)
  - Gin Framework
  - GORM

- **Database:**

  - PostgreSQL

- **Authentication:**
  - JWT (JSON Web Tokens)

## Cara Instalasi

**Clone repository:**

```sh
git clone https://github.com/agungramananda/workout-tracker
```

### Backend

1. Install dependencies:

   ```sh
   cd/be
   go mod tidy
   ```

2. Setup environment variables:
   Gunakan file `env.example` yang telah disediakan dan tambahkan variabel-variabel berikut:

   ```env
   HOST=localhost
   PORT=8080
   SECRET_KEY=your_jwt_secret_key
   REFRESH_TOKEN_SECRET=your_refresh_token_secret_key

   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

### Frontend

1. Clone repository:

   ```sh
   git clone https://github.com/username/capstone-project-rg.git
   cd capstone-project-rg/fe/workout-tracker/frontend
   ```

2. Install dependencies:

   ```sh
   npm install
   ```

3. Setup environment variables:
   Buat file `.env` dan tambahkan variabel-variabel berikut:

   ```env
   REACT_APP_API_URL=http://localhost:8080
   ```

4. Run the application:
   ```sh
   npm start
   ```

## Daftar API Endpoints

### 1. User Authentication

- **Endpoint:** `/register`

  - **Method:** POST
  - **Deskripsi:** Mendaftarkan pengguna baru.
  - **Request Body:**
    ```json
    {
      "username": "string",
      "email": "string",
      "password": "string",
      "fullname": "string"
    }
    ```
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Register successful",
      "data": {
        "access_token": "string"
      }
    }
    ```

- **Endpoint:** `/login`
  - **Method:** POST
  - **Deskripsi:** Login pengguna.
  - **Request Body:**
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Login successful",
      "data": {
        "access_token": "string"
      }
    }
    ```

### 2. Workout Management

- **Endpoint:** `/workouts`

  - **Method:** GET
  - **Deskripsi:** Mendapatkan daftar semua workout.
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Workout retrieved successfully",
      "data": [
        {
          "id": "integer",
          "user_id": "integer",
          "name": "string",
          "description": "string",
          "date": "string",
          "time": "string",
          "rest_between_exercises": "integer",
          "is_completed": "boolean",
          "exercises_plan": [
            {
              "exercise_id": "integer",
              "exercise_name": "string",
              "sets": "integer",
              "reps": "integer",
              "weight": "integer",
              "rest_time": "integer",
              "order": "integer",
              "is_completed": "boolean"
            }
          ],
          "comments": [
            {
              "user_id": "integer",
              "comment": "string"
            }
          ]
        }
      ]
    }
    ```

- **Endpoint:** `/workouts`

  - **Method:** POST
  - **Deskripsi:** Menambahkan workout baru.
  - **Request Body:**
    ```json
    {
      "name": "string",
      "description": "string",
      "date": "string",
      "time": "string",
      "rest_between_exercises": "integer",
      "is_completed": "boolean",
      "exercises_plan": [
        {
          "exercise_id": "integer",
          "sets": "integer",
          "reps": "integer",
          "weight": "integer",
          "rest_time": "integer",
          "order": "integer",
          "is_completed": "boolean"
        },
        {
          "exercise_id": "integer",
          "sets": "integer",
          "reps": "integer",
          "weight": "integer",
          "rest_time": "integer",
          "order": "integer",
          "is_completed": "boolean"
        }
      ],
      "comments": [
        {
          "user_id": "integer",
          "comment": "string"
        }
      ]
    }
    ```
  - **Response:**
    ```json
    {
      "message": "Workout added successfully"
    }
    ```

- **Endpoint:** `/workouts/{id}`

  - **Method:** PUT
  - **Deskripsi:** Mengedit workout berdasarkan ID.
  - **Request Body:**
    ```json
    {
      "name": "string",
      "description": "string",
      "date": "string",
      "time": "string",
      "rest_between_exercises": "integer",
      "is_completed": "boolean",
      "exercises_plan": [
        {
          "exercise_id": "integer",
          "sets": "integer",
          "reps": "integer",
          "weight": "integer",
          "rest_time": "integer",
          "order": "integer",
          "is_completed": "boolean"
        }
      ],
      "comments": [
        {
          "user_id": "integer",
          "comment": "string"
        }
      ]
    }
    ```
  - **Response:**
    ```json
    {
      "message": "Workout updated successfully"
    }
    ```

- **Endpoint:** `/workouts/{id}`

  - **Method:** DELETE
  - **Deskripsi:** Menghapus workout berdasarkan ID.
  - **Response:**
    ```json
    {
      "message": "Workout deleted successfully"
    }
    ```

- **Endpoint:** `/workouts/{id}`
  - **Method:** GET
  - **Deskripsi:** Mendapatkan workout berdasarkan ID.
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Workout retrieved successfully",
      "data": {
        "id": "integer",
        "user_id": "integer",
        "name": "string",
        "description": "string",
        "date": "string",
        "time": "string",
        "rest_between_exercises": "integer",
        "is_completed": "boolean",
        "exercises_plan": [
          {
            "exercise_id": "integer",
            "exercise_name": "string",
            "sets": "integer",
            "reps": "integer",
            "weight": "integer",
            "rest_time": "integer",
            "order": "integer",
            "is_completed": "boolean"
          }
        ],
        "comments": [
          {
            "user_id": "integer",
            "comment": "string"
          }
        ]
      }
    }
    ```

### 3. Exercise Management

- **Endpoint:** `/exercises`

  - **Method:** GET
  - **Deskripsi:** Mendapatkan daftar semua exercise.
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Exercises retrieved successfully",
      "data": [
        {
          "id": "integer",
          "name": "string",
          "description": "string"
        }
      ]
    }
    ```

- **Endpoint:** `/exercises/{id}`
  - **Method:** GET
  - **Deskripsi:** Mendapatkan exercise berdasarkan ID.
  - **Response:**
    ```json
    {
      "status": 200,
      "message": "Exercise retrieved successfully",
      "data": {
        "id": "integer",
        "name": "string",
        "description": "string"
      }
    }
    ```

## Struktur Database

| Table Name     | Column Name            | Data Type | Description |
| -------------- | ---------------------- | --------- | ----------- |
| users          | id                     | INTEGER   | Primary Key |
|                | username               | TEXT      |             |
|                | email                  | TEXT      |             |
|                | password               | TEXT      |             |
| workouts       | id                     | INTEGER   | Primary Key |
|                | user_id                | INTEGER   | Foreign Key |
|                | name                   | TEXT      |             |
|                | description            | TEXT      |             |
|                | date                   | DATE      |             |
|                | time                   | TIME      |             |
|                | rest_between_exercises | INTEGER   |             |
|                | is_completed           | BOOLEAN   |             |
| exercises      | id                     | INTEGER   | Primary Key |
|                | name                   | TEXT      |             |
|                | description            | TEXT      |             |
| comments       | id                     | INTEGER   | Primary Key |
|                | workout_id             | INTEGER   | Foreign Key |
|                | user_id                | INTEGER   | Foreign Key |
|                | comment                | TEXT      |             |
| exercise_plans | id                     | INTEGER   | Primary Key |
|                | workout_id             | INTEGER   | Foreign Key |
|                | exercise_id            | INTEGER   | Foreign Key |
|                | sets                   | INTEGER   |             |
|                | reps                   | INTEGER   |             |
|                | weight                 | INTEGER   |             |
|                | rest_time              | INTEGER   |             |
|                | order                  | INTEGER   |             |
|                | is_completed           | BOOLEAN   |             |

## Dokumentasi Flow Interaksi Frontend dengan Backend

1. **User Registration:**

   - Pengguna mengisi form pendaftaran di frontend.
   - Frontend mengirimkan data ke endpoint `/register`.
   - Backend memproses data dan mengembalikan respons sukses atau gagal.

2. **User Login:**

   - Pengguna mengisi form login di frontend.
   - Frontend mengirimkan data ke endpoint `/login`.
   - Backend memverifikasi data dan mengembalikan token jika berhasil.

3. **Fetching Workouts:**

   - Frontend mengirimkan permintaan GET ke endpoint `/workouts`.
   - Backend mengembalikan daftar workout yang terkait dengan pengguna.

4. **Adding Workout:**

   - Pengguna mengisi form workout di frontend.
   - Frontend mengirimkan data ke endpoint `/workouts`.
   - Backend memproses data dan mengembalikan respons sukses atau gagal.

5. **Editing Workout:**

   - Pengguna mengisi form edit workout di frontend.
   - Frontend mengirimkan data ke endpoint `/workouts/{id}`.
   - Backend memproses data dan mengembalikan respons sukses atau gagal.

6. **Deleting Workout:**

   - Pengguna memilih workout yang ingin dihapus di frontend.
   - Frontend mengirimkan permintaan DELETE ke endpoint `/workouts/{id}`.
   - Backend memproses permintaan dan mengembalikan respons sukses atau gagal.

7. **Fetching Workout by ID:**

   - Frontend mengirimkan permintaan GET ke endpoint `/workouts/{id}`.
   - Backend mengembalikan data workout yang sesuai dengan ID.

8. **Fetching Exercises:**

   - Frontend mengirimkan permintaan GET ke endpoint `/exercises`.
   - Backend mengembalikan daftar exercise.

9. **Fetching Exercise by ID:**
   - Frontend mengirimkan permintaan GET ke endpoint `/exercises/{id}`.
   - Backend mengembalikan data exercise yang sesuai dengan ID.
