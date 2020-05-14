const content = document.getElementById('content')
const input = document.getElementById('input')
const like_button_inline = document.querySelectorAll('#like_button_inline')
const like_button_baloon = document.getElementById('like_button_baloon')
const likes_inline = document.querySelectorAll('#likes_inline')
const likes_baloon = document.getElementById('likes_baloon')
const knowledge_id = document.getElementById('knowledge_id').value
let snsArea = document.querySelectorAll('.sns-area');
let title = document.getElementById('title').innerHTML;
let shareUrl = location.href; // 現在のページURLを使用する場合 location.href;
let shareText = title+'\n#駆け出しエンジニアと繋がりたい\n#プログラミング初心者'; // 現在のページタイトルを使用する場合 document.title;

document.addEventListener('DOMContentLoaded', function () {
    snsArea.forEach(function(Area){
        generate_share_button(Area, shareUrl, shareText,title);
    }) 
   
    if (localStorage.getItem('noLoginLike')) {
        let value = localStorage.getItem('noLoginLike')
        let values = value.split(',')
        for (let i = 0; i < values.length; i++) {
            if (values[i] == knowledge_id) {
                for (let j = 0;j < like_button_inline.length; j++){
                    like_button_inline[j].textContent = 'LIKED'
                    like_button_inline[j].classList.add('liked-button')   
                }
                like_button_baloon.textContent = 'LIKED'
                like_button_baloon.classList.add('liked-button')   
                break
            }
        }
    }

    content.innerHTML = input.value.replace(/<table/g, "<div class='scroll-table'><table").replace(/<\/table>/g, "</table></div>")
    var p_table_items = document.getElementById("p_table_items"); // 目次を追加する先(table of contents)
    var p_table_items_devise = document.getElementById("p_table_items_devise"); // 目次を追加する先(table of contents)
    var div = document.createElement('div'); // 作成する目次のコンテナ要素
    // .entry-content配下のh2、h3要素を全て取得する
    var matches = document.querySelectorAll('.content h2,.content h3');
    // .entry-content配下のh2、h3要素を全て取得する
    // 取得した見出しタグ要素の数だけ以下の操作を繰り返す
    matches.forEach(function (value, i) {
        // 見出しタグ要素のidを取得し空の場合は内容をidにする
        var id = value.id;
        if (id === '') {
            id = value.textContent;
            value.id = id;
        }
        // 要素がh2タグの場合
        if (value.tagName === 'H2') {
            var ul = document.createElement('ul');
            var li = document.createElement('li');
            var a = document.createElement('a');
            // 追加する<ul><li><a>タイトル</a></li></ul>を準備する
            a.innerHTML = value.textContent;
            a.href = '#' + value.id;
            a.className = "h2 sidenav-close"
            li.appendChild(a)
            ul.appendChild(li);
            // コンテナ要素である<div>の中に要素を追加する
            div.appendChild(ul);
        }
        // 要素がh3タグの場合
        if (value.tagName === 'H3') {
            var ul = document.createElement('ul');
            var li = document.createElement('li');
            var a = document.createElement('a');
            // コンテナ要素である<div>の中から最後の<li>を取得する
            var lastUl = div.lastElementChild;
            var lastLi = lastUl.lastElementChild;
            // 追加する<ul><li><a>タイトル</a></li></ul>を準備する
            a.innerHTML = '&nbsp; ->' + value.textContent;
            a.href = '#' + value.id;
            a.className = "h3 sidenav-close"
            li.appendChild(a)
            ul.appendChild(li);
            // 最後の<li>の中に要素を追加する
            lastLi.appendChild(ul);
        }
    });
    p_table_items_devise.appendChild(div);
    p_table_items.innerHTML = p_table_items_devise.innerHTML
    var elems = document.querySelectorAll('.sidenav');
    var instances = M.Sidenav.init(elems,{draggable:true,edge:'right'});
    smoothScroll();
    code_pen_init();
}); 
 
function code_pen_init(){
    let code_pen = document.getElementsByTagName('a')
    let code_pen_link= [];
    for(let i =0;i<code_pen.length;i++){
        if(code_pen[i].text=="codepen"){
            code_pen_link[i] = code_pen[i].getAttribute("href"); 
            code_pen[i].innerHTML = `<p class="codepen" data-height="265" data-theme-id="dark" data-default-tab="html,js,css,result" data-user="codedatabase" data-slug-hash="NWqpdOv" style="height: 265px; box-sizing: border-box; display: flex; align-items: center; justify-content: center; border: 2px solid; margin: 1em 0; padding: 1em;" data-pen-title="id_41">
            <span>See the Pen <a href=`+code_pen_link[i]+`>
            id_41</a> by Code Database team
            on <a href="https://codepen.io">CodePen</a>.</span>
            </p>`
        }
    }
}
// シェアボタンを生成する関数
function generate_share_button(area, url, text,title) {
    // シェアボタンの作成
    let twBtn = document.createElement('div');
    twBtn.className = 'twitter-btn';
    // 各シェアボタンのリンク先
    let twHref = 'https://twitter.com/share?text='+encodeURIComponent(text)+'&url='+encodeURIComponent(url);
    // シェアボタンにリンクを追加
    let twLink = '<a href="' + twHref + '" ' + 'target="_blank"'+ ' class = "twitter"><img src="/static/public/twitter.png" ><div class="tweet-text hide-on-small-only">Tweet</div></a>';
    twBtn.innerHTML = twLink;
    // シェアボタンを表示
    area.appendChild(twBtn);
}

