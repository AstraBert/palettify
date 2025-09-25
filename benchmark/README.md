# Benchmarking the color extraction API

Simple script to benchmark the color extraction API by sending 100 requests and tracking the average response time, the number of failures and successes.

Prepare the environment with:

```bash
cd benchmark/
uv sync
source .venv/bin/activate
```

Run the code from the parent directory with:

```bash
python3 benchmark/main.py
```

