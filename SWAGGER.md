# Swagger Documentation

This project now includes Swagger/OpenAPI documentation for the API.

## Accessing Swagger UI

Once the server is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Generating Documentation

After making changes to API handlers or adding new endpoints with Swagger annotations, regenerate the documentation:

```bash
~/go/bin/swag init -g cmd/main.go -o docs
```

Or if `swag` is in your PATH:

```bash
swag init -g cmd/main.go -o docs
```

## Swagger Annotations

The API documentation is generated from Go comments in the handler files. Example:

```go
// CreateCampaign godoc
// @Summary Create a new campaign
// @Description Create a new marketing campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateCampaignRequest true "Campaign request"
// @Success 200 {object} dto.CampaignResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /campaign [post]
func (h *CampaignHandler) CreateCampaign(g *gin.Context) {
    // handler code
}
```

## API Overview

The API is organized into the following sections:

### User Management
- `POST /api/v1/user/register` - Register new user
- `POST /api/v1/user/login` - Login user
- `POST /api/v1/user/logout` - Logout user
- `GET /api/v1/user/me` - Get current user info
- `POST /api/v1/user/market-credential` - Save marketplace credentials
- `GET /api/v1/user/market-credential/{platform}` - Check marketplace credentials
- `DELETE /api/v1/user/market-credential/{platform}` - Delete marketplace credentials

### Product Management
- `POST /api/v1/product` - Add new product from marketplace URL
- `GET /api/v1/product` - Get all user products
- `GET /api/v1/product/{productId}` - Get product by ID
- `GET /api/v1/product/{productId}/offer` - Get product offers
- `DELETE /api/v1/product/{productId}` - Delete product

### Campaign Management
- `POST /api/v1/campaign` - Create new campaign
- `GET /api/v1/campaign` - Get user campaigns
- `GET /api/v1/campaign/available` - Get public campaigns
- `DELETE /api/v1/campaign/{campaign_id}` - Delete campaign

### Link Management
- `POST /api/v1/link` - Create affiliate link
- `GET /api/v1/link/{link_id}` - Get link by ID
- `GET /api/v1/link/short-code/{short_code}` - Get link by short code
- `GET /api/v1/link/campaign/{campaignId}` - Get links by campaign
- `GET /api/v1/link/redirect/{short_code}` - Redirect to affiliate URL
- `DELETE /api/v1/link/{link_id}` - Delete link

### Dashboard
- `GET /api/v1/dashboard/metrics` - Get dashboard analytics

### Public Redirect
- `GET /go/{short_code}` - Public redirect endpoint (also tracks clicks)

## Authentication

Most endpoints require authentication using a session cookie. After successful login or registration, a session cookie is set automatically.

For Swagger UI testing, you can use the "Authorize" button at the top of the page to enter your bearer token.

## Dependencies

- [swaggo/swag](https://github.com/swaggo/swag) - Swagger documentation generator
- [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) - Gin middleware for Swagger
- [swaggo/files](https://github.com/swaggo/files) - Embedded Swagger files

## Installation

The Swagger dependencies are already added to `go.mod`. To install the CLI tool:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Docker Support

### Building with Docker

The Dockerfile automatically generates Swagger documentation during the build process. Simply build the image:

```bash
docker build -t market-place-affiliate-api .
```

### Running with Docker Compose

Use the provided `docker-compose.yml` to run the API with PostgreSQL:

```bash
docker-compose up -d
```

Access Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

### Environment Variables for Docker

Create a `.env` file in the project root:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=your-secure-password
DB_NAME=affiliate
JWT_SECRET=your-jwt-secret-key
PASSWORD_SECRET=12345678901234567890123456789012
HTTP_HOST=0.0.0.0
HTTP_PORT=80
```

### Manual Swagger Generation with Make

```bash
make swagger
```

This will regenerate the Swagger docs in the `docs/` directory.
