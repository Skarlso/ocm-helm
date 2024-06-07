# Build the component

ocm add componentversions --create  components.yaml -f

# Transfer

ocm transfer component ./transport-archive ghcr.io/skarlso/helm-test -f