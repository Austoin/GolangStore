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
