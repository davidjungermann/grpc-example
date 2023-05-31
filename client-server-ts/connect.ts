import { ConnectRouter } from "@bufbuild/connect";
import { HealthService } from "./gen/proto/health/v1/health_connect";

export default (router: ConnectRouter) =>
  // registers HealthService
  router.service(HealthService, {
    // implements rpc Say
    async checkHealth() {
      return {
        status: "OK",
      };
    },
  });
