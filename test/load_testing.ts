import http from "k6/http";
import { check, sleep } from "k6";
import { SharedArray } from "k6/data";

import SensorData from "./interface";

const sensorData: SensorData[] = new SharedArray(
  "load_test_data.json",
  function () {
    return JSON.parse(open("./load_test_data.json"));
  }
);

export const options = {
  stages: [
    { duration: "5m", target: 25 },
    { duration: "30m", target: 25 },
    { duration: "5m", target: 0 },
  ],
};

export default function () {
  const data = sensorData[Math.floor(Math.random() * sensorData.length)];
  console.log(data);
  const payload = JSON.stringify({
    timestamp: data.ts,
    device_id: data.device,
    co: data.co,
    humid: data.humidity,
    temp: data.temp,
    lpg: data.lpg,
    smoke: data.smoke,
    light: data.light,
    motion: data.motion,
  });

  const headers = { "Content-Type": "application/json" };

  const res = http.post(`${__ENV.GATEWAY_ENDPOINT}/sensors`, payload, {
    headers,
  });

  check(res, {
    "Res status is 200": (r) => res.status === 200,
    "Res Content-Type header": (r) =>
      res.headers["Content-Type"] === "application/json",
  });

  sleep(1);
}
