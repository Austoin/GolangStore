import type { ButtonHTMLAttributes, ReactNode } from "react";

type PrimaryButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  children: ReactNode;
};

export function PrimaryButton({ children, ...props }: PrimaryButtonProps) {
  return (
    <button className="primary-button" {...props}>
      {children}
    </button>
  );
}
