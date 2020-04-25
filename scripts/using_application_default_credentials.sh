#!/bin/bash

gcloud auth application-default login \
  --client-id-file=client_secret.json \
  --scopes \
https://www.googleapis.com/auth/gmail.labels,\
https://www.googleapis.com/auth/gmail.settings.basic
