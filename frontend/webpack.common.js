module.exports = {
  target: 'web',
  externals: [],
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    modules: ['node_modules'],
    extensions: ['.js', '.ts'],
    fallback: { util: false },
  },
}
