# Library Management System

The Library Management System is a comprehensive solution designed to manage library operations efficiently. It facilitates the registration and login of administrators, allows users to browse a catalog of books, request book issuances, and track their issued books.

#### NOTE: 
This project was given as an assignment to SDSLabs and other sister Clubs. 

#### Future Enhancements

- **Containerization**: I plan to dockerize the application to streamline deployment processes and ensure consistency across different environments. This will facilitate easier scaling, development, and testing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Apache2
- Go Lang
- Go Lang Migrate tool

### Installation

1. **Clone the repository**

   Start by cloning the repository to your local machine.

   ```sh
   git clone <repository-url>
   cd <repository-name>
   ```

2. **Set up Go modules**
   Install the project dependencies using Go modules.
   ```sh
   go mod vendor
   go mod tidy 
   ```
3. **Environment Configuration**
   Copy the example environment configuration file to create your own `.env` file.
   ```sh
   cp .env.example .env
   ```
### Database Setup
1. Install Go Lang Migrate tool. You can find installation instructions for your specific OS at [Go Lang Migrate GitHub](https://github.com/golang-migrate/migrate).
   
2. To apply migrations to your database, run:
   ```
   migrate -path database/migration -database "yourdatabaseurl" up
   ```
   Replace `"yourdatabaseurl"` with your actual database connection string. This command applies all available migrations.

3. To roll back the last migration, use:
   ``` 
   migrate -path database/migration -database "yourdatabaseurl" down
   ```
   This command rolls back the most recently applied migration.

Ensure you have your database running and accessible through the connection string you provide to the migrate tool.

## Development

### Live Reloading

For a better development experience with live reloading, we use `nodemon` to monitor changes in Go files and automatically restart the server. This setup requires `nodemon` to be installed globally on your machine. If you haven't installed `nodemon`, you can do so by running:

```sh
npm install -g nodemon
```
Once `nodemon` is installed, you can start the development server with live reloading enabled by running:
```sh
make server
```
This command watches for changes in all `.go` files within the project directory and subdirectories. If a change is detected, `nodemon` will gracefully restart the Go application by sending a `SIGTERM` signal and then executing go run `cmd/main.go`.

Ensure you have all necessary Go dependencies installed and your `.env` file configured before running the server.

## Hosting and running the application
1. Install apache2: `sudo apt install apache2`
2. `sudo a2enmod proxy proxy_http`
3. `cd /etc/apache2/sites-available`
4. `sudo nano mvc.sdslabs.local.conf`
Add: 
```
<VirtualHost *:80>
	ServerName mvc.sdslabs.local
	ServerAdmin youremailid
	ProxyPreserveHost On
	ProxyPass / http://127.0.0.1:8000/
	ProxyPassReverse / http://127.0.0.1:8000/
	TransferLog /var/log/apache2/mvc_access.log
	ErrorLog /var/log/apache2/mvc_error.log
</VirtualHost>
```
5. `sudo a2ensite mvc.sdslabs.local.conf`
6. add `127.0.0.1	mvc.sdslabs.local` to `/etc/hosts`
7. `sudo a2dissite 000-default.conf`
8. `sudo apache2ctl configtest `
9. `sudo systemctl restart apache2`
10. `sudo systemctl status apache2`
11. Check `mvc.sdslabs.local` on your browser


## Features

- **Admin Registration and Login**: Navigate to `/admin/register` to register as an admin and `/admin/login` to log in.
- **User Registration and Login**: The home page provides options for user sign-up and login.
- **Book Catalog**: Once logged in, users can browse the catalog of books.
- **Book Issuance Requests**: Users can request the issuance of books from the admin.
- **Track Issued Books**: Users can view the books they have issued.