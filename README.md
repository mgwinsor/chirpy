# chirpy

This is a lightweight backend for a web service similar to twitter (X). But
instead of tweets, users post chirps!

Some key features:
- User authentication management through JWTs and refresh tokens
- Create, read, and delete chirps
- Basic metrics (hit count)

Please note that this is not a production server.

## Installation

Install the binary to your machine with `go install`.

```bash
go install github.com/mgwinsor/chirpy@latest
```

You will also need a PostgeSQL server running locally. Be sure to run the goose
migrations on the database before starting the server!

```bash
goose postgres "postgres://user@localhost:5432/chirpy" up
```

An `.env` file will also be required for the server to run. It must be placed
in the same location as the chirpy binary. The following variables need to be
defined:
- `DB_URL`: the URL for the PostgeSQL database
    - Example: `"postgres://user@localhost:5432/chirpy?sslmode=disable"`
- `PLATFORM`: the environment where the code runs (recommended value is `dev`)
- `JWT_SECRET`: the secret key for generating the JWT tokens
    - You can generate one with `openssl rand -base64 64`
- `POLKA_KEY`: the key for the imaginary Polka webhook

## Usage

There are different endpoints available:

- `/api/healthz` ([GET](./docs/api_endpoints.md#get-apihealthz))
- `/api/polka/webhooks` ([POST](./docs/api_endpoints.md/post-apipolkawebhooks))
- `/api/chirps` ([POST](./docs/api_endpoints.md/post-apichirps),
[GET](./docs/api_endpoints.md/get-apichirps),
[GET](./docs/api_endpoints.md#get-apichirpschirpid),
[DELETE](./docs/api_endpoints.md#delete-apichirpschirpid))
- `/api/users` ([POST](./docs/api_endpoints.md#post-apiusers),
[PUT](./docs/api_endpoints.md#put-apiusers))
- `/api/login` ([PUT](./docs/api_endpoints.md#post-apilogin),
[POST](./docs/api_endpoints.md#put-apilogin))
- `/api/refresh` ([POST](./docs/api_endpoints.md#post-apirefresh))
- `/api/revoke` ([POST](./docs/api_endpoints.md#post-apirevoke))
- `/admin/metrics` ([GET](./docs/api_endpoints.md#get-apimetrics))
- `/admin/reset` ([GET](./docs/api_endpoints.md#get-apireset))
