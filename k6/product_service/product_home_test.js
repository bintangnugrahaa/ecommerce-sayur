import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 20,
  duration: "1m",
};

export default function () {
  const BASE_URL = "http://localhost:8082";

  const url = `${BASE_URL}/products/home`;

  const res = http.get(url);

  check(res, {
    "status is 200": (r) => r.status === 200,
    "message is success": (r) => JSON.parse(r.body).message === "success",
    "contains list": (r) => Array.isArray(JSON.parse(r.body).data),
  });

  sleep(1);
}
