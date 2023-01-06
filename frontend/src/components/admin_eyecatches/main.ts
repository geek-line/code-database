import 'materialize-css/dist/css/materialize.min.css'
import './style.css'

const submit_btn_post = document.getElementById('submit_btn_post') as HTMLInputElement
const submit_btn_put = document.querySelectorAll<HTMLInputElement>('#submit_btn_put')
const submit_btn_delete = document.querySelectorAll<HTMLInputElement>('#submit_btn_delete')
const file_uploader_post = document.getElementById('file_uploader_post') as HTMLInputElement
const file_uploader_put = document.querySelectorAll<HTMLInputElement>('#file_uploader_put')
const file_previews = document.querySelectorAll<HTMLImageElement>('#file_preview')
const forms = document.querySelectorAll<HTMLFormElement>('#form')
import { s3, albumBucketName } from '../../helpers/aws_init'

submit_btn_post.addEventListener('click', function () {
  const timestamp = new Date().getTime()
  const filename = file_uploader_post.files && 'file-' + timestamp + '-' + file_uploader_post.files[0].name
  if (file_uploader_post.files === null) {
    return
  }
  s3.putObject(
    {
      Bucket: albumBucketName,
      Key: 'eyecatches/' + filename,
      ContentType: file_uploader_post.files[0].type,
      Body: file_uploader_post.files[0],
      ACL: 'public-read',
    },
    function (err, data) {
      if (data !== null) {
        ;(document.getElementById('src_post') as HTMLInputElement).value =
          'https://image.code-database.com/' + 'eyecatches/' + filename
        const formdata = new FormData(document.getElementById('form_post') as HTMLFormElement)
        const XHR = new XMLHttpRequest()
        XHR.open('POST', '/admin/eyecatches/')
        XHR.send(formdata)
        XHR.onreadystatechange = function () {
          if (XHR.readyState === 4) {
            if (XHR.status === 200) {
              alert('データが更新されました')
              location.href = '/admin/eyecatches/'
            } else {
              alert('データが正常に送れませんでした')
            }
          }
        }
      } else {
        alert('アップロード失敗.')
      }
    }
  )
})
for (let i = 0; i < file_uploader_put.length; i++) {
  file_uploader_put[i].addEventListener('change', function () {
    if (this.files === null) {
      return
    }
    const file = this.files[0]
    // ファイルのブラウザ上でのURLを取得する
    const blobUrl = window.URL.createObjectURL(file)
    file_previews[i].src = blobUrl
  })
}
for (let i = 0; i < forms.length; i++) {
  submit_btn_put[i].addEventListener('click', function () {
    const files = file_uploader_put[i].files
    if (files && files[0]) {
      const timestamp = new Date().getTime()
      const filename = 'file-' + timestamp + '-' + files[0].name
      const key = document
        .querySelectorAll<HTMLInputElement>('#src_put')
        [i].value.replace(/^https:\/\/code-database-images\.s3-ap-northeast-1\.amazonaws\.com\//, '')
      s3.putObject(
        { Bucket: albumBucketName, Key: 'eyecatches/' + filename, ContentType: files[0].type, Body: files[0], ACL: 'public-read' },
        function (err, data) {
          if (data !== null) {
            document.querySelectorAll<HTMLInputElement>('#src_put')[i].value =
              'https://image.code-database.com/' + 'eyecatches/' + filename
            const formdata = new FormData(forms[i])
            const XHR = new XMLHttpRequest()
            XHR.open('PUT', '/admin/eyecatches/')
            XHR.send(formdata)
            XHR.onreadystatechange = function () {
              if (XHR.readyState === 4) {
                if (XHR.status === 200) {
                  s3.deleteObject({ Bucket: albumBucketName, Key: key }, function (err) {
                    if (err != null) {
                      alert('データの削除に失敗しました')
                      return
                    }
                    alert('データが更新されました')
                    location.href = '/admin/eyecatches/'
                  })
                } else {
                  alert('データが正常に送れませんでした')
                  return
                }
              }
            }
          } else {
            alert('アップロード失敗.')
          }
        }
      )
    } else {
      const formdata = new FormData(forms[i])
      const XHR = new XMLHttpRequest()
      XHR.open('PUT', '/admin/eyecatches/')
      XHR.send(formdata)
      XHR.onreadystatechange = function () {
        if (XHR.readyState === 4) {
          if (XHR.status === 200) {
            alert('データが更新されました')
            location.href = '/admin/eyecatches/'
          } else {
            alert('データが正常に送れませんでした')
          }
        }
      }
    }
  })
  submit_btn_delete[i].addEventListener('click', function () {
    const formdata = new FormData(forms[i])
    const XHR = new XMLHttpRequest()
    const key = document
      .querySelectorAll<HTMLInputElement>('#src_put')
      [i].value.replace(/^https:\/\/code-database-images.s3-ap-northeast-1\.amazonaws\.com\//, '')
    s3.deleteObject({ Bucket: albumBucketName, Key: key }, function () {
      XHR.open('DELETE', '/admin/eyecatches/')
      XHR.send(formdata)
      XHR.onreadystatechange = function () {
        if (XHR.readyState === 4) {
          if (XHR.status === 200) {
            alert('データが更新されました')
            location.href = '/admin/eyecatches/'
          } else {
            alert('データが正常に送れませんでした')
          }
        }
      }
    })
  })
}
