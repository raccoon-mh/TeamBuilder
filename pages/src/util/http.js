import axios from 'axios';

// Axios 인스턴스 생성
const apiClient = axios.create({
  baseURL: 'http://183.99.38.116:8000',
  withCredentials: true
});

// 토큰을 포함한 GET 요청 예시
apiClient.get('/test')
  .then(response => {
    console.log('Response:', response.data);
  })
  .catch(error => {
    console.error('Error:', error);
  });