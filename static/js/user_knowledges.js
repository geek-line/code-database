const search_input = document.querySelectorAll('#search_input')
const search_submit = document.querySelectorAll('#search_submit')

for (let i = 0; i < search_submit.length; i++) {
  search_input[i].addEventListener('keydown', function () {
    if (window.event.keyCode == 13) {
      search_submit[i].click()
    }
  })
  search_submit[i].addEventListener('click', function (e) {
    if (!search_input[i].value) {
      e.preventDefault()
      return
    }
    const XHR = new XMLHttpRequest()
    let queries = search_input[i].value.split(/\s+/g)
    for (let j = 0; j < queries.length; j++) {
      queries[j] = encodeURIComponent(queries[j])
    }
    const qvalue = queries.join('+')
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
