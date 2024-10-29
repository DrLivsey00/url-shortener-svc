#!/bin/sh

/usr/local/bin/url-shortener-svc run service &
sleep 2
/usr/local/bin/url-shortener-svc migrate down &
sleep 2
/usr/local/bin/url-shortener-svc migrate up


wait