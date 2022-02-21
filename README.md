# Getting Started

### Starting the application with docker
```zsh
  make dev-docker-up
```

### Shutting application wih docker
```zsh
   make dev-docker-down
```

### Starting the locally with Go and Docker
```zsh
  make dev-local
```

The server will be listening on localhost:8080.
# End Points

### The complete API documentation is available at /swagger.

Demo: http://localhost:8080/docs/v1/starwars/swagger/index.html

### `/api/v1/starwars - POST`

```json
{
    "name": "coruscant",
    "climate": "clim1,clim2",
    "terrain": "terrain1,terrain2"
}
```

### `/api/v1/starwars - Get`

### Response

```json
{
    "data": [
        {
            "id": "6213aa3022a3324f016ea902",
            "name": "coruscant",
            "climate": "clim1,clim2",
            "terrain": "terrain1,terrain2",
            "quantity_film_appearances": 4,
            "created_at": "2022-02-21T15:05:19.249Z",
            "updated_at": "2022-02-21T15:05:19.249Z"
        },
        {
            "id": "6213aa3322a3324f016ea903",
            "name": "test2",
            "climate": "clim1,clim2",
            "terrain": "terrain1,terrain2",
            "quantity_film_appearances": 0,
            "created_at": "2022-02-21T15:05:22.786Z",
            "updated_at": "2022-02-21T15:05:22.786Z"
        }
    ]
}
```

### `/api/v1/starwars/{id} - Get`

### Response

```json
{
    "id": "6213aa3022a3324f016ea902",
    "name": "coruscant",
    "climate": "clim1,clim2",
    "terrain": "terrain1,terrain2",
    "quantity_film_appearances": 4,
    "created_at": "2022-02-21T15:05:19.249Z",
    "updated_at": "2022-02-21T15:05:19.249Z"
}
```

### `/api/v1/starwars/name/{name} - POST`

### Response

```json
{
    "id": "6213aa3022a3324f016ea902",
    "name": "coruscant",
    "climate": "clim1,clim2",
    "terrain": "terrain1,terrain2",
    "quantity_film_appearances": 4,
    "created_at": "2022-02-21T15:05:19.249Z",
    "updated_at": "2022-02-21T15:05:19.249Z"
}
```

### `/api/v1/starwars/{id} - DELETE`

## Testing

### Run tests
```zsh
  make test
```

