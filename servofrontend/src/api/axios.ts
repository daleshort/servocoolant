import axios from "axios"


export const axiosPublic = axios.create({
    baseURL: import.meta.env.BASE_URL
})