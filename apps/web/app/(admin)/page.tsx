import React from "react";
import Section from "../_components/Section";

import { chartData } from "@/lib/constant/dummy";

import Card from "../_components/Card";

import { LatestTable, StatusPieChart } from "../_sections/Datas";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Job Trakcer",
  description: "Job Tracker Web Application",
};

type Props = {};

export default function page({}: Props) {
  return (
    <Section classname="flex flex-col  gap-4 max-w-screen ">
      <div className="flex flex-wrap gap-4">
        {chartData.map((_, i) => (
          <Card label={_.status} {..._} key={i} />
        ))}
      </div>
      <div className="flex flex-wrap gap-4">
        <div className="flex-2 w-64">
          <LatestTable />
        </div>
        <div className="flex-1 w-40 ">
          <StatusPieChart />
        </div>
      </div>
      <div className="flex-1"></div>
    </Section>
  );
}
