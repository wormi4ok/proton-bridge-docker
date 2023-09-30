FROM nginx:bookworm

WORKDIR /protonmail

COPY gpgparams install.sh VERSION ./

RUN bash install.sh

EXPOSE 25/tcp
EXPOSE 143/tcp
COPY nginx.conf /etc/nginx/nginx.conf

COPY entrypoint.sh .

ENTRYPOINT ["bash", "/protonmail/entrypoint.sh"]
