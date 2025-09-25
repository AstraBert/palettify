import requests
import time
from statistics import mean, stdev

def main():
    failed = 0
    success = 0
    times = []
    for i in range(100):
        if (i+1) % 10 == 0:
            print(f"Processed {i+1} requests")
        with open("testfiles/gopher.png", "rb") as f:
            files = {
                "image": ("docker.png", f, "image/png")
            }
            start = time.time()
            response = requests.post("http://localhost:8000/colors", files=files)
        time_delta = time.time() - start
        if "An error occurred while extracting the palette colors, try again with a different image" in response.text:
            failed += 1
        else:
            success += 1
        times.append(time_delta)
    print("Total successes:", success)
    print("Total failed:", failed)
    print("Average time:", mean(times))
    print("Stdev:", stdev(times))

if __name__ == "__main__":
    main()