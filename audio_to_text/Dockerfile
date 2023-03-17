FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-audio-to-text.json ./mockoon-audio-to-text.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-audio-to-text.json
USER mockoon

EXPOSE 3004

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-audio-to-text.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test