# pienso app

A go http serve built to serve personal blog content using Google Cloud Storage as the file repository.

After the blog has been launched the first post will detail my experience, gotchas, and other tips.

## Starting the Server

`go run app/main.go`

### Testing
curl localhost:3333/

### RoadMap
- Build project
- Simple Markdown File Parser
- LightWeight Front End (Choose a Framework like Vue.js or React) 
- Build out API endpoints
- Deploy server
- Host frontend as Isomorphic Application

### External Depenencies:
This excercise is mainly to demonstrate and learn how to build a go package using a minimal amount
of external dependencies. That being said, the purpose of the project is not to reinvent the wheel either.
Lightweight http router: github.com/go-chi/chi
Easy access to Google Cloud Storage: cloud.google.com/go/storage