const smoothScroll = () =>{
    let links = document.querySelectorAll('.item_devise a[href^="#"]');
    const speed = 3000;          // スクロールスピード   
    const divisor = 100;        // 分割数
    const tolerance = 5;        // 許容誤差
    const headerHeight = 40;     // 固定ヘッダーがある場合はその高さ分ずらす
    const interval = divisor / speed;
    for(let i = 0; i < links.length; i++){
      links[i].addEventListener('click',(e)=>{
        e.preventDefault();
        let nowY = window.pageYOffset;
        const href = e.currentTarget.getAttribute('href');   //href取得
        const splitHref = href.split('#');
        const targetID = splitHref[1];
        const target = document.getElementById(targetID);  
        if( target != null){
          const targetRectTop = target.getBoundingClientRect().top;
          const targetY = targetRectTop + nowY - headerHeight;
          const minY = Math.abs((targetY - nowY)/divisor);
          doScroll(minY,nowY,targetY,tolerance,interval);
        }
      });
    }
}
  
const doScroll = (minY,nowY,targetY,tolerance,interval) =>{
    let toY ;
    if( targetY < nowY ){
        toY = nowY - minY;
    }else{
        toY = nowY + minY;
    }
    window.scrollTo(0, toY);
    if( targetY - tolerance > toY || toY > targetY + tolerance){
    window.setTimeout(doScroll,interval,minY,toY,targetY,tolerance,interval);
    }else{
    return false;
    }
}

function sendLikeFromBaloon() {
    let values = []
    let value = ''
    if (localStorage.getItem('noLoginLike')) {
        value = localStorage.getItem('noLoginLike')
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
    let formdata = new FormData(document.getElementById('like_form_baloon'))
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
                        likes_inline[i].textContent = Number(likes_inline[i].textContent) - 1
                        like_button_inline[i].textContent = 'LIKE'
                        like_button_inline[i].classList.remove('liked-button')
                    }
                    likes_baloon.textContent = Number(likes_baloon.textContent) - 1
                    like_button_baloon.textContent = 'LIKE'
                    like_button_baloon.classList.remove('liked-button')
                } else {
                    for (let i = 0; i < likes_inline.length; i++) {
                        likes_inline[i].textContent = Number(likes_inline[i].textContent) + 1
                        like_button_inline[i].textContent = 'LIKED'
                        like_button_inline[i].classList.add('liked-button')
                    }
                    likes_baloon.textContent = Number(likes_baloon.textContent) + 1
                    like_button_baloon.textContent = 'LIKED'
                    like_button_baloon.classList.add('liked-button')
                }
                localStorage.setItem('noLoginLike', value)
            } else {
                alert('データが正常に送れませんでした')
            }
        }
    }
    XHR.send(formdata)
}
function sendLikeFromInline() {
    let values = []
    let value = ''
    if (localStorage.getItem('noLoginLike')) {
        value = localStorage.getItem('noLoginLike')
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
    let formdata = new FormData(document.getElementById('like_form_inline'))
    if (isFound) {
        XHR.open('PUT', '/knowledges/like')
    } else {
        XHR.open('POST', '/knowledges/like')
    }
    XHR.onreadystatechange = function () {
        if (XHR.readyState === 4) {
            if (XHR.status === 200) {
                if (isFound) {
                    for (let i = 0; i < likes_inline.length; i++){
                        likes_inline[i].textContent = Number(likes_inline[i].textContent) - 1
                        like_button_inline[i].textContent = 'LIKE'
                        like_button_inline[i].classList.remove('liked-button')
                    }
                    likes_baloon.textContent = Number(likes_baloon.textContent) - 1
                    like_button_baloon.textContent = 'LIKE'
                    like_button_baloon.classList.remove('liked-button')
                } else {
                    for (let i = 0; i < likes_inline.length; i++){
                        likes_inline[i].textContent = Number(likes_inline[i].textContent) + 1
                        like_button_inline[i].textContent = 'LIKED'
                        like_button_inline[i].classList.add('liked-button')
                    }
                    likes_baloon.textContent = Number(likes_baloon.textContent) + 1
                    like_button_baloon.textContent = 'LIKED'
                    like_button_baloon.classList.add('liked-button')
                }
                localStorage.setItem('noLoginLike', value)
            } else {
                alert('データが正常に送れませんでした')
            }
        }
    }
    XHR.send(formdata)
}