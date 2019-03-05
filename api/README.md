# API Document


## 文章列表

URL: `api/articles`

Method: `GET`

Response:

```json=
{
    "data": [{
        "title": "標題",
        "content": "內容",
        "created_at": "2013-09-12T22:50:20+08:00",
        "updated_at": "2013-09-12T22:50:20+08:00",
        "likes_count": 3
    }]
    "code": 0,
    "message": "success"
}
```

## 建立文章

URL: `api/articles`

Method: `POST`

Request: 

| key | type | Required | description |
| --- | --- |  --- | --- |
| title | string | Required | 標題 |
| content | string | Required | 內容 |

```json=
{
    "title": "標題",
    "content": "這是測試的內容"
}
```

Response:

```json=
{
    "data": null,
    "code": 0,
    "message": "success"
}
```

## 按讚

URL: `api/articles/{{id}}/like`

Method: `PATCH`

Response:

```json=
{
    "data": null,
    "code": 0,
    "message": "success"
}
```