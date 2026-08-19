import math
import sys

def calculate(op: str, a: float, b: float = None) -> float:
    """Simple scientific calculator function."""
    if op == "add":
        return a + b
    elif op == "subtract":
        return a - b
    elif op == "multiply":
        return a * b
    elif op == "divide":
        return a / b
    elif op == "power":
        return a ** b
    elif op == "sqrt":
        return math.sqrt(a)
    elif op == "sin":
        return math.sin(a)
    elif op == "cos":
        return math.cos(a)
    elif op == "tan":
        return math.tan(a)
    elif op == "log":
        return math.log10(a)
    elif op == "ln":
        return math.log(a)
    else:
        raise ValueError(f"Unknown operation: {op}")


def main():
    if len(sys.argv) < 3:
        print("Usage: python calculator.py <op> <a> [b]")
        print("Operations: add, subtract, multiply, divide, power, sqrt, sin, cos, tan, log, ln")
        print("Examples:")
        print("  python calculator.py add 2 3")
        print("  python calculator.py sqrt 16")
        print("  python calculator.py sin 1.5708")
        sys.exit(1)

    op = sys.argv[1]
    a = float(sys.argv[2])
    b = float(sys.argv[3]) if len(sys.argv) > 3 else None

    try:
        result = calculate(op, a, b)
        print(result)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    # Demo
    # print("add(2, 3) =", calculate("add", 2, 3))
    # print("sqrt(16) =", calculate("sqrt", 16))
    # print("sin(π/2) =", calculate("sin", math.pi / 2))
    # print("log(100) =", calculate("log", 100))
    # print("ln(e) =", calculate("ln", math.e))
    main()