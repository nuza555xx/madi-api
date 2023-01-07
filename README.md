# Madi API
  ของดี

> ทำวันนี้ให้แย่กว่าเมื่อวาน พร่ามมม

| Services | Method |  Endpoint |  Description  |
| -------- | ------ | --------- |  ----------   |
| Authentication | POST | auth/register         | สมัครสมาชิกไงกำ                 |
| Authentication | GET  | auth/login            | เข้าสู่ระบบ                      |
| Authentication | POST | auth/register-token   | สมัคร fcm messaging            |
| Authentication | POST | auth/unregister-token | ยกเลิกสมัคร fcm messaging.      |
| User           | GET  | user/:userId          | แสดงข้อมูลผู้ใช้                   |
| User           | GET  | user/my-car           | แสดงข้อมูลรถโง่ๆ                 |
| Insurance      | POST | insure/car            | เพิ่มรถของตัวเอง เช่น MG          |
| Insurance      | PUT  | insure/car            | แก้ไขรถตัวเอง เช่น เหลือแค่ผ่อนกุญแจ |
| Insurance      | GET  | insure/search/:carNo  | ค้นหารถไอ้เหี้ยนั้น                 |
| Notification   | GET  | notification/list     | แจ้งเตือนโง่ๆ                    |


```mermaid
sequenceDiagram
    Client->>+Rest: [POST] auth/register
    Rest->>+DB: create user
    DB-->>-Rest: return user
    Rest-->>-Client: response accessToken

    Client->>+Rest: [POST] auth/login
    Rest->>+DB: query user
    DB-->>-Rest: return user
    Rest-->>-Client: response accessToken

    Client->>+Rest: [POST] auth/register-token
    Rest->>-DB: update user

    Client->>+Rest: [POST] auth/unregister-token
    Rest->>-DB: update user

    Client->>+Rest: [GET] user/userId
    Rest->>+DB: query user
    DB-->>-Rest: return user
    Rest-->>-Client: response user

    Client->>+Rest: [GET] user/my-car
    Rest->>+DB: query car
    DB-->>-Rest: return car
    Rest-->>-Client: response car[]

    Client->>+Rest: [POST] insure/car
    Rest->>+DB: create car
    DB-->>-Rest: return car
    Rest-->>-Client: response car

    Client->>+Rest: [PUT] insure/car
    Rest->>+DB: update car
    DB-->>-Rest: return car
    Rest-->>-Client: response car

    Client->>+Rest: [PUT] insure/search/carNo
    Rest->>+DB: query car
    DB-->>-Rest: return car
    Rest-->>-Client: response car

```

