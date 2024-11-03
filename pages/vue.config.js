const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  devServer: {
    host: '0.0.0.0', // 모든 IP에서 접근 허용
    allowedHosts: 'all', // 모든 호스트 허용
  },
  publicPath: process.env.NODE_ENV === 'production' ? '/TeamBuilder/' : '/',
  transpileDependencies: true
})
