FROM concourse/busyboxplus:git

# satisfy go crypto/x509
RUN curl -s https://pki.goog/roots.pem | cat - /etc/ssl/certs/*.pem > /etc/ssl/certs/ca-certificates.crt

ADD assets/ /opt/resource/
