import type { Product } from "../types";

const PRODUCT_BASE_URL = "http://127.0.0.1:8081";

export async function listProductsHttp(): Promise<Product[]> {
  const response = await fetch(`${PRODUCT_BASE_URL}/products`, { cache: "no-store" });
  if (!response.ok) {
    throw new Error(`failed to list products: ${response.status}`);
  }

  return response.json();
}

export async function getProductDetailHttp(id: number): Promise<Product> {
  const response = await fetch(`${PRODUCT_BASE_URL}/products/${id}`, { cache: "no-store" });
  if (!response.ok) {
    throw new Error(`failed to get product detail: ${response.status}`);
  }

  return response.json();
}

const CART_BASE_URL = "http://127.0.0.1:8083";
const ORDER_BASE_URL = "http://127.0.0.1:8082";

export async function listCartItemsHttp(userId: number) {
  const response = await fetch(`${CART_BASE_URL}/carts/${userId}`, { cache: "no-store" });
  if (!response.ok) throw new Error(`failed to list cart items: ${response.status}`);
  return response.json();
}

export async function createOrderFromCartHttp(userId: number) {
  const response = await fetch(`${ORDER_BASE_URL}/orders/from-cart`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId }),
    cache: "no-store",
  });
  if (!response.ok) throw new Error(`failed to create order: ${response.status}`);
  return response.json();
}

export async function listOrdersHttp() {
  const response = await fetch(`${ORDER_BASE_URL}/orders`, { cache: "no-store" });
  if (!response.ok) throw new Error(`failed to list orders: ${response.status}`);
  return response.json();
}
