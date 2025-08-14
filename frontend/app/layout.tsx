import type { Metadata } from "next";
import "./globals.css";
import { Header } from "../components/Header";

export const metadata: Metadata = {
  title: {
    default: "tamaco-blog",
    template: "%s | tamaco-blog",
  },
  description: "tamaco-blog",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Header />
        <main className="mx-auto max-w-screen-lg px-4 py-6">
          {children}
        </main>
      </body>
    </html>
  );
}
