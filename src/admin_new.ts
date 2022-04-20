const form = document.getElementById('form') as HTMLFormElement
const add_tag_button = document.getElementById('add_tag_button') as HTMLButtonElement
const submit_button = document.getElementById('submit_button') as HTMLInputElement
const select_eyecatch = document.getElementById('select_eyecatch') as HTMLSelectElement
const file_preview = document.getElementById('file_preview') as HTMLImageElement
const select_display = document.getElementById('select_display') as HTMLDivElement
window.addEventListener('DOMContentLoaded', function () {
  select_display.textContent = null
  file_preview.src = select_eyecatch.value
})
add_tag_button.addEventListener('click', function () {
  const selectElement = document.createElement('div')
  selectElement.innerHTML = add_tag_button.value
  select_display.appendChild(selectElement)
  selectElement.childNodes[1].addEventListener('click', function () {
    selectElement.parentNode && selectElement.parentNode.removeChild(selectElement)
  })
})
select_eyecatch.addEventListener('change', function (e) {
  file_preview.src = this.value
})
submit_button.addEventListener('click', function (e) {
  const contentWindow = (document.getElementById('tinymce_body_ifr') as HTMLIFrameElement).contentWindow
  const content = contentWindow && (document.getElementById('tinymce') as HTMLElement).innerHTML
  if (content === null) {
    return
  }
  const rowContent = content.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g, '').replace(/\n/g, '')
  const elem_tags = document.getElementsByClassName('elem_tag') as HTMLCollectionOf<HTMLSelectElement>
  if ((document.getElementById('form-title') as HTMLInputElement).value == '') {
    alert('タイトルを入力してください')
    e.preventDefault()
    return
  }
  let arr: { [key: string]: boolean } = {}
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
  let formdata = new FormData(form)
  formdata.append('content', content)
  formdata.append('row_content', rowContent)
  formdata.append('tags', tags)
  const XHR = new XMLHttpRequest()
  XHR.open('POST', '/admin/save/')
  XHR.send(formdata)
  XHR.onreadystatechange = function () {
    if (XHR.readyState === 4) {
      if (XHR.status === 200) {
        location.href = '/admin/knowledges'
      } else {
        alert('データが正常に送れませんでした')
      }
    }
  }
})

export {}
