FROM node:14-alpine

RUN npm install -g @mockoon/cli@2.4.0
COPY mockoon-speakerclassifier.json ./mockoon-speakerclassifier.json

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mockoon-speakerclassifier.json
USER mockoon

EXPOSE 3007

ENTRYPOINT ["mockoon-cli", "start", "--daemon-off", "--data", "mockoon-speakerclassifier.json", "--container"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test