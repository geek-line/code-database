const content = document.getElementById('content') as HTMLDivElement
const input = document.getElementById('input') as HTMLInputElement
const like_button_inline = document.querySelectorAll<HTMLButtonElement>('#like_button_inline')
const like_button_baloon = document.getElementById('like_button_baloon') as HTMLButtonElement
const likes_inline = document.querySelectorAll<HTMLSpanElement>('#likes_inline')
const likes_baloon = document.getElementById('likes_baloon') as HTMLSpanElement
const knowledge_id = (document.getElementById('knowledge_id') as HTMLInputElement).value
let snsArea = document.querySelectorAll('.sns-area')
let title = (document.getElementById('title') as HTMLDivElement).innerHTML
let shareUrl = location.href // 現在のページURLを使用する場合 location.href;
let shareText = title + '\n#駆け出しエンジニアと繋がりたい\n#プログラミング初心者' // 現在のページタイトルを使用する場合 document.title;

document.addEventListener('DOMContentLoaded', function () {
  content.innerHTML = content.innerHTML.replace(/<table/g, "<div class='scroll-table'><table").replace(/<\/table>/g, '</table></div>')

  snsArea.forEach(function (Area) {
    generate_share_button(Area, shareUrl, shareText, title)
  })

  if (localStorage.getItem('noLoginLike')) {
    let value = localStorage.getItem('noLoginLike')
    if (value === null) {
      return
    }
    let values = value.split(',')
    for (let i = 0; i < values.length; i++) {
      if (values[i] == knowledge_id) {
        for (let j = 0; j < like_button_inline.length; j++) {
          like_button_inline[j].textContent = 'LIKED'
          like_button_inline[j].classList.add('liked-button')
        }
        like_button_baloon.textContent = 'LIKED'
        like_button_baloon.classList.add('liked-button')
        break
      }
    }
  }
  let p_table_items = document.getElementById('p_table_items') as HTMLDivElement
  let p_table_items_devise = document.getElementById('p_table_items_devise') as HTMLDivElement
  let div = document.createElement('div')
  let matches = document.querySelectorAll<HTMLElement>('.content h2,.content h3')
  matches.forEach(function (value) {
    let id = value.id
    if (id === '') {
      id = value.textContent || ''
      value.id = id
    }
    if (value.tagName === 'H2') {
      let ul = document.createElement('ul')
      let li = document.createElement('li')
      let a = document.createElement('a')
      a.innerHTML = value.textContent || ''
      a.href = '#' + value.id
      a.className = 'h2 sidenav-close'
      li.appendChild(a)
      ul.appendChild(li)
      div.appendChild(ul)
    }
    if (value.tagName === 'H3') {
      let ul = document.createElement('ul')
      let li = document.createElement('li')
      let a = document.createElement('a')
      let lastUl = div.lastElementChild
      if (lastUl === null) {
        return
      }
      let lastLi = lastUl.lastElementChild
      if (lastLi === null) {
        return
      }
      a.innerHTML = '&nbsp; ->' + value.textContent
      a.href = '#' + value.id
      a.className = 'h3 sidenav-close'
      li.appendChild(a)
      ul.appendChild(li)
      lastLi.appendChild(ul)
    }
  })
  p_table_items_devise.appendChild(div)
  p_table_items.innerHTML = p_table_items_devise.innerHTML
  smoothScroll()
  code_pen_init()
  updateCodeSnippet()
})

