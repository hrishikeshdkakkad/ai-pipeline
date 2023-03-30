FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-asrlatest.json ./mockoon-asrlatest.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-asrlatest.json
USER mockoon

EXPOSE 3008

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-asrlatest.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test