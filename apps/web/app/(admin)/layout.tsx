import React from "react";
import Navbar from "../_components/Navbar";

type Props = {
  children: React.ReactNode;
};

export default function DashboardLayout({ children }: Props) {
  return (
    <div>
      <Navbar />
      <section className="bg-neutral-100 h-screen pt-20">{children}</section>
    </div>
  );
}
