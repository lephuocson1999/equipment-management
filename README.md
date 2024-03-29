# equipment-management

## Database
1. Create postgresql server
```
docker compose up -d
```

2. Create a new database
```
CREATE DATABASE equipments;
```

3. Import table: Import structure data in equipment.sql file

4. Import example data from import_data.csv file

## How to use
- Create .env from .env.example
- Fetch dependencies : `go mod tidy` `go mod download`
- Build app : `go build -o app`
- Run Server: `./app serve`

Coding convention
1. Project structure
```
cmd // Chứa main run service
    └── server.go
    └── root.go
di // Dependency Injection
    └── devicefx
        └── initialize.go
    └── postgresfx
        └── initialize.go
pkg // Chứa common func
    dto
        └── base.go
    constant
        └── constant.go
    middleware
        └── authenticate.go
service // Chứa business logic tương ứng từng service
    device
        router // handler nhận request, validate input
            └── device.go
        usecase // Phần xử lý logic, get/mix data từ db, api
            └── device_ucase.go
            └── device_ucase_test.go
        repository // Tương tác DB, 3rd api
            └── device.go
        builder // Xử lý các func builder support cho các main logic 
            └── device.go
            └── device_test.go
        model 
            └── device.go
        constant 
            └── device.go
```

2. HTTP response
HTTP code: `2xx`:success; `4xx`:client input error; `5xx`: server err

## Step coding API
Viết theo thứ tự phụ thuộc: repositories -> usecases -> routers.

## CURL
1. Search name và description: q=Aga 
2. Pagination: page=1&limit=2
3. Sort: screen DESC
4. Search from, to với chi phí: from_expense=1000&to_expense=4000
5. Search from, to với ngày sản xuất (timestamp): from_date_of_manufacture=1709830800&to_date_of_manufacture=1710435600
6. Filter multi value với các field khác: working_radius=30m&working_radius=20m

```
curl --location 'localhost:8000/api/v1/devices?page=1&limit=2&sort=screen%20DESC&q=Aga' \
--header 'x-api-key: 123456789'
```
