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


# ER Diagram

``` mermaid
erDiagram
    User }|..|{ Car : has
    User }|..|{ Device : has
    User }|..|{ Notification : has
    Car }|..|{ InsuranceProvider : has
    User {
        OID id
        string phone
        string name
        object setting
        object geolocation
        date createdAt
        date updatedAt
    }

    Device {
        OID id
        OID userId
        string deviceName
        string deviceUUID
        string firebaseToken
        date createdAt
        date updatedAt
    }

    Car {
        OID id
        OID userId
        OID insureId
        string brandName
        string genNo
        string carNo
        number expiredYear
        string insureRangeAmount
        date createdAt
        date updatedAt
    }

    Notification {
        OID id
        OID userId
        string title
        string message
        boolean read
        date readAt
        object ref
        date createdAt
        date updatedAt
    }

    InsuranceProvider {
        OID id 
        string name
        string phone
        string address
        date createdAt
        date updatedAt
    }

```
# Sequence Diagram

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

# Tech Stacks
| Name | Meaning | Role
| -------- | ------ |
| IONIC Framework | Using application cross platform is easy. | Frontend
| Golang | Using api prove business logic is the fast. | Backend
| MongoDB | Using store data. | Store
| Redis | Using cache data is stable. | Store
| Firebase | Using make push notification to device | Provide





