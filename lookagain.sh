#!/bin/bash
find . -name '*.sh' | sed  's#/##g' | sed  's/test//g' | cut -f2 -d '.'