# Journal API

Base URL: `/api`

Authentication: JWT Bearer token required for all authenticated endpoints

## Notes (authenticated)

### Get All Notes

- URL: `/notes`
- Method: GET
- Response:

```json
[
    {
        uuid:
        user_id
        note
        created_at
        updated_at
    }
]
```

### Create Note

- URL: `/notes`
- Method: GET
- Body: `{note: }`
- Response: `{uuid: }`
