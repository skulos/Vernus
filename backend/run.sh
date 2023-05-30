#!/bin/bash

sleep 5s

echo '/register {"name":"echo-a", "version":"v1"}'
echo ""
curl http://localhost:8080/register -d '{"name":"echo-a", "version":"v1"}'
echo ""

sleep 5s


echo '/register {"name":"echo-b", "version":"v2"}'
echo ""
curl http://localhost:8080/register -d '{"name":"echo-b", "version":"v2"}'
echo ""

sleep 5s

echo '/deploy {"name":"echo-b", "version":"v2"}'
echo ""
curl http://localhost:8080/deploy -d '{"name":"echo-b", "version":"v2"}'
echo ""


# curl -X POST http://localhost:8080/destroy
