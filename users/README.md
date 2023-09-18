# Courtesy

https://github.com/golang-standards/project-layout - Community standard project directory
https://go.dev/talks/2014/names.slide#1 - Proper naming in Go project
https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2 - Explaination about
https://go.dev/doc/effective_go#names - Effective go, useful official Go documentation

## Dependency
`sudo snap install sqlc` - Install sqlc CLI
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate - Install database migration

## Development Notes
### Docker
1. Translate docker cli to docker compose file https://www.composerize.com/
2. Explaination
```
version: '3'
services:
  db:
    container_name: users_db
    image: postgres:latest
    restart: always
    env_file:
      - .env
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data/
    networks:
      - solo-project
    healthcheck:
      test: ["CMD-SHELL", "sh -c 'pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB}'"]
      interval: 10s
      timeout: 3s
      retries: 3

volumes:
  postgres_data:

networks:
  solo-project:
    driver: bridge  
```
`version` is current docker compose version
`services` define service container
`db` the default container name
`container_name` naming the container, optional
`restart` 
`env_file` get the .env file instead of defining in docker compose
`ports` `host:container` define container ports
`volumes` mounting directory from host to container
`networks` define container networks, necessary if you want to make sure container interact within same networks
`healthcheck` checking the container status
`test` exec command line to check the container
`interval` interval to check health status
`timeout` duration of waiting
`retries` times how status will check
