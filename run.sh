#!/bin/bash
env LISTEN_ADDRESS=$(hostname -I | awk '{print $2}') \
    LISTEN_PORT=3050 \
    SLACK_TOKEN=xoxb- \
    SLACK_CHANNEL=zabbix_alerts \
    AGG_MAX_MESSAGES=10 \
    AGG_TIME_LIMIT=30 \
    NOTIFY_USERS=ekhabarov,pivolan,tenshi,artem.alemasov \
    API_ADDRESS="https://botsmetrics.appgrowth.com/details" \
    API_TOKEN_REFRESH_INTERVAL=86400 \
    API_TOKEN_REFRESH_ADDRESS="https://api.appgrowth.com/token/refresh/" \
    MONITOR_INTERVAL=300 \
    MONITOR_MAX_STORED_ITEMS=3 \
    NOTIFY_MONITOR_USERS=ekhabarov \
    FKP_THRESHOLD=1000 \
    /opt/apps/mag/bin/mag
