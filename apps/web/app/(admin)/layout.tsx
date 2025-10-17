import React from "react";
import Navbar from "../_components/Navbar";
import PrivateRoute from "../_components/PrivateRoute";

type Props = {
  children: React.ReactNode;
};

export default function DashboardLayout({ children }: Props) {
  return (
    <PrivateRoute>
      <Navbar />
      <section className="bg-neutral-100 h-screen pt-20">{children}</section>
    </PrivateRoute>
  );
}
