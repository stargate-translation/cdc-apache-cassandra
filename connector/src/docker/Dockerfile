ARG LUNASTREAMING_IMAGE
ARG LUNASTREAMING_VERSION

# Install pulsar cassandra source connector
FROM ${LUNASTREAMING_IMAGE}:${LUNASTREAMING_VERSION}
ARG BUILD_VERSION
COPY pulsar-cassandra-source-${BUILD_VERSION}.nar /pulsar/connectors/

# Install pulsar http sink connector
USER root
RUN apt-get clean && apt-get update && apt-get install -y wget
RUN wget https://archive.apache.org/dist/pulsar/pulsar-3.0.0/connectors/pulsar-io-http-3.0.0.nar -O /tmp/pulsar-io-http-3.0.0.nar
RUN mv /tmp/pulsar-io-http-3.0.0.nar /pulsar/connectors
USER 10000
