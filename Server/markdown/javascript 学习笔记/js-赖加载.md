# js-赖加载

```
function check() {
  let windowHeight = $(window).height();
  let windowScrollTop = $(window).scrollTop();
  let divH = $('#box img').offset().top;
  if (divH < (windowHeight + windowScrollTop) && divH > windowScrollTop) {
    return true;
  }
  return;
}
$(window).on('scroll', () => {
  changeShow()
})

function changeShow() {
  let url = $('img').each(function () {
    let tmpImg = $(this)
    if (isLoaded(tmpImg)) { return };
    if (check(tmpImg)) {
      setTimeout(function () {
        showImg(tmpImg);
      }, 1000)
    }
  })
}
function isLoaded($img) {
  console.log($img.attr('data-src'));
  return $img.attr('data-src') === $img.attr('src') ? true : false //如果data-src和src相等那么就是已经加载过了
}

function showImg(imgName) {
  console.log(imgName)
  imgName.attr('src', imgName.attr('data-src'));
}
```
