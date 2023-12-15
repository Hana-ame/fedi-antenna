import '@/app/ui/global.css'; // add this line to import global CSS
import { inter } from '@/app/ui/fonts'; // add this line to import fonts

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${inter.className} antialiased`}>{children}</body>
      {/* <body>{children}</body> */}
    </html>
  );
}
