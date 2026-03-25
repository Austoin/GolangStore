import { mockCartItems, mockOrders, mockProducts } from "../mock/data";
import type { CartItem, Order, Product } from "../types";

export function listProducts(): Product[] {
  return mockProducts;
}

export function getProductDetail(id: number): Product | undefined {
  return mockProducts.find((item) => item.id === id);
}

export function listCartItems(userId: number): CartItem[] {
  return mockCartItems.filter((item) => item.userId === userId);
}

export function listOrders(): Order[] {
  return mockOrders;
}
