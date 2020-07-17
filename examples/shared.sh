#!/bin/bash

function bold()
{
    echo -e "\033[1m${1}\033[22m"
}

function dim()
{
    echo -e "\033[2m${1}\033[22m"
}

function italic()
{
    echo -e "\033[3m${1}\033[23m"
}

function underline()
{
    echo -e "\033[4m${1}\033[24m"
}

function red()
{
    echo -e "\033[31m${1}\033[39m"
}

function green()
{
    echo -e "\033[32m${1}\033[39m"
}
