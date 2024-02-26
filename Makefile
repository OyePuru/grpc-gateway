name: Generate Stubs

on:
  push:
    paths:
      - 'proto/**'
  pull_request:
    paths:
      - 'proto/**'

jobs:
  generate-stubs:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout repository
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.17'

    - name: Install protoc
      run: |
        sudo apt-get update
        sudo apt-get install -y protobuf-compiler

    - name: Install protoc-gen-go
      run: |
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

    - name: Install protoc-gen-go-grpc
      run: |
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    - name: Install protoc-gen-grpc-gateway
      run: |
        go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

    - name: Generate stubs
      run: |
        make generate-stubs
        git diff --quiet || git commit -am "Update generated stubs" && git push
