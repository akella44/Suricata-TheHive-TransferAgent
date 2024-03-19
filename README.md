### Description
This app is a link between Suricata and TheHive, whose job is to forward alerts to TheHive dashboard.
#### Requirements
Suricata should have EVE JSON log output configured in the redis channel. 
To configure the app use env variables such as "REDIS_HOST, REDIS_PASSWORD, REDIS_CHANNEL, THEHIVE_API_HOST, THEHIVE_API_TOKEN".
