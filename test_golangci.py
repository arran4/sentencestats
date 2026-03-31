import json
with open('extracted_blocks.json') as f:
    blocks = json.load(f)
for i, block in enumerate(blocks):
    if 'golangci.yml' in block or 'run:' in block:
        print(f"--- Block {i} ---")
        print(block)
