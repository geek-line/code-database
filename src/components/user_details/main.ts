import iziToast from 'izitoast'
import './style.css'
import 'materialize-css/dist/css/materialize.min.css'
import 'prismjs/themes/prism.css'
import 'izitoast/dist/css/iziToast.min.css'

const content = document.getElementById('content') as HTMLDivElement
const like_button_inlines = document.querySelectorAll<HTMLButtonElement>('#like_button_inline')
const like_button_baloon = document.getElementById('like_button_baloon') as HTMLButtonElement
const likes_inline = document.querySelectorAll<HTMLSpanElement>('#likes_inline')
const likes_baloon = document.getElementById('likes_baloon') as HTMLSpanElement
const knowledge_id = (document.getElementById('knowledge_id') as HTMLInputElement).value
const snsArea = document.querySelectorAll('.sns-area')
const title = (document.getElementById('title') as HTMLDivElement).innerHTML
const shareUrl = location.href // 現在のページURLを使用する場合 location.href;
const shareText = title + '\n#駆け出しエンジニアと繋がりたい\n#プログラミング初心者' // 現在のページタイトルを使用する場合 document.title;

document.addEventListener('DOMContentLoaded', function () {
  content.innerHTML = content.innerHTML.replace(/<table/g, "<div class='scroll-table'><table").replace(/<\/table>/g, '</table></div>')

  snsArea.forEach(function (Area) {
    generate_share_button(Area, shareUrl, shareText)
  })

  if (localStorage.getItem('noLoginLike')) {
    const value = localStorage.getItem('noLoginLike')
    if (value === null) {
      return
    }
    const values = value.split(',')
    for (let i = 0; i < values.length; i++) {
      if (values[i] == knowledge_id) {
        for (let j = 0; j < like_button_inlines.length; j++) {
          like_button_inlines[j].textContent = 'LIKED'
          like_button_inlines[j].classList.add('liked-button')
        }
        like_button_baloon.textContent = 'LIKED'
        like_button_baloon.classList.add('liked-button')
        break
      }
    }
  }
  const p_table_items = document.getElementById('p_table_items') as HTMLDivElement
  const p_table_items_devise = document.getElementById('p_table_items_devise') as HTMLDivElement
  const div = document.createElement('div')
  const matches = document.querySelectorAll<HTMLElement>('.content h2,.content h3')
  matches.forEach(function (value) {
    let id = value.id
    if (id === '') {
      id = value.textContent || ''
      value.id = id
    }
    if (value.tagName === 'H2') {
      const ul = document.createElement('ul')
      const li = document.createElement('li')
      const a = document.createElement('a')
      a.innerHTML = value.textContent || ''
      a.href = '#' + value.id
      a.className = 'h2 sidenav-close'
      li.appendChild(a)
      ul.appendChild(li)
      div.appendChild(ul)
    }
    if (value.tagName === 'H3') {
      const ul = document.createElement('ul')
      const li = document.createElement('li')
      const a = document.createElement('a')
      const lastUl = div.lastElementChild
      if (lastUl === null) {
        return
      }
      const lastLi = lastUl.lastElementChild
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
  const code_pen = document.getElementsByTagName('a')
  for (let i = 0; i < code_pen.length; i++) {
    const pattern = /https:\/\/codepen.io\/(.*?)/
    if (code_pen[i].href.match(pattern)) {
      code_pen[i].target = '_blank'
      code_pen[i].rel = 'noreferrer noopener'
      const codePen = document.createElement('p')
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
        const a = document.createElement('a')
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
  const attachments = document.querySelectorAll<HTMLElement>('pre[class*=language-]')
  for (const attachment of attachments) {
    const pre = document.createElement('pre')
    pre.className = attachment.className
    pre.innerHTML = attachment.innerHTML
    const previousElementSibling = attachment.previousElementSibling
    if (previousElementSibling === null) {
      return
    }
    const matchedPattern = (previousElementSibling.textContent || '').match(/(^タイトル:)+.*/)
    if (matchedPattern) {
      const title = matchedPattern[0].substr(5)
      attachment.innerHTML =
        "<span class='code_title'>" +
        title +
        '</span><br/>' +
        `<button id="copy_button" class='copy_clipboard'>copy</button>` +
        "<pre class='code_displey'>" +
        pre.innerHTML +
        '</pre>'
      previousElementSibling.remove()
    } else {
      attachment.innerHTML =
        "<span class='code_notitle'>" +
        '</span>' +
        `<button id="copy_button" class='copy_clipboard'>copy</button>` +
        "<pre class='code_displey'>" +
        attachment.innerHTML +
        '</pre>'
    }
  }
  const copy_buttons = document.querySelectorAll<HTMLButtonElement>('#copy_button')
  copy_buttons.forEach((button) => {
    button.addEventListener('click', copy)
  })
}

function copy(e: MouseEvent) {
  let pre = document.createElement('pre') as Node
  const path = e.composedPath()[1] as HTMLElement
  const text = (path.childNodes[0] as HTMLElement).outerHTML
  if (text.match(/<span class="code_title">+.*/)) {
    iziToast.success({ title: 'Copied' })

    pre = path.childNodes[3]
  } else if (text.match(/<span class="code_notitle">+.*/)) {
    pre = path.childNodes[2]
    iziToast.success({ title: 'Copied' })
  }
  const selection = document.getSelection()
  selection && selection.selectAllChildren(pre)
  document.execCommand('copy')
  selection && selection.empty()
}

function generate_share_button(area: Element, url: string, text: string) {
  const twBtn = document.createElement('div')
  twBtn.className = 'twitter-btn'
  const twHref = 'https://twitter.com/share?text=' + encodeURIComponent(text) + '&url=' + encodeURIComponent(url)
  const twLink =
    '<a href="' +
    twHref +
    '" ' +
    'target="_blank"' +
    ' class = "twitter"><img src="/public/twitter.png" ><div class="tweet-text hide-on-small-only">Tweet</div></a>'
  twBtn.innerHTML = twLink
  area.appendChild(twBtn)
}
const smoothScroll = () => {
  const links = document.querySelectorAll('.item_devise a[href^="#"]')
  const speed = 3000 // スクロールスピード
  const divisor = 100 // 分割数
  const tolerance = 5 // 許容誤差
  const headerHeight = 40 // 固定ヘッダーがある場合はその高さ分ずらす
  const interval = divisor / speed
  for (let i = 0; i < links.length; i++) {
    links[i].addEventListener('click', (e) => {
      e.preventDefault()
      const nowY = window.pageYOffset
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

like_button_baloon.addEventListener('click', sendLikeFromBaloon)
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
  const formdata = new FormData(document.getElementById('like_form_baloon') as HTMLFormElement)
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
            like_button_inlines[i].textContent = 'LIKE'
            like_button_inlines[i].classList.remove('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) - 1}`
          like_button_baloon.textContent = 'LIKE'
          like_button_baloon.classList.remove('liked-button')
        } else {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) + 1}`
            like_button_inlines[i].textContent = 'LIKED'
            like_button_inlines[i].classList.add('liked-button')
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

like_button_inlines.forEach((button) => {
  button.addEventListener('click', sendLikeFromInline)
})
function sendLikeFromInline() {
  const value = localStorage.getItem('noLoginLike')
  const values = value ? value.split(',') : []
  let isFound = false
  let newValue: string
  for (let i = 0; i < values.length; i++) {
    if (values[i] == knowledge_id) {
      values.splice(i, 1)
      newValue = values.join()
      isFound = true
      break
    }
  }
  if (!isFound) {
    values.push(knowledge_id)
    newValue = values.join()
  }
  const XHR = new XMLHttpRequest()
  const formdata = new FormData(document.getElementById('like_form_inline') as HTMLFormElement)
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
            like_button_inlines[i].textContent = 'LIKE'
            like_button_inlines[i].classList.remove('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) - 1}`
          like_button_baloon.textContent = 'LIKE'
          like_button_baloon.classList.remove('liked-button')
        } else {
          for (let i = 0; i < likes_inline.length; i++) {
            likes_inline[i].textContent = `${Number(likes_inline[i].textContent) + 1}`
            like_button_inlines[i].textContent = 'LIKED'
            like_button_inlines[i].classList.add('liked-button')
          }
          likes_baloon.textContent = `${Number(likes_baloon.textContent) + 1}`
          like_button_baloon.textContent = 'LIKED'
          like_button_baloon.classList.add('liked-button')
        }
        localStorage.setItem('noLoginLike', newValue)
      } else {
        alert('データが正常に送れませんでした')
      }
    }
  }
  XHR.send(formdata)
}
