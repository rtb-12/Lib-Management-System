# mvc-lecture
- Clone the repo. From the root directory:
```
- go mod vendor
- go mod tidy
- cp .env.example .env
```


- MYSQL:
1. `mysql -u root -p` : and enter password
2. Create a new database 'books': `CREATE DATABASE books;`
3. Connect to the database: `USE books;`
4. Import the sql dump file into your database: `mysql -u root -p books < dump.sql`


- Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

- Hosting
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
11. Check mvc.sdslabs.local on your browser

