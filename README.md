# go-and-api

> یک برنامه ساده برای تمرین و کار با زبان گولنگ
> می توانید
> [کد](https://github.com/kubucation/go-rollercoaster-api)
> و
> [ویدیو](https://www.youtube.com/watch?v=2v11Ym6Ct9Q)
> اصلی را نیز ببینید

## Routes and Requirements

- `GET /users` برگرداندن همه کاربر های موجود به صورت لیستی از جیسون ها
- `GET /users/{id}` برگرداندن یک کاربر که شناسه آن را وارد کردید به صورت جیسون
- `POST /users/` ایجاد یک کاربر جدید در لیست داده هایمان
- `POST /users/` اگر نوع داده در هدر ارسالی به صورت جیسون نبود خطا بدهد
- `GET /admin` چک کردن که آیا کاربر درست درخواست داده است
- `GET /users/random` یک کاربر را به صورت شانسی انتخاب کرده و به مسیری که اطلاعات آن است برود

### Data Types

> **Content-Type:** application/json

```go
// User data type
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
}
```

And **JSON** output:

```json
{
  "id": "10000",
  "name": "Sajad",
  "family": "Fahimian",
  "email": "test@github.com",
  "age": 27,
}
```
