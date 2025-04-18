import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  vus: 5,
  duration: "1m",
};

export default function () {
  const res = http.get("https://quickpizza.grafana.com");
  check(res, { "status was 200": (r) => r.status == 200 });

  sleep(1);
}
