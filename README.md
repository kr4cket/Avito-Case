# Запуск dev-среды

##### Код запуска dev-среды:
```
 docker-compose up -d --build avito-app
 migrate -database postgres://postgres:task@localhost:5436/?sslmode=disable -path ./schema up
```
---

#### Пример работы:

##### Запрос
`http://localhost:8081/api/service/cost/1`
##### Ответ
**JSON**
```
{
    "location_id": "10",
    "matrix_id": "100",
    "microcategory_id": "5",
    "price": "1000",
    "user_segment_id": "5
}
```