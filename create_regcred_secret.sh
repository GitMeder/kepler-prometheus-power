#!/bin/bash
set -eux
kubectl create --namespace kepler secret docker-registry gitlab-registry-creds --docker-server=gitlab.lrz.de:5005 --docker-username=regcred --docker-password=$1 --docker-email=johannes.ebke@hm.edu
