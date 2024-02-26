import axios from "axios"

const BASE_URL = 'http://192.168.1.202:8080'


export const axiosPublic = axios.create({
    baseURL: BASE_URL
})