# Message Aggregation Service

## Overview
Service made as a simple TCP server which accepts messages on address and port 
provided by ENV variables. All messages are grouped by some attribute and send
to adapters, which, in turn, translate it to developers via Slack, EMail etc.

## Documetation
```bash 
godoc -http=:8080
```

## Usage
```golang
	import "github.com/mdspinc/mag/magclient"
	
	err := magclient.Setup("127.0.0.1", 3040)
	if err != nil {
		//...
	}

	magclient.Send("Some string")

```

## Application parts
### Endpoint
Represents TCP server. 

### Handlers
It handels incoming messages based on their type. Each message have header with 
its type.

### Aggregators
Accomulate incoming messages by type and then send its to sender. There are two 
cases when sending happen:
* Number of messages of one of type reached limit. In this case just one message 
will be sent. 
* Period of time is up. In this case all collected messages will be sent.

### Senders 
Each sender is an adapter to some service like Slack, Email(not implemented yet), etc.


## Environment variables 
| Variable | Type | Description | 
| :---: | :---: | :--- |
| `LSITEN_ADDRESS` | string | Address to listen to, e.g. 127.0.0.1. Default is 0.0.0.0 | 
| `LISTEN_PORT` | integer | Port number to listen to. Default is 3050 |
| `SLACK_TOKEN` | string | Token for accessing Slack from [this](https://api.slack.com/docs/oauth-test-tokens) page.|
| `SLACK_CHANNEL` | string | Channel for posting messages. |
| `AGG_MAX_MESSAGES` | integer | Number of messages to send. |
| `AGG_TIME_LIMIT` | integer | Send any number (less than `AGG_MAX_MESSAGES` ) of messages each `AGG_TIME_LIMIT` seconds. |
| `NOTIFY_USERS` | string | List of user for mention in Slack for default messages. |
| `BOTSMETRICS_API_ADDRESS` | string | API address with campaign statistics. |
| `MONITOR_INTERVAL` | integer | Number of seconds between requests to BOTSMETRIC API. |
| `MONITOR_MAX_STORED_ITEMS` | integer | Number of items to store and analyze. |
| `NOTIFY_MONITOR_USERS` | string | List of users for mention in Slack for monitoring messages. |
| `FKP_THRESHOLD` | integer | Send notification if number of errors of type `FREQUENCY_KEY_PRESENT` for last `MONITOR_INTERVAL` seconds is less than threshold value. |

