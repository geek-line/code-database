import './style.css'
import 'materialize-css/dist/css/materialize.min.css'

const forms = document.querySelectorAll<HTMLFormElement>('#form')
const update_buttons = document.querySelectorAll('#update_button')
const delete_buttons = document.querySelectorAll('#delete_button')
const file_preview = document.querySelectorAll<HTMLImageElement>('#file_preview')
const select_eyecatch = document.querySelectorAll<HTMLSelectElement>('#select_eyecatch')
window.addEventListener('DOMContentLoaded', function () {
  file_preview[0].src = select_eyecatch[0].value
})
for (let i = 0; i < select_eyecatch.length; i++) {
  select_eyecatch[i].addEventListener('change', function () {
    file_preview[i].src = this.value
  })
}
for (let i = 0; i < forms.length; i++) {
  update_buttons[i].addEventListener('click', function (e) {
    if (document.querySelectorAll<HTMLInputElement>('#form-name')[i].value == '') {
      alert('名前を入力してください')
      e.preventDefault()
      return
    }
    const formdata = new FormData(forms[i])
    const XHR = new XMLHttpRequest()
    XHR.open('PUT', '/admin/categories/')
    XHR.send(formdata)
    XHR.onreadystatechange = function () {
      if (XHR.readyState === 4) {
        if (XHR.status === 200) {
          alert('データが更新されました')
          location.href = '/admin/categories/'
        } else {
          alert('データが正常に送れませんでした')
        }
      }
    }
  })
  delete_buttons[i].addEventListener('click', function () {
    const formdata = new FormData(forms[i])
    const XHR = new XMLHttpRequest()
    XHR.open('DELETE', '/admin/categories/')
    XHR.send(formdata)
    XHR.onreadystatechange = function () {
      if (XHR.readyState === 4) {
        if (XHR.status === 200) {
          alert('データが更新されました')
          location.href = '/admin/categories/'
        } else {
          alert('データが正常に送れませんでした')
        }
      }
    }
  })
}
