import { ApplicationStatus } from "./applications.type";
import { User } from "./auth,type";

export type AppProps = {
    sidebarOpen: boolean;
    sidebarToggle: () => void;
    applicationStatusFilter: ApplicationStatus|"all";
    applicationStatusFilterChange: (e: ApplicationStatus|"all") => void;
    notifications: string[];
    addNotification: (msg:string) => void;
    user: User;
    setUser: (e:User) => void
}