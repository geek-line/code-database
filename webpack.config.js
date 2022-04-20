const path = require('path')
const TerserPlugin = require("terser-webpack-plugin");

const getConfig = (name) => {
    return {
        target: 'web',
        entry:{
            [name]: path.resolve(__dirname, `./src/${name}.ts`)
        },
        output: {
            library: {
                type: "commonjs2"
            },
            path: path.resolve(__dirname, 'static/js'),  //バンドルしたファイルの出力先のパスを指定
            filename: '[name].js' //出力時のファイル名の指定
        },
        externals: [],
        module: {
            rules: [
              {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
              },
            ],
        },
        resolve: {
            modules: ["node_modules"],
            extensions: ['.js', '.ts'],
            fallback: { "util": false }
        },
        optimization: {
            minimize: false,
            minimizer: [
              new TerserPlugin({
                parallel: true,
              }),
            ],
        },
    }
}
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
    getConfig('user_knowledges')
]