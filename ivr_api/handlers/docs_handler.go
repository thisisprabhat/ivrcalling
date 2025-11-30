package handlers

import (
	"github.com/gin-gonic/gin"
)

// DocsHandler serves the Scalar API documentation
func DocsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>IVR Calling System API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>
        body {
            margin: 0;
            padding: 0;
        }
    </style>
</head>
<body>
    <script id="api-reference" data-url="/docs/swagger.yaml"></script>
    <script>
        var configuration = {
            theme: 'purple',
            layout: 'modern',
            showSidebar: true,
            darkMode: true,
            metadata: {
                title: 'IVR Calling System API',
                description: 'Comprehensive API documentation for the IVR calling system',
                favicon: 'data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><text y=".9em" font-size="90">📞</text></svg>'
            },
            searchHotKey: 'k',
            hideModels: false,
            hideDownloadButton: false,
            authentication: {
                preferredSecurityScheme: 'ApiKeyAuth',
                apiKey: {
                    token: ''
                }
            }
        }

        document.getElementById('api-reference').dataset.configuration = JSON.stringify(configuration)
    </script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>`
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html)
	}
}

// SwaggerYAMLHandler serves the OpenAPI/Swagger YAML file
func SwaggerYAMLHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File("./docs/swagger.yaml")
	}
}
