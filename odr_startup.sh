#!/kaniko/bin/sh

# Add any extra certificates in the environment variable EXTRA_CERTS to the certs store
echo "${EXTRA_CERTS}" >> /kaniko/ssl/certs/ca-certificates.crt

/kaniko/waypoint $@
