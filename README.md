# Message Aggregation Service

## Overview
Service made as simple TCP server which accepts messages on address and port 
provided by ENV variables. All messages are grouped by some attribute and send
to adapters, which, in turn, translate it to developers via Slack, EMail etc.

## Environment variagbles 
| Variable | Type | Description | 
| :---: | :---: | :---|
| `LSITEN_ADDRESS` | string | Address to listen to, e.g. 127.0.0.1. Default is 0.0.0.0 | 
| `LISTEN_PORT` | integer | Port number to listen to. Default is 3050 |


