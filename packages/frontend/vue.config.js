module.exports = {
  transpileDependencies: ["vuetify"],
  configureWebpack: {
    devServer: {
      hot: false,
      inline: false
    }
  }
};
