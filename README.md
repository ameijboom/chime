# Chime

A Discord bot application for managing and scheduling gaming events with seamless integration between Discord and a web interface.

## Overview

Chime helps gaming communities organize events by providing event management tools directly within Discord, backed by a modern web interface. Track attendees, manage party compositions, and keep your events organized.

## Features

- **Discord Slash Commands** - Create and manage events directly from Discord
- **Event Scheduling** - Schedule events with support for different party types (Light/Full)
- **Attendee Tracking** - Track members joining events with their job roles (Tank, Healer, DPS)
- **Party Size Management** - Automatic enforcement of party size limits (4 for Light parties, 8 for Full parties)
- **Event Status Tracking** - Monitor event states (Ongoing, Ended, On Hold)
- **Web Dashboard** - Modern web interface for event management (coming soon)

## Architecture

Chime is built with a three-tier architecture:

```
Frontend (SvelteKit) ←→ Backend (Go) ←→ Discord API
                            ↓
                      SQLite Database
```

- **Backend**: Go application using `discordgo` for Discord integration and gRPC for API communication
- **Frontend**: SvelteKit with TypeScript and Tailwind CSS
- **Communication**: gRPC with Protocol Buffers for type-safe client-server communication
- **Database**: SQLite with GORM for persistence

## Tech Stack

### Backend
- Go 1.24.4
- [discordgo](https://github.com/bwmarrin/discordgo) - Discord API client
- gRPC & Protocol Buffers
- GORM with SQLite

### Frontend
- SvelteKit
- TypeScript
- Tailwind CSS
- Vite

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Node.js 18+ and pnpm
- Discord Bot Token and Application ID

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run cmd/server/main.go \
  --token YOUR_DISCORD_BOT_TOKEN \
  --guild YOUR_GUILD_ID \
  --database chime.db \
  --grpc-port :50051
```

#### Available Flags

- `--token` - Discord bot token (required)
- `--guild` - Guild ID for command registration (required)
- `--database` - Database file path (default: `chime.db`)
- `--grpc-port` - gRPC server port (default: `:50051`)
- `--remove-commands` - Remove all commands on shutdown

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
pnpm install
```

3. Start the development server:
```bash
pnpm dev
```

The web interface will be available at `http://localhost:5173`

## Project Structure

```
chime/
├── backend/                    # Go backend application
│   ├── cmd/server/            # Application entry point
│   ├── internal/
│   │   ├── bot/              # Discord bot logic
│   │   │   ├── commands/     # Slash command definitions
│   │   │   └── handlers/     # Event handlers
│   │   ├── repository/       # Database layer
│   │   └── server/           # gRPC server
│   └── generated/            # Generated protobuf code
├── frontend/                  # SvelteKit application
│   └── src/
│       ├── routes/           # Page routes
│       └── lib/              # Shared components
├── proto/                    # Protocol Buffer definitions
│   └── events.proto
└── state.db                  # SQLite database
```

## Data Models

### Event
- Guild Event ID (Discord)
- Message ID (Discord)
- Status (Ongoing, Ended, On Hold)
- Party Type (Light, Full)
- Attendees

### Attendee
- User ID (Discord)
- Event ID
- Job Role (Tank, Healer, DPS)

## Development Status

Chime is currently in active development. The core Discord bot functionality and database layer are implemented, with the web interface coming soon.

### Recent Updates
- ✅ Base bot structure and gRPC server
- ✅ Event creation and management
- ✅ Attendee tracking with role assignment
- ✅ Database persistence with SQLite
- 🚧 Web dashboard interface (in progress)

## Contributing

This is a personal project, but suggestions and feedback are welcome!
