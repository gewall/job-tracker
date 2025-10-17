import { create } from "zustand";
import { AppProps } from "../types/app.type";
import { User } from "../types/auth,type";


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
    },
    user: {},
    setUser: (user:User) =>{
        set(s => ({
            user: {
                access_token: user.access_token,
                userId: user.userId
            }
        }))
    }
}))