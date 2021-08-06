#!/bin/sh

echo "$(<LICENSE)"

SITEMAP=http://localhost:8001/sitemap
PORT=8002

./bin/stapafor serve --sitemap=$SITEMAP --port=$PORT
