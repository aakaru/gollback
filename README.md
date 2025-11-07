# n8n Backup Tool

A Go-based tool to backup n8n workflows with version control.

## Setup

1. Start n8n: `docker compose up -d`
2. Access n8n: http://localhost:5678
3. Run backup tool: `go run main.go`

## Features

- Automatic workflow backup
- Version tracking
- Easy restore
