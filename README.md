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

```json
{
  "id": "someid",
  "name": "name of the coaster",
  "inPark": "the amusement park the ride is in",
  "manufacturer": "name of the manufacturer",
  "height": 27,
}
```
