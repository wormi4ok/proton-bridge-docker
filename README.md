# Proton Mail Bridge in a Docker container

This is yet another unofficial Docker image to run [Proton Mail Bridge](https://github.com/ProtonMail/proton-bridge).
Use it at your own risk.
Inspired by [shenxn/protonmail-bridge-docker](https://github.com/shenxn/protonmail-bridge-docker)

## How to start

Initialize Proton Mail Bridge

```bash
docker compose run --rm protonmail-bridge init
login # enter your credentials
# optionally, switch to split mode if you have more that one email address
# in split mode every email address has own local credentials
# change mode 0
info 0 # View information to connect to your accounts
exit
```

Start Proton mail bridge in the background:

```bash
docker compose up -d
```

It is intended to in a local secure network, so default proxy config doesn't use encryption.
To enable encryption, you can mount a custom `nginx.conf` file to the [protonmail-bridge](docker-compose.yml) container
at `/etc/nginx/nginx.conf` and follow [the official guide from NGINX][] to configure TLS/SSL.

[the official guide from NGINX]: https://docs.nginx.com/nginx/admin-guide/mail-proxy/mail-proxy/#setting-up-ssltls-for-a-mail-proxy
