import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 25,
  duration: "1m",
};

export default function () {
  const BASE_URL = "http://localhost:8082";

  const query = {
    page: 1,
    perPage: 10,
    orderBy: "created_at",
    orderType: "desc",
    price: "",
    search: "",
  };

  const url =
    `${BASE_URL}/products/shop?` +
    `page=${query.page}&perPage=${query.perPage}` +
    `&orderBy=${query.orderBy}&orderType=${query.orderType}` +
    `&price=${encodeURIComponent(query.price)}` +
    `&search=${query.search}`;

  const res = http.get(url);

  check(res, {
    "status is 200": (r) => r.status === 200,
    "message is success": (r) => JSON.parse(r.body).message === "success",
    "has pagination": (r) => JSON.parse(r.body).pagination !== null,
    "data is array": (r) => Array.isArray(JSON.parse(r.body).data),
  });

  sleep(1);
}
