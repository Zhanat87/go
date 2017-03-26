#!/usr/bin/env bash

# http://how-to.wikia.com/wiki/How_to_read_command_line_arguments_in_a_bash_script
# usage: sh update.sh last_changes
git add . && git commit -m $1 && git push && rm bin/go && go get -u github.com/Zhanat87/go