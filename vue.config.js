module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],

  devServer: {
    open: true,
    host: 'localhost',
    port: 8080,
    https: false,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8888/',
        ws: true,
        changOrigin: true,
        pathRewrite: {
          '^/api': '/api'
        }
      }

    }
  },

  publicPath: ''
}