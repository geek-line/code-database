const path = require('path')
const { merge } = require('webpack-merge')
const common = require('./webpack.common.js')

const getConfig = (name) =>
  merge(common, {
    mode: 'development',
    entry: {
      [name]: path.resolve(__dirname, `./src/${name}.ts`),
    },
    output: {
      library: {
        type: 'commonjs2',
      },
      path: path.resolve(__dirname, 'static/js'), //バンドルしたファイルの出力先のパスを指定
      filename: '[name].js', //出力時のファイル名の指定
    },
  })

module.exports = [
  getConfig('admin_categories'),
  getConfig('admin_edit'),
  getConfig('admin_eyecatches'),
  getConfig('admin_knowledges'),
  getConfig('admin_new'),
  getConfig('admin_tags'),
  getConfig('aws_init'),
  getConfig('tinymce_init'),
  getConfig('user_details'),
  getConfig('user_knowledges'),
]
