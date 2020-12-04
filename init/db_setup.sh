#!/usr/bin/env bash

echo 'Creating application user and db'

mongo rest_api \
        --host localhost \
        --port 27027 \
        -u root \
        -p supersecret \
        --authenticationDatabase admin \
        --eval "db.createUser({user: 'ala', pwd: 'secret', roles:[{role:'dbOwner', db: 'rest_api'}]});"

