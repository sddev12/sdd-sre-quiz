# Local MongoDB for SRE Quiz

This directory contains resources to run a local MongoDB instance for development and testing.

## Usage

1. **Start MongoDB:**
   ```bash
   docker-compose up -d
   ```
2. **Stop MongoDB:**
   ```bash
   docker-compose down
   ```

- The database will be available at `mongodb://localhost:27017/sre_quiz`.
- You can add `.js` or `.sh` scripts to the `init/` directory to initialize collections or seed data.

## Data Persistence

- MongoDB data is stored in a Docker volume (`mongo_data`) and will persist between restarts.