function code_pen_init() {
  let code_pen = document.getElementsByTagName('a')
  for (let i = 0; i < code_pen.length; i++) {
    const pattern = /https:\/\/codepen.io\/(.*?)/
    if (code_pen[i].href.match(pattern)) {
      code_pen[i].target = '_blank'
      code_pen[i].rel = 'noreferrer noopener'
      let codePen = document.createElement('p')
      codePen.className = 'codepen'
      codePen.setAttribute('data-height', '395')
      codePen.setAttribute(
        'style',
        'height: 265px; box-sizing: border-box; display: flex; align-items: center; justify-content: center; border: 2px solid; margin: 1em 0; padding: 1em ;margin-top: 4rem;'
      )
      const matchedPattern = code_pen[i].href.match('https:\\/\\/codepen.io\\/(.*?)\\/')
      if (matchedPattern) {
        const userName = matchedPattern.map((s) => s.slice(19).slice(0, -1))[0]
        const title = code_pen[i].href.slice(19 + userName.length + 5)
        codePen.setAttribute('data-slug-hash', title)
        codePen.setAttribute('data-user', userName)
        codePen.setAttribute('data-default-tab', 'js,result')
        let a = document.createElement('a')
        a.href = code_pen[i].href
        codePen.appendChild(a)
        const parentNode = code_pen[i].parentNode
        parentNode && parentNode.replaceChild(codePen, code_pen[i])
        const script = document.createElement('script')
        script.async = true
        script.type = 'text/javascript'
        script.src = 'https://production-assets.codepen.io/assets/embed/ei.js'
        document.head.appendChild(script)
      }
    }
  }
}
function updateCodeSnippet() {
  let attachments = document.querySelectorAll<HTMLElement>('pre[class*=language-]')
  for (let attachment of attachments) {
    let pre = document.createElement('pre')
    pre.className = attachment.className
    pre.innerHTML = attachment.innerHTML
    const previousElementSibling = attachment.previousElementSibling
    if (previousElementSibling === null) {
      return
    }
    const matchedPattern = previousElementSibling.textContent || ''.match(/(^タイトル:)+.*/)
    if (matchedPattern) {
      let title = matchedPattern[0].substr(5)
      attachment.innerHTML =
        "<span class='code_title'>" +
        title +
        '</span><br/>' +
        `<button class='copy_clipboard' onclick='copy(event)'>copy</button>` +
        "<pre class='code_displey'>" +
        pre.innerHTML +
        '</pre>'
      previousElementSibling.remove()
    } else {
      attachment.innerHTML =
        "<span class='code_notitle'>" +
        '</span>' +
        `<button class='copy_clipboard' onclick='copy(event)'>copy</button>` +
        "<pre class='code_displey'>" +
        attachment.innerHTML +
        '</pre>'
    }
  }
}

// function copy(e) {
//   let pre = document.createElement('pre')
//   let text = e.path[1].childNodes[0].outerHTML
//   if (text.match(/<span class="code_title">+.*/)) {
//     iziToast.success({ title: 'Copied' })

//     pre = e.path[1].childNodes[3]
//   } else if (text.match(/<span class="code_notitle">+.*/)) {
//     pre = e.path[1].childNodes[2]
//     iziToast.success({ title: 'Copied' })
//   }
//   document.getSelection().selectAllChildren(pre)
//   document.execCommand('copy')
//   document.getSelection().empty(pre)
// }

