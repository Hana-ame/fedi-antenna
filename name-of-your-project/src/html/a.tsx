import { ReactNode } from "react";
import { Link } from "react-router-dom";

interface Props {
  children?: ReactNode;
  href: string;
  title?: string;
}

export default function({ children, href, title }: Props) {
  return (
    <Link to={href} title={title}>
      {children}
    </Link>
  )
}