#!/bin/sh

echo "$(<LICENSE)"

SITEMAP=localhost:8000/sitemap
PORT=8080

./bin/stapafor serve --sitemap=$SITEMAP --port=$PORT
