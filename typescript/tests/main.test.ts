import { describe, expect, it } from "vitest";

import * as main from "../src/main.js";

describe("main", () => {
  it("myFunction1 returns expected greeting", () => {
    expect(main.myFunction1()).toBe("Hello, World!");
  });

  it("myFunction2 returns expected number", () => {
    expect(main.myFunction2()).toBe(42);
  });

  it("run succeeds with default behavior", () => {
    expect(() => main.run()).not.toThrow();
  });
});
