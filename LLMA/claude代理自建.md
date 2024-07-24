```
version: "3.8"
services:
  fuclaude:
    image: pengzhile/fuclaude:latest
    container_name: Fuclaude
    restart: always
    ports:
      - "14300:8181"
    environment:
      - FUCLAUDE_SIGNUP_ENABLED=true
```
