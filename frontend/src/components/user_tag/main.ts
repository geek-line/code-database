import 'materialize-css/dist/css/materialize.min.css'
import './style.css'
import { textToQueryValue } from '../../helpers/shared_logic'

const search_input = document.querySelectorAll<HTMLInputElement>('#search_input')
const search_submit = document.querySelectorAll<HTMLInputElement>('#search_submit')

for (let i = 0; i < search_submit.length; i++) {
  // TODO: Enterキーで検索できる機能を実装する
  // search_input[i].addEventListener('keydown', function () {
  //   if (window.event.keyCode == 13) {
  //     search_submit[i].click()
  //   }
  // })
  search_submit[i].addEventListener('click', function (e) {
    if (!search_input[i].value) {
      e.preventDefault()
      return
    }
    const XHR = new XMLHttpRequest()
    const qvalue = textToQueryValue(search_input[i].value)
    XHR.open('GET', '/search?q=' + qvalue)
    XHR.onreadystatechange = function () {
      if (XHR.readyState === 4) {
        if (XHR.status === 200) {
          location.href = '/search?q=' + qvalue
        } else {
          alert('キーワードを正常に送信できませんでした。')
        }
      }
    }
    XHR.send()
  })
}
