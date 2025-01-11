jwt
===
jwt is a command line tool for generating and parsing JWT tokens in Go.


### Install:

```shell
go install github.com/qshuai/jwt@latest
```

### Usage:

```
$ jwt --help
A command-line tool for signing and parsing JWT

Usage:
  jwt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  keygen      Generate private key and save to the file specified. Support ECDSA and RAS.
  parse       Parse JWT and output formated information to console
  sign        Sign JWT using alg, secret and claims

Flags:
  -h, --help   help for jwt

Use "jwt [command] --help" for more information about a command.
```

Please refer to the command-line help documentation for this tool, as it can address most of your questions. If you still have any uncertainties, feel free to submit an issue.

### Examples: 

- Sign and Parse JWT token using HMAC sign method:

  ```shell
  $ jwt sign --alg HS256 --secret=helloworld --claims iss:somecorp --claims sub:usercenter --claims exp:1736754916000 --claims iat:1736554948000
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoic29tZWNvcnAiLCJzdWIiOiJ1c2VyY2VudGVyIn0.nh794O0Sox53dcVeMZJgA5bjJUS2VPqWLxwtJnkAyKQ
  
  $ jwt parse -s helloworld eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoic29tZWNvcnAiLCJzdWIiOiJ1c2VyY2VudGVyIn0.nh794O0Sox53dcVeMZJgA5bjJUS2VPqWLxwtJnkAyKQ
  Token Validation: true
  HEADER:
  	"alg": HS256
  	"typ": JWT
  
  PAYLOAD:
  	"iss": somecorp
  	"sub": usercenter
  	"exp": 1736754916000
  	"iat": 1736554948000
  ```

- Sign and Parse JWT token using ECDSA sign method:

  ```shell
  $ jwt keygen ecdsa --key-file=private-key.ecdsa.pem --bit-size=256
  $ ls 
  private-key.ecdsa.pem
  
  $ jwt sign --alg ES256 --key-file private-key.ecdsa.pem --claims iss:jidu --claims sub:usercenter --claims exp:1736754916000 --claims iat:1736554948000
  eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoiamlkdSIsInN1YiI6InVzZXJjZW50ZXIifQ.omxqh5Dhl5hgP1A2PFbiBMAvDpLeRlDm2E1ZOYFPcOoOkg1ZV-1Ks-bwv-ce9EYgW1GAWNWuMacyYCZD9zDLyw
  
  $ jwt parse --key-file=private-key.ecdsa.pem eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoiamlkdSIsInN1YiI6InVzZXJjZW50ZXIifQ.omxqh5Dhl5hgP1A2PFbiBMAvDpLeRlDm2E1ZOYFPcOoOkg1ZV-1Ks-bwv-ce9EYgW1GAWNWuMacyYCZD9zDLyw
  Token Validation: true
  HEADER:
  	"alg": ES256
  	"typ": JWT
  
  PAYLOAD:
  	"iat": 1736554948000
  	"iss": jidu
  	"sub": usercenter
  	"exp": 1736754916000
  ```

- Sign and Parse JWT token using RSA sign method:

  ```shell
  $ jwt keygen rsa --key-file=private-key.rsa.pem --bit-size=2048
  $ ls
  private-key.rsa.pem
  
  $ jwt sign --alg RS256 --key-file private-key.rsa.pem --claims iss:jidu --claims sub:usercenter --claims exp:1736754916000 --claims iat:1736554948000
  eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoiamlkdSIsInN1YiI6InVzZXJjZW50ZXIifQ.XXVDjnVVXzLYrpz0v0ByxnMP5-s0asJKhKuFcG3YezLWrKqo8C3NpGytnve57cPG76D331TqsKgZv-YJJ8NcudgmdHxooH6reQ5gCzI32EEj18XTOCkaKTSRjV09lbD1vKqqWXAyyfefyP0t5S6OCGQ66EJS5SN08boNGaysY5BtaPvkL__kqFt4kojTVxJrBemhyI0WBOc7SJhbnPSrX2NXT0-Y2suwZakxx5UWGv5oupezT7fkr0zhAnrkYQc25D7jo0utjhC2BgNnvkRg_QV-ii2HntQSvnn2Qf00EZGW7i_DvYlNEFNUMZFNPYPWczN7q2DvIHK6aLCM26vILg
  
  $ jwt parse --key-file=private-key.rsa.pem eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NTQ5MTYwMDAsImlhdCI6MTczNjU1NDk0ODAwMCwiaXNzIjoiamlkdSIsInN1YiI6InVzZXJjZW50ZXIifQ.XXVDjnVVXzLYrpz0v0ByxnMP5-s0asJKhKuFcG3YezLWrKqo8C3NpGytnve57cPG76D331TqsKgZv-YJJ8NcudgmdHxooH6reQ5gCzI32EEj18XTOCkaKTSRjV09lbD1vKqqWXAyyfefyP0t5S6OCGQ66EJS5SN08boNGaysY5BtaPvkL__kqFt4kojTVxJrBemhyI0WBOc7SJhbnPSrX2NXT0-Y2suwZakxx5UWGv5oupezT7fkr0zhAnrkYQc25D7jo0utjhC2BgNnvkRg_QV-ii2HntQSvnn2Qf00EZGW7i_DvYlNEFNUMZFNPYPWczN7q2DvIHK6aLCM26vILg
  Token Validation: true
  HEADER:
  	"alg": RS256
  	"typ": JWT
  
  PAYLOAD:
  	"exp": 1736754916000
  	"iat": 1736554948000
  	"iss": jidu
  	"sub": usercenter
  ```
