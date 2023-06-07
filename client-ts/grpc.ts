import { createPromiseClient } from "@bufbuild/connect";
import { HealthService } from "./gen/proto/health/v1/health_connect";
import { createGrpcTransport } from "@bufbuild/connect-node";

const transport = createGrpcTransport({
  // Requests will be made to <baseUrl>/<package>.<service>/method
  baseUrl: "http://localhost:50051",

  // You have to tell the Node.js http API which HTTP version to use.
  httpVersion: "2",

  // Interceptors apply to all calls running through this transport.
  interceptors: [],
});

async function main() {
  const client = createPromiseClient(HealthService, transport);
  const res = await client.checkHealth({});
  console.log(res);
}
void main();
