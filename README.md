# URL Shortener Service
This application provides a URL shortening service through an HTTP endpoint. It accepts long URLs, shortens them, and redirects users from short URLs to the original URLs. The application is built with Go and uses in-memory storage to manage URL mappings.

### Prerequisites:
go
docker

### Getting Started:
#### Clone the Repository:
```bash
git clone https://github.com/HemanthSanju/URL_Shortner.git
```
```bash
cd URL_Shortner
```

#### Build the Docker Image for the application: 
docker build -t urlshortener .

#### Run the URL shortener application in a Docker container, exposing port 8080 for accessing the API:
docker run -p 8080:8080 urlshortener

#### Send an HTTP GET request to the shorten endpoint with the URL as a query parameter:
curl "http://localhost:8080/shorten?url=http://example.com"

## Metrics
curl http://localhost:8080/metrics