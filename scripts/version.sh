#!/bin/sh

echo v$(date +'%Y%m%d')-$(git rev-parse --short HEAD)
