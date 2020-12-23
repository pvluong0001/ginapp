import axios from "axios";

const config = {
    baseURL: 'http://localhost:3456/api'
}

const axiosClient = axios.create(config);

export default axiosClient;
