import type { Metadata } from "next";
import "./globals.css";
import { ToastProvider } from "./ui/toast";
import '@fontsource/vazirmatn'

export const metadata: Metadata = {
  title: "فروشگاه اینترنتی اکسسوری",
  description: "Transform your body with elite training",
  icons: {
    icon: '/etiky.svg',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="fa">
      <head>
      </head>
      <body>
        <ToastProvider>
          <main>
            {children}
          </main>
          <footer>
          </footer>
        </ToastProvider>
      </body>
    </html>
  );
}