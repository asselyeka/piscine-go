#!/bin/bash
find . -name '*.sh' | sed  's#/##g' | sed  's/.sh//' | cut -f2 -d '.' 