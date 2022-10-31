#!/bin/sh

gcloud compute start-iap-tunnel bbgame-mdb805-dev 3306 --local-host-port=localhost:3306 --zone=asia-east1-a --project=drd-project-201807
