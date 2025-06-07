#!/usr/bin/env bash

go build -ldflags="-w -s" -o $HOME/bin/grug .

chmod +x $HOME/bin/grug
