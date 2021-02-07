# nebula-http-gateway

Gateway to provide a http interface for the Nebula Graph service.

## Build

```bash
$ cd /path/to/nebula-http-gateway
$ make
```

## Run

```bash
$ ./nebula-httpd
```

## Required

- Go 11+
- [beego](https://beego.me/)

## User Guide

#### API Definition
| Name       | Path               | Method |
|------------|--------------------|--------|
| connect    | /api/db/connect    | POST   |
| exec       | /api/db/exec       | POST   |
| disconnect | /api/db/disconnect | POST   |

> Connect API

The requested json body
```json
{
  "username": "user",
  "password": "password",
  "address": "192.168.8.26",
  "port": 9669
}
```
The description of the parameters is as follows.

| Field    | Description                                                                                                                 |
|----------|-----------------------------------------------------------------------------------------------------------------------------|
| username | Sets the username of your Nebula Graph account. Before enabling authentication, you can use any characters as the username. |
| password | Sets the password of your Nebula Graph account. Before enabling authentication, you can use any characters as the password. |
| address  | Sets the IP address of the graphd service.                                                                                  |
| port     | Sets the port number of the graphd service. The default port number is 9669.                                                |

```bash
$ curl -i  -X POST -d '{"username":"user","password":"password","address":"192.168.8.26","port":9669}' http://127.0.0.1:8080/api/db/connect
HTTP/1.1 200 OK
Content-Length: 100
Content-Type: application/json; charset=utf-8
Server: beegoServer:1.12.3
Set-Cookie: nsid=bec2e665ba62a13554b617d70de8b9b9; Path=/; HttpOnly
Set-Cookie: Secure=true; Path=/
Set-Cookie: SameSite=None; Path=/
Date: Fri, 02 Apr 2021 08:49:18 GMT

{
  "code": 0,
  "data": "5e18fa40-5343-422f-84e3-e7f9cad6b735",
  "message": "Login successfully"
}
```

Notice:

The response data nsid "5e18fa40-5343-422f-84e3-e7f9cad6b735" is encoded by HMAC-SH256 encryption algorithm，so it's not the same as what you get from a cookie.
If you connect the graphd service succeed, remember to save the *NSID*, it's important for the *exec api* to execute nGQL.
If you restart the gateway server, all authenticated session will be lost, please noticed.

> Exec API

The requested json body
```json
{
"gql": "show spaces;"
}
```
The request header is necessary to request exec api, cookies must be set to get results.

```bash
$ curl -H "Cookie: SameSite=None; nsid=bec2e665ba62a13554b617d70de8b9b9" -H "nsid: bec2e665ba62a13554b617d70de8b9b9" -X POST -d '{"gql": "show spaces;"}' http://127.0.0.1:8080/api/db/exec
  {
    "code": 0,
    "data": {
      "headers": [
        "Name"
      ],
      "tables": [
        {
          "Name": "nba"
        }
      ],
      "timeCost": 4232
    },
    "message": ""
```

> Disconnect API

```bash
$ curl -X POST http://127.0.0.1:8080/api/db/disconnect
{
  "code": 0,
  "data": null,
  "message": "Disconnect successfully"
```
