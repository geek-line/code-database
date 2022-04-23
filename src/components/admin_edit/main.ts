import 'materialize-css/dist/css/materialize.min.css'
import 'prismjs/themes/prism.css'
import './style.css'
import { initTinyMce } from '../../helpers/tinymce_init'

const form = document.getElementById('form') as HTMLFormElement
const submit_btn = document.getElementById('submit_btn') as HTMLButtonElement
const select_eyecatch = document.getElementById('select_eyecatch') as HTMLSelectElement
const file_preview = document.getElementById('file_preview') as HTMLImageElement
const add_tag_button = document.getElementById('add_tag_button') as HTMLButtonElement
const select_display = document.getElementById('select_display') as HTMLDivElement
window.addEventListener('DOMContentLoaded', function () {
  initTinyMce()
  if (select_display) {
    select_display.textContent = null
  }
  const selectedTagsValue = (document.getElementsByName('selectedTagsID')[0] as HTMLInputElement).value
  const selectedTags = selectedTagsValue.split(',')
  selectedTags.pop()
  for (let i = 0; i < selectedTags.length; i++) {
    const selectElement = document.createElement('div')
    selectElement.innerHTML = add_tag_button.value
    select_display.appendChild(selectElement)
    selectElement.childNodes[1].addEventListener('click', function () {
      selectElement.parentNode && selectElement.parentNode.removeChild(selectElement)
    })
    for (let j = 0; j < selectElement.childNodes[0].childNodes.length; j++) {
      if ((selectElement.childNodes[0].childNodes[j] as HTMLOptionElement).value === selectedTags[i]) {
        ;(selectElement.childNodes[0].childNodes[j] as HTMLOptionElement).selected = true
        break
      }
    }
  }
})
add_tag_button.addEventListener('click', function () {
  const selectElement = document.createElement('div')
  selectElement.innerHTML = add_tag_button.value
  select_display.appendChild(selectElement)
  selectElement.childNodes[1].addEventListener('click', function () {
    selectElement.parentNode && selectElement.parentNode.removeChild(selectElement)
  })
})
select_eyecatch.addEventListener('change', function () {
  file_preview.src = this.value
})
submit_btn.addEventListener('click', function (e) {
  sendData(e)
})
function sendData(e: MouseEvent) {
  const contentWindow = (document.getElementById('tinymce_body_ifr') as HTMLIFrameElement).contentWindow
  const content = contentWindow && (contentWindow.document.getElementById('tinymce') as HTMLElement).innerHTML
  if (content === null) {
    return
  }
  const rowContent = content && content.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g, '').replace(/\n/g, '')
  const elem_tags = document.getElementsByClassName('elem_tag') as HTMLCollectionOf<HTMLSelectElement>
  if ((document.getElementById('form-title') as HTMLInputElement).value == '') {
    alert('タイトルを入力してください')
    e.preventDefault()
    return
  }
  const arr: { [key: string]: boolean } = {}
  let tags = ''
  for (let i = 0; i < elem_tags.length; i++) {
    if (arr[elem_tags[i].value]) {
      alert('タグが重複しています')
      e.preventDefault()
      return
    }
    arr[elem_tags[i].value] = true
    tags += elem_tags[i].value
    tags += ','
  }
  tags = tags.slice(0, -1)
  const formdata = new FormData(form)
  formdata.append('content', content)
  formdata.append('row_content', rowContent)
  formdata.append('tags', tags)
  const XHR = new XMLHttpRequest()
  XHR.open('PUT', '/admin/save/')
  XHR.send(formdata)
  XHR.onreadystatechange = function () {
    if (XHR.readyState === 4) {
      if (XHR.status === 200) {
        alert('データが更新されました')
        location.href = '/admin/knowledges'
      } else {
        alert('データが正常に送れませんでした')
      }
    }
  }
}

export {}
