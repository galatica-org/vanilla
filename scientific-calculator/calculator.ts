export function calculate(op: string, a: number, b?: number): number {
  switch (op) {
    case "add": return a + b!;
    case "subtract": return a - b!;
    case "multiply": return a * b!;
    case "divide": return a / b!;
    case "power": return Math.pow(a, b!);
    case "sqrt": return Math.sqrt(a);
    case "sin": return Math.sin(a);
    case "cos": return Math.cos(a);
    case "tan": return Math.tan(a);
    case "log": return Math.log10(a);
    case "ln": return Math.log(a);
    default: throw new Error(`Unknown operation: ${op}`);
  }
}

function main() {
  const args = process.argv.slice(2);
  if (args.length < 2) {
    console.log("Usage: node calculator.js <op> <a> [b]");
    console.log("Operations: add, subtract, multiply, divide, power, sqrt, sin, cos, tan, log, ln");
    console.log("Examples:");
    console.log("  node calculator.js add 2 3");
    console.log("  node calculator.js sqrt 16");
    console.log("  node calculator.js sin 1.5708");
    process.exit(1);
  }

  const op = args[0];
  const a = parseFloat(args[1]);
  const b = args.length > 2 ? parseFloat(args[2]) : undefined;

  try {
    const result = calculate(op, a, b);
    console.log(result);
  } catch (e) {
    console.error("Error:", (e as Error).message);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}