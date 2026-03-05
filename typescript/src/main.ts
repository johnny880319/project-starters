export function myFunction1(): string {
  return "Hello, World!";
}

export function myFunction2(): number {
  return 42;
}

export function run(): void {
  if (myFunction1() !== "Hello, World!") {
    throw new Error("unexpected result from myFunction1");
  }

  if (myFunction2() !== 42) {
    throw new Error("unexpected result from myFunction2");
  }
}
