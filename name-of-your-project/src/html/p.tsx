import { ReactNode } from "react";

export default function({ children }: { children?: ReactNode}) {
  return (
    <p>
      {children}
    </p>
  )
}