```
➜  tls git:(audit) openssl genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus
.......................................................................................+++
..................+++
e is 65537 (0x10001)
➜  tls git:(audit) ✗ openssl req -new -x509 -key server.key -out server.crt -days 3650
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [XX]:CN
State or Province Name (full name) []:cq
Locality Name (eg, city) [Default City]:cq
Organization Name (eg, company) [Default Company Ltd]:gkcs
Organizational Unit Name (eg, section) []:gkcs
Common Name (eg, your name or your server's hostname) []:apis
Email Address []:1@1.cn
➜  tls git:(audit) ✗ ls
server.crt  server.key
➜  tls git:(audit) ✗
```