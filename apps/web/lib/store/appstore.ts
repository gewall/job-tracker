import { create } from "zustand";
import { AppProps } from "../types/app.type";


export const useAppStore = create<AppProps>((set) =>({
    sidebarOpen: true,
    sidebarToggle: () => {
        set((s) => ({sidebarOpen: !s.sidebarOpen}))
    },
    applicationStatusFilter: "all",
    applicationStatusFilterChange:(e) => {
        set(s => ({applicationStatusFilter: e}))
    },
    notifications: [],
    addNotification: (e) => {
        set(s => ({notifications: [...s.notifications, e]}))
    }
}))