# uptime-dash

Free, open source, self-hosted Uptime Monitoring for everyone

## Local development (Docker)

```sh
docker-compose up -d
docker exec -it timescaledb psql -U postgres
\c uptime_dash
\d
```
