const path = require('path')
const fs = require('fs')
const { merge } = require('webpack-merge')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const TerserPlugin = require('terser-webpack-plugin')
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin')
const common = require('./webpack.common.js')

const componentNames = fs.readdirSync(path.resolve(__dirname, 'src/components'))

const getComponentConfig = (componentNames) => {
  const entries = {}
  componentNames.forEach((name) => {
    entries[name] = path.resolve(__dirname, `./src/components/${name}/main.ts`)
  })
  return merge(common, {
    mode: 'production',
    entry: entries,
    plugins: [
      new MiniCssExtractPlugin({ filename: 'css/[name].css' }),
      ...componentNames.map(
        (componentName) =>
          new HtmlWebpackPlugin({
            inject: /^_/.test(componentName) ? false : 'body',
            chunks: [componentName],
            template: path.resolve(__dirname, `src/components/${componentName}/index.html`),
            filename: path.resolve(__dirname, `../backend/dist/template/${componentName}.html`),
          })
      ),
    ],
    module: {
      rules: [
        {
          test: /skin\.css$/i,
          use: [MiniCssExtractPlugin.loader, 'css-loader'],
        },
        {
          test: /content\.css$/i,
          use: ['css-loader'],
        },
        {
          test: /style\.css$/i,
          use: [MiniCssExtractPlugin.loader, 'css-loader'],
          exclude: /node_modules/,
        },
        {
          test: /materialize\.min\.css$/i,
          use: [MiniCssExtractPlugin.loader, 'css-loader'],
        },
        {
          test: /prism\.css$/i,
          use: [MiniCssExtractPlugin.loader, 'css-loader'],
        },
        {
          test: /iziToast\.min\.css/i,
          use: [MiniCssExtractPlugin.loader, 'css-loader'],
        },
      ],
    },
    optimization: {
      splitChunks: {
        chunks: 'all',
        cacheGroups: {
          tinymceVendor: {
            test: /[\\/]node_modules[\\/](tinymce)[\\/](.*js|.*skin.css)|[\\/]plugins[\\/]/,
            name: 'tinymce',
          },
          materializeVendor: {
            test: /node_modules\/materialize\-css\/dist\/css\/materialize\.min\.css/,
            name: 'materialize',
          },
        },
      },
      minimize: true,
      minimizer: [
        new TerserPlugin({
          parallel: true,
        }),
        new CssMinimizerPlugin({
          parallel: true,
        }),
      ],
    },
    output: {
      publicPath: '/dist',
      path: path.resolve(__dirname, '../backend/dist'), //バンドルしたファイルの出力先のパスを指定
      filename: 'js/[name].js', //出力時のファイル名の指定
    },
  })
}

module.exports = getComponentConfig(componentNames)
