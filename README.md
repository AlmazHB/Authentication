Athentication in Go 

I am using

1.Gin

2.MongoDB

3.JWT

Handlers

	
POST("/v1/api/register", h.register)  //for register user

GET("/v1/api/login", h.login)  //for login from user

GET("/v1/api/welcome", h.authenticate(h.welcome))  // Welecom JWT 

