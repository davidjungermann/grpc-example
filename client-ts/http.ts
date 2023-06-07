import axios from "axios";

async function main() {
  try {
    const response = await axios.get("http://localhost:8080/v1/health", {});
    console.log(response.data);
  } catch (error) {
    console.error(error);
  }
}

main();
