import axios from 'axios';

// Axios 인스턴스 생성
const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_HOST,
  withCredentials: true
});

export async function apiGet(endpoint) {
  try {
    const response = await apiClient.get(endpoint);
    return {status: response.status, data: response.data};
  } catch (error) {
    return {status: error.status, data: error};
  }
}