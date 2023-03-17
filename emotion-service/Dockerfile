FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-emotions-model.json ./mockoon-emotions-model.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-emotions-model.json
USER mockoon

EXPOSE 3002

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-emotions-model.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test