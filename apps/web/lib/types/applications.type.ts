
export type ApplicationStatus = 
  | "saved"       // disimpan, belum apply
  | "applied"     // sudah apply
  | "interview"   // sudah ada jadwal interview
  | "offer"       // dapat tawaran kerja
  | "rejected";   // ditolak

export interface JobApplication {
  id: string;                // ID unik
  companyName: string;       // Nama perusahaan
  position: string;          // Posisi yang dilamar
  applicationDate: Date;     // Tanggal melamar
  status: ApplicationStatus; // Status lamaran

  interviewDate?: Date;      // Jadwal interview
  interviewLocation?: string;// Lokasi / link interview

  contactHR?: string;        // Kontak HR (nama, email, atau nomor telepon dalam 1 string)

  salaryExpectation?: number; // Gaji yang diharapkan
  salaryOffered?: number;     // Gaji yang ditawarkan (jika ada)

  location?: string;          // Lokasi kerja
  jobType?: "full-time" | "part-time" | "internship" | "contract"; // jenis kerja

  feedback?: string;          // Feedback dari recruiter
  notes?: string;             // Catatan pribadi
  resumeVersion?: string;     // CV/Cover Letter yang dipakai
}
