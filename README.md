# Market Place Affiliate API

A RESTful API service for managing marketplace affiliate links from Lazada and Shopee platforms. Built with Go, Gin, and GORM.

Base backend demo api : https://api-production-5a6a.up.railway.app

Swagger : https://api-production-5a6a.up.railway.app/swagger/index.html

![Simple archietecture](zdemo/Screenshot%202568-11-16%20at%2007.22.51.png)


## 🚀 Features

- **User Management** - Registration, authentication, and profile management
- **Product Management** - Import products from Lazada/Shopee URLs with automatic offer tracking
- **Campaign Management** - Create and manage marketing campaigns with UTM tracking
- **Affiliate Link Generation** - Generate short affiliate links with click tracking
- **Dashboard Analytics** - View performance metrics, top products, and click statistics
- **Marketplace Integration** - Support for Lazada and Shopee affiliate APIs
- **Swagger Documentation** - Interactive API documentation at `/swagger/index.html`

## 🛠️ Tech Stack

- **Language**: Go 1.25
- **Web Framework**: Gin
- **ORM**: GORM with PostgreSQL
- **Authentication**: JWT tokens with cookie-based sessions
- **Documentation**: Swagger/OpenAPI 3.0
- **External APIs**: Lazada Affiliate API, Shopee Affiliate API
- **Testing**: Testify with mock repositories

## 📁 Project Structure

```
.
├── cmd/
│   ├── main.go              # Application entry point
│   └── httpserver/          # HTTP server setup
├── config/                  # Configuration management
├── infrastructures/         # Database and external services
├── internal/
│   ├── core/
│   │   ├── domains/         # Domain models
│   │   ├── dto/             # Data transfer objects
│   │   ├── ports/           # Interfaces/contracts
│   │   └── services/        # Business logic
│   ├── handlers/            # HTTP request handlers
│   └── repositories/
│       ├── db/              # Database implementations
│       └── mocks/           # Test mocks
├── docs/                    # Swagger documentation
└── pkg/                     # Utility packages
```

## 🏃 Getting Started

### Prerequisites

- Go 1.25 or higher
- PostgreSQL 15+
- Lazada/Shopee Affiliate API credentials (optional for testing)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/market-place-affiliate/api.git
   cd api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run database migrations**
   ```bash
   go run cmd/main.go
   # Migrations run automatically on startup
   ```

5. **Start the server**
   ```bash
   make start
   # or with hot reload
   make dev
   ```

6. **Access the API**
   - API: `http://localhost:8080`
   - Swagger UI: `http://localhost:8080/swagger/index.html`

### Docker Deployment

1. **Using Docker Compose** (Recommended)
   ```bash
   docker-compose up -d
   ```

2. **Manual Docker Build**
   ```bash
   docker build -t market-place-affiliate-api .
   docker run -p 8080:80 market-place-affiliate-api
   ```

## 📚 API Documentation

Full API documentation is available via Swagger UI at `/swagger/index.html` when the server is running.

### Main Endpoints

#### Authentication
- `POST /api/v1/user/register` - Register new user
- `POST /api/v1/user/login` - Login user
- `GET /api/v1/user/me` - Get current user info

#### Products
- `POST /api/v1/product` - Import product from marketplace URL
- `GET /api/v1/product` - List user's products
- `GET /api/v1/product/{id}/offer` - Get product offers

#### Campaigns
- `POST /api/v1/campaign` - Create campaign
- `GET /api/v1/campaign` - List campaigns
- `DELETE /api/v1/campaign/{id}` - Delete campaign

#### Links
- `POST /api/v1/link` - Generate affiliate link
- `GET /api/v1/link/campaign/{id}` - Get campaign links
- `GET /go/{short_code}` - Redirect (tracks clicks)

#### Dashboard
- `GET /api/v1/dashboard/metrics` - Get analytics

## 🧪 Testing

Run all tests:
```bash
go test ./... -v
```

Run tests with coverage:
```bash
go test ./... -cover
```

Run specific test suite:
```bash
go test ./internal/core/services/... -v
```

## 🔧 Development

### Generate Swagger Documentation
```bash
make swagger
# or
swag init -g cmd/main.go -o docs
```

### Build the Application
```bash
make build
```

### Clean Build Artifacts
```bash
make clean
```

## 🌐 Marketplace Integration

### Lazada Setup
1. Create a developer account at [Lazada Open Platform](https://open.lazada.com)
2. Register your application and get API credentials
3. Save credentials via `POST /api/v1/user/market-credential`

### Shopee Setup
1. Register at [Shopee Open Platform](https://open.shopee.com)
2. Create an affiliate app and get credentials
3. Save credentials via `POST /api/v1/user/market-credential`

## 🔒 Security

- Passwords are hashed using AES-256 encryption
- JWT tokens for stateless authentication
- Session cookies with HTTPOnly flag
- CORS configuration for cross-origin requests
- SQL injection prevention via GORM parameterized queries

## 📊 Database Schema

Main entities:
- **Users** - User accounts and credentials
- **Products** - Imported marketplace products
- **Offers** - Product offers from marketplaces
- **Campaigns** - Marketing campaigns
- **Links** - Generated affiliate links
- **Clicks** - Click tracking records
- **MarketplaceCredentials** - User's API credentials

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Environment Variables

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=your-password
DB_NAME=affiliate

# Security
JWT_SECRET=your-jwt-secret
PASSWORD_SECRET=your-32-byte-password-secret

# Server
HTTP_HOST=0.0.0.0
HTTP_PORT=8080
```

## 📄 License

This project is licensed under the MIT License.

## 🙋 Support

For support and questions:
- Open an issue on GitHub
- Check Swagger documentation
- Review [SWAGGER.md](./SWAGGER.md) for API details

## 🎯 Roadmap

- [ ] Support for more marketplace platforms
- [ ] Advanced analytics and reporting
- [ ] Webhook notifications for clicks
- [ ] Bulk link generation
- [ ] Rate limiting and API quotas
- [ ] Admin dashboard
- [ ] Multi-language support
