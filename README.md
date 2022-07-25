# Squeeze APIs

## Quickstart

1. create `deployments/.env` (from [temple](deployments/.env.sample))

2. run docker compose

   ```bash
   cd deployments && docker compose up -d
   ```

## Services

1. Poller (read smart contract, update redis)
2. API (go-fiber)
   - [GET] /api/topics
     - sort: count
     - filter: poll, question
   - [GET] /api/topic/:address
