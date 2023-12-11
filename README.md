## Catatan Keuangan


### Database
```sql
CREATE DATABASE IF NOT EXISTS catatan_keuangan_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE transaction_type AS ENUM ('CREDIT', 'DEBIT');

CREATE TABLE expenses (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    date DATE NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    transaction_type transaction_type,
    balance DOUBLE PRECISION NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
```

### API Documentation

1. `CREATE` pengeluaran
    - `POST` -> `/api/v1/expenses`
    - Request:
      ```json
      {
         "amount": 250000,
         "transactionType": "CREDIT",
         "description": "Tambahan jajan"
      }
      ``` 
    - Response:
      ```json
      "status": {
         "code": 201,
         "message": "Created"
      },
      "data": {
         "id": "a81bc81b-dead-4e5d-abff-90865d1e13b1",
         "date": "2023-12-08 05:17:42.583767+07",
         "amount": 250000,
         "transactionType": "CREDIT",
         "description": "Tambahan jajan",
         "createdAt": "2023-12-08 05:17:42.583767+07",
         "updatedAt": "2023-12-08 05:17:42.583767+07"
      }
      ``` 
3. `LIST` Pengeluaran (pagination dan atau bisa filter dengan range tanggal)
    - `GET` -> `/api/v1/expenses`
    - Query Params:
      ```
      ?page=1&size=5
      ?page=1&size=5&startDate=2023-12-08&endDate=2023-12-08
      ``` 
    - Response:
      ```json
      "status": {
         "code": 200,
         "message": "Ok"
      },
      "data": [
        {
           "id": "a81bc81b-dead-4e5d-abff-90865d1e13b1",
           "date": "2023-12-08 05:17:42.583767+07",
           "amount": 250000,
           "transactionType": "CREDIT",
           "balance": 50000000,
           "description": "Tambahan jajan",
           "createdAt": "2023-12-08 05:17:42.583767+07",
           "updatedAt": "2023-12-08 05:17:42.583767+07"
        }
      ],
      "paging": {
         "page": 1,
         "totalPages": 2,
         "totalRows": 10,
         "rowsPerPage": 5
      }
      ``` 
5. `GET` by ID
    - `GET` -> `/api/v1/expenses/a81bc81b-dead-4e5d-abff-90865d1e13b1`
    - Response:
      ```json
      "status": {
         "code": 200,
         "message": "Created"
      },
      "data": {
         "id": "a81bc81b-dead-4e5d-abff-90865d1e13b1",
         "date": "2023-12-08 05:17:42.583767+07",
         "amount": 250000,
         "transactionType": "CREDIT",
         "balance": 50000000,
         "description": "Tambahan jajan",
         "createdAt": "2023-12-08 05:17:42.583767+07",
         "updatedAt": "2023-12-08 05:17:42.583767+07"
      }
7. `GET` berdasarkan tipe transaksi
    - `GET` -> `/api/v1/expenses/type/:type`
    - Params:
      ```
      /CREDIT
      /DEBIT
      ``` 
    - Response:
      ```json
      "status": {
         "code": 200,
         "message": "Ok"
      },
      "data": [
        {
           "id": "a81bc81b-dead-4e5d-abff-90865d1e13b1",
           "date": "2023-12-08 05:17:42.583767+07",
           "amount": 250000,
           "transactionType": "CREDIT",
           "balance": 50000000,
           "description": "Tambahan jajan",
           "createdAt": "2023-12-08 05:17:42.583767+07",
           "updatedAt": "2023-12-08 05:17:42.583767+07"
        }
      ],
      "paging": {
         "page": 1,
         "totalPages": 2,
         "totalRows": 10,
         "rowsPerPage": 5
      }
