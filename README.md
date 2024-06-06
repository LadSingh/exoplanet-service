# Exoplanet Microservice

## Description

This microservice manages exoplanets and provides fuel estimation for space voyages.

## Endpoints

- `POST /exoplanets`: Add a new exoplanet.
- `GET /exoplanets`: List all exoplanets.
- `GET /exoplanets/{id}`: Get exoplanet by ID.
- `PUT /exoplanets/{id}`: Update exoplanet.
- `DELETE /exoplanets/{id}`: Delete exoplanet.
- `GET /exoplanets/{id}/fuel?crew_capacity={int}`: Get fuel estimation for a trip to an exoplanet.

## Running the Service

### Using Docker

```sh
docker build -t exoplanet-service .
docker run -p 8080:8080 exoplanet-service


```local run
go mod tidy
go run ./cmd/exoplanet-service
