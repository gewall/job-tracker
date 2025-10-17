import React from "react";
import Form from "./_component/Form";
import { Metadata } from "next";
type Props = {};

export const metadata: Metadata = {
  title: "Job Trakcer",
  description: "Job Tracker Web Application",
};

export default function page({}: Props) {
  return (
    <div className="flex justify-center items-center w-screen h-screen">
      <div className=" w-80 h-96 rounded-2xl overflow-hidden border-2">
        <div className="bg-indigo-500 text-center text-white py-4">
          <h4 className="text-xl font-medium">Sign In</h4>
        </div>
        <div className="px-4 py-8">
          <Form />
        </div>
      </div>
    </div>
  );
}
