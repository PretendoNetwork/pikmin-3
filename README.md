# PIKMIN 3 (Wii U) replacement server

Server includes both the authentication and secure servers

## Configuration
All configuration options are handled via environment variables

| Name                                    | Description                                                                                                            | Required                                      |
|-----------------------------------------|------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `PN_PIKMIN3_POSTGRES_URI`               | Fully qualified URI to your Postgres server (Example `postgres://username:password@localhost/pikmin3?sslmode=disable`) | Yes                                           |
| `PN_PIKMIN3_KERBEROS_PASSWORD`          | Password used as part of the internal server data in Kerberos tickets                                                  | No (Default password `password` will be used) |
| `PN_PIKMIN3_AUTHENTICATION_SERVER_PORT` | Port for the authentication server                                                                                     | Yes                                           |
| `PN_PIKMIN3_SECURE_SERVER_HOST`         | Host name for the secure server (should point to the same address as the authentication server)                        | Yes                                           |
| `PN_PIKMIN3_SECURE_SERVER_PORT`         | Port for the secure server                                                                                             | Yes                                           |
| `PN_PIKMIN3_ACCOUNT_GRPC_HOST`          | Host name for your account server gRPC service                                                                         | Yes                                           |
| `PN_PIKMIN3_ACCOUNT_GRPC_PORT`          | Port for your account server gRPC service                                                                              | Yes                                           |
| `PN_PIKMIN3_ACCOUNT_GRPC_API_KEY`       | API key for your account server gRPC service                                                                           | No (Assumed to be an open gRPC API)           |