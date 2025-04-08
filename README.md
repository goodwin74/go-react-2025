# Go 1.24.2 and React 19 Starter Kit (template empty) with Dev mode(HMR) Vite
# Стартовый шаблон для Go 1.24.2 (как backend) и React 19 (как frontend) с режимом HMR от сборщика Vite

The starter pack includes everything you need to start a REST API:
- Controllers
- Routing
- Basic functions for working with PostgreSQL
- Integration with React
  
---

**Remember: after downloading the project, install dependencies for react and vite.**\
**To do this, run the command `npm install` in the /web folder.**

## Run in dev mode for testing (with HMR from Vite)

Set APPMODE=dev in .env file 

Run in terminal 1 (in "/web" dir):\
`npm run dev`

Run in terminal 2 (in root dir):\
`go run cmd/main.go`

## Build & Run in prod mode

Set APPMODE=prod in .env file 

Run in terminal for build (step-by-step):\
`cd web`\
`npm run build`\
`cd ..`\
`go build cmd/main.go`\
Run the compiled Go project
`./main.exe`\

