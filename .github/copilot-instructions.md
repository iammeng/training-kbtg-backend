# GitHub Copilot Instructions for Training KBTG Backend

## Project Overview
This is a Go-based RESTful API backend service built with Fiber framework, featuring JWT authentication, SQLite database, and comprehensive profile management. The project follows clean architecture principles with proper separation of concerns.

## Code Style Guidelines

### Go Language Standards
- Follow standard Go conventions and idioms
- Use `gofmt` formatting
- Prefer explicit error handling over panic
- Use meaningful variable and function names
- Keep functions small and focused (single responsibility)
- Add godoc comments for exported functions and types

### Project Structure
```
├── main.go                 # Main application entry point
├── models/                 # Data models and DTOs
├── handlers/               # HTTP request handlers
├── middleware/             # HTTP middleware functions
├── database/               # Database connection and migrations
├── docs/                   # Auto-generated Swagger documentation
└── detail.md              # Technical documentation
```

### Naming Conventions
- **Files**: Use snake_case (e.g., `user_profile.go`)
- **Functions**: Use camelCase (e.g., `getUserProfile`)
- **Constants**: Use UPPER_SNAKE_CASE (e.g., `JWT_SECRET`)
- **Structs**: Use PascalCase (e.g., `UserProfile`)
- **Variables**: Use camelCase (e.g., `userID`)

## Database Guidelines

### GORM Best Practices
- Always use struct tags for database fields
- Include `json:"-"` for sensitive fields like passwords
- Use soft deletes with `gorm.DeletedAt`
- Include `created_at` and `updated_at` timestamps
- Use unique indexes for email and membership_id fields

### Example Model Structure
```go
type User struct {
    ID            uint           `gorm:"primarykey" json:"id"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
    Email         string         `gorm:"uniqueIndex;not null" json:"email"`
    Password      string         `gorm:"not null" json:"-"`
    // ... other fields
}
```

## API Development Standards

### Handler Functions
- Always validate input data before processing
- Use proper HTTP status codes (200, 201, 400, 401, 404, 409, 500)
- Return consistent JSON error responses
- Include comprehensive Swagger documentation comments

### Error Response Format
```go
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    "error": "Descriptive error message",
})
```

### Success Response Format
```go
// For single resource
return c.JSON(models.UserResponse{
    User: user,
})

// For simple messages
return c.JSON(fiber.Map{
    "message": "Success message",
})
```

## Security Guidelines

### Authentication & Authorization
- Always hash passwords using bcrypt with appropriate cost
- Use JWT tokens with reasonable expiration times (24 hours)
- Validate JWT tokens in middleware before accessing protected routes
- Never log or expose sensitive information (passwords, tokens)

### Input Validation
- Validate all input data (email format, password length, required fields)
- Sanitize input to prevent injection attacks
- Use proper struct tags for validation

## Swagger Documentation Standards

### API Documentation Comments
```go
// FunctionName godoc
// @Summary Brief description of the endpoint
// @Description Detailed description of what the endpoint does
// @Tags TagName
// @Accept json
// @Produce json
// @Param paramName body/path/query models.RequestType true "Parameter description"
// @Success 200 {object} models.ResponseType
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /endpoint/path [method]
```

### Required Documentation Elements
- Always include @Summary and @Description
- Specify correct @Tags for grouping
- Document all parameters with proper types
- Include all possible response codes
- Add @Security BearerAuth for protected endpoints

## Middleware Development

### JWT Middleware Pattern
```go
func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extract and validate token
        // Set user context
        c.Locals("user_id", userID)
        c.Locals("email", email)
        return c.Next()
    }
}
```

### Middleware Guidelines
- Keep middleware functions focused and reusable
- Always call `c.Next()` for successful cases
- Return early for validation failures
- Use `c.Locals()` to pass data between middleware and handlers

## Testing Considerations

### API Testing
- Test all endpoints with valid and invalid data
- Verify authentication and authorization
- Test error cases and edge conditions
- Use proper HTTP methods and status codes

### Example Test Cases
- Register with valid/invalid email formats
- Login with correct/incorrect credentials
- Access protected routes with/without valid tokens
- Update profile with partial/complete data

## Development Workflow

### Code Organization
1. Define models in `models/` package first
2. Implement database operations
3. Create handlers with proper validation
4. Add middleware for common functionality
5. Generate and update Swagger documentation
6. Test all endpoints thoroughly

### Git Commit Messages
- Use descriptive commit messages
- Include affected components (e.g., "Add profile management API endpoints")
- Mention breaking changes if any
- Reference issues or features being implemented

## Performance Guidelines

### Database Optimization
- Use proper indexes for frequently queried fields
- Implement pagination for list endpoints
- Use select specific fields instead of `SELECT *`
- Consider connection pooling for production

### API Response Optimization
- Return only necessary data fields
- Implement proper caching headers
- Use compression middleware for large responses
- Limit request/response sizes

## Production Readiness

### Environment Configuration
- Use environment variables for configuration
- Never hardcode secrets or credentials
- Implement proper logging levels
- Add health check endpoints

### Security Hardening
- Use strong JWT secrets
- Implement rate limiting
- Add request validation middleware
- Enable CORS with appropriate origins
- Use HTTPS in production

## Code Examples to Follow

### Complete Handler Example
```go
// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user's profile information
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body models.UpdateProfileRequest true "Profile data"
// @Success 200 {object} models.ProfileResponse
// @Failure 400 {object} map[string]string
// @Router /profile [put]
func UpdateProfile(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uint)
    
    var req models.UpdateProfileRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }
    
    // Validation logic here
    
    // Database operations here
    
    return c.JSON(models.ProfileResponse{
        User: user,
    })
}
```

## Common Patterns to Use

### Database Transaction Pattern
```go
tx := database.DB.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Error; err != nil {
    return err
}

// Database operations...

if err := tx.Commit().Error; err != nil {
    tx.Rollback()
    return err
}
```

### Error Handling Pattern
```go
if err != nil {
    log.Printf("Error description: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "User-friendly error message",
    })
}
```

## Documentation Requirements

### Code Comments
- Add comments for complex business logic
- Explain non-obvious code decisions
- Document any workarounds or temporary solutions
- Keep comments up-to-date with code changes

### API Documentation
- Maintain accurate Swagger documentation
- Update README.md with new endpoints
- Include example requests/responses
- Document authentication requirements

## Quality Checklist

Before submitting code, ensure:
- [ ] All functions have appropriate error handling
- [ ] Input validation is implemented
- [ ] Swagger documentation is complete and accurate
- [ ] No sensitive information is logged or exposed
- [ ] Database migrations work correctly
- [ ] All endpoints return proper HTTP status codes
- [ ] JWT authentication works as expected
- [ ] Code follows project naming conventions
- [ ] No hardcoded values or credentials
- [ ] Performance considerations are addressed

## Troubleshooting Common Issues

### Database Issues
- Check database connection and migration status
- Verify GORM model definitions and tags
- Ensure proper foreign key relationships

### Authentication Issues
- Verify JWT token generation and validation
- Check middleware order and implementation
- Validate user permissions and roles

### API Issues
- Confirm proper HTTP methods and routes
- Check request/response serialization
- Verify CORS configuration for frontend integration

Remember to always prioritize security, maintainability, and performance in your code. When in doubt, refer to Go best practices and the existing codebase patterns.
