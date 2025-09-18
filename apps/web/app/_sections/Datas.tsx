"use client";

import React from "react";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@workspace/ui/components/tooltip";
import { SectionContent } from "../_components/Section";
import { CircleAlert } from "lucide-react";
import { DataTable } from "../_components/Datatable";
import { applicationsColumns } from "../_datatable/columns";

import { applicationColumnsData, chartData } from "@/lib/constant/dummy";
import {
  type ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@workspace/ui/components/chart";
import { PieChart } from "recharts";
import { Pie } from "recharts";
import { Cell } from "recharts";

type Props = {};

export function LatestTable({}: Props) {
  return (
    <SectionContent>
      <div className="flex items-center py-2 justify-between">
        <h5 className="text-lg font-light ">Job Applications</h5>
        <Tooltip>
          <TooltipTrigger className="text-neutral-400 hover:text-neutral-600">
            <CircleAlert className="stroke-1 w-5 h-5" />
          </TooltipTrigger>
          <TooltipContent>
            <p>Your latest job application</p>
          </TooltipContent>
        </Tooltip>
      </div>
      <DataTable columns={applicationsColumns} data={applicationColumnsData} />
    </SectionContent>
  );
}

const chartConfig = {
  applied: {
    label: "Applied",
  },
  saved: {
    label: "Saved",
  },
  offer: {
    label: "Offer",
  },
  rejected: {
    label: "Rejected",
  },
  interview: {
    label: "interview",
  },
  count: {
    label: "Total Applications",
  },
} satisfies ChartConfig;

const BLUE_PALETTE = [
  "#0B76FF", // electric blue
  "#1E90FF", // dodger
  "#3EA0FF",
  "#6CB8FF",
  "#A9D3FF",
];

export function StatusPieChart() {
  return (
    <SectionContent classname="flex flex-col h-full">
      <div className="flex items-center justify-between my-2">
        <h5 className="font-light text-lg flex-none">Applications Status</h5>
        <Tooltip>
          <TooltipTrigger className="text-neutral-400 hover:text-neutral-600">
            <CircleAlert className="stroke-1 w-5 h-5" />
          </TooltipTrigger>
          <TooltipContent>
            <p>The overall status of your job application</p>
          </TooltipContent>
        </Tooltip>
      </div>
      <ChartContainer
        config={chartConfig}
        className="min-h-[200px] w-full flex-1"
      >
        <PieChart>
          <ChartTooltip
            content={
              <ChartTooltipContent
                indicator="dashed"
                labelKey="count"
                nameKey="status"
              />
            }
          />
          <ChartLegend
            content={
              <ChartLegendContent nameKey="status" className="flex-wrap" />
            }
          />
          <Pie dataKey="count" data={chartData}>
            {chartData.map((entry, index) => (
              <Cell
                key={`cell-${entry.status}`}
                fill={BLUE_PALETTE[index % BLUE_PALETTE.length]}
              />
            ))}
          </Pie>
        </PieChart>
      </ChartContainer>
    </SectionContent>
  );
}
