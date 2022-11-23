# Hospital Management System API

## Introduction
**Hospital Management System API** is an API for an HMS app (desktop/tablet/web based) that focuses on the low-end health facilities (small hospitals, puskesmas, klinik, etc). This app need to accommodate various roles in health facilities such as:
- Medical staffs - are the main users of the app, who inputs the patient data and process any inquiry of the patient. Usually medical staffs are located in the front lobby (Admission).
- Nurses and Doctors - are the secondary users of the app, who need updated information about their patients to plan desired acts. The subroles of this user are consultants (specialists), registrars (seniors), and residents.

## How to use ?
1. Clone this repository
    ```shell
    git clone https://github.com/Alterra-Batch3-Kelompok-11/hms-backend.git
    ```
2. Create branch development
   ```shell
   git branch development
   ```
3. Pull from development
   ```shell
   git pull origin development
   ```
4. Sync dependencies
    ```shell
    go mod tidy
    ```
5. Create database `hospital_mng`
6. Make sure your configuration in file `.env` 
7. Run program
    ```shell
    go run main.go
    ```

## Specification
Run the API , and then go to http://localhost:8080/swagger/index.html

