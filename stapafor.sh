#!/bin/sh

echo "$(<LICENSE)"

APP_PORT=8080
SITEMAP=https://b-nova.com/sitemaps/sitemap.xml

./bin/stapafor --port=$APP_PORT --sitemap=SITEMAP
