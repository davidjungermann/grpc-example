import { createPromiseClient } from "@bufbuild/connect";
import { HealthService } from "./gen/proto/health/v1/health_connect";
import { createConnectTransport } from "@bufbuild/connect-node";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  httpVersion: "1.1",
});

async function main() {
  const client = createPromiseClient(HealthService, transport);
  const res = await client.checkHealth({});
  console.log(res);
}
void main();
