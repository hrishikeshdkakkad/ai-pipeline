FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-slp-keyword.json ./mockoon-slp-keyword.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-slp-keyword.json
USER mockoon

EXPOSE 3003

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-slp-keyword.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test