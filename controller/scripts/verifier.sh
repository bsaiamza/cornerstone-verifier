#!/bin/bash
docker stop cornerstone-verifier-agent
docker rm cornerstone-verifier-agent

docker run -p 10000:10000 -p 10001:10001  --name cornerstone-verifier-agent --network=bridge -it bcgovimages/aries-cloudagent:py36-1.16-1_0.7.3 start \
-l 'Iamza Cornerstone Verifier' \
-it http 0.0.0.0 10000 \
-ot http \
--admin 0.0.0.0 10001 \
--admin-insecure-mode \
-e http://172.17.0.2:10000/ \
--genesis-url http://172.20.0.1:9000/genesis \
--seed IamzaCornerstoneVerifier00000000  \
--wallet-type indy \
--wallet-name verifier_wallet \
--wallet-key verifier_secret \
--log-level 'info'  \
--auto-provision \
--auto-accept-invites \
--auto-accept-requests \
--auto-ping-connection 