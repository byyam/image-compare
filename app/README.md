


nginx config:

``` shell
server{
  listen 18080;
  server_name localhost;

  location / {
    root <app-path>;
    index index.html;
  }
}
```


