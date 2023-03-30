FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-speechactivityrecognition.json ./mockoon-speechactivityrecognition.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-speechactivityrecognition.json
USER mockoon

EXPOSE 3009

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-speechactivityrecognition.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test