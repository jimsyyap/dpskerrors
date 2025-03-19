import axios from 'axios';
const api = axios.create({ baseURL: process.env.REACT_APP_BACKEND_URL });
api.interceptors.request.use(config => {
  config.headers.Authorization = localStorage.getItem('token');
  return config;
});
export default api;
