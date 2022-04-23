import tinymce from 'tinymce'
import { s3, albumBucketName } from './aws_init'
/* Default icons are required for TinyMCE 5.3 or above */
import 'tinymce/icons/default'
/* A theme is also required */
import 'tinymce/themes/silver'
/* Import the skin */
import 'tinymce/skins/ui/oxide/skin.css'
/* Import plugins */
import 'tinymce/plugins/codesample'

import 'tinymce/plugins/link'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/table'
import 'tinymce/plugins/image'

// TODO: requireでなくともimportできるようにする
/* Import content css */
// eslint-disable-next-line @typescript-eslint/no-var-requires
const contentUiCss = require('tinymce/skins/ui/oxide/content.css').default
// eslint-disable-next-line @typescript-eslint/no-var-requires
const contentCss = require('tinymce/skins/content/default/content.css').default

export const initTinyMce = () => {
  console.log('called')
  tinymce.init({
    selector: '#tinymce_body',
    branding: false, // クレジットの削除
    height: '640',
    tinycomments_mode: 'embedded',
    tinycomments_author: 'Author name',
    plugins: 'link image lists table codesample',
    skin: false,
    content_css: false,
    content_style: contentUiCss.toString() + '\n' + contentCss.toString(),
    codesample_languages: [
      { text: 'HTML/XML', value: 'markup' },
      { text: 'JavaScript', value: 'javascript' },
      { text: 'CSS', value: 'css' },
      { text: 'PHP', value: 'php' },
      { text: 'Ruby', value: 'ruby' },
      { text: 'Python', value: 'python' },
      { text: 'Java', value: 'java' },
      { text: 'C', value: 'c' },
      { text: 'C#', value: 'csharp' },
      { text: 'C++', value: 'cpp' },
      { text: 'Golang', value: 'go' },
    ],
    toolbar: 'undo redo | styleselect | link bold italic | image codesample | numlist bullist | table tabledelete',
    images_upload_handler: function (blobInfo, success, failure) {
      setTimeout(function () {
        const timestamp = new Date().getTime()
        const filename = 'file' + timestamp + blobInfo.filename()
        // TODO: バックエンドに画像アップロード用のエンドポイントを作成する
        s3.putObject(
          { Bucket: albumBucketName, Key: 'uploads/' + filename, ContentType: blobInfo.blob().type, Body: blobInfo.blob(), ACL: 'public-read' },
          function (err, data) {
            if (data !== null) {
              const srcURL = 'https://code-database-images.s3-ap-northeast-1.amazonaws.com/' + 'uploads/' + filename
              success(srcURL)
            } else {
              failure('アップロード失敗.')
            }
          }
        )
      }, 2000)
    },
  })
}
