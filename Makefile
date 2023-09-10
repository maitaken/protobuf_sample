.PHONY: protoc
protoc:
	protoc -I=./proto/ --go_out=./proto/ --go_opt=paths=source_relative ./proto/*.proto
	protoc -I=./proto/ --python_out=./proto/ ./proto/*.proto

setup:
	python3 -m venv .venv
	.venv/bin/pip install -r requirements.txt
