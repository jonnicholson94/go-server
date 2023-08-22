
# A straightforward server, using net/http, and written in Go

I've put together a basic Go server, that handles some simple requests and mimicking the CRUD actions in a todo app.

It contains three main directories:

- Handlers -> the primary route functions
- Middleware -> some core functionality that should be called on every request
- Utils -> in this case, just the database setup

Behind the scenes, I provisioned a Postgres database, which in reality is only containing one table.

As a result, I'm using a pg driver alongside the typical SQL package in Go.

Overall, I really enjoyed writing this. Compared to my experience writing servers in other languages, Go was super simple and performant. 

When I next build a server from scratch outside of my portfolio, I'll definitely try to use Go!