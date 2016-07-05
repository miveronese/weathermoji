#!/bin/bash

go build . && export $(cat .env | xargs) && ./weathermoji