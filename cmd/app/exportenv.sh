#!/bin/bash
export $(cat ../../.env | xargs)
export DEV=1
