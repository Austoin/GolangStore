import type { CartItem, Order, Product } from "../types";

export const mockProducts: Product[] = [
  { id: 101, name: "Phone X", description: "旗舰手机", price: 199900, stock: 8, status: "在售" },
  { id: 102, name: "Cable", description: "快充数据线", price: 4900, stock: 25, status: "在售" },
  { id: 103, name: "Mouse", description: "无线鼠标", price: 12900, stock: 4, status: "低库存" },
];

export const mockCartItems: CartItem[] = [
  { userId: 1, productId: 101, productName: "Phone X", price: 199900, quantity: 2, checked: true },
];

export const mockOrders: Order[] = [
  {
    orderNo: "O1001",
    userId: 1,
    totalAmount: 399800,
    status: "待支付",
    items: [{ productId: 101, productName: "Phone X", price: 199900, quantity: 2 }],
  },
];
