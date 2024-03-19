### Description
This app is a link between Suricata and TheHive, whose job is to forward alerts to TheHive dashboard.
### Requirements
Suricata should have EVE JSON log output configured in the redis channel. 
To configure the app use env variables such as "REDIS_HOST, REDIS_PASSWORD, REDIS_CHANNEL, THEHIVE_API_HOST, THEHIVE_API_TOKEN".
### Example of usage
Here is example of how it work for [test alert](https://docs.suricata.io/en/suricata-6.0.9/quickstart.html#alerting)
![Screenshot_5](https://github.com/akella44/TheHiveSender/assets/61851015/bcf76fa8-4745-44a8-bdd3-7bff924c4657)
