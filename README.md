# goblog

A sample of Go language web server - Weblog

## External Packages
* [Echo](https://github.com/labstack/echo): A fast go web server framework
* [Gorm](https://gorm.io): An ORM framework

## Internal Packages
The blog is operable with various internal packages. This packages are separated pieces of logic to increase readability and maintainability.

### Core
Core packages consists of main functionalities such as database connection, migrations and static settings.

### Routers
It's the package that responsible for adding the routes to web server and applying the middlewares to requests and responses.

### Models
This package contains all the models in application that will map to a table in database using Gorm.

### Logics
In all the application, only this package allowed to communicate directly with database. It's a business logic layer.

### Controllers
This package is responsible to receiving requests, process them and generate corresponding response.

### Middlewares
Middleware package contains all of required middleware for web server to apply on requests and responses.

### Pipelines
It's that package that responsible to containing blind and independent concurrent pipelines for various of tasks (Concurrent pipeline design pattern)

### Serializers
For more readability and increase control of user inputs, all user inputs that aren't based on a model, will handle with a serializer.

### Utils
Utils package contains general functionalities in application.
