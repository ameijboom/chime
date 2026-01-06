# Chime

A Discord bot for managing Free Company activities with a web dashboard.

## Structure

- **bot/** - Discord bot (TypeScript, discord.js)
- **api/** - Backend API (Rust, Axum)
- **app/** - Web dashboard (SvelteKit)

## Setup

### Bot

```bash
cd bot
cp .env.example .env
pnpm install
pnpm dev
```

### API

```bash
cd api
cargo run
```

### Frontend

```bash
cd app
pnpm install
pnpm dev
```

## Environment Variables

See `bot/.env.example` for required Discord credentials.

## License

MIT
