# Squeeze APIs

## Quickstart

1. create `.env` (from [temple](.env.sample))

2. run docker compose

   ```bash
   docker compose up -d
   ```

## Services

1. Poller (read smart contract, update redis)
2. API (go-fiber)
   - [GET] /api/topics
   - [GET] /api/topic/:address
