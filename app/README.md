


nginx config:

``` shell
upstream test_servers {
  server localhost:3000;
}

server{
  listen 18080;
  server_name localhost;
  keepalive_timeout  70;

  location ^~ /app/ {
    root /home/test/image-compare/;
  }
  error_page  405     =200 $uri;

  location / {
    proxy_pass http://test_servers;
    proxy_read_timeout 300s;

    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

    proxy_http_version 1.1;
  }
}

```


