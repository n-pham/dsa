from concurrent.futures import InterpreterPoolExecutor
import freecodecamp_202510

if __name__ == "__main__":
    with InterpreterPoolExecutor() as pool:
        results = pool.map(freecodecamp_202510.run_test, freecodecamp_202510.tests)

    all_passed = True
    for res in results:
        print(f"{res[0]}: {res[1]}")
        if res[1] == "FAILED":
            all_passed = False
            print(f"  Error: {res[2]}")

    if all_passed:
        print("\nAll tests passed.")
    else:
        print("\nSome tests failed.")