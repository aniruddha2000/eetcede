## eetcede

### Add the SSL cerificate

```bash
$ openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
```

```bash
$ openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
```

### Run the project

```bash
$ go buiild

# To run the server in in memory storage mode
$ ./eetcede

# To run the server with disk storage mode 
$ ./eetcede -storage-type=disk
```
