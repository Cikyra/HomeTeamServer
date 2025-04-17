# HomeTeamServer

## Getting Started 

Run this on Mac and Windows:
```bash
docker-compose up
```

### Windows
Run in command prompt
```bash
call env.bat
go run main.go
```

### Unix/Linux
```bash
# Run this line if first run
cp .env.example .env

# Run these two lines every time
source .env
go run main.go
```