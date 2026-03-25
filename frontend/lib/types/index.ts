export type Product = {
  id: number;
  name: string;
  description: string;
  price: number;
  stock: number;
  status: string;
};

export type CartItem = {
  userId: number;
  productId: number;
  productName: string;
  price: number;
  quantity: number;
  checked: boolean;
};

export type OrderItem = {
  productId: number;
  productName: string;
  price: number;
  quantity: number;
};

export type Order = {
  orderNo: string;
  userId: number;
  totalAmount: number;
  status: string;
  items: OrderItem[];
};
