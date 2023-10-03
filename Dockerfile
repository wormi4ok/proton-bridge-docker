FROM nginx:bookworm

LABEL org.opencontainers.image.source="https://github.com/wormi4ok/proton-bridge-docker" \
      org.opencontainers.image.title="Proton Mail Bridge" \
      org.opencontainers.image.desctiption="Docker image to run Proton Mail Bridge" \
      org.opencontainers.image.authors="wormi4ok"

WORKDIR /protonmail

COPY gpgparams install.sh VERSION ./

RUN bash install.sh

EXPOSE 25/tcp
EXPOSE 143/tcp
COPY nginx.conf /etc/nginx/nginx.conf

COPY entrypoint.sh .

ENTRYPOINT ["bash", "/protonmail/entrypoint.sh"]
