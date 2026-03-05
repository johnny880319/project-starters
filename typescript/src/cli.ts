import { run } from "./main.js";

try {
  run();
  console.log("All tests passed successfully!");
} catch (error) {
  console.error("Error:", error);
  process.exit(1);
}
