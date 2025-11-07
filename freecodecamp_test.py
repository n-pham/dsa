import threading
import freecodecamp_202510

def run_test_wrapper(test, results, index):
    """Wrapper to run a test and store the result in a list at a specific index."""
    results[index] = freecodecamp_202510.run_test(test)

if __name__ == "__main__":
    tests = freecodecamp_202510.tests
    results = [None] * len(tests)
    threads = []

    for i, test in enumerate(tests):
        thread = threading.Thread(target=run_test_wrapper, args=(test, results, i))
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()

    all_passed = True
    for res in results:
        if res:
            print(f"{res[0]}: {res[1]}")
            if res[1] == "FAILED":
                all_passed = False
                print(f"  Error: {res[2]}")

    if all_passed:
        print("\nAll tests passed.")
    else:
        print("\nSome tests failed.")