function generate_share_button(area: Element, url: string, text: string, title: string) {
  let twBtn = document.createElement('div')
  twBtn.className = 'twitter-btn'
  let twHref = 'https://twitter.com/share?text=' + encodeURIComponent(text) + '&url=' + encodeURIComponent(url)
  let twLink =
    '<a href="' +
    twHref +
    '" ' +
    'target="_blank"' +
    ' class = "twitter"><img src="/static/public/twitter.png" ><div class="tweet-text hide-on-small-only">Tweet</div></a>'
  twBtn.innerHTML = twLink
  area.appendChild(twBtn)
}
const smoothScroll = () => {
  let links = document.querySelectorAll('.item_devise a[href^="#"]')
  const speed = 3000 // スクロールスピード
  const divisor = 100 // 分割数
  const tolerance = 5 // 許容誤差
  const headerHeight = 40 // 固定ヘッダーがある場合はその高さ分ずらす
  const interval = divisor / speed
  for (let i = 0; i < links.length; i++) {
    links[i].addEventListener('click', (e) => {
      e.preventDefault()
      let nowY = window.pageYOffset
      const href = links[i].getAttribute('href') //href取得
      if (href) {
        const splitHref = href.split('#')
        const targetID = splitHref[1]
        const target = document.getElementById(targetID)
        if (target !== null) {
          const targetRectTop = target.getBoundingClientRect().top
          const targetY = targetRectTop + nowY - headerHeight
          const minY = Math.abs((targetY - nowY) / divisor)
          doScroll(minY, nowY, targetY, tolerance, interval)
        }
      }
    })
  }
}
const doScroll = (minY: number, nowY: number, targetY: number, tolerance: number, interval: number) => {
  let toY
  if (targetY < nowY) {
    toY = nowY - minY
  } else {
    toY = nowY + minY
  }
  window.scrollTo(0, toY)
  if (targetY - tolerance > toY || toY > targetY + tolerance) {
    window.setTimeout(doScroll, interval, minY, toY, targetY, tolerance, interval)
  } else {
    return false
  }
}
function sendLikeFromBaloon() {
  let values: string[] = []
  let value = localStorage.getItem('noLoginLike')
  if (value) {
    values = value.split(',')
  }
  let isFound = false
  for (let i = 0; i < values.length; i++) {
    if (values[i] == knowledge_id) {
      values.splice(i, 1)
      value = values.join()
      isFound = true
      break
    }
  }
  if (!isFound) {
    values.push(knowledge_id)
    value = values.join()
  }
  const XHR = new XMLHttpRequest()
  let formdata = new FormData(document.getElementById('like_form_baloon') as HTMLFormElement)
  if (isFound) {
    XHR.open('PUT', '/knowledges/like')
  } else {
    XHR.open('POST', '/knowledges/like')
  }
  XHR.onreadystatechange = function () {
    if (XHR.readyState === 4) {
      if (XHR.status === 200) {
        if (isFound) {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) - 1}`
            like_button_inline[i].textContent = 'LIKE'
            like_button_inline[i].classList.remove('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) - 1}`
          like_button_baloon.textContent = 'LIKE'
          like_button_baloon.classList.remove('liked-button')
        } else {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) + 1}`
            like_button_inline[i].textContent = 'LIKED'
            like_button_inline[i].classList.add('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) + 1}`
          like_button_baloon.textContent = 'LIKED'
          like_button_baloon.classList.add('liked-button')
        }
        value && localStorage.setItem('noLoginLike', value)
      } else {
        alert('データが正常に送れませんでした')
      }
    }
  }
  XHR.send(formdata)
}
function sendLikeFromInline() {
  let values: string[] = []
  let value = localStorage.getItem('noLoginLike')
  if (value) {
    values = value.split(',')
  }
  let isFound = false
  for (let i = 0; i < values.length; i++) {
    if (values[i] == knowledge_id) {
      values.splice(i, 1)
      value = values.join()
      isFound = true
      break
    }
  }
  if (!isFound) {
    values.push(knowledge_id)
    value = values.join()
  }
  const XHR = new XMLHttpRequest()
  let formdata = new FormData(document.getElementById('like_form_inline') as HTMLFormElement)
  if (isFound) {
    XHR.open('PUT', '/knowledges/like')
  } else {
    XHR.open('POST', '/knowledges/like')
  }
  XHR.onreadystatechange = function () {
    if (XHR.readyState === 4) {
      if (XHR.status === 200) {
        if (isFound) {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) - 1}`
            like_button_inline[i].textContent = 'LIKE'
            like_button_inline[i].classList.remove('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) - 1}`
          like_button_baloon.textContent = 'LIKE'
          like_button_baloon.classList.remove('liked-button')
        } else {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) + 1}`
            like_button_inline[i].textContent = 'LIKED'
            like_button_inline[i].classList.add('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) + 1}`
          like_button_baloon.textContent = 'LIKED'
          like_button_baloon.classList.add('liked-button')
        }
        value && localStorage.setItem('noLoginLike', value)
      } else {
        alert('データが正常に送れませんでした')
      }
    }
  }
  XHR.send(formdata)
}

export {}
