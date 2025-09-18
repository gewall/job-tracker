"use client";

import { ApplicationStatus } from "@/lib/types/applications.type";
import { ColumnDef } from "@tanstack/react-table";

export type ApplicationColumns = {
  companyName: string;
  position: string;
  contactHR: string;
  interviewDate: string;
  status: ApplicationStatus;
};

export const applicationsColumns: ColumnDef<ApplicationColumns>[] = [
  {
    accessorKey: "companyName",
    header: "Company",
  },
  {
    accessorKey: "position",
    header: "Position",
  },
  {
    accessorKey: "contactHR",
    header: "Contact",
  },
  {
    accessorKey: "interviewDate",
    header: "Interview Date",
  },
  {
    accessorKey: "status",
    header: "Status",
  },
];
