import { ApplicationListColumns } from "@/app/(admin)/application/_datatable/columns";
import { ApplicationStatus } from "../types/applications.type";
import { ApplicationColumns } from "@/app/_datatable/columns";

export const applicationListColumnsData: ApplicationListColumns[] =[
  {
    id: "1",
    companyName: "PT Tech Nusantara",
    position: "Frontend Developer",
    applicationDate: "01 Sep 2025",
    contactHR: "Ibu Rina - rina.hr@technusantara.com",
    interviewDate: "20 Sep 2025, 10:00",
    notes: "HR responsif, jadwal interview sudah fix.",
    status: "interview" as ApplicationStatus,
  },
  {
    id: "2",
    companyName: "Startup Kreatif",
    position: "UI/UX Designer",
    applicationDate: "28 Aug 2025",
    contactHR: "Pak Andi - andi@startupkreatif.id",
    interviewDate: "-",
    notes: "Menunggu balasan HR, follow-up minggu depan.",
    status: "applied" as ApplicationStatus,
  },
  {
    id: "3",
    companyName: "Global Corp",
    position: "Backend Engineer",
    applicationDate: "15 Aug 2025",
    contactHR: "HRD - careers@globalcorp.com",
    interviewDate: "-",
    notes: "Ditolak karena kurang pengalaman Go.",
    status: "rejected" as ApplicationStatus,
  },
  {
    id: "4",
    companyName: "Digital Vision",
    position: "Mobile Developer (React Native)",
    applicationDate: "05 Sep 2025",
    contactHR: "Bu Sinta - sinta@digitalvision.co.id",
    interviewDate: "-",
    notes: "Sudah ada offer, sedang negosiasi gaji.",
    status: "offer" as ApplicationStatus,
  },
  {
    id: "5",
    companyName: "EduTech Indonesia",
    position: "Data Analyst",
    applicationDate: "10 Sep 2025",
    contactHR: "Pak Budi - hrd@edutech.id",
    interviewDate: "-",
    notes: "Belum apply, masih siapkan CV versi data science.",
    status: "saved" as ApplicationStatus,
  },
   {
    id: "6",
    companyName: "FinTech Global",
    position: "DevOps Engineer",
    applicationDate: "02 Sep 2025",
    contactHR: "Rani - rani@fintechglobal.com",
    interviewDate: "22 Sep 2025, 14:00",
    notes: "Interview tahap 1 via Zoom.",
    status: "interview" as ApplicationStatus,
  },
  {
    id: "7",
    companyName: "HealthCare Solutions",
    position: "Fullstack Developer",
    applicationDate: "18 Aug 2025",
    contactHR: "HR Team - careers@hcsolutions.com",
    interviewDate: "-",
    notes: "Sudah apply, belum ada respon.",
    status: "applied" as ApplicationStatus,
  },
  {
    id: "8",
    companyName: "RetailMart",
    position: "Data Engineer",
    applicationDate: "12 Aug 2025",
    contactHR: "Pak Joko - joko@retailmart.co.id",
    interviewDate: "-",
    notes: "Ditolak setelah screening CV.",
    status: "rejected" as ApplicationStatus,
  },
  {
    id: "9",
    companyName: "Creative Agency X",
    position: "Graphic Designer",
    applicationDate: "07 Sep 2025",
    contactHR: "Dina - dina@cax.id",
    interviewDate: "25 Sep 2025, 09:00",
    notes: "Tes portofolio sebelum interview HR.",
    status: "interview" as ApplicationStatus,
  },
  {
    id: "10",
    companyName: "Global E-Commerce",
    position: "Product Manager",
    applicationDate: "03 Sep 2025",
    contactHR: "HRD - jobs@globalecom.com",
    interviewDate: "-",
    notes: "Mendapatkan offer awal, menunggu kontrak resmi.",
    status: "offer" as ApplicationStatus,
  },
];

export const applicationColumnsData: ApplicationColumns[] = [
  {
    companyName: "PT Tech Nusantara",
    position: "Frontend Developer",
    contactHR: "Ibu Rina - rina.hr@technusantara.com",
    interviewDate: "20 Sep 2025, 10:00",
    status: "interview" as ApplicationStatus,
  },
  {
    companyName: "Startup Kreatif",
    position: "UI/UX Designer",
    contactHR: "Pak Andi - andi@startupkreatif.id",
    interviewDate: "-",
    status: "applied" as ApplicationStatus,
  },
  {
    companyName: "Global Corp",
    position: "Backend Engineer",
    contactHR: "HRD - careers@globalcorp.com",
    interviewDate: "-",
    status: "rejected" as ApplicationStatus,
  },
  {
    companyName: "Digital Vision",
    position: "Mobile Developer (React Native)",
    contactHR: "Bu Sinta - sinta@digitalvision.co.id",
    interviewDate: "-",
    status: "offer" as ApplicationStatus,
  },
  {
    companyName: "EduTech Indonesia",
    position: "Data Analyst",
    contactHR: "Pak Budi - hrd@edutech.id",
    interviewDate: "-",
    status: "saved" as ApplicationStatus,
  },
];

export const chartData = [
    {
        status: "saved" as ApplicationStatus,
        count: 2
    },
    {
        status: "applied" as ApplicationStatus,
        count: 3
    },
    {
        status: "offer" as ApplicationStatus,
        count: 1
    },
    {
        status: "rejected" as ApplicationStatus,
        count: 2
    },
    {
        status: "interview" as ApplicationStatus,
        count: 1
    }
]