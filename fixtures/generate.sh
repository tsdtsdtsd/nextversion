#!/bin/bash

directories=(no-valid-tags valid-tag)

create_first_commit() {
    echo "content" > myfile
    git add myfile
    git commit -m "feat: add file"
}

# Cleanup target directory if necessary
for dir in "${directories[@]}"; do
    # Check if the directory exists
    if [ -d "$dir" ]; then
        rm -rf $dir
    fi
        
    mkdir $dir
done

# Config

git config user.email "tester@example.com"
git config user.name "Tester Tony"
git config init.defaultBranch main

#

cd valid-tag

git init

create_first_commit

git tag -a -m "v0.1.0" v0.1.0

echo "content" > myfile2
git add myfile2
git commit -m "feat: add file2"

# 

cd ../no-valid-tags

git init

create_first_commit

git tag -a -m "should-fail" "hello"