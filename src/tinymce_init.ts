import tinymce from 'tinymce'
import { s3, albumBucketName } from './aws_init'

tinymce.init({
  selector: '#tinymce_body',
  branding: false, // クレジットの削除
  height: '640',
  tinycomments_mode: 'embedded',
  tinycomments_author: 'Author name',
  plugins: 'link image lists table codesample',
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
      const file = blobInfo.blob()
      const timestamp = new Date().getTime()
      const filename = 'file' + timestamp + blobInfo.name
      s3.putObject(
        { Bucket: albumBucketName, Key: 'uploads/' + filename, ContentType: blobInfo.blob().type, Body: blobInfo.blob(), ACL: 'public-read' },
        function (err, data) {
          if (data !== null) {
            const srcHTML = 'https://code-database-images.s3-ap-northeast-1.amazonaws.com/' + 'uploads/' + filename
            success(srcHTML)
          } else {
            alert('アップロード失敗.')
          }
        }
      )
    }, 2000)
  },
})

export {}